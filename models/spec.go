// Copyright 2020 Google Inc. All Rights Reserved.

package models

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	rpc "apigov.dev/flame/rpc"
	ptypes "github.com/golang/protobuf/ptypes"
)

// Spec ...
type Spec struct {
	ProjectID   string    // Uniquely identifies a project.
	ProductID   string    // Uniquely identifies a product within a project.
	VersionID   string    // Uniquely identifies a version within a product.
	SpecID      string    // Uniquely identifies a spec within a version.
	DisplayName string    // A human-friendly name.
	Description string    // A detailed description.
	CreateTime  time.Time // Creation time.
	UpdateTime  time.Time // Time of last change.
	Style       string    // Specification format.
}

// NewSpecFromResourceName parses resource names and returns an initialized spec.
func NewSpecFromResourceName(name string) (*Spec, error) {
	spec := &Spec{}
	r := regexp.MustCompile("^/projects/" + nameRegex + "/products/" + nameRegex + "/versions/" + nameRegex + "/specs/" + nameRegex + "$")
	m := r.FindAllStringSubmatch(name, -1)
	if m == nil {
		return nil, errors.New("invalid spec name")
	}
	spec.ProjectID = m[0][1]
	spec.ProductID = m[0][2]
	spec.VersionID = m[0][3]
	spec.SpecID = m[0][4]
	return spec, nil
}

// NewSpecFromMessage returns an initialized spec from a message.
func NewSpecFromMessage(message *rpc.Spec) (*Spec, error) {
	spec, err := NewSpecFromResourceName(message.GetName())
	if err != nil {
		return nil, err
	}
	spec.DisplayName = message.GetDisplayName()
	spec.Description = message.GetDescription()
	//spec.Availability = message.GetAvailability()
	//spec.RecommendedVersion = message.GetRecommendedVersion()
	return spec, nil
}

// ResourceName generates the resource name of a spec.
func (spec *Spec) ResourceName() string {
	return fmt.Sprintf("/projects/%s/products/%s/versions/%s/specs/%s", spec.ProjectID, spec.ProductID, spec.VersionID, spec.SpecID)
}

// Message returns a message representing a spec.
func (spec *Spec) Message() (message *rpc.Spec, err error) {
	message = &rpc.Spec{}
	message.Name = spec.ResourceName()
	message.DisplayName = spec.DisplayName
	message.Description = spec.Description
	message.CreateTime, err = ptypes.TimestampProto(spec.CreateTime)
	message.UpdateTime, err = ptypes.TimestampProto(spec.UpdateTime)
	//message.Availability = spec.Availability
	//message.RecommendedVersion = spec.RecommendedVersion
	return message, err
}

// Update modifies a spec using the contents of a message.
func (spec *Spec) Update(message *rpc.Spec) error {
	spec.UpdateTime = spec.CreateTime
	return nil
}