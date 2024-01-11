/*
Copyright © 2024 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package openapiv3utils

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func BuildOpenAPIV3Template(openAPIV3Schema map[string]interface{}) (templateValue interface{}, err error) {
	objType, typeExist := openAPIV3Schema[string(TypeKey)]

	if !typeExist {
		return nil, errors.New("Type doesn't exist")
	}

	switch objType.(string) {
	case string(ObjectType):
		templateValue = map[string]interface{}{}

		if objSchema, ok := openAPIV3Schema[string(PropertiesKey)]; ok {
			for k, v := range objSchema.(map[string]interface{}) {
				templateValue.(map[string]interface{})[k], err = BuildOpenAPIV3Template(v.(map[string]interface{}))

				if err != nil {
					err = errors.Wrapf(err, "Error in key '%s'", k)

					return nil, err
				}
			}
		} else {
			templateValue.(map[string]interface{})["custom_key"], err = BuildOpenAPIV3Template(openAPIV3Schema[string(AdditionalPropertiesKey)].(map[string]interface{}))

			if err != nil {
				err = errors.Wrapf(err, "Error in custom key")

				return nil, err
			}
		}
	case string(ArrayType):
		var templateItems interface{}

		templateValue = make([]interface{}, 0)

		templateItems, err = BuildOpenAPIV3Template(openAPIV3Schema[string(ItemsKey)].(map[string]interface{}))

		if err != nil {
			err = errors.Wrapf(err, "Error in array definition")

			return nil, err
		}

		templateValue = append(templateValue.([]interface{}), templateItems)
	case string(StringType):
		templateValue = "String"

		if regexPattern, ok := openAPIV3Schema[string(PatternKey)]; ok {
			templateValue = fmt.Sprintf("%s (regex: %s)", templateValue, regexPattern)
		}

		if minLen, ok := openAPIV3Schema[string(MinLengthKey)]; ok {
			templateValue = fmt.Sprintf("%s (minLen: %v)", templateValue, minLen)
		}

		if maxLen, ok := openAPIV3Schema[string(MaxLengthKey)]; ok {
			templateValue = fmt.Sprintf("%s (maxLen: %v)", templateValue, maxLen)
		}
	case string(BooleanType):
		templateValue = false
	case string(IntegerType):
		templateValue = 1

		if minValue, ok := openAPIV3Schema[string(MinimumKey)]; ok {
			templateValue = minValue
		} else if maxValue, ok := openAPIV3Schema[string(MaximumKey)]; ok {
			templateValue = maxValue
		}
	case string(NumberType):
		templateValue = 0.5
		templateValue = fmt.Sprintf("%.2f", templateValue)
		templateValue, _ = strconv.ParseFloat(templateValue.(string), 64)
	}

	return templateValue, err
}
