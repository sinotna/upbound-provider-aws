/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BucketWebsiteConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type BucketWebsiteConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	ErrorDocument []ErrorDocumentParameters `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// +kubebuilder:validation:Optional
	IndexDocument []IndexDocumentParameters `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// +kubebuilder:validation:Optional
	RedirectAllRequestsTo []RedirectAllRequestsToParameters `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RoutingRule []RoutingRuleParameters `json:"routingRule,omitempty" tf:"routing_rule,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	HTTPErrorCodeReturnedEquals *string `json:"httpErrorCodeReturnedEquals,omitempty" tf:"http_error_code_returned_equals,omitempty"`

	// +kubebuilder:validation:Optional
	KeyPrefixEquals *string `json:"keyPrefixEquals,omitempty" tf:"key_prefix_equals,omitempty"`
}

type ErrorDocumentObservation struct {
}

type ErrorDocumentParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`
}

type IndexDocumentObservation struct {
}

type IndexDocumentParameters struct {

	// +kubebuilder:validation:Required
	Suffix *string `json:"suffix" tf:"suffix,omitempty"`
}

type RedirectAllRequestsToObservation struct {
}

type RedirectAllRequestsToParameters struct {

	// +kubebuilder:validation:Required
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RedirectObservation struct {
}

type RedirectParameters struct {

	// +kubebuilder:validation:Optional
	HTTPRedirectCode *string `json:"httpRedirectCode,omitempty" tf:"http_redirect_code,omitempty"`

	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	ReplaceKeyPrefixWith *string `json:"replaceKeyPrefixWith,omitempty" tf:"replace_key_prefix_with,omitempty"`

	// +kubebuilder:validation:Optional
	ReplaceKeyWith *string `json:"replaceKeyWith,omitempty" tf:"replace_key_with,omitempty"`
}

type RoutingRuleObservation struct {
}

type RoutingRuleParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Redirect []RedirectParameters `json:"redirect" tf:"redirect,omitempty"`
}

// BucketWebsiteConfigurationSpec defines the desired state of BucketWebsiteConfiguration
type BucketWebsiteConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketWebsiteConfigurationParameters `json:"forProvider"`
}

// BucketWebsiteConfigurationStatus defines the observed state of BucketWebsiteConfiguration.
type BucketWebsiteConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketWebsiteConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketWebsiteConfiguration is the Schema for the BucketWebsiteConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketWebsiteConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketWebsiteConfigurationSpec   `json:"spec"`
	Status            BucketWebsiteConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketWebsiteConfigurationList contains a list of BucketWebsiteConfigurations
type BucketWebsiteConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketWebsiteConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketWebsiteConfiguration_Kind             = "BucketWebsiteConfiguration"
	BucketWebsiteConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketWebsiteConfiguration_Kind}.String()
	BucketWebsiteConfiguration_KindAPIVersion   = BucketWebsiteConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketWebsiteConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketWebsiteConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketWebsiteConfiguration{}, &BucketWebsiteConfigurationList{})
}