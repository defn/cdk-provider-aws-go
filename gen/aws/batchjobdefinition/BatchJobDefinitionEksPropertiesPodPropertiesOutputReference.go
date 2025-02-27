package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/batchjobdefinition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchJobDefinitionEksPropertiesPodPropertiesOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Containers() BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference
	ContainersInput() *BatchJobDefinitionEksPropertiesPodPropertiesContainers
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsPolicy() *string
	SetDnsPolicy(val *string)
	DnsPolicyInput() *string
	// Experimental.
	Fqn() *string
	HostNetwork() interface{}
	SetHostNetwork(val interface{})
	HostNetworkInput() interface{}
	InternalValue() *BatchJobDefinitionEksPropertiesPodProperties
	SetInternalValue(val *BatchJobDefinitionEksPropertiesPodProperties)
	Metadata() BatchJobDefinitionEksPropertiesPodPropertiesMetadataOutputReference
	MetadataInput() *BatchJobDefinitionEksPropertiesPodPropertiesMetadata
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Volumes() BatchJobDefinitionEksPropertiesPodPropertiesVolumesList
	VolumesInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutContainers(value *BatchJobDefinitionEksPropertiesPodPropertiesContainers)
	PutMetadata(value *BatchJobDefinitionEksPropertiesPodPropertiesMetadata)
	PutVolumes(value interface{})
	ResetDnsPolicy()
	ResetHostNetwork()
	ResetMetadata()
	ResetServiceAccountName()
	ResetVolumes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionEksPropertiesPodPropertiesOutputReference
type jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) Containers() BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesContainersOutputReference
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ContainersInput() *BatchJobDefinitionEksPropertiesPodPropertiesContainers {
	var returns *BatchJobDefinitionEksPropertiesPodPropertiesContainers
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) DnsPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) DnsPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) HostNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) HostNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) InternalValue() *BatchJobDefinitionEksPropertiesPodProperties {
	var returns *BatchJobDefinitionEksPropertiesPodProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) Metadata() BatchJobDefinitionEksPropertiesPodPropertiesMetadataOutputReference {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) MetadataInput() *BatchJobDefinitionEksPropertiesPodPropertiesMetadata {
	var returns *BatchJobDefinitionEksPropertiesPodPropertiesMetadata
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) Volumes() BatchJobDefinitionEksPropertiesPodPropertiesVolumesList {
	var returns BatchJobDefinitionEksPropertiesPodPropertiesVolumesList
	_jsii_.Get(
		j,
		"volumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) VolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumesInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionEksPropertiesPodPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BatchJobDefinitionEksPropertiesPodPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionEksPropertiesPodPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference{}

	_jsii_.Create(
		"aws.batchJobDefinition.BatchJobDefinitionEksPropertiesPodPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionEksPropertiesPodPropertiesOutputReference_Override(b BatchJobDefinitionEksPropertiesPodPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.batchJobDefinition.BatchJobDefinitionEksPropertiesPodPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetDnsPolicy(val *string) {
	if err := j.validateSetDnsPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPolicy",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetHostNetwork(val interface{}) {
	if err := j.validateSetHostNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostNetwork",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetInternalValue(val *BatchJobDefinitionEksPropertiesPodProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) PutContainers(value *BatchJobDefinitionEksPropertiesPodPropertiesContainers) {
	if err := b.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putContainers",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) PutMetadata(value *BatchJobDefinitionEksPropertiesPodPropertiesMetadata) {
	if err := b.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putMetadata",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) PutVolumes(value interface{}) {
	if err := b.validatePutVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVolumes",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ResetDnsPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetDnsPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ResetHostNetwork() {
	_jsii_.InvokeVoid(
		b,
		"resetHostNetwork",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		b,
		"resetMetadata",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		b,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ResetVolumes() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEksPropertiesPodPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

