// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gapic_test

import (
	"context"

	gapic "github.com/apigee/registry/gapic"
	rpcpb "github.com/apigee/registry/rpc"
	"google.golang.org/api/iterator"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func ExampleNewRegistryClient() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleRegistryClient_GetStatus() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb#Empty.
	}
	resp, err := c.GetStatus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_ListProjects() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ListProjectsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ListProjectsRequest.
	}
	it := c.ListProjects(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_GetProject() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetProjectRequest.
	}
	resp, err := c.GetProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_CreateProject() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.CreateProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#CreateProjectRequest.
	}
	resp, err := c.CreateProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_UpdateProject() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.UpdateProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#UpdateProjectRequest.
	}
	resp, err := c.UpdateProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_DeleteProject() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.DeleteProjectRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#DeleteProjectRequest.
	}
	err = c.DeleteProject(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRegistryClient_ListApis() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ListApisRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ListApisRequest.
	}
	it := c.ListApis(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_GetApi() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetApiRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetApiRequest.
	}
	resp, err := c.GetApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_CreateApi() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.CreateApiRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#CreateApiRequest.
	}
	resp, err := c.CreateApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_UpdateApi() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.UpdateApiRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#UpdateApiRequest.
	}
	resp, err := c.UpdateApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_DeleteApi() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.DeleteApiRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#DeleteApiRequest.
	}
	err = c.DeleteApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRegistryClient_ListApiVersions() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ListApiVersionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ListApiVersionsRequest.
	}
	it := c.ListApiVersions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_GetApiVersion() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetApiVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetApiVersionRequest.
	}
	resp, err := c.GetApiVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_CreateApiVersion() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.CreateApiVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#CreateApiVersionRequest.
	}
	resp, err := c.CreateApiVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_UpdateApiVersion() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.UpdateApiVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#UpdateApiVersionRequest.
	}
	resp, err := c.UpdateApiVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_DeleteApiVersion() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.DeleteApiVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#DeleteApiVersionRequest.
	}
	err = c.DeleteApiVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRegistryClient_ListApiSpecs() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ListApiSpecsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ListApiSpecsRequest.
	}
	it := c.ListApiSpecs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_GetApiSpec() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetApiSpecRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetApiSpecRequest.
	}
	resp, err := c.GetApiSpec(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_GetApiSpecContents() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetApiSpecContentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetApiSpecContentsRequest.
	}
	resp, err := c.GetApiSpecContents(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_CreateApiSpec() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.CreateApiSpecRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#CreateApiSpecRequest.
	}
	resp, err := c.CreateApiSpec(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_UpdateApiSpec() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.UpdateApiSpecRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#UpdateApiSpecRequest.
	}
	resp, err := c.UpdateApiSpec(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_DeleteApiSpec() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.DeleteApiSpecRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#DeleteApiSpecRequest.
	}
	err = c.DeleteApiSpec(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRegistryClient_TagApiSpecRevision() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.TagApiSpecRevisionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#TagApiSpecRevisionRequest.
	}
	resp, err := c.TagApiSpecRevision(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_ListApiSpecRevisions() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ListApiSpecRevisionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ListApiSpecRevisionsRequest.
	}
	it := c.ListApiSpecRevisions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_RollbackApiSpec() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.RollbackApiSpecRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#RollbackApiSpecRequest.
	}
	resp, err := c.RollbackApiSpec(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_DeleteApiSpecRevision() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.DeleteApiSpecRevisionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#DeleteApiSpecRevisionRequest.
	}
	err = c.DeleteApiSpecRevision(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRegistryClient_ListArtifacts() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ListArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ListArtifactsRequest.
	}
	it := c.ListArtifacts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleRegistryClient_GetArtifact() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetArtifactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetArtifactRequest.
	}
	resp, err := c.GetArtifact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_GetArtifactContents() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.GetArtifactContentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#GetArtifactContentsRequest.
	}
	resp, err := c.GetArtifactContents(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_CreateArtifact() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.CreateArtifactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#CreateArtifactRequest.
	}
	resp, err := c.CreateArtifact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_ReplaceArtifact() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.ReplaceArtifactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#ReplaceArtifactRequest.
	}
	resp, err := c.ReplaceArtifact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegistryClient_DeleteArtifact() {
	ctx := context.Background()
	c, err := gapic.NewRegistryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &rpcpb.DeleteArtifactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/apigee/registry/rpc#DeleteArtifactRequest.
	}
	err = c.DeleteArtifact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}