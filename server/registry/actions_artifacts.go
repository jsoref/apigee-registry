// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"context"
	"fmt"
	"strings"

	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/registry/internal/storage"
	"github.com/apigee/registry/server/registry/internal/storage/models"
	"github.com/apigee/registry/server/registry/names"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type artifactParent interface {
	Artifact(string) names.Artifact
}

func parseArtifactParent(name string) (artifactParent, error) {
	if s, err := names.ParseSpec(name); err == nil {
		return s, nil
	} else if v, err := names.ParseVersion(name); err == nil {
		return v, nil
	} else if d, err := names.ParseDeployment(name); err == nil {
		return d, nil
	} else if a, err := names.ParseApi(name); err == nil {
		return a, nil
	} else if p, err := names.ParseProjectWithLocation(name); err == nil {
		return p, nil
	}

	return nil, fmt.Errorf("invalid artifact parent %q", name)
}

// CreateArtifact handles the corresponding API request.
func (s *RegistryServer) CreateArtifact(ctx context.Context, req *rpc.CreateArtifactRequest) (*rpc.Artifact, error) {
	// Parent name must be valid.
	parent, err := parseArtifactParent(req.GetParent())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	// Artifact body must be nonempty.
	if req.GetArtifact() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid artifact %+v: body must be provided", req.GetArtifact())
	}
	// Artifact name must be valid.
	name := parent.Artifact(req.GetArtifactId())
	if err := name.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var response *rpc.Artifact
	if err := s.runInTransaction(ctx, func(ctx context.Context, db *storage.Client) error {
		// Creation should only succeed when the parent exists.
		var err error
		switch parent := parent.(type) {
		case names.Project:
			_, err = db.GetProject(ctx, parent)
		case names.Api:
			_, err = db.GetApi(ctx, parent)
		case names.Version:
			_, err = db.GetVersion(ctx, parent)
		case names.Spec:
			_, err = db.GetSpec(ctx, parent)
		case names.Deployment:
			_, err = db.GetDeployment(ctx, parent)
		}
		if err != nil {
			return err
		}
		artifact, err := models.NewArtifact(name, req.GetArtifact())
		if err != nil {
			return err
		}
		if err := db.CreateArtifact(ctx, artifact); err != nil {
			return err
		}
		if err := db.SaveArtifactContents(ctx, artifact, req.Artifact.GetContents()); err != nil {
			return err
		}
		response = artifact.Message()
		return nil
	}); err != nil {
		return nil, err
	}

	s.notify(ctx, rpc.Notification_CREATED, response.GetName())
	return response, nil
}

// DeleteArtifact handles the corresponding API request.
func (s *RegistryServer) DeleteArtifact(ctx context.Context, req *rpc.DeleteArtifactRequest) (*emptypb.Empty, error) {
	// Artifact name must be valid.
	name, err := names.ParseArtifact(req.GetName())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.runInTransaction(ctx, func(ctx context.Context, db *storage.Client) error {
		return db.DeleteArtifact(ctx, name)
	}); err != nil {
		return nil, err
	}
	s.notify(ctx, rpc.Notification_DELETED, req.GetName())
	return &emptypb.Empty{}, nil
}

// GetArtifact handles the corresponding API request.
func (s *RegistryServer) GetArtifact(ctx context.Context, req *rpc.GetArtifactRequest) (*rpc.Artifact, error) {
	db, err := s.getStorageClient(ctx)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	name, err := names.ParseArtifact(req.GetName())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	artifact, err := db.GetArtifact(ctx, name)
	if err != nil {
		return nil, err
	}

	return artifact.Message(), nil
}

// GetArtifactContents handles the corresponding API request.
func (s *RegistryServer) GetArtifactContents(ctx context.Context, req *rpc.GetArtifactContentsRequest) (*httpbody.HttpBody, error) {
	db, err := s.getStorageClient(ctx)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	name, err := names.ParseArtifact(req.GetName())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	artifact, err := db.GetArtifact(ctx, name)
	if err != nil {
		return nil, err
	}

	blob, err := db.GetArtifactContents(ctx, name)
	if err != nil {
		return nil, err
	}

	if strings.Contains(artifact.MimeType, "+gzip") {
		artifact.MimeType = strings.ReplaceAll(artifact.MimeType, "+gzip", "")
		blob.Contents, err = models.GUnzippedBytes(blob.Contents)
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "failed to unzip contents with gzip MIME type: %s", err)
		}
	}

	return &httpbody.HttpBody{
		ContentType: artifact.MimeType,
		Data:        blob.Contents,
	}, nil
}

// ListArtifacts handles the corresponding API request.
func (s *RegistryServer) ListArtifacts(ctx context.Context, req *rpc.ListArtifactsRequest) (*rpc.ListArtifactsResponse, error) {
	db, err := s.getStorageClient(ctx)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	if req.GetPageSize() < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid page_size %d: must not be negative", req.GetPageSize())
	} else if req.GetPageSize() > 1000 {
		req.PageSize = 1000
	} else if req.GetPageSize() == 0 {
		req.PageSize = 50
	}

	parent, err := parseArtifactParent(req.GetParent())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var listing storage.ArtifactList
	switch parent := parent.(type) {
	case names.Project:
		listing, err = db.ListProjectArtifacts(ctx, parent, storage.PageOptions{
			Size:   req.GetPageSize(),
			Filter: req.GetFilter(),
			Order:  req.GetOrderBy(),
			Token:  req.GetPageToken(),
		})
	case names.Api:
		listing, err = db.ListApiArtifacts(ctx, parent, storage.PageOptions{
			Size:   req.GetPageSize(),
			Filter: req.GetFilter(),
			Order:  req.GetOrderBy(),
			Token:  req.GetPageToken(),
		})
	case names.Version:
		listing, err = db.ListVersionArtifacts(ctx, parent, storage.PageOptions{
			Size:   req.GetPageSize(),
			Filter: req.GetFilter(),
			Order:  req.GetOrderBy(),
			Token:  req.GetPageToken(),
		})
	case names.Spec:
		listing, err = db.ListSpecArtifacts(ctx, parent, storage.PageOptions{
			Size:   req.GetPageSize(),
			Filter: req.GetFilter(),
			Order:  req.GetOrderBy(),
			Token:  req.GetPageToken(),
		})
	case names.Deployment:
		listing, err = db.ListDeploymentArtifacts(ctx, parent, storage.PageOptions{
			Size:   req.GetPageSize(),
			Filter: req.GetFilter(),
			Order:  req.GetOrderBy(),
			Token:  req.GetPageToken(),
		})
	}
	if err != nil {
		return nil, err
	}

	response := &rpc.ListArtifactsResponse{
		Artifacts:     make([]*rpc.Artifact, len(listing.Artifacts)),
		NextPageToken: listing.Token,
	}

	for i, artifact := range listing.Artifacts {
		response.Artifacts[i] = artifact.Message()
	}

	return response, nil
}

// ReplaceArtifact handles the corresponding API request.
func (s *RegistryServer) ReplaceArtifact(ctx context.Context, req *rpc.ReplaceArtifactRequest) (*rpc.Artifact, error) {
	db, err := s.getStorageClient(ctx)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	name, err := names.ParseArtifact(req.Artifact.GetName())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var artifact *models.Artifact
	err = db.Transaction(ctx, func(ctx context.Context, db *storage.Client) error {
		db.LockArtifacts(ctx)
		// Replacement should only succeed on artifacts that currently exist.
		if _, err = db.GetArtifact(ctx, name); err != nil {
			return err
		}
		artifact, err = models.NewArtifact(name, req.GetArtifact())
		if err != nil {
			return err
		}
		if err := db.SaveArtifact(ctx, artifact); err != nil {
			return err
		}
		return db.SaveArtifactContents(ctx, artifact, req.Artifact.GetContents())
	})
	if err != nil {
		return nil, err
	}

	s.notify(ctx, rpc.Notification_UPDATED, name.String())
	return artifact.Message(), nil
}
