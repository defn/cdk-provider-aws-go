package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/defn/cdktf-provider-aws-go/gen/aws/jsii"

	"github.com/defn/cdktf-provider-aws-go/gen/aws/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference interface {
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
	Slot() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentSlotList
	SlotInput() interface{}
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
	PutSlot(value interface{})
	ResetName()
	ResetSlot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference
type jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) Slot() Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentSlotList {
	var returns Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentSlotList
	_jsii_.Get(
		j,
		"slot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) SlotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference{}

	_jsii_.Create(
		"aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference_Override(l Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"aws.lexv2ModelsIntent.Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) PutSlot(value interface{}) {
	if err := l.validatePutSlotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlot",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		l,
		"resetName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) ResetSlot() {
	_jsii_.InvokeVoid(
		l,
		"resetSlot",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntentConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
