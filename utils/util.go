/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package utils provides utility functions for various purposes.
package utils

import (
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

// LoadFromYaml reads a YAML file from the given path and unmarshals it into the provided interface.
func LoadFromYaml(path string, cfg interface{}) error {
	b, err := os.ReadFile(path) // #nosec G304
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, cfg)
}

// EscapePgsqlValue escape pg sql vale.
func EscapePgsqlValue(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, `%`, `\%`)
	value = strings.ReplaceAll(value, `_`, `\_`)
	value = strings.ReplaceAll(value, `'`, `\'`)
	value = strings.ReplaceAll(value, `"`, `\"`)
	value = strings.ReplaceAll(value, `[`, `\[`)
	value = strings.ReplaceAll(value, `]`, `\]`)
	value = strings.ReplaceAll(value, `^`, `\^`)
	return value
}
