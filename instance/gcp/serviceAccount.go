package gcp

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var AccountID pulumi.String = "pipeline-service"
var DisplayName pulumi.String = "Linux Service Account"

func CreateServiceAccount(ctx *pulumi.Context) (*serviceaccount.Account, error) {
	return serviceaccount.NewAccount(ctx, "default", &serviceaccount.AccountArgs{
		AccountId:   AccountID,
		DisplayName: DisplayName,
	})
}
