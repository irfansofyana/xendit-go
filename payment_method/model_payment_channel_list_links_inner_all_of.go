/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the PaymentChannelListLinksInnerAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentChannelListLinksInnerAllOf{}

// PaymentChannelListLinksInnerAllOf struct for PaymentChannelListLinksInnerAllOf
type PaymentChannelListLinksInnerAllOf struct {
	// Target URI that should contain a target to Internationalized Resource Identifiers (IRI)
	Href *string `json:"href,omitempty"`
	// The link relation type described how the current context (source) is related to target resource
	Rel *string `json:"rel,omitempty"`
	// The HTTP method to be used to call `href`
	Method *string `json:"method,omitempty"`
}

// NewPaymentChannelListLinksInnerAllOf instantiates a new PaymentChannelListLinksInnerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentChannelListLinksInnerAllOf() *PaymentChannelListLinksInnerAllOf {
	this := PaymentChannelListLinksInnerAllOf{}
	return &this
}

// NewPaymentChannelListLinksInnerAllOfWithDefaults instantiates a new PaymentChannelListLinksInnerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentChannelListLinksInnerAllOfWithDefaults() *PaymentChannelListLinksInnerAllOf {
	this := PaymentChannelListLinksInnerAllOf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PaymentChannelListLinksInnerAllOf) GetHref() string {
	if o == nil || utils.IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelListLinksInnerAllOf) GetHrefOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PaymentChannelListLinksInnerAllOf) HasHref() bool {
	if o != nil && !utils.IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PaymentChannelListLinksInnerAllOf) SetHref(v string) {
	o.Href = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *PaymentChannelListLinksInnerAllOf) GetRel() string {
	if o == nil || utils.IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelListLinksInnerAllOf) GetRelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Rel) {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *PaymentChannelListLinksInnerAllOf) HasRel() bool {
	if o != nil && !utils.IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *PaymentChannelListLinksInnerAllOf) SetRel(v string) {
	o.Rel = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentChannelListLinksInnerAllOf) GetMethod() string {
	if o == nil || utils.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelListLinksInnerAllOf) GetMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentChannelListLinksInnerAllOf) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentChannelListLinksInnerAllOf) SetMethod(v string) {
	o.Method = &v
}

func (o PaymentChannelListLinksInnerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentChannelListLinksInnerAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !utils.IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullablePaymentChannelListLinksInnerAllOf struct {
	value *PaymentChannelListLinksInnerAllOf
	isSet bool
}

func (v NullablePaymentChannelListLinksInnerAllOf) Get() *PaymentChannelListLinksInnerAllOf {
	return v.value
}

func (v *NullablePaymentChannelListLinksInnerAllOf) Set(val *PaymentChannelListLinksInnerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannelListLinksInnerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannelListLinksInnerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannelListLinksInnerAllOf(val *PaymentChannelListLinksInnerAllOf) *NullablePaymentChannelListLinksInnerAllOf {
	return &NullablePaymentChannelListLinksInnerAllOf{value: val, isSet: true}
}

func (v NullablePaymentChannelListLinksInnerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannelListLinksInnerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

