// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pulumi/provider.proto

package pulumirpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResourceProviderClient is the client API for ResourceProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceProviderClient interface {
	// Parameterize takes either a string array of command line inputs or a value embedded from sdk generation.
	//
	// Providers can be parameterized with either multiple extension packages (which don't define their own provider
	// resources), or with a replacement package (which does define its own provider resource).
	//
	// Parameterize may be called multiple times for extension packages, but for a replacement package it will only be
	// called once. Extension packages may even be called multiple times for the same package name, but with different
	// versions.
	//
	// Parameterize should work the same for both the `ParametersArgs` input and the `ParametersValue` input. Either way
	// should return the sub-package name and version (which for `ParametersValue` should match the given input).
	//
	// For extension resources their CRUD operations will include the version of which sub-package they correspond to.
	Parameterize(ctx context.Context, in *ParameterizeRequest, opts ...grpc.CallOption) (*ParameterizeResponse, error)
	// GetSchema fetches the schema for this resource provider.
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error)
	// CheckConfig validates the configuration for this resource provider.
	CheckConfig(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// DiffConfig checks the impact a hypothetical change to this provider's configuration will have on the provider.
	DiffConfig(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	// Configure configures the resource provider with "globals" that control its behavior.
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	// Invoke dynamically executes a built-in function in the provider.
	Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error)
	// StreamInvoke dynamically executes a built-in function in the provider, which returns a stream
	// of responses.
	StreamInvoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (ResourceProvider_StreamInvokeClient, error)
	// Call dynamically executes a method in the provider associated with a component resource.
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
	// Check validates that the given property bag is valid for a resource of the given type and returns the inputs
	// that should be passed to successive calls to Diff, Create, or Update for this resource. As a rule, the provider
	// inputs returned by a call to Check should preserve the original representation of the properties as present in
	// the program inputs. Though this rule is not required for correctness, violations thereof can negatively impact
	// the end-user experience, as the provider inputs are using for detecting and rendering diffs.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	// Diff checks what impacts a hypothetical update will have on the resource's properties.
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transactional").
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read the current live state associated with a resource.  Enough state must be include in the inputs to uniquely
	// identify the resource; this is typically just the resource ID, but may also include some properties.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update updates an existing resource with new values.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Construct creates a new instance of the provided component resource and returns its state.
	Construct(ctx context.Context, in *ConstructRequest, opts ...grpc.CallOption) (*ConstructResponse, error)
	// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
	// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either return a
	// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
	// to the host to decide how long to wait after Cancel is called before (e.g.)
	// hard-closing any gRPC connection.
	Cancel(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginInfo, error)
	// Attach sends the engine address to an already running plugin.
	Attach(ctx context.Context, in *PluginAttach, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetMapping fetches the mapping for this resource provider, if any. A provider should return an empty
	// response (not an error) if it doesn't have a mapping for the given key.
	GetMapping(ctx context.Context, in *GetMappingRequest, opts ...grpc.CallOption) (*GetMappingResponse, error)
	// GetMappings is an optional method that returns what mappings (if any) a provider supports. If a provider does not
	// implement this method the engine falls back to the old behaviour of just calling GetMapping without a name.
	// If this method is implemented than the engine will then call GetMapping only with the names returned from this method.
	GetMappings(ctx context.Context, in *GetMappingsRequest, opts ...grpc.CallOption) (*GetMappingsResponse, error)
}

type resourceProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceProviderClient(cc grpc.ClientConnInterface) ResourceProviderClient {
	return &resourceProviderClient{cc}
}

func (c *resourceProviderClient) Parameterize(ctx context.Context, in *ParameterizeRequest, opts ...grpc.CallOption) (*ParameterizeResponse, error) {
	out := new(ParameterizeResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Parameterize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error) {
	out := new(GetSchemaResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) CheckConfig(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/CheckConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) DiffConfig(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/DiffConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Invoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (*InvokeResponse, error) {
	out := new(InvokeResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) StreamInvoke(ctx context.Context, in *InvokeRequest, opts ...grpc.CallOption) (ResourceProvider_StreamInvokeClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResourceProvider_ServiceDesc.Streams[0], "/pulumirpc.ResourceProvider/StreamInvoke", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceProviderStreamInvokeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceProvider_StreamInvokeClient interface {
	Recv() (*InvokeResponse, error)
	grpc.ClientStream
}

type resourceProviderStreamInvokeClient struct {
	grpc.ClientStream
}

func (x *resourceProviderStreamInvokeClient) Recv() (*InvokeResponse, error) {
	m := new(InvokeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *resourceProviderClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Construct(ctx context.Context, in *ConstructRequest, opts ...grpc.CallOption) (*ConstructResponse, error) {
	out := new(ConstructResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Construct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Cancel(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) GetPluginInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Attach(ctx context.Context, in *PluginAttach, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/Attach", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) GetMapping(ctx context.Context, in *GetMappingRequest, opts ...grpc.CallOption) (*GetMappingResponse, error) {
	out := new(GetMappingResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/GetMapping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) GetMappings(ctx context.Context, in *GetMappingsRequest, opts ...grpc.CallOption) (*GetMappingsResponse, error) {
	out := new(GetMappingsResponse)
	err := c.cc.Invoke(ctx, "/pulumirpc.ResourceProvider/GetMappings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceProviderServer is the server API for ResourceProvider service.
// All implementations must embed UnimplementedResourceProviderServer
// for forward compatibility
type ResourceProviderServer interface {
	// Parameterize takes either a string array of command line inputs or a value embedded from sdk generation.
	//
	// Providers can be parameterized with either multiple extension packages (which don't define their own provider
	// resources), or with a replacement package (which does define its own provider resource).
	//
	// Parameterize may be called multiple times for extension packages, but for a replacement package it will only be
	// called once. Extension packages may even be called multiple times for the same package name, but with different
	// versions.
	//
	// Parameterize should work the same for both the `ParametersArgs` input and the `ParametersValue` input. Either way
	// should return the sub-package name and version (which for `ParametersValue` should match the given input).
	//
	// For extension resources their CRUD operations will include the version of which sub-package they correspond to.
	Parameterize(context.Context, *ParameterizeRequest) (*ParameterizeResponse, error)
	// GetSchema fetches the schema for this resource provider.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
	// CheckConfig validates the configuration for this resource provider.
	CheckConfig(context.Context, *CheckRequest) (*CheckResponse, error)
	// DiffConfig checks the impact a hypothetical change to this provider's configuration will have on the provider.
	DiffConfig(context.Context, *DiffRequest) (*DiffResponse, error)
	// Configure configures the resource provider with "globals" that control its behavior.
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	// Invoke dynamically executes a built-in function in the provider.
	Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error)
	// StreamInvoke dynamically executes a built-in function in the provider, which returns a stream
	// of responses.
	StreamInvoke(*InvokeRequest, ResourceProvider_StreamInvokeServer) error
	// Call dynamically executes a method in the provider associated with a component resource.
	Call(context.Context, *CallRequest) (*CallResponse, error)
	// Check validates that the given property bag is valid for a resource of the given type and returns the inputs
	// that should be passed to successive calls to Diff, Create, or Update for this resource. As a rule, the provider
	// inputs returned by a call to Check should preserve the original representation of the properties as present in
	// the program inputs. Though this rule is not required for correctness, violations thereof can negatively impact
	// the end-user experience, as the provider inputs are using for detecting and rendering diffs.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	// Diff checks what impacts a hypothetical update will have on the resource's properties.
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transactional").
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read the current live state associated with a resource.  Enough state must be include in the inputs to uniquely
	// identify the resource; this is typically just the resource ID, but may also include some properties.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update updates an existing resource with new values.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	// Construct creates a new instance of the provided component resource and returns its state.
	Construct(context.Context, *ConstructRequest) (*ConstructResponse, error)
	// Cancel signals the provider to gracefully shut down and abort any ongoing resource operations.
	// Operations aborted in this way will return an error (e.g., `Update` and `Create` will either return a
	// creation error or an initialization error). Since Cancel is advisory and non-blocking, it is up
	// to the host to decide how long to wait after Cancel is called before (e.g.)
	// hard-closing any gRPC connection.
	Cancel(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// GetPluginInfo returns generic information about this plugin, like its version.
	GetPluginInfo(context.Context, *emptypb.Empty) (*PluginInfo, error)
	// Attach sends the engine address to an already running plugin.
	Attach(context.Context, *PluginAttach) (*emptypb.Empty, error)
	// GetMapping fetches the mapping for this resource provider, if any. A provider should return an empty
	// response (not an error) if it doesn't have a mapping for the given key.
	GetMapping(context.Context, *GetMappingRequest) (*GetMappingResponse, error)
	// GetMappings is an optional method that returns what mappings (if any) a provider supports. If a provider does not
	// implement this method the engine falls back to the old behaviour of just calling GetMapping without a name.
	// If this method is implemented than the engine will then call GetMapping only with the names returned from this method.
	GetMappings(context.Context, *GetMappingsRequest) (*GetMappingsResponse, error)
	mustEmbedUnimplementedResourceProviderServer()
}

// UnimplementedResourceProviderServer must be embedded to have forward compatible implementations.
type UnimplementedResourceProviderServer struct {
}

func (UnimplementedResourceProviderServer) Parameterize(context.Context, *ParameterizeRequest) (*ParameterizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parameterize not implemented")
}
func (UnimplementedResourceProviderServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedResourceProviderServer) CheckConfig(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckConfig not implemented")
}
func (UnimplementedResourceProviderServer) DiffConfig(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiffConfig not implemented")
}
func (UnimplementedResourceProviderServer) Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedResourceProviderServer) Invoke(context.Context, *InvokeRequest) (*InvokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (UnimplementedResourceProviderServer) StreamInvoke(*InvokeRequest, ResourceProvider_StreamInvokeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamInvoke not implemented")
}
func (UnimplementedResourceProviderServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedResourceProviderServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedResourceProviderServer) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedResourceProviderServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourceProviderServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedResourceProviderServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResourceProviderServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResourceProviderServer) Construct(context.Context, *ConstructRequest) (*ConstructResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Construct not implemented")
}
func (UnimplementedResourceProviderServer) Cancel(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedResourceProviderServer) GetPluginInfo(context.Context, *emptypb.Empty) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}
func (UnimplementedResourceProviderServer) Attach(context.Context, *PluginAttach) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attach not implemented")
}
func (UnimplementedResourceProviderServer) GetMapping(context.Context, *GetMappingRequest) (*GetMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMapping not implemented")
}
func (UnimplementedResourceProviderServer) GetMappings(context.Context, *GetMappingsRequest) (*GetMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMappings not implemented")
}
func (UnimplementedResourceProviderServer) mustEmbedUnimplementedResourceProviderServer() {}

// UnsafeResourceProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceProviderServer will
// result in compilation errors.
type UnsafeResourceProviderServer interface {
	mustEmbedUnimplementedResourceProviderServer()
}

func RegisterResourceProviderServer(s grpc.ServiceRegistrar, srv ResourceProviderServer) {
	s.RegisterService(&ResourceProvider_ServiceDesc, srv)
}

func _ResourceProvider_Parameterize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParameterizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Parameterize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Parameterize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Parameterize(ctx, req.(*ParameterizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_CheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).CheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/CheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).CheckConfig(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_DiffConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).DiffConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/DiffConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).DiffConfig(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Invoke(ctx, req.(*InvokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_StreamInvoke_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InvokeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceProviderServer).StreamInvoke(m, &resourceProviderStreamInvokeServer{stream})
}

type ResourceProvider_StreamInvokeServer interface {
	Send(*InvokeResponse) error
	grpc.ServerStream
}

type resourceProviderStreamInvokeServer struct {
	grpc.ServerStream
}

func (x *resourceProviderStreamInvokeServer) Send(m *InvokeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ResourceProvider_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Construct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConstructRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Construct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Construct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Construct(ctx, req.(*ConstructRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Cancel(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).GetPluginInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Attach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginAttach)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Attach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/Attach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Attach(ctx, req.(*PluginAttach))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_GetMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).GetMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/GetMapping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).GetMapping(ctx, req.(*GetMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_GetMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).GetMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pulumirpc.ResourceProvider/GetMappings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).GetMappings(ctx, req.(*GetMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceProvider_ServiceDesc is the grpc.ServiceDesc for ResourceProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pulumirpc.ResourceProvider",
	HandlerType: (*ResourceProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parameterize",
			Handler:    _ResourceProvider_Parameterize_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _ResourceProvider_GetSchema_Handler,
		},
		{
			MethodName: "CheckConfig",
			Handler:    _ResourceProvider_CheckConfig_Handler,
		},
		{
			MethodName: "DiffConfig",
			Handler:    _ResourceProvider_DiffConfig_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _ResourceProvider_Configure_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _ResourceProvider_Invoke_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _ResourceProvider_Call_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _ResourceProvider_Check_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _ResourceProvider_Diff_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceProvider_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ResourceProvider_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceProvider_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceProvider_Delete_Handler,
		},
		{
			MethodName: "Construct",
			Handler:    _ResourceProvider_Construct_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _ResourceProvider_Cancel_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _ResourceProvider_GetPluginInfo_Handler,
		},
		{
			MethodName: "Attach",
			Handler:    _ResourceProvider_Attach_Handler,
		},
		{
			MethodName: "GetMapping",
			Handler:    _ResourceProvider_GetMapping_Handler,
		},
		{
			MethodName: "GetMappings",
			Handler:    _ResourceProvider_GetMappings_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamInvoke",
			Handler:       _ResourceProvider_StreamInvoke_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pulumi/provider.proto",
}
