/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.3 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// ExtrasCustomFieldsListUiVisibleParameter the model 'ExtrasCustomFieldsListUiVisibleParameter'
type ExtrasCustomFieldsListUiVisibleParameter string

// List of extras_custom_fields_list_ui_visible_parameter
const (
	EXTRASCUSTOMFIELDSLISTUIVISIBLEPARAMETER_ALWAYS ExtrasCustomFieldsListUiVisibleParameter = "always"
	EXTRASCUSTOMFIELDSLISTUIVISIBLEPARAMETER_HIDDEN ExtrasCustomFieldsListUiVisibleParameter = "hidden"
	EXTRASCUSTOMFIELDSLISTUIVISIBLEPARAMETER_IF_SET ExtrasCustomFieldsListUiVisibleParameter = "if-set"
)

// All allowed values of ExtrasCustomFieldsListUiVisibleParameter enum
var AllowedExtrasCustomFieldsListUiVisibleParameterEnumValues = []ExtrasCustomFieldsListUiVisibleParameter{
	"always",
	"hidden",
	"if-set",
}

func (v *ExtrasCustomFieldsListUiVisibleParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtrasCustomFieldsListUiVisibleParameter(value)
	for _, existing := range AllowedExtrasCustomFieldsListUiVisibleParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtrasCustomFieldsListUiVisibleParameter", value)
}

// NewExtrasCustomFieldsListUiVisibleParameterFromValue returns a pointer to a valid ExtrasCustomFieldsListUiVisibleParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtrasCustomFieldsListUiVisibleParameterFromValue(v string) (*ExtrasCustomFieldsListUiVisibleParameter, error) {
	ev := ExtrasCustomFieldsListUiVisibleParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtrasCustomFieldsListUiVisibleParameter: valid values are %v", v, AllowedExtrasCustomFieldsListUiVisibleParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtrasCustomFieldsListUiVisibleParameter) IsValid() bool {
	for _, existing := range AllowedExtrasCustomFieldsListUiVisibleParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to extras_custom_fields_list_ui_visible_parameter value
func (v ExtrasCustomFieldsListUiVisibleParameter) Ptr() *ExtrasCustomFieldsListUiVisibleParameter {
	return &v
}

type NullableExtrasCustomFieldsListUiVisibleParameter struct {
	value *ExtrasCustomFieldsListUiVisibleParameter
	isSet bool
}

func (v NullableExtrasCustomFieldsListUiVisibleParameter) Get() *ExtrasCustomFieldsListUiVisibleParameter {
	return v.value
}

func (v *NullableExtrasCustomFieldsListUiVisibleParameter) Set(val *ExtrasCustomFieldsListUiVisibleParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableExtrasCustomFieldsListUiVisibleParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableExtrasCustomFieldsListUiVisibleParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtrasCustomFieldsListUiVisibleParameter(val *ExtrasCustomFieldsListUiVisibleParameter) *NullableExtrasCustomFieldsListUiVisibleParameter {
	return &NullableExtrasCustomFieldsListUiVisibleParameter{value: val, isSet: true}
}

func (v NullableExtrasCustomFieldsListUiVisibleParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtrasCustomFieldsListUiVisibleParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}