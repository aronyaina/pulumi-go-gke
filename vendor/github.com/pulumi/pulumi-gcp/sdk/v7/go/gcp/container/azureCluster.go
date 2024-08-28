// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Anthos cluster running on Azure.
//
// For more information, see:
// * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
// ## Example Usage
//
// ### Basic_azure_cluster
// A basic example of a containerazure azure cluster
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			versions, err := container.GetAzureVersions(ctx, &container.GetAzureVersionsArgs{
//				Project:  pulumi.StringRef("my-project-name"),
//				Location: pulumi.StringRef("us-west1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			basic, err := container.NewAzureClient(ctx, "basic", &container.AzureClientArgs{
//				ApplicationId: pulumi.String("12345678-1234-1234-1234-123456789111"),
//				Location:      pulumi.String("us-west1"),
//				Name:          pulumi.String("client-name"),
//				TenantId:      pulumi.String("12345678-1234-1234-1234-123456789111"),
//				Project:       pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = container.NewAzureCluster(ctx, "primary", &container.AzureClusterArgs{
//				Authorization: &container.AzureClusterAuthorizationArgs{
//					AdminUsers: container.AzureClusterAuthorizationAdminUserArray{
//						&container.AzureClusterAuthorizationAdminUserArgs{
//							Username: pulumi.String("mmv2@google.com"),
//						},
//					},
//					AdminGroups: container.AzureClusterAuthorizationAdminGroupArray{
//						&container.AzureClusterAuthorizationAdminGroupArgs{
//							Group: pulumi.String("group@domain.com"),
//						},
//					},
//				},
//				AzureRegion: pulumi.String("westus2"),
//				Client: basic.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("projects/my-project-number/locations/us-west1/azureClients/%v", name), nil
//				}).(pulumi.StringOutput),
//				ControlPlane: &container.AzureClusterControlPlaneArgs{
//					SshConfig: &container.AzureClusterControlPlaneSshConfigArgs{
//						AuthorizedKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8yaayO6lnb2v+SedxUMa2c8vtIEzCzBjM3EJJsv8Vm9zUDWR7dXWKoNGARUb2mNGXASvI6mFIDXTIlkQ0poDEPpMaXR0g2cb5xT8jAAJq7fqXL3+0rcJhY/uigQ+MrT6s+ub0BFVbsmGHNrMQttXX9gtmwkeAEvj3mra9e5pkNf90qlKnZz6U0SVArxVsLx07vHPHDIYrl0OPG4zUREF52igbBPiNrHJFDQJT/4YlDMJmo/QT/A1D6n9ocemvZSzhRx15/Arjowhr+VVKSbaxzPtEfY0oIg2SrqJnnr/l3Du5qIefwh5VmCZe4xopPUaDDoOIEFriZ88sB+3zz8ib8sk8zJJQCgeP78tQvXCgS+4e5W3TUg9mxjB6KjXTyHIVhDZqhqde0OI3Fy1UuVzRUwnBaLjBnAwP5EoFQGRmDYk/rEYe7HTmovLeEBUDQocBQKT4Ripm/xJkkWY7B07K/tfo56dGUCkvyIVXKBInCh+dLK7gZapnd4UWkY0xBYcwo1geMLRq58iFTLA2j/JmpmHXp7m0l7jJii7d44uD3tTIFYThn7NlOnvhLim/YcBK07GMGIN7XwrrKZKmxXaspw6KBWVhzuw1UPxctxshYEaMLfFg/bwOw8HvMPr9VtrElpSB7oiOh91PDIPdPBgHCi7N2QgQ5l/ZDBHieSpNrQ== thomasrodgers"),
//					},
//					SubnetId: pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet/subnets/default"),
//					Version:  pulumi.String(versions.ValidVersions[0]),
//				},
//				Fleet: &container.AzureClusterFleetArgs{
//					Project: pulumi.String("my-project-number"),
//				},
//				Location: pulumi.String("us-west1"),
//				Name:     pulumi.String("name"),
//				Networking: &container.AzureClusterNetworkingArgs{
//					PodAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("10.200.0.0/16"),
//					},
//					ServiceAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("10.32.0.0/24"),
//					},
//					VirtualNetworkId: pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet"),
//				},
//				ResourceGroupId: pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-cluster"),
//				Project:         pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Beta_basic_enum_azure_cluster
// A basic example of a containerazure azure cluster with lowercase enums (beta)
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			versions, err := container.GetAzureVersions(ctx, &container.GetAzureVersionsArgs{
//				Project:  pulumi.StringRef("my-project-name"),
//				Location: pulumi.StringRef("us-west1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			basic, err := container.NewAzureClient(ctx, "basic", &container.AzureClientArgs{
//				ApplicationId: pulumi.String("12345678-1234-1234-1234-123456789111"),
//				Location:      pulumi.String("us-west1"),
//				Name:          pulumi.String("client-name"),
//				TenantId:      pulumi.String("12345678-1234-1234-1234-123456789111"),
//				Project:       pulumi.String("my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = container.NewAzureCluster(ctx, "primary", &container.AzureClusterArgs{
//				Authorization: &container.AzureClusterAuthorizationArgs{
//					AdminUsers: container.AzureClusterAuthorizationAdminUserArray{
//						&container.AzureClusterAuthorizationAdminUserArgs{
//							Username: pulumi.String("mmv2@google.com"),
//						},
//					},
//				},
//				AzureRegion: pulumi.String("westus2"),
//				Client: basic.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("projects/my-project-number/locations/us-west1/azureClients/%v", name), nil
//				}).(pulumi.StringOutput),
//				ControlPlane: &container.AzureClusterControlPlaneArgs{
//					SshConfig: &container.AzureClusterControlPlaneSshConfigArgs{
//						AuthorizedKey: pulumi.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8yaayO6lnb2v+SedxUMa2c8vtIEzCzBjM3EJJsv8Vm9zUDWR7dXWKoNGARUb2mNGXASvI6mFIDXTIlkQ0poDEPpMaXR0g2cb5xT8jAAJq7fqXL3+0rcJhY/uigQ+MrT6s+ub0BFVbsmGHNrMQttXX9gtmwkeAEvj3mra9e5pkNf90qlKnZz6U0SVArxVsLx07vHPHDIYrl0OPG4zUREF52igbBPiNrHJFDQJT/4YlDMJmo/QT/A1D6n9ocemvZSzhRx15/Arjowhr+VVKSbaxzPtEfY0oIg2SrqJnnr/l3Du5qIefwh5VmCZe4xopPUaDDoOIEFriZ88sB+3zz8ib8sk8zJJQCgeP78tQvXCgS+4e5W3TUg9mxjB6KjXTyHIVhDZqhqde0OI3Fy1UuVzRUwnBaLjBnAwP5EoFQGRmDYk/rEYe7HTmovLeEBUDQocBQKT4Ripm/xJkkWY7B07K/tfo56dGUCkvyIVXKBInCh+dLK7gZapnd4UWkY0xBYcwo1geMLRq58iFTLA2j/JmpmHXp7m0l7jJii7d44uD3tTIFYThn7NlOnvhLim/YcBK07GMGIN7XwrrKZKmxXaspw6KBWVhzuw1UPxctxshYEaMLfFg/bwOw8HvMPr9VtrElpSB7oiOh91PDIPdPBgHCi7N2QgQ5l/ZDBHieSpNrQ== thomasrodgers"),
//					},
//					SubnetId: pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet/subnets/default"),
//					Version:  pulumi.String(versions.ValidVersions[0]),
//				},
//				Fleet: &container.AzureClusterFleetArgs{
//					Project: pulumi.String("my-project-number"),
//				},
//				Location: pulumi.String("us-west1"),
//				Name:     pulumi.String("name"),
//				Networking: &container.AzureClusterNetworkingArgs{
//					PodAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("10.200.0.0/16"),
//					},
//					ServiceAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("10.32.0.0/24"),
//					},
//					VirtualNetworkId: pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-byo/providers/Microsoft.Network/virtualNetworks/my--dev-vnet"),
//				},
//				ResourceGroupId: pulumi.String("/subscriptions/12345678-1234-1234-1234-123456789111/resourceGroups/my--dev-cluster"),
//				Project:         pulumi.String("my-project-name"),
//				LoggingConfig: &container.AzureClusterLoggingConfigArgs{
//					ComponentConfig: &container.AzureClusterLoggingConfigComponentConfigArgs{
//						EnableComponents: pulumi.StringArray{
//							pulumi.String("system_components"),
//							pulumi.String("workloads"),
//						},
//					},
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
//
// ## Import
//
// Cluster can be imported using any of these accepted formats:
//
// * `projects/{{project}}/locations/{{location}}/azureClusters/{{name}}`
//
// * `{{project}}/{{location}}/{{name}}`
//
// * `{{location}}/{{name}}`
//
// When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example:
//
// ```sh
// $ pulumi import gcp:container/azureCluster:AzureCluster default projects/{{project}}/locations/{{location}}/azureClusters/{{name}}
// ```
//
// ```sh
// $ pulumi import gcp:container/azureCluster:AzureCluster default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
// $ pulumi import gcp:container/azureCluster:AzureCluster default {{location}}/{{name}}
// ```
type AzureCluster struct {
	pulumi.CustomResourceState

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization AzureClusterAuthorizationOutput `pulumi:"authorization"`
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion pulumi.StringOutput `pulumi:"azureRegion"`
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication AzureClusterAzureServicesAuthenticationPtrOutput `pulumi:"azureServicesAuthentication"`
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client pulumi.StringPtrOutput `pulumi:"client"`
	// Configuration related to the cluster control plane.
	ControlPlane AzureClusterControlPlaneOutput `pulumi:"controlPlane"`
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveAnnotations pulumi.MapOutput       `pulumi:"effectiveAnnotations"`
	// Output only. The endpoint of the cluster's API server.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Fleet configuration.
	Fleet AzureClusterFleetOutput `pulumi:"fleet"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Logging configuration.
	LoggingConfig AzureClusterLoggingConfigOutput `pulumi:"loggingConfig"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Cluster-wide networking configuration.
	Networking AzureClusterNetworkingOutput `pulumi:"networking"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	State pulumi.StringOutput `pulumi:"state"`
	// Output only. A globally unique identifier for the cluster.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time at which this cluster was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs AzureClusterWorkloadIdentityConfigArrayOutput `pulumi:"workloadIdentityConfigs"`
}

// NewAzureCluster registers a new resource with the given unique name, arguments, and options.
func NewAzureCluster(ctx *pulumi.Context,
	name string, args *AzureClusterArgs, opts ...pulumi.ResourceOption) (*AzureCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorization == nil {
		return nil, errors.New("invalid value for required argument 'Authorization'")
	}
	if args.AzureRegion == nil {
		return nil, errors.New("invalid value for required argument 'AzureRegion'")
	}
	if args.ControlPlane == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlane'")
	}
	if args.Fleet == nil {
		return nil, errors.New("invalid value for required argument 'Fleet'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Networking == nil {
		return nil, errors.New("invalid value for required argument 'Networking'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureCluster
	err := ctx.RegisterResource("gcp:container/azureCluster:AzureCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureCluster gets an existing AzureCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureClusterState, opts ...pulumi.ResourceOption) (*AzureCluster, error) {
	var resource AzureCluster
	err := ctx.ReadResource("gcp:container/azureCluster:AzureCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureCluster resources.
type azureClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization *AzureClusterAuthorization `pulumi:"authorization"`
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion *string `pulumi:"azureRegion"`
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication *AzureClusterAzureServicesAuthentication `pulumi:"azureServicesAuthentication"`
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client *string `pulumi:"client"`
	// Configuration related to the cluster control plane.
	ControlPlane *AzureClusterControlPlane `pulumi:"controlPlane"`
	// Output only. The time at which this cluster was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description          *string                `pulumi:"description"`
	EffectiveAnnotations map[string]interface{} `pulumi:"effectiveAnnotations"`
	// Output only. The endpoint of the cluster's API server.
	Endpoint *string `pulumi:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Fleet configuration.
	Fleet *AzureClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Logging configuration.
	LoggingConfig *AzureClusterLoggingConfig `pulumi:"loggingConfig"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// Cluster-wide networking configuration.
	Networking *AzureClusterNetworking `pulumi:"networking"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling *bool `pulumi:"reconciling"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	State *string `pulumi:"state"`
	// Output only. A globally unique identifier for the cluster.
	Uid *string `pulumi:"uid"`
	// Output only. The time at which this cluster was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs []AzureClusterWorkloadIdentityConfig `pulumi:"workloadIdentityConfigs"`
}

type AzureClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	Authorization AzureClusterAuthorizationPtrInput
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion pulumi.StringPtrInput
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication AzureClusterAzureServicesAuthenticationPtrInput
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client pulumi.StringPtrInput
	// Configuration related to the cluster control plane.
	ControlPlane AzureClusterControlPlanePtrInput
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringPtrInput
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description          pulumi.StringPtrInput
	EffectiveAnnotations pulumi.MapInput
	// Output only. The endpoint of the cluster's API server.
	Endpoint pulumi.StringPtrInput
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Fleet configuration.
	Fleet AzureClusterFleetPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Logging configuration.
	LoggingConfig AzureClusterLoggingConfigPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// Cluster-wide networking configuration.
	Networking AzureClusterNetworkingPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolPtrInput
	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
	ResourceGroupId pulumi.StringPtrInput
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	State pulumi.StringPtrInput
	// Output only. A globally unique identifier for the cluster.
	Uid pulumi.StringPtrInput
	// Output only. The time at which this cluster was last updated.
	UpdateTime pulumi.StringPtrInput
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs AzureClusterWorkloadIdentityConfigArrayInput
}

func (AzureClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureClusterState)(nil)).Elem()
}

type azureClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization AzureClusterAuthorization `pulumi:"authorization"`
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion string `pulumi:"azureRegion"`
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication *AzureClusterAzureServicesAuthentication `pulumi:"azureServicesAuthentication"`
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client *string `pulumi:"client"`
	// Configuration related to the cluster control plane.
	ControlPlane AzureClusterControlPlane `pulumi:"controlPlane"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// Fleet configuration.
	Fleet AzureClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Logging configuration.
	LoggingConfig *AzureClusterLoggingConfig `pulumi:"loggingConfig"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// Cluster-wide networking configuration.
	Networking AzureClusterNetworking `pulumi:"networking"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
	ResourceGroupId string `pulumi:"resourceGroupId"`
}

// The set of arguments for constructing a AzureCluster resource.
type AzureClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effectiveAnnotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	Authorization AzureClusterAuthorizationInput
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion pulumi.StringInput
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication AzureClusterAzureServicesAuthenticationPtrInput
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client pulumi.StringPtrInput
	// Configuration related to the cluster control plane.
	ControlPlane AzureClusterControlPlaneInput
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// Fleet configuration.
	Fleet AzureClusterFleetInput
	// The location for the resource
	Location pulumi.StringInput
	// Logging configuration.
	LoggingConfig AzureClusterLoggingConfigPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// Cluster-wide networking configuration.
	Networking AzureClusterNetworkingInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
	ResourceGroupId pulumi.StringInput
}

func (AzureClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureClusterArgs)(nil)).Elem()
}

type AzureClusterInput interface {
	pulumi.Input

	ToAzureClusterOutput() AzureClusterOutput
	ToAzureClusterOutputWithContext(ctx context.Context) AzureClusterOutput
}

func (*AzureCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCluster)(nil)).Elem()
}

func (i *AzureCluster) ToAzureClusterOutput() AzureClusterOutput {
	return i.ToAzureClusterOutputWithContext(context.Background())
}

func (i *AzureCluster) ToAzureClusterOutputWithContext(ctx context.Context) AzureClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureClusterOutput)
}

// AzureClusterArrayInput is an input type that accepts AzureClusterArray and AzureClusterArrayOutput values.
// You can construct a concrete instance of `AzureClusterArrayInput` via:
//
//	AzureClusterArray{ AzureClusterArgs{...} }
type AzureClusterArrayInput interface {
	pulumi.Input

	ToAzureClusterArrayOutput() AzureClusterArrayOutput
	ToAzureClusterArrayOutputWithContext(context.Context) AzureClusterArrayOutput
}

type AzureClusterArray []AzureClusterInput

func (AzureClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureCluster)(nil)).Elem()
}

func (i AzureClusterArray) ToAzureClusterArrayOutput() AzureClusterArrayOutput {
	return i.ToAzureClusterArrayOutputWithContext(context.Background())
}

func (i AzureClusterArray) ToAzureClusterArrayOutputWithContext(ctx context.Context) AzureClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureClusterArrayOutput)
}

// AzureClusterMapInput is an input type that accepts AzureClusterMap and AzureClusterMapOutput values.
// You can construct a concrete instance of `AzureClusterMapInput` via:
//
//	AzureClusterMap{ "key": AzureClusterArgs{...} }
type AzureClusterMapInput interface {
	pulumi.Input

	ToAzureClusterMapOutput() AzureClusterMapOutput
	ToAzureClusterMapOutputWithContext(context.Context) AzureClusterMapOutput
}

type AzureClusterMap map[string]AzureClusterInput

func (AzureClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureCluster)(nil)).Elem()
}

func (i AzureClusterMap) ToAzureClusterMapOutput() AzureClusterMapOutput {
	return i.ToAzureClusterMapOutputWithContext(context.Background())
}

func (i AzureClusterMap) ToAzureClusterMapOutputWithContext(ctx context.Context) AzureClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureClusterMapOutput)
}

type AzureClusterOutput struct{ *pulumi.OutputState }

func (AzureClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureCluster)(nil)).Elem()
}

func (o AzureClusterOutput) ToAzureClusterOutput() AzureClusterOutput {
	return o
}

func (o AzureClusterOutput) ToAzureClusterOutputWithContext(ctx context.Context) AzureClusterOutput {
	return o
}

// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
// `effectiveAnnotations` for all of the annotations present on the resource.
func (o AzureClusterOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Configuration related to the cluster RBAC settings.
func (o AzureClusterOutput) Authorization() AzureClusterAuthorizationOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterAuthorizationOutput { return v.Authorization }).(AzureClusterAuthorizationOutput)
}

// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
func (o AzureClusterOutput) AzureRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.AzureRegion }).(pulumi.StringOutput)
}

// Azure authentication configuration for management of Azure resources
func (o AzureClusterOutput) AzureServicesAuthentication() AzureClusterAzureServicesAuthenticationPtrOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterAzureServicesAuthenticationPtrOutput {
		return v.AzureServicesAuthentication
	}).(AzureClusterAzureServicesAuthenticationPtrOutput)
}

// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
// `AzureCluster`. `AzureClient` names are formatted as
// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
func (o AzureClusterOutput) Client() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringPtrOutput { return v.Client }).(pulumi.StringPtrOutput)
}

// Configuration related to the cluster control plane.
func (o AzureClusterOutput) ControlPlane() AzureClusterControlPlaneOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterControlPlaneOutput { return v.ControlPlane }).(AzureClusterControlPlaneOutput)
}

// Output only. The time at which this cluster was created.
func (o AzureClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
func (o AzureClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AzureClusterOutput) EffectiveAnnotations() pulumi.MapOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.MapOutput { return v.EffectiveAnnotations }).(pulumi.MapOutput)
}

// Output only. The endpoint of the cluster's API server.
func (o AzureClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o AzureClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Fleet configuration.
func (o AzureClusterOutput) Fleet() AzureClusterFleetOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterFleetOutput { return v.Fleet }).(AzureClusterFleetOutput)
}

// The location for the resource
func (o AzureClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Logging configuration.
func (o AzureClusterOutput) LoggingConfig() AzureClusterLoggingConfigOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterLoggingConfigOutput { return v.LoggingConfig }).(AzureClusterLoggingConfigOutput)
}

// The name of this resource.
func (o AzureClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Cluster-wide networking configuration.
func (o AzureClusterOutput) Networking() AzureClusterNetworkingOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterNetworkingOutput { return v.Networking }).(AzureClusterNetworkingOutput)
}

// The project for the resource
func (o AzureClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. If set, there are currently changes in flight to the cluster.
func (o AzureClusterOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
func (o AzureClusterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
func (o AzureClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Output only. A globally unique identifier for the cluster.
func (o AzureClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time at which this cluster was last updated.
func (o AzureClusterOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureCluster) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Output only. Workload Identity settings.
func (o AzureClusterOutput) WorkloadIdentityConfigs() AzureClusterWorkloadIdentityConfigArrayOutput {
	return o.ApplyT(func(v *AzureCluster) AzureClusterWorkloadIdentityConfigArrayOutput { return v.WorkloadIdentityConfigs }).(AzureClusterWorkloadIdentityConfigArrayOutput)
}

type AzureClusterArrayOutput struct{ *pulumi.OutputState }

func (AzureClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureCluster)(nil)).Elem()
}

func (o AzureClusterArrayOutput) ToAzureClusterArrayOutput() AzureClusterArrayOutput {
	return o
}

func (o AzureClusterArrayOutput) ToAzureClusterArrayOutputWithContext(ctx context.Context) AzureClusterArrayOutput {
	return o
}

func (o AzureClusterArrayOutput) Index(i pulumi.IntInput) AzureClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureCluster {
		return vs[0].([]*AzureCluster)[vs[1].(int)]
	}).(AzureClusterOutput)
}

type AzureClusterMapOutput struct{ *pulumi.OutputState }

func (AzureClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureCluster)(nil)).Elem()
}

func (o AzureClusterMapOutput) ToAzureClusterMapOutput() AzureClusterMapOutput {
	return o
}

func (o AzureClusterMapOutput) ToAzureClusterMapOutputWithContext(ctx context.Context) AzureClusterMapOutput {
	return o
}

func (o AzureClusterMapOutput) MapIndex(k pulumi.StringInput) AzureClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureCluster {
		return vs[0].(map[string]*AzureCluster)[vs[1].(string)]
	}).(AzureClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureClusterInput)(nil)).Elem(), &AzureCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureClusterArrayInput)(nil)).Elem(), AzureClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureClusterMapInput)(nil)).Elem(), AzureClusterMap{})
	pulumi.RegisterOutputType(AzureClusterOutput{})
	pulumi.RegisterOutputType(AzureClusterArrayOutput{})
	pulumi.RegisterOutputType(AzureClusterMapOutput{})
}
