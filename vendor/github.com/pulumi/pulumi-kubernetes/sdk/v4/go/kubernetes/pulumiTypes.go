// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Options to configure the Helm Release resource.
type HelmReleaseSettings struct {
	// The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
	Driver *string `pulumi:"driver"`
	// The path to the helm plugins directory.
	PluginsPath *string `pulumi:"pluginsPath"`
	// The path to the registry config file.
	RegistryConfigPath *string `pulumi:"registryConfigPath"`
	// The path to the directory containing cached repository indexes.
	RepositoryCache *string `pulumi:"repositoryCache"`
	// The path to the file containing repository names and URLs.
	RepositoryConfigPath *string `pulumi:"repositoryConfigPath"`
}

// Defaults sets the appropriate defaults for HelmReleaseSettings
func (val *HelmReleaseSettings) Defaults() *HelmReleaseSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Driver == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_DRIVER"); d != nil {
			driver_ := d.(string)
			tmp.Driver = &driver_
		}
	}
	if tmp.PluginsPath == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_PLUGINS_PATH"); d != nil {
			pluginsPath_ := d.(string)
			tmp.PluginsPath = &pluginsPath_
		}
	}
	if tmp.RegistryConfigPath == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_REGISTRY_CONFIG_PATH"); d != nil {
			registryConfigPath_ := d.(string)
			tmp.RegistryConfigPath = &registryConfigPath_
		}
	}
	if tmp.RepositoryCache == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_REPOSITORY_CACHE"); d != nil {
			repositoryCache_ := d.(string)
			tmp.RepositoryCache = &repositoryCache_
		}
	}
	if tmp.RepositoryConfigPath == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_REPOSITORY_CONFIG_PATH"); d != nil {
			repositoryConfigPath_ := d.(string)
			tmp.RepositoryConfigPath = &repositoryConfigPath_
		}
	}
	return &tmp
}

// HelmReleaseSettingsInput is an input type that accepts HelmReleaseSettingsArgs and HelmReleaseSettingsOutput values.
// You can construct a concrete instance of `HelmReleaseSettingsInput` via:
//
//	HelmReleaseSettingsArgs{...}
type HelmReleaseSettingsInput interface {
	pulumi.Input

	ToHelmReleaseSettingsOutput() HelmReleaseSettingsOutput
	ToHelmReleaseSettingsOutputWithContext(context.Context) HelmReleaseSettingsOutput
}

// Options to configure the Helm Release resource.
type HelmReleaseSettingsArgs struct {
	// The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
	Driver pulumi.StringPtrInput `pulumi:"driver"`
	// The path to the helm plugins directory.
	PluginsPath pulumi.StringPtrInput `pulumi:"pluginsPath"`
	// The path to the registry config file.
	RegistryConfigPath pulumi.StringPtrInput `pulumi:"registryConfigPath"`
	// The path to the directory containing cached repository indexes.
	RepositoryCache pulumi.StringPtrInput `pulumi:"repositoryCache"`
	// The path to the file containing repository names and URLs.
	RepositoryConfigPath pulumi.StringPtrInput `pulumi:"repositoryConfigPath"`
}

// Defaults sets the appropriate defaults for HelmReleaseSettingsArgs
func (val *HelmReleaseSettingsArgs) Defaults() *HelmReleaseSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Driver == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_DRIVER"); d != nil {
			tmp.Driver = pulumi.StringPtr(d.(string))
		}
	}
	if tmp.PluginsPath == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_PLUGINS_PATH"); d != nil {
			tmp.PluginsPath = pulumi.StringPtr(d.(string))
		}
	}
	if tmp.RegistryConfigPath == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_REGISTRY_CONFIG_PATH"); d != nil {
			tmp.RegistryConfigPath = pulumi.StringPtr(d.(string))
		}
	}
	if tmp.RepositoryCache == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_REPOSITORY_CACHE"); d != nil {
			tmp.RepositoryCache = pulumi.StringPtr(d.(string))
		}
	}
	if tmp.RepositoryConfigPath == nil {
		if d := utilities.GetEnvOrDefault(nil, nil, "PULUMI_K8S_HELM_REPOSITORY_CONFIG_PATH"); d != nil {
			tmp.RepositoryConfigPath = pulumi.StringPtr(d.(string))
		}
	}
	return &tmp
}
func (HelmReleaseSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmReleaseSettings)(nil)).Elem()
}

func (i HelmReleaseSettingsArgs) ToHelmReleaseSettingsOutput() HelmReleaseSettingsOutput {
	return i.ToHelmReleaseSettingsOutputWithContext(context.Background())
}

func (i HelmReleaseSettingsArgs) ToHelmReleaseSettingsOutputWithContext(ctx context.Context) HelmReleaseSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmReleaseSettingsOutput)
}

func (i HelmReleaseSettingsArgs) ToHelmReleaseSettingsPtrOutput() HelmReleaseSettingsPtrOutput {
	return i.ToHelmReleaseSettingsPtrOutputWithContext(context.Background())
}

func (i HelmReleaseSettingsArgs) ToHelmReleaseSettingsPtrOutputWithContext(ctx context.Context) HelmReleaseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmReleaseSettingsOutput).ToHelmReleaseSettingsPtrOutputWithContext(ctx)
}

// HelmReleaseSettingsPtrInput is an input type that accepts HelmReleaseSettingsArgs, HelmReleaseSettingsPtr and HelmReleaseSettingsPtrOutput values.
// You can construct a concrete instance of `HelmReleaseSettingsPtrInput` via:
//
//	        HelmReleaseSettingsArgs{...}
//
//	or:
//
//	        nil
type HelmReleaseSettingsPtrInput interface {
	pulumi.Input

	ToHelmReleaseSettingsPtrOutput() HelmReleaseSettingsPtrOutput
	ToHelmReleaseSettingsPtrOutputWithContext(context.Context) HelmReleaseSettingsPtrOutput
}

type helmReleaseSettingsPtrType HelmReleaseSettingsArgs

func HelmReleaseSettingsPtr(v *HelmReleaseSettingsArgs) HelmReleaseSettingsPtrInput {
	return (*helmReleaseSettingsPtrType)(v)
}

func (*helmReleaseSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmReleaseSettings)(nil)).Elem()
}

func (i *helmReleaseSettingsPtrType) ToHelmReleaseSettingsPtrOutput() HelmReleaseSettingsPtrOutput {
	return i.ToHelmReleaseSettingsPtrOutputWithContext(context.Background())
}

func (i *helmReleaseSettingsPtrType) ToHelmReleaseSettingsPtrOutputWithContext(ctx context.Context) HelmReleaseSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HelmReleaseSettingsPtrOutput)
}

// Options to configure the Helm Release resource.
type HelmReleaseSettingsOutput struct{ *pulumi.OutputState }

func (HelmReleaseSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HelmReleaseSettings)(nil)).Elem()
}

func (o HelmReleaseSettingsOutput) ToHelmReleaseSettingsOutput() HelmReleaseSettingsOutput {
	return o
}

func (o HelmReleaseSettingsOutput) ToHelmReleaseSettingsOutputWithContext(ctx context.Context) HelmReleaseSettingsOutput {
	return o
}

func (o HelmReleaseSettingsOutput) ToHelmReleaseSettingsPtrOutput() HelmReleaseSettingsPtrOutput {
	return o.ToHelmReleaseSettingsPtrOutputWithContext(context.Background())
}

func (o HelmReleaseSettingsOutput) ToHelmReleaseSettingsPtrOutputWithContext(ctx context.Context) HelmReleaseSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HelmReleaseSettings) *HelmReleaseSettings {
		return &v
	}).(HelmReleaseSettingsPtrOutput)
}

// The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
func (o HelmReleaseSettingsOutput) Driver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmReleaseSettings) *string { return v.Driver }).(pulumi.StringPtrOutput)
}

// The path to the helm plugins directory.
func (o HelmReleaseSettingsOutput) PluginsPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmReleaseSettings) *string { return v.PluginsPath }).(pulumi.StringPtrOutput)
}

// The path to the registry config file.
func (o HelmReleaseSettingsOutput) RegistryConfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmReleaseSettings) *string { return v.RegistryConfigPath }).(pulumi.StringPtrOutput)
}

// The path to the directory containing cached repository indexes.
func (o HelmReleaseSettingsOutput) RepositoryCache() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmReleaseSettings) *string { return v.RepositoryCache }).(pulumi.StringPtrOutput)
}

// The path to the file containing repository names and URLs.
func (o HelmReleaseSettingsOutput) RepositoryConfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HelmReleaseSettings) *string { return v.RepositoryConfigPath }).(pulumi.StringPtrOutput)
}

type HelmReleaseSettingsPtrOutput struct{ *pulumi.OutputState }

func (HelmReleaseSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HelmReleaseSettings)(nil)).Elem()
}

func (o HelmReleaseSettingsPtrOutput) ToHelmReleaseSettingsPtrOutput() HelmReleaseSettingsPtrOutput {
	return o
}

func (o HelmReleaseSettingsPtrOutput) ToHelmReleaseSettingsPtrOutputWithContext(ctx context.Context) HelmReleaseSettingsPtrOutput {
	return o
}

func (o HelmReleaseSettingsPtrOutput) Elem() HelmReleaseSettingsOutput {
	return o.ApplyT(func(v *HelmReleaseSettings) HelmReleaseSettings {
		if v != nil {
			return *v
		}
		var ret HelmReleaseSettings
		return ret
	}).(HelmReleaseSettingsOutput)
}

// The backend storage driver for Helm. Values are: configmap, secret, memory, sql.
func (o HelmReleaseSettingsPtrOutput) Driver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmReleaseSettings) *string {
		if v == nil {
			return nil
		}
		return v.Driver
	}).(pulumi.StringPtrOutput)
}

// The path to the helm plugins directory.
func (o HelmReleaseSettingsPtrOutput) PluginsPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmReleaseSettings) *string {
		if v == nil {
			return nil
		}
		return v.PluginsPath
	}).(pulumi.StringPtrOutput)
}

// The path to the registry config file.
func (o HelmReleaseSettingsPtrOutput) RegistryConfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmReleaseSettings) *string {
		if v == nil {
			return nil
		}
		return v.RegistryConfigPath
	}).(pulumi.StringPtrOutput)
}

// The path to the directory containing cached repository indexes.
func (o HelmReleaseSettingsPtrOutput) RepositoryCache() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmReleaseSettings) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryCache
	}).(pulumi.StringPtrOutput)
}

// The path to the file containing repository names and URLs.
func (o HelmReleaseSettingsPtrOutput) RepositoryConfigPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HelmReleaseSettings) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryConfigPath
	}).(pulumi.StringPtrOutput)
}

// Options for tuning the Kubernetes client used by a Provider.
type KubeClientSettings struct {
	// Maximum burst for throttle. Default value is 120.
	Burst *int `pulumi:"burst"`
	// Maximum queries per second (QPS) to the API server from this client. Default value is 50.
	Qps *float64 `pulumi:"qps"`
	// Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
	Timeout *int `pulumi:"timeout"`
}

// Defaults sets the appropriate defaults for KubeClientSettings
func (val *KubeClientSettings) Defaults() *KubeClientSettings {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Burst == nil {
		if d := utilities.GetEnvOrDefault(nil, utilities.ParseEnvInt, "PULUMI_K8S_CLIENT_BURST"); d != nil {
			burst_ := d.(int)
			tmp.Burst = &burst_
		}
	}
	if tmp.Qps == nil {
		if d := utilities.GetEnvOrDefault(nil, utilities.ParseEnvFloat, "PULUMI_K8S_CLIENT_QPS"); d != nil {
			qps_ := d.(float64)
			tmp.Qps = &qps_
		}
	}
	if tmp.Timeout == nil {
		if d := utilities.GetEnvOrDefault(nil, utilities.ParseEnvInt, "PULUMI_K8S_CLIENT_TIMEOUT"); d != nil {
			timeout_ := d.(int)
			tmp.Timeout = &timeout_
		}
	}
	return &tmp
}

// KubeClientSettingsInput is an input type that accepts KubeClientSettingsArgs and KubeClientSettingsOutput values.
// You can construct a concrete instance of `KubeClientSettingsInput` via:
//
//	KubeClientSettingsArgs{...}
type KubeClientSettingsInput interface {
	pulumi.Input

	ToKubeClientSettingsOutput() KubeClientSettingsOutput
	ToKubeClientSettingsOutputWithContext(context.Context) KubeClientSettingsOutput
}

// Options for tuning the Kubernetes client used by a Provider.
type KubeClientSettingsArgs struct {
	// Maximum burst for throttle. Default value is 120.
	Burst pulumi.IntPtrInput `pulumi:"burst"`
	// Maximum queries per second (QPS) to the API server from this client. Default value is 50.
	Qps pulumi.Float64PtrInput `pulumi:"qps"`
	// Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
	Timeout pulumi.IntPtrInput `pulumi:"timeout"`
}

// Defaults sets the appropriate defaults for KubeClientSettingsArgs
func (val *KubeClientSettingsArgs) Defaults() *KubeClientSettingsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Burst == nil {
		if d := utilities.GetEnvOrDefault(nil, utilities.ParseEnvInt, "PULUMI_K8S_CLIENT_BURST"); d != nil {
			tmp.Burst = pulumi.IntPtr(d.(int))
		}
	}
	if tmp.Qps == nil {
		if d := utilities.GetEnvOrDefault(nil, utilities.ParseEnvFloat, "PULUMI_K8S_CLIENT_QPS"); d != nil {
			tmp.Qps = pulumi.Float64Ptr(d.(float64))
		}
	}
	if tmp.Timeout == nil {
		if d := utilities.GetEnvOrDefault(nil, utilities.ParseEnvInt, "PULUMI_K8S_CLIENT_TIMEOUT"); d != nil {
			tmp.Timeout = pulumi.IntPtr(d.(int))
		}
	}
	return &tmp
}
func (KubeClientSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeClientSettings)(nil)).Elem()
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsOutput() KubeClientSettingsOutput {
	return i.ToKubeClientSettingsOutputWithContext(context.Background())
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsOutputWithContext(ctx context.Context) KubeClientSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeClientSettingsOutput)
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return i.ToKubeClientSettingsPtrOutputWithContext(context.Background())
}

func (i KubeClientSettingsArgs) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeClientSettingsOutput).ToKubeClientSettingsPtrOutputWithContext(ctx)
}

// KubeClientSettingsPtrInput is an input type that accepts KubeClientSettingsArgs, KubeClientSettingsPtr and KubeClientSettingsPtrOutput values.
// You can construct a concrete instance of `KubeClientSettingsPtrInput` via:
//
//	        KubeClientSettingsArgs{...}
//
//	or:
//
//	        nil
type KubeClientSettingsPtrInput interface {
	pulumi.Input

	ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput
	ToKubeClientSettingsPtrOutputWithContext(context.Context) KubeClientSettingsPtrOutput
}

type kubeClientSettingsPtrType KubeClientSettingsArgs

func KubeClientSettingsPtr(v *KubeClientSettingsArgs) KubeClientSettingsPtrInput {
	return (*kubeClientSettingsPtrType)(v)
}

func (*kubeClientSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeClientSettings)(nil)).Elem()
}

func (i *kubeClientSettingsPtrType) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return i.ToKubeClientSettingsPtrOutputWithContext(context.Background())
}

func (i *kubeClientSettingsPtrType) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeClientSettingsPtrOutput)
}

// Options for tuning the Kubernetes client used by a Provider.
type KubeClientSettingsOutput struct{ *pulumi.OutputState }

func (KubeClientSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeClientSettings)(nil)).Elem()
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsOutput() KubeClientSettingsOutput {
	return o
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsOutputWithContext(ctx context.Context) KubeClientSettingsOutput {
	return o
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return o.ToKubeClientSettingsPtrOutputWithContext(context.Background())
}

func (o KubeClientSettingsOutput) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeClientSettings) *KubeClientSettings {
		return &v
	}).(KubeClientSettingsPtrOutput)
}

// Maximum burst for throttle. Default value is 120.
func (o KubeClientSettingsOutput) Burst() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeClientSettings) *int { return v.Burst }).(pulumi.IntPtrOutput)
}

// Maximum queries per second (QPS) to the API server from this client. Default value is 50.
func (o KubeClientSettingsOutput) Qps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v KubeClientSettings) *float64 { return v.Qps }).(pulumi.Float64PtrOutput)
}

// Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
func (o KubeClientSettingsOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v KubeClientSettings) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

type KubeClientSettingsPtrOutput struct{ *pulumi.OutputState }

func (KubeClientSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeClientSettings)(nil)).Elem()
}

func (o KubeClientSettingsPtrOutput) ToKubeClientSettingsPtrOutput() KubeClientSettingsPtrOutput {
	return o
}

func (o KubeClientSettingsPtrOutput) ToKubeClientSettingsPtrOutputWithContext(ctx context.Context) KubeClientSettingsPtrOutput {
	return o
}

func (o KubeClientSettingsPtrOutput) Elem() KubeClientSettingsOutput {
	return o.ApplyT(func(v *KubeClientSettings) KubeClientSettings {
		if v != nil {
			return *v
		}
		var ret KubeClientSettings
		return ret
	}).(KubeClientSettingsOutput)
}

// Maximum burst for throttle. Default value is 120.
func (o KubeClientSettingsPtrOutput) Burst() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeClientSettings) *int {
		if v == nil {
			return nil
		}
		return v.Burst
	}).(pulumi.IntPtrOutput)
}

// Maximum queries per second (QPS) to the API server from this client. Default value is 50.
func (o KubeClientSettingsPtrOutput) Qps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *KubeClientSettings) *float64 {
		if v == nil {
			return nil
		}
		return v.Qps
	}).(pulumi.Float64PtrOutput)
}

// Maximum time in seconds to wait before cancelling a HTTP request to the Kubernetes server. Default value is 32.
func (o KubeClientSettingsPtrOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *KubeClientSettings) *int {
		if v == nil {
			return nil
		}
		return v.Timeout
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HelmReleaseSettingsInput)(nil)).Elem(), HelmReleaseSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HelmReleaseSettingsPtrInput)(nil)).Elem(), HelmReleaseSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeClientSettingsInput)(nil)).Elem(), KubeClientSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeClientSettingsPtrInput)(nil)).Elem(), KubeClientSettingsArgs{})
	pulumi.RegisterOutputType(HelmReleaseSettingsOutput{})
	pulumi.RegisterOutputType(HelmReleaseSettingsPtrOutput{})
	pulumi.RegisterOutputType(KubeClientSettingsOutput{})
	pulumi.RegisterOutputType(KubeClientSettingsPtrOutput{})
}
