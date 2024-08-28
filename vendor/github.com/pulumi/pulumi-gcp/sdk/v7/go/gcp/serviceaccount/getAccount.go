// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serviceaccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the service account from a project. For more information see
// the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := serviceaccount.LookupAccount(ctx, &serviceaccount.LookupAccountArgs{
//				AccountId: "object-viewer",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Save Key In Kubernetes Secret
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
//	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
//	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myaccount, err := serviceaccount.LookupAccount(ctx, &serviceaccount.LookupAccountArgs{
//				AccountId: "myaccount-id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			mykey, err := serviceaccount.NewKey(ctx, "mykey", &serviceaccount.KeyArgs{
//				ServiceAccountId: pulumi.String(myaccount.Name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = corev1.NewSecret(ctx, "google-application-credentials", &corev1.SecretArgs{
//				Metadata: &metav1.ObjectMetaArgs{
//					Name: pulumi.String("google-application-credentials"),
//				},
//				Data: pulumi.StringMap{
//					"json": std.Base64decodeOutput(ctx, std.Base64decodeOutputArgs{
//						Input: mykey.PrivateKey,
//					}, nil).ApplyT(func(invoke std.Base64decodeResult) (*string, error) {
//						return invoke.Result, nil
//					}).(pulumi.StringPtrOutput),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("gcp:serviceaccount/getAccount:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccount.
type LookupAccountArgs struct {
	// The Google service account ID. This be one of:
	//
	// * The name of the service account within the project (e.g. `my-service`)
	//
	// * The fully-qualified path to a service account resource (e.g.
	// `projects/my-project/serviceAccounts/...`)
	//
	// * The email address of the service account (e.g.
	// `my-service@my-project.iam.gserviceaccount.com`)
	AccountId string `pulumi:"accountId"`
	// The ID of the project that the service account is present in.
	// Defaults to the provider project configuration.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getAccount.
type LookupAccountResult struct {
	AccountId string `pulumi:"accountId"`
	// The display name for the service account.
	DisplayName string `pulumi:"displayName"`
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email string `pulumi:"email"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
	Member string `pulumi:"member"`
	// The fully-qualified name of the service account.
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The unique id of the service account.
	UniqueId string `pulumi:"uniqueId"`
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

// A collection of arguments for invoking getAccount.
type LookupAccountOutputArgs struct {
	// The Google service account ID. This be one of:
	//
	// * The name of the service account within the project (e.g. `my-service`)
	//
	// * The fully-qualified path to a service account resource (e.g.
	// `projects/my-project/serviceAccounts/...`)
	//
	// * The email address of the service account (e.g.
	// `my-service@my-project.iam.gserviceaccount.com`)
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// The ID of the project that the service account is present in.
	// Defaults to the provider project configuration.
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

// A collection of values returned by getAccount.
type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// The display name for the service account.
func (o LookupAccountResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The e-mail address of the service account. This value
// should be referenced from any `organizations.getIAMPolicy` data sources
// that would grant the service account privileges.
func (o LookupAccountResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Email }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
func (o LookupAccountResultOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Member }).(pulumi.StringOutput)
}

// The fully-qualified name of the service account.
func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

// The unique id of the service account.
func (o LookupAccountResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
