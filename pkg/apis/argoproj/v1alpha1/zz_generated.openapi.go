// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/argoproj/v1alpha1.ArgoCD":       schema_pkg_apis_argoproj_v1alpha1_ArgoCD(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDSpec":   schema_pkg_apis_argoproj_v1alpha1_ArgoCDSpec(ref),
		"./pkg/apis/argoproj/v1alpha1.ArgoCDStatus": schema_pkg_apis_argoproj_v1alpha1_ArgoCDStatus(ref),
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCD(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCD is the Schema for the argocds API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/argoproj/v1alpha1.ArgoCDSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/argoproj/v1alpha1.ArgoCDStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/argoproj/v1alpha1.ArgoCDSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDSpec defines the desired state of ArgoCD",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the ArgoCD container image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version is the tag to use with the ArgoCD container image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tls": {
						SchemaProps: spec.SchemaProps{
							Description: "TLS defines the TLS options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDTLSSpec"),
						},
					},
					"dex": {
						SchemaProps: spec.SchemaProps{
							Description: "Dex defines the Dex server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDDexSpec"),
						},
					},
					"grafana": {
						SchemaProps: spec.SchemaProps{
							Description: "Grafana defines the Grafana server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDGrafanaSpec"),
						},
					},
					"prometheus": {
						SchemaProps: spec.SchemaProps{
							Description: "Prometheus defines the Prometheus server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDPrometheusSpec"),
						},
					},
					"redis": {
						SchemaProps: spec.SchemaProps{
							Description: "Redis defines the Redis server options for ArgoCD.",
							Ref:         ref("./pkg/apis/argoproj/v1alpha1.ArgoCDRedisSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/argoproj/v1alpha1.ArgoCDDexSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDGrafanaSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDPrometheusSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDRedisSpec", "./pkg/apis/argoproj/v1alpha1.ArgoCDTLSSpec"},
	}
}

func schema_pkg_apis_argoproj_v1alpha1_ArgoCDStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ArgoCDStatus defines the observed state of ArgoCD",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
