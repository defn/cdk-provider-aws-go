package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference interface {
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
	Condition() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchConditionList
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepList
	NextStepInput() interface{}
	Response() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchResponseList
	ResponseInput() interface{}
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
	PutCondition(value interface{})
	PutNextStep(value interface{})
	PutResponse(value interface{})
	ResetCondition()
	ResetNextStep()
	ResetResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference
type jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) Condition() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchConditionList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchConditionList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) NextStep() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchNextStepList
	_jsii_.Get(
		j,
		"nextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) NextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) Response() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchResponseList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchResponseList
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference{}

	_jsii_.Create(
		"aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference_Override(l Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) PutCondition(value interface{}) {
	if err := l.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCondition",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) PutNextStep(value interface{}) {
	if err := l.validatePutNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNextStep",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) PutResponse(value interface{}) {
	if err := l.validatePutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putResponse",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		l,
		"resetCondition",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ResetNextStep() {
	_jsii_.InvokeVoid(
		l,
		"resetNextStep",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		l,
		"resetResponse",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalConditionalBranchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
