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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: google/cloud/apigeeregistry/applications/v1alpha1/registry_lint.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Lint wraps the result of a linter run for an API.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type Lint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the result.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The lint results for each file examined.
	Files []*LintFile `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Lint) Reset() {
	*x = Lint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lint) ProtoMessage() {}

func (x *Lint) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lint.ProtoReflect.Descriptor instead.
func (*Lint) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{0}
}

func (x *Lint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lint) GetFiles() []*LintFile {
	if x != nil {
		return x.Files
	}
	return nil
}

// LintFile wraps the result of a linter run for a file.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type LintFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path of the file linted.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// The problems found when the file was linted.
	Problems []*LintProblem `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
}

func (x *LintFile) Reset() {
	*x = LintFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintFile) ProtoMessage() {}

func (x *LintFile) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintFile.ProtoReflect.Descriptor instead.
func (*LintFile) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{1}
}

func (x *LintFile) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *LintFile) GetProblems() []*LintProblem {
	if x != nil {
		return x.Problems
	}
	return nil
}

// LintProblem represents a problem found by a linter.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type LintProblem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A message describing the problem.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// An identifier for a related design rule.
	RuleId string `protobuf:"bytes,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	// A link for a related design rule.
	RuleDocUri string `protobuf:"bytes,3,opt,name=rule_doc_uri,json=ruleDocUri,proto3" json:"rule_doc_uri,omitempty"`
	// A suggestion for resolving the problem.
	Suggestion string `protobuf:"bytes,4,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	// The location in the file of the problem.
	Location *LintLocation `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *LintProblem) Reset() {
	*x = LintProblem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintProblem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintProblem) ProtoMessage() {}

func (x *LintProblem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintProblem.ProtoReflect.Descriptor instead.
func (*LintProblem) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{2}
}

func (x *LintProblem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LintProblem) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *LintProblem) GetRuleDocUri() string {
	if x != nil {
		return x.RuleDocUri
	}
	return ""
}

func (x *LintProblem) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

func (x *LintProblem) GetLocation() *LintLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

// LintLocation represents a range of text in a file.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type LintLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The initial position of a problem.
	StartPosition *LintPosition `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"start_position,omitempty"`
	// The end position of a problem.
	EndPosition *LintPosition `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"end_position,omitempty"`
}

func (x *LintLocation) Reset() {
	*x = LintLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintLocation) ProtoMessage() {}

func (x *LintLocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintLocation.ProtoReflect.Descriptor instead.
func (*LintLocation) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{3}
}

func (x *LintLocation) GetStartPosition() *LintPosition {
	if x != nil {
		return x.StartPosition
	}
	return nil
}

func (x *LintLocation) GetEndPosition() *LintPosition {
	if x != nil {
		return x.EndPosition
	}
	return nil
}

// LintPosition represents a point in a file.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type LintPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A line number in a file.
	LineNumber int32 `protobuf:"varint,1,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// A column number in a file.
	ColumnNumber int32 `protobuf:"varint,2,opt,name=column_number,json=columnNumber,proto3" json:"column_number,omitempty"`
}

func (x *LintPosition) Reset() {
	*x = LintPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintPosition) ProtoMessage() {}

func (x *LintPosition) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintPosition.ProtoReflect.Descriptor instead.
func (*LintPosition) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{4}
}

func (x *LintPosition) GetLineNumber() int32 {
	if x != nil {
		return x.LineNumber
	}
	return 0
}

func (x *LintPosition) GetColumnNumber() int32 {
	if x != nil {
		return x.ColumnNumber
	}
	return 0
}

// LintStats summarizes linter results.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type LintStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A sum of the total operations (GET, POST, PUT, DELETE).
	OperationCount int32 `protobuf:"varint,1,opt,name=operation_count,json=operationCount,proto3" json:"operation_count,omitempty"`
	// The number of schemas contained under this resource.
	SchemaCount int32 `protobuf:"varint,2,opt,name=schema_count,json=schemaCount,proto3" json:"schema_count,omitempty"`
	// Problems found in linting.
	ProblemCounts []*LintProblemCount `protobuf:"bytes,3,rep,name=problem_counts,json=problemCounts,proto3" json:"problem_counts,omitempty"`
}

func (x *LintStats) Reset() {
	*x = LintStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintStats) ProtoMessage() {}

func (x *LintStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintStats.ProtoReflect.Descriptor instead.
func (*LintStats) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{5}
}

func (x *LintStats) GetOperationCount() int32 {
	if x != nil {
		return x.OperationCount
	}
	return 0
}

func (x *LintStats) GetSchemaCount() int32 {
	if x != nil {
		return x.SchemaCount
	}
	return 0
}

func (x *LintStats) GetProblemCounts() []*LintProblemCount {
	if x != nil {
		return x.ProblemCounts
	}
	return nil
}

// LintProblemCount represents the number of times a problem was found in linting.
// (-- api-linter: core::0123::resource-annotation=disabled
//     aip.dev/not-precedent: This message is not currently used in an API. --)
type LintProblemCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of occurrences of the problem.
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// An identifier for the related design rule.
	RuleId string `protobuf:"bytes,2,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	// A link for the related design rule.
	RuleDocUri string `protobuf:"bytes,3,opt,name=rule_doc_uri,json=ruleDocUri,proto3" json:"rule_doc_uri,omitempty"`
}

func (x *LintProblemCount) Reset() {
	*x = LintProblemCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LintProblemCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintProblemCount) ProtoMessage() {}

func (x *LintProblemCount) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintProblemCount.ProtoReflect.Descriptor instead.
func (*LintProblemCount) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP(), []int{6}
}

func (x *LintProblemCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *LintProblemCount) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *LintProblemCount) GetRuleDocUri() string {
	if x != nil {
		return x.RuleDocUri
	}
	return ""
}

var File_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto protoreflect.FileDescriptor

var file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x6d, 0x0a, 0x04, 0x4c, 0x69,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x08, 0x4c, 0x69,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x5a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x22,
	0xdf, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x44, 0x6f,
	0x63, 0x55, 0x72, 0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x66, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x0c, 0x65, 0x6e,
	0x64, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54,
	0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0xc3, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x6a,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x63, 0x0a, 0x10, 0x4c, 0x69,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x55, 0x72, 0x69, 0x42,
	0x70, 0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x4c, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x3b, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescOnce sync.Once
	file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescData = file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDesc
)

func file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescGZIP() []byte {
	file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescOnce.Do(func() {
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescData)
	})
	return file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDescData
}

var file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_goTypes = []interface{}{
	(*Lint)(nil),             // 0: google.cloud.apigeeregistry.applications.v1alpha1.Lint
	(*LintFile)(nil),         // 1: google.cloud.apigeeregistry.applications.v1alpha1.LintFile
	(*LintProblem)(nil),      // 2: google.cloud.apigeeregistry.applications.v1alpha1.LintProblem
	(*LintLocation)(nil),     // 3: google.cloud.apigeeregistry.applications.v1alpha1.LintLocation
	(*LintPosition)(nil),     // 4: google.cloud.apigeeregistry.applications.v1alpha1.LintPosition
	(*LintStats)(nil),        // 5: google.cloud.apigeeregistry.applications.v1alpha1.LintStats
	(*LintProblemCount)(nil), // 6: google.cloud.apigeeregistry.applications.v1alpha1.LintProblemCount
}
var file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_depIdxs = []int32{
	1, // 0: google.cloud.apigeeregistry.applications.v1alpha1.Lint.files:type_name -> google.cloud.apigeeregistry.applications.v1alpha1.LintFile
	2, // 1: google.cloud.apigeeregistry.applications.v1alpha1.LintFile.problems:type_name -> google.cloud.apigeeregistry.applications.v1alpha1.LintProblem
	3, // 2: google.cloud.apigeeregistry.applications.v1alpha1.LintProblem.location:type_name -> google.cloud.apigeeregistry.applications.v1alpha1.LintLocation
	4, // 3: google.cloud.apigeeregistry.applications.v1alpha1.LintLocation.start_position:type_name -> google.cloud.apigeeregistry.applications.v1alpha1.LintPosition
	4, // 4: google.cloud.apigeeregistry.applications.v1alpha1.LintLocation.end_position:type_name -> google.cloud.apigeeregistry.applications.v1alpha1.LintPosition
	6, // 5: google.cloud.apigeeregistry.applications.v1alpha1.LintStats.problem_counts:type_name -> google.cloud.apigeeregistry.applications.v1alpha1.LintProblemCount
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_init() }
func file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_init() {
	if File_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintFile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintProblem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintPosition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LintProblemCount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_goTypes,
		DependencyIndexes: file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_depIdxs,
		MessageInfos:      file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_msgTypes,
	}.Build()
	File_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto = out.File
	file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_rawDesc = nil
	file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_goTypes = nil
	file_google_cloud_apigeeregistry_applications_v1alpha1_registry_lint_proto_depIdxs = nil
}