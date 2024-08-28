package kubernetes

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
)

func GenerateDeployment(ctx *pulumi.Context, namespace *corev1.Namespace, appLabels pulumi.StringMap, deployment_name string, container_name string, image_name string, k8sProvider *kubernetes.Provider) (deployment *appsv1.Deployment, error error) {

	deployment, err := appsv1.NewDeployment(ctx, deployment_name, &appsv1.DeploymentArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Namespace: namespace.Metadata.Name(),
			Labels:    appLabels,
		},
		Spec: appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: appLabels,
			},
			Replicas: pulumi.Int(1),
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: appLabels,
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:  pulumi.String(container_name),
							Image: pulumi.String(image_name),
							Ports: corev1.ContainerPortArray{
								corev1.ContainerPortArgs{
									ContainerPort: pulumi.Int(80),
								},
							},
						}},
				},
			},
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return nil, err
	}
	return deployment, nil
}
