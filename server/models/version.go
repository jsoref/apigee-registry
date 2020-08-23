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

package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/apigee/registry/rpc"
	"github.com/apigee/registry/server/names"
	ptypes "github.com/golang/protobuf/ptypes"
)

// VersionEntityName is used to represent versions in storage.
const VersionEntityName = "Version"

// Version ...
type Version struct {
	ProjectID   string    // Uniquely identifies a project.
	ApiID       string    // Uniquely identifies a api within a project.
	VersionID   string    // Uniquely identifies a version wihtin a api.
	DisplayName string    // A human-friendly name.
	Description string    // A detailed description.
	CreateTime  time.Time // Creation time.
	UpdateTime  time.Time // Time of last change.
	State       string    // Lifecycle stage.
}

// ParseParentApi ...
func ParseParentApi(parent string) ([]string, error) {
	r := regexp.MustCompile("^projects/" + names.NameRegex +
		"/apis/" + names.NameRegex +
		"$")
	m := r.FindAllStringSubmatch(parent, -1)
	if m == nil {
		return nil, fmt.Errorf("invalid parent '%s'", parent)
	}
	return m[0], nil
}

// NewVersionFromParentAndVersionID returns an initialized api for a specified parent and apiID.
func NewVersionFromParentAndVersionID(parent string, versionID string) (*Version, error) {
	r := regexp.MustCompile("^projects/" + names.NameRegex + "/apis/" + names.NameRegex + "$")
	m := r.FindAllStringSubmatch(parent, -1)
	if m == nil {
		return nil, fmt.Errorf("invalid api '%s'", parent)
	}
	if err := names.ValidateID(versionID); err != nil {
		return nil, err
	}
	version := &Version{}
	version.ProjectID = m[0][1]
	version.ApiID = m[0][2]
	version.VersionID = versionID
	return version, nil
}

// NewVersionFromResourceName parses resource names and returns an initialized version.
func NewVersionFromResourceName(name string) (*Version, error) {
	version := &Version{}
	m := names.VersionRegexp().FindAllStringSubmatch(name, -1)
	if m == nil {
		return nil, fmt.Errorf("invalid version name (%s)", name)
	}
	version.ProjectID = m[0][1]
	version.ApiID = m[0][2]
	version.VersionID = m[0][3]
	return version, nil
}

// NewVersionFromMessage returns an initialized version from a message.
func NewVersionFromMessage(message *rpc.Version) (*Version, error) {
	version, err := NewVersionFromResourceName(message.GetName())
	if err != nil {
		return nil, err
	}
	version.DisplayName = message.GetDisplayName()
	version.Description = message.GetDescription()
	version.State = message.GetState()
	return version, nil
}

// ResourceName generates the resource name of a version.
func (version *Version) ResourceName() string {
	return fmt.Sprintf("projects/%s/apis/%s/versions/%s", version.ProjectID, version.ApiID, version.VersionID)
}

// Message returns a message representing a version.
func (version *Version) Message() (message *rpc.Version, err error) {
	message = &rpc.Version{}
	message.Name = version.ResourceName()
	message.DisplayName = version.DisplayName
	message.Description = version.Description
	message.CreateTime, err = ptypes.TimestampProto(version.CreateTime)
	message.UpdateTime, err = ptypes.TimestampProto(version.UpdateTime)
	message.State = version.State
	return message, err
}

// Update modifies a version using the contents of a message.
func (version *Version) Update(message *rpc.Version) error {
	version.DisplayName = message.GetDisplayName()
	version.Description = message.GetDescription()
	version.State = message.GetState()
	version.UpdateTime = time.Now()
	return nil
}
