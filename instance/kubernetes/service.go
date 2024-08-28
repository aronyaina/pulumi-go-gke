package kubernetes

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GenerateService(ctx *pulumi.Context, namespace *corev1.Namespace, appLabels pulumi.StringMap, k8sProvider *kubernetes.Provider) (service *corev1.Service, error error) {
	service, err := corev1.NewService(ctx, "app-service", &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Namespace: namespace.Metadata.Name(),
			Labels:    appLabels,
		},
		Spec: &corev1.ServiceSpecArgs{
			Ports: corev1.ServicePortArray{
				corev1.ServicePortArgs{
					Port:       pulumi.Int(80),
					TargetPort: pulumi.Int(80),
				},
			},
			Selector: appLabels,
			Type:     pulumi.String("LoadBalancer"),
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return nil, err
	}
	return service, nil
}
