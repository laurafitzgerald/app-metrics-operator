// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsService":       schema_pkg_apis_metrics_v1alpha1_AppMetricsService(ref),
		"github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsServiceSpec":   schema_pkg_apis_metrics_v1alpha1_AppMetricsServiceSpec(ref),
		"github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsServiceStatus": schema_pkg_apis_metrics_v1alpha1_AppMetricsServiceStatus(ref),
	}
}

func schema_pkg_apis_metrics_v1alpha1_AppMetricsService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AppMetricsService is the Schema for the appmetricsservices API",
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
							Ref: ref("github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsServiceSpec", "github.com/aerogear/app-metrics-operator/pkg/apis/metrics/v1alpha1.AppMetricsServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_metrics_v1alpha1_AppMetricsServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AppMetricsServiceSpec defines the desired state of AppMetricsService",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_metrics_v1alpha1_AppMetricsServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AppMetricsServiceStatus defines the observed state of AppMetricsService",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{},
	}
}
