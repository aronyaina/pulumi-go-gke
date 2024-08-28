package gcp

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var MachineType pulumi.String = "e2-micro"
var Zone pulumi.String = "us-central1"
var NodePoolName pulumi.String = "k8s-node-pool"
var NodeCount pulumi.Int = 2
var MinNodeCount pulumi.Int = 1
var MaxNodeCount pulumi.Int = 2

func CreateNodePool(ctx *pulumi.Context, cluster *container.Cluster, serviceAccountEmail pulumi.StringInput) error {
	_, err := container.NewNodePool(ctx, "primary_preemptible_nodes", &container.NodePoolArgs{
		Name:      NodePoolName,
		Location:  Zone,
		Cluster:   cluster.ID(),
		NodeCount: NodeCount,
		NodeConfig: &container.NodePoolNodeConfigArgs{
			Preemptible:    pulumi.Bool(true),
			MachineType:    MachineType,
			ServiceAccount: serviceAccountEmail,
			OauthScopes: pulumi.StringArray{
				pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
				pulumi.String("https://www.googleapis.com/auth/compute"),
				pulumi.String("https://www.googleapis.com/auth/devstorage.read_only"),
				pulumi.String("https://www.googleapis.com/auth/logging.write"),
				pulumi.String("https://www.googleapis.com/auth/monitoring"),
			},
		},
		Autoscaling: &container.NodePoolAutoscalingArgs{
			MinNodeCount: MinNodeCount,
			MaxNodeCount: MaxNodeCount,
		},
	})
	return err
}
