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

// checks if the FrontPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrontPortTemplateRequest{}

// FrontPortTemplateRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type FrontPortTemplateRequest struct {
	DeviceType NullableDeviceTypeRequest `json:"device_type,omitempty"`
	ModuleType NullableModuleTypeRequest `json:"module_type,omitempty"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name"`
	// Physical label
	Label                *string                 `json:"label,omitempty"`
	Type                 FrontPortTypeValue      `json:"type"`
	Color                *string                 `json:"color,omitempty"`
	RearPort             RearPortTemplateRequest `json:"rear_port"`
	RearPortPosition     *int32                  `json:"rear_port_position,omitempty"`
	Description          *string                 `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FrontPortTemplateRequest FrontPortTemplateRequest

// NewFrontPortTemplateRequest instantiates a new FrontPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPortTemplateRequest(name string, type_ FrontPortTypeValue, rearPort RearPortTemplateRequest) *FrontPortTemplateRequest {
	this := FrontPortTemplateRequest{}
	this.Name = name
	this.Type = type_
	this.RearPort = rearPort
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// NewFrontPortTemplateRequestWithDefaults instantiates a new FrontPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortTemplateRequestWithDefaults() *FrontPortTemplateRequest {
	this := FrontPortTemplateRequest{}
	var rearPortPosition int32 = 1
	this.RearPortPosition = &rearPortPosition
	return &this
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPortTemplateRequest) GetDeviceType() DeviceTypeRequest {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret DeviceTypeRequest
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplateRequest) GetDeviceTypeOk() (*DeviceTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *FrontPortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableDeviceTypeRequest and assigns it to the DeviceType field.
func (o *FrontPortTemplateRequest) SetDeviceType(v DeviceTypeRequest) {
	o.DeviceType.Set(&v)
}

// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *FrontPortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *FrontPortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FrontPortTemplateRequest) GetModuleType() ModuleTypeRequest {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret ModuleTypeRequest
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FrontPortTemplateRequest) GetModuleTypeOk() (*ModuleTypeRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *FrontPortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableModuleTypeRequest and assigns it to the ModuleType field.
func (o *FrontPortTemplateRequest) SetModuleType(v ModuleTypeRequest) {
	o.ModuleType.Set(&v)
}

// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *FrontPortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *FrontPortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetName returns the Name field value
func (o *FrontPortTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPortTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value
func (o *FrontPortTemplateRequest) GetType() FrontPortTypeValue {
	if o == nil {
		var ret FrontPortTypeValue
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetTypeOk() (*FrontPortTypeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FrontPortTemplateRequest) SetType(v FrontPortTypeValue) {
	o.Type = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *FrontPortTemplateRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *FrontPortTemplateRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *FrontPortTemplateRequest) SetColor(v string) {
	o.Color = &v
}

// GetRearPort returns the RearPort field value
func (o *FrontPortTemplateRequest) GetRearPort() RearPortTemplateRequest {
	if o == nil {
		var ret RearPortTemplateRequest
		return ret
	}

	return o.RearPort
}

// GetRearPortOk returns a tuple with the RearPort field value
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetRearPortOk() (*RearPortTemplateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RearPort, true
}

// SetRearPort sets field value
func (o *FrontPortTemplateRequest) SetRearPort(v RearPortTemplateRequest) {
	o.RearPort = v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *FrontPortTemplateRequest) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *FrontPortTemplateRequest) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *FrontPortTemplateRequest) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o FrontPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrontPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	toSerialize["rear_port"] = o.RearPort
	if !IsNil(o.RearPortPosition) {
		toSerialize["rear_port_position"] = o.RearPortPosition
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrontPortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"rear_port",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFrontPortTemplateRequest := _FrontPortTemplateRequest{}

	err = json.Unmarshal(data, &varFrontPortTemplateRequest)

	if err != nil {
		return err
	}

	*o = FrontPortTemplateRequest(varFrontPortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		delete(additionalProperties, "color")
		delete(additionalProperties, "rear_port")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFrontPortTemplateRequest struct {
	value *FrontPortTemplateRequest
	isSet bool
}

func (v NullableFrontPortTemplateRequest) Get() *FrontPortTemplateRequest {
	return v.value
}

func (v *NullableFrontPortTemplateRequest) Set(val *FrontPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortTemplateRequest(val *FrontPortTemplateRequest) *NullableFrontPortTemplateRequest {
	return &NullableFrontPortTemplateRequest{value: val, isSet: true}
}

func (v NullableFrontPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}