package kubernetes

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

var AppLabel pulumi.String = "app-nginx"

func GenerateAppLabel() *pulumi.StringMap {
	appLabel := pulumi.StringMap{
		"app": AppLabel,
	}
	return &appLabel
}
