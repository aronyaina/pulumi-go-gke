package kubernetes

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var Namespace string = "staging-ns"

func GenerateNamespace(ctx *pulumi.Context, k8sProvider *kubernetes.Provider) (new_namespace *corev1.Namespace, err error) {
	namespace, err := corev1.NewNamespace(ctx, Namespace, &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(Namespace),
		},
	}, pulumi.Provider(k8sProvider))
	if err != nil {
		return nil, err
	}
	return namespace, nil
}
