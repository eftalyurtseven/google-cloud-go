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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newDisksClientHook clientHook

// DisksCallOptions contains the retry settings for each method of DisksClient.
type DisksCallOptions struct {
	AddResourcePolicies    []gax.CallOption
	AggregatedList         []gax.CallOption
	CreateSnapshot         []gax.CallOption
	Delete                 []gax.CallOption
	Get                    []gax.CallOption
	GetIamPolicy           []gax.CallOption
	Insert                 []gax.CallOption
	List                   []gax.CallOption
	RemoveResourcePolicies []gax.CallOption
	Resize                 []gax.CallOption
	SetIamPolicy           []gax.CallOption
	SetLabels              []gax.CallOption
	TestIamPermissions     []gax.CallOption
}

// internalDisksClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalDisksClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddResourcePolicies(context.Context, *computepb.AddResourcePoliciesDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListDisksRequest, ...gax.CallOption) (*computepb.DiskAggregatedList, error)
	CreateSnapshot(context.Context, *computepb.CreateSnapshotDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Delete(context.Context, *computepb.DeleteDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetDiskRequest, ...gax.CallOption) (*computepb.Disk, error)
	GetIamPolicy(context.Context, *computepb.GetIamPolicyDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	Insert(context.Context, *computepb.InsertDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListDisksRequest, ...gax.CallOption) (*computepb.DiskList, error)
	RemoveResourcePolicies(context.Context, *computepb.RemoveResourcePoliciesDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	Resize(context.Context, *computepb.ResizeDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetIamPolicy(context.Context, *computepb.SetIamPolicyDiskRequest, ...gax.CallOption) (*computepb.Policy, error)
	SetLabels(context.Context, *computepb.SetLabelsDiskRequest, ...gax.CallOption) (*computepb.Operation, error)
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsDiskRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// DisksClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Disks API.
type DisksClient struct {
	// The internal transport-dependent client.
	internalClient internalDisksClient

	// The call options for this service.
	CallOptions *DisksCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DisksClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DisksClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *DisksClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddResourcePolicies adds existing resource policies to a disk. You can only add one policy which will be applied to this disk for scheduling snapshot creation.
func (c *DisksClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddResourcePolicies(ctx, req, opts...)
}

// AggregatedList retrieves an aggregated list of persistent disks.
func (c *DisksClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListDisksRequest, opts ...gax.CallOption) (*computepb.DiskAggregatedList, error) {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// CreateSnapshot creates a snapshot of a specified persistent disk.
func (c *DisksClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.CreateSnapshot(ctx, req, opts...)
}

// Delete deletes the specified persistent disk. Deleting a disk removes its data permanently and is irreversible. However, deleting a disk does not delete any snapshots previously made from the disk. You must separately delete snapshots.
func (c *DisksClient) Delete(ctx context.Context, req *computepb.DeleteDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns a specified persistent disk. Gets a list of available persistent disks by making a list() request.
func (c *DisksClient) Get(ctx context.Context, req *computepb.GetDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *DisksClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// Insert creates a persistent disk in the specified project using the data in the request. You can create a disk from a source (sourceImage, sourceSnapshot, or sourceDisk) or create an empty 500 GB data disk by omitting all properties. You can also create a disk that is larger than the default size by specifying the sizeGb property.
func (c *DisksClient) Insert(ctx context.Context, req *computepb.InsertDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of persistent disks contained within the specified zone.
func (c *DisksClient) List(ctx context.Context, req *computepb.ListDisksRequest, opts ...gax.CallOption) (*computepb.DiskList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// RemoveResourcePolicies removes resource policies from a disk.
func (c *DisksClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveResourcePolicies(ctx, req, opts...)
}

// Resize resizes the specified persistent disk. You can only increase the size of the disk.
func (c *DisksClient) Resize(ctx context.Context, req *computepb.ResizeDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Resize(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *DisksClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// SetLabels sets the labels on a disk. To learn more about labels, read the Labeling Resources documentation.
func (c *DisksClient) SetLabels(ctx context.Context, req *computepb.SetLabelsDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *DisksClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type disksRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDisksRESTClient creates a new disks rest client.
//
// The Disks API.
func NewDisksRESTClient(ctx context.Context, opts ...option.ClientOption) (*DisksClient, error) {
	clientOpts := append(defaultDisksRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &disksRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &DisksClient{internalClient: c, CallOptions: &DisksCallOptions{}}, nil
}

func defaultDisksRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *disksRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *disksRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *disksRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddResourcePolicies adds existing resource policies to a disk. You can only add one policy which will be applied to this disk for scheduling snapshot creation.
func (c *disksRESTClient) AddResourcePolicies(ctx context.Context, req *computepb.AddResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetDisksAddResourcePoliciesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/addResourcePolicies", req.GetProject(), req.GetZone(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// AggregatedList retrieves an aggregated list of persistent disks.
func (c *disksRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListDisksRequest, opts ...gax.CallOption) (*computepb.DiskAggregatedList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/disks", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.IncludeAllScopes != nil {
		params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.DiskAggregatedList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// CreateSnapshot creates a snapshot of a specified persistent disk.
func (c *disksRESTClient) CreateSnapshot(ctx context.Context, req *computepb.CreateSnapshotDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetSnapshotResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/createSnapshot", req.GetProject(), req.GetZone(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.GuestFlush != nil {
		params.Add("guestFlush", fmt.Sprintf("%v", req.GetGuestFlush()))
	}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified persistent disk. Deleting a disk removes its data permanently and is irreversible. However, deleting a disk does not delete any snapshots previously made from the disk. You must separately delete snapshots.
func (c *disksRESTClient) Delete(ctx context.Context, req *computepb.DeleteDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v", req.GetProject(), req.GetZone(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get returns a specified persistent disk. Gets a list of available persistent disks by making a list() request.
func (c *disksRESTClient) Get(ctx context.Context, req *computepb.GetDiskRequest, opts ...gax.CallOption) (*computepb.Disk, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v", req.GetProject(), req.GetZone(), req.GetDisk())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Disk{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// GetIamPolicy gets the access control policy for a resource. May be empty if no such policy or resource exists.
func (c *disksRESTClient) GetIamPolicy(ctx context.Context, req *computepb.GetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/getIamPolicy", req.GetProject(), req.GetZone(), req.GetResource())

	params := url.Values{}
	if req != nil && req.OptionsRequestedPolicyVersion != nil {
		params.Add("optionsRequestedPolicyVersion", fmt.Sprintf("%v", req.GetOptionsRequestedPolicyVersion()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Policy{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a persistent disk in the specified project using the data in the request. You can create a disk from a source (sourceImage, sourceSnapshot, or sourceDisk) or create an empty 500 GB data disk by omitting all properties. You can also create a disk that is larger than the default size by specifying the sizeGb property.
func (c *disksRESTClient) Insert(ctx context.Context, req *computepb.InsertDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetDiskResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks", req.GetProject(), req.GetZone())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req != nil && req.SourceImage != nil {
		params.Add("sourceImage", fmt.Sprintf("%v", req.GetSourceImage()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// List retrieves a list of persistent disks contained within the specified zone.
func (c *disksRESTClient) List(ctx context.Context, req *computepb.ListDisksRequest, opts ...gax.CallOption) (*computepb.DiskList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks", req.GetProject(), req.GetZone())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.DiskList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// RemoveResourcePolicies removes resource policies from a disk.
func (c *disksRESTClient) RemoveResourcePolicies(ctx context.Context, req *computepb.RemoveResourcePoliciesDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetDisksRemoveResourcePoliciesRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/removeResourcePolicies", req.GetProject(), req.GetZone(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Resize resizes the specified persistent disk. You can only increase the size of the disk.
func (c *disksRESTClient) Resize(ctx context.Context, req *computepb.ResizeDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetDisksResizeRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/resize", req.GetProject(), req.GetZone(), req.GetDisk())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// SetIamPolicy sets the access control policy on the specified resource. Replaces any existing policy.
func (c *disksRESTClient) SetIamPolicy(ctx context.Context, req *computepb.SetIamPolicyDiskRequest, opts ...gax.CallOption) (*computepb.Policy, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetZoneSetPolicyRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/setIamPolicy", req.GetProject(), req.GetZone(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Policy{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// SetLabels sets the labels on a disk. To learn more about labels, read the Labeling Resources documentation.
func (c *disksRESTClient) SetLabels(ctx context.Context, req *computepb.SetLabelsDiskRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetZoneSetLabelsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/setLabels", req.GetProject(), req.GetZone(), req.GetResource())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *disksRESTClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsDiskRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTestPermissionsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/disks/%v/testIamPermissions", req.GetProject(), req.GetZone(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.TestPermissionsResponse{}

	return rsp, unm.Unmarshal(buf, rsp)
}