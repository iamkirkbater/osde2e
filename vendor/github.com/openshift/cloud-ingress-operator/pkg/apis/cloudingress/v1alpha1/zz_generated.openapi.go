// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APIScheme":       schema_pkg_apis_cloudingress_v1alpha1_APIScheme(ref),
		"github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeSpec":   schema_pkg_apis_cloudingress_v1alpha1_APISchemeSpec(ref),
		"github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeStatus": schema_pkg_apis_cloudingress_v1alpha1_APISchemeStatus(ref),
	}
}

func schema_pkg_apis_cloudingress_v1alpha1_APIScheme(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APIScheme is the Schema for the APISchemes API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
							Ref: ref("github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeSpec", "github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cloudingress_v1alpha1_APISchemeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APISchemeSpec defines the desired state of APIScheme",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"managementAPIServerIngress": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.ManagementAPIServerIngress"),
						},
					},
				},
				Required: []string{"managementAPIServerIngress"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.ManagementAPIServerIngress"},
	}
}

func schema_pkg_apis_cloudingress_v1alpha1_APISchemeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "APISchemeStatus defines the observed state of APIScheme",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cloudLoadBalancerDNSName": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeCondition"),
									},
								},
							},
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/cloud-ingress-operator/pkg/apis/cloudingress/v1alpha1.APISchemeCondition"},
	}
}