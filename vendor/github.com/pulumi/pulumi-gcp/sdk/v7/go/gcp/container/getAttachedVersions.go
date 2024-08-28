// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides access to available platform versions in a location for a given project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := container.GetAttachedVersions(ctx, &container.GetAttachedVersionsArgs{
//				Location: "us-west1",
//				Project:  "my-project",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("firstAvailableVersion", versions.ValidVersions[0])
//			return nil
//		})
//	}
//
// ```
func GetAttachedVersions(ctx *pulumi.Context, args *GetAttachedVersionsArgs, opts ...pulumi.InvokeOption) (*GetAttachedVersionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAttachedVersionsResult
	err := ctx.Invoke("gcp:container/getAttachedVersions:getAttachedVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAttachedVersions.
type GetAttachedVersionsArgs struct {
	// The location to list versions for.
	Location string `pulumi:"location"`
	// ID of the project to list available platform versions for. Should match the project the cluster will be deployed to.
	// Defaults to the project that the provider is authenticated with.
	Project string `pulumi:"project"`
}

// A collection of values returned by getAttachedVersions.
type GetAttachedVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	Project  string `pulumi:"project"`
	// A list of versions available for use with this project and location.
	ValidVersions []string `pulumi:"validVersions"`
}

func GetAttachedVersionsOutput(ctx *pulumi.Context, args GetAttachedVersionsOutputArgs, opts ...pulumi.InvokeOption) GetAttachedVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAttachedVersionsResult, error) {
			args := v.(GetAttachedVersionsArgs)
			r, err := GetAttachedVersions(ctx, &args, opts...)
			var s GetAttachedVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAttachedVersionsResultOutput)
}

// A collection of arguments for invoking getAttachedVersions.
type GetAttachedVersionsOutputArgs struct {
	// The location to list versions for.
	Location pulumi.StringInput `pulumi:"location"`
	// ID of the project to list available platform versions for. Should match the project the cluster will be deployed to.
	// Defaults to the project that the provider is authenticated with.
	Project pulumi.StringInput `pulumi:"project"`
}

func (GetAttachedVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAttachedVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getAttachedVersions.
type GetAttachedVersionsResultOutput struct{ *pulumi.OutputState }

func (GetAttachedVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAttachedVersionsResult)(nil)).Elem()
}

func (o GetAttachedVersionsResultOutput) ToGetAttachedVersionsResultOutput() GetAttachedVersionsResultOutput {
	return o
}

func (o GetAttachedVersionsResultOutput) ToGetAttachedVersionsResultOutputWithContext(ctx context.Context) GetAttachedVersionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAttachedVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachedVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAttachedVersionsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachedVersionsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetAttachedVersionsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetAttachedVersionsResult) string { return v.Project }).(pulumi.StringOutput)
}

// A list of versions available for use with this project and location.
func (o GetAttachedVersionsResultOutput) ValidVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAttachedVersionsResult) []string { return v.ValidVersions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAttachedVersionsResultOutput{})
}
