/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// ZoomUsApplication Schema for the Zoom app (key name: `zoomus`)  To create a Zoom app, use the [Create an Application](/openapi/okta-management/management/tag/Application/#tag/Application/operation/createApplication) request with the following parameters in the request body. > **Note:** The Zoom app only supports `SAML_2_0` sign-on mode. 
type ZoomUsApplication struct {
	Accessibility *ApplicationAccessibility `json:"accessibility,omitempty"`
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// User-defined display name for app
	Label string `json:"label"`
	Licensing *ApplicationLicensing `json:"licensing,omitempty"`
	Name string `json:"name"`
	// Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps)
	Profile map[string]map[string]interface{} `json:"profile,omitempty"`
	SignOnMode *string `json:"signOnMode,omitempty"`
	// App instance status
	Status *string `json:"status,omitempty"`
	Visibility *ApplicationVisibility `json:"visibility,omitempty"`
	Settings ZoomUsApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _ZoomUsApplication ZoomUsApplication

// NewZoomUsApplication instantiates a new ZoomUsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoomUsApplication(label string, name string, settings ZoomUsApplicationSettings) *ZoomUsApplication {
	this := ZoomUsApplication{}
	this.Label = label
	this.Name = name
	this.Settings = settings
	return &this
}

// NewZoomUsApplicationWithDefaults instantiates a new ZoomUsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoomUsApplicationWithDefaults() *ZoomUsApplication {
	this := ZoomUsApplication{}
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetAccessibility() ApplicationAccessibility {
	if o == nil || o.Accessibility == nil {
		var ret ApplicationAccessibility
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given ApplicationAccessibility and assigns it to the Accessibility field.
func (o *ZoomUsApplication) SetAccessibility(v ApplicationAccessibility) {
	o.Accessibility = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *ZoomUsApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetLabel returns the Label field value
func (o *ZoomUsApplication) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ZoomUsApplication) SetLabel(v string) {
	o.Label = v
}

// GetLicensing returns the Licensing field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetLicensing() ApplicationLicensing {
	if o == nil || o.Licensing == nil {
		var ret ApplicationLicensing
		return ret
	}
	return *o.Licensing
}

// GetLicensingOk returns a tuple with the Licensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetLicensingOk() (*ApplicationLicensing, bool) {
	if o == nil || o.Licensing == nil {
		return nil, false
	}
	return o.Licensing, true
}

// HasLicensing returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasLicensing() bool {
	if o != nil && o.Licensing != nil {
		return true
	}

	return false
}

// SetLicensing gets a reference to the given ApplicationLicensing and assigns it to the Licensing field.
func (o *ZoomUsApplication) SetLicensing(v ApplicationLicensing) {
	o.Licensing = &v
}

// GetName returns the Name field value
func (o *ZoomUsApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZoomUsApplication) SetName(v string) {
	o.Name = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetProfile() map[string]map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetProfileOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]map[string]interface{} and assigns it to the Profile field.
func (o *ZoomUsApplication) SetProfile(v map[string]map[string]interface{}) {
	o.Profile = v
}

// GetSignOnMode returns the SignOnMode field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetSignOnMode() string {
	if o == nil || o.SignOnMode == nil {
		var ret string
		return ret
	}
	return *o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetSignOnModeOk() (*string, bool) {
	if o == nil || o.SignOnMode == nil {
		return nil, false
	}
	return o.SignOnMode, true
}

// HasSignOnMode returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasSignOnMode() bool {
	if o != nil && o.SignOnMode != nil {
		return true
	}

	return false
}

// SetSignOnMode gets a reference to the given string and assigns it to the SignOnMode field.
func (o *ZoomUsApplication) SetSignOnMode(v string) {
	o.SignOnMode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ZoomUsApplication) SetStatus(v string) {
	o.Status = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ZoomUsApplication) GetVisibility() ApplicationVisibility {
	if o == nil || o.Visibility == nil {
		var ret ApplicationVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetVisibilityOk() (*ApplicationVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ZoomUsApplication) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given ApplicationVisibility and assigns it to the Visibility field.
func (o *ZoomUsApplication) SetVisibility(v ApplicationVisibility) {
	o.Visibility = &v
}

// GetSettings returns the Settings field value
func (o *ZoomUsApplication) GetSettings() ZoomUsApplicationSettings {
	if o == nil {
		var ret ZoomUsApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *ZoomUsApplication) GetSettingsOk() (*ZoomUsApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *ZoomUsApplication) SetSettings(v ZoomUsApplicationSettings) {
	o.Settings = v
}

func (o ZoomUsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if o.Licensing != nil {
		toSerialize["licensing"] = o.Licensing
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.SignOnMode != nil {
		toSerialize["signOnMode"] = o.SignOnMode
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if true {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ZoomUsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varZoomUsApplication := _ZoomUsApplication{}

	err = json.Unmarshal(bytes, &varZoomUsApplication)
	if err == nil {
		*o = ZoomUsApplication(varZoomUsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accessibility")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "label")
		delete(additionalProperties, "licensing")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "signOnMode")
		delete(additionalProperties, "status")
		delete(additionalProperties, "visibility")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableZoomUsApplication struct {
	value *ZoomUsApplication
	isSet bool
}

func (v NullableZoomUsApplication) Get() *ZoomUsApplication {
	return v.value
}

func (v *NullableZoomUsApplication) Set(val *ZoomUsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableZoomUsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableZoomUsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoomUsApplication(val *ZoomUsApplication) *NullableZoomUsApplication {
	return &NullableZoomUsApplication{value: val, isSet: true}
}

func (v NullableZoomUsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoomUsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

