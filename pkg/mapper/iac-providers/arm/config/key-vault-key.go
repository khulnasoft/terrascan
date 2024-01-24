/*
    Copyright (C) 2022 Tenable, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package config

import (
	"time"

	"github.com/khulnasoft/terrascan/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrascan/pkg/mapper/iac-providers/arm/types"
)

const armAttributes = "attributes"
const tfExpirationDate = "expiration_date"

// KeyVaultKeyConfig returns config for azurerm_key_vault_key
func KeyVaultKeyConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	attr := convert.ToMap(r.Properties, armAttributes)
	if i := attr["exp"]; i != nil {
		t := time.Unix(int64(i.(float64)), 0)
		cf[tfExpirationDate] = t.Format(time.RFC3339)
	}
	return cf
}
