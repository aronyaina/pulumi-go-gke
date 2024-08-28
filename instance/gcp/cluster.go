package gcp

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const CLUSTER_NAME = "primary"

var Name pulumi.String = "my-gke-cluster"
var Location pulumi.String = "us-central1"
var RemoveDefaultNodePool pulumi.Bool = true
var InitialNodeCount pulumi.Int = 1
var DeletionProtection pulumi.Bool = false

func CreateCluster(ctx *pulumi.Context) (*container.Cluster, error) {
	return container.NewCluster(ctx, CLUSTER_NAME, &container.ClusterArgs{
		Name:                  Name,
		Location:              Location,
		RemoveDefaultNodePool: RemoveDefaultNodePool,
		InitialNodeCount:      InitialNodeCount,
		DeletionProtection:    DeletionProtection,
	})
}
