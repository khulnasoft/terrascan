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
	"strings"

	"github.com/khulnasoft/terrascan/pkg/mapper/convert"
	fn "github.com/khulnasoft/terrascan/pkg/mapper/iac-providers/arm/functions"
	"github.com/khulnasoft/terrascan/pkg/mapper/iac-providers/arm/types"
)

const publicAccess = "publicAccess"
const tfContainerAccessType = "container_access_type"

// StorageContainerConfig returns config for azurerm_storage_container
func StorageContainerConfig(r types.Resource, params map[string]interface{}) map[string]interface{} {
	cf := map[string]interface{}{
		tfLocation: fn.LookUpString(nil, params, r.Location),
		tfName:     fn.LookUpString(nil, params, r.Name),
		tfTags:     r.Tags,
	}

	access := fn.LookUpString(nil, params, convert.ToString(r.Properties, publicAccess))
	if strings.ToUpper(access) == "NONE" {
		access = "private"
	}
	cf[tfContainerAccessType] = access
	return cf
}
