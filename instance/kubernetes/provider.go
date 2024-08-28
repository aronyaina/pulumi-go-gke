package kubernetes

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GenerateProvider(ctx *pulumi.Context, name string, cluster *container.Cluster) (*kubernetes.Provider, error) {
	k8sProvider, err := kubernetes.NewProvider(ctx, "k8sprovider", &kubernetes.ProviderArgs{
		Kubeconfig: GenerateKubeconfig(cluster.Endpoint, cluster.Name, cluster.MasterAuth),
	}, pulumi.DependsOn([]pulumi.Resource{cluster}))
	if err != nil {
		return nil, err
	}
	return k8sProvider, err
}
