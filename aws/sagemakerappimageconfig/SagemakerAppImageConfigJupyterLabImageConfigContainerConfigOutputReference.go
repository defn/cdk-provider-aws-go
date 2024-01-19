package sagemakerappimageconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/sagemakerappimageconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference interface {
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
	ContainerArguments() *[]*string
	SetContainerArguments(val *[]*string)
	ContainerArgumentsInput() *[]*string
	ContainerEntrypoint() *[]*string
	SetContainerEntrypoint(val *[]*string)
	ContainerEntrypointInput() *[]*string
	ContainerEnvironmentVariables() *map[string]*string
	SetContainerEnvironmentVariables(val *map[string]*string)
	ContainerEnvironmentVariablesInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig
	SetInternalValue(val *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetContainerArguments()
	ResetContainerEntrypoint()
	ResetContainerEnvironmentVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference
type jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ContainerArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ContainerArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ContainerEntrypoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ContainerEntrypointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ContainerEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"containerEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ContainerEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"containerEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) InternalValue() *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig {
	var returns *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference{}

	_jsii_.Create(
		"aws.sagemakerAppImageConfig.SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference_Override(s SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.sagemakerAppImageConfig.SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetContainerArguments(val *[]*string) {
	if err := j.validateSetContainerArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerArguments",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetContainerEntrypoint(val *[]*string) {
	if err := j.validateSetContainerEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEntrypoint",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetContainerEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetContainerEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetInternalValue(val *SagemakerAppImageConfigJupyterLabImageConfigContainerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ResetContainerArguments() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerArguments",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ResetContainerEntrypoint() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerEntrypoint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ResetContainerEnvironmentVariables() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerEnvironmentVariables",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerAppImageConfigJupyterLabImageConfigContainerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
