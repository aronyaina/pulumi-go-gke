package main

import (
	"k8s/instance/gcp"
	"k8s/instance/kubernetes"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		serviceAccount, err := gcp.CreateServiceAccount(ctx)
		if err != nil {
			return err
		}

		cluster, err := gcp.CreateCluster(ctx)
		if err != nil {
			return err
		}

		err = gcp.CreateNodePool(ctx, cluster, serviceAccount.Email)
		if err != nil {
			return err
		}

		provider, err := kubernetes.GenerateProvider(ctx, "gke-provider", cluster)
		if err != nil {
			return err
		}

		namespace, err := kubernetes.GenerateNamespace(ctx, provider)
		if err != nil {
			return err
		}
		appLabel := kubernetes.GenerateAppLabel()

		_, err = kubernetes.GenerateDeployment(ctx, namespace, *appLabel, "demo-deployment", "demo-container", "nginx", provider)
		if err != nil {
			return err
		}

		service, err := kubernetes.GenerateService(ctx, namespace, *appLabel, provider)
		if err != nil {
			return err
		}

		ctx.Export("ExternalUrl", service.Status.ApplyT(func(status *corev1.ServiceStatus) *string {
			ingress := status.LoadBalancer.Ingress[0]
			if ingress.Hostname != nil {
				return ingress.Hostname
			}
			return ingress.Ip
		}))

		ctx.Export("AccountMail", serviceAccount.Email)
		ctx.Export("AccountName", serviceAccount.Name)
		ctx.Export("ClusterId", cluster.ID())
		ctx.Export("ClusterIP", cluster.ClusterIpv4Cidr)
		ctx.Export("ClusterEndPoint", cluster.Endpoint)
		ctx.Export("Provider", provider)

		return nil
	})
}
