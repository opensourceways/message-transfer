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

// Difference 函数返回 sliceA 中存在但 sliceB 中不存在的元素
func Difference(sliceA, sliceB []string) []string {
	// 创建一个映射，用于存储 sliceB 中的元素
	setB := make(map[string]struct{})
	for _, item := range sliceB {
		setB[item] = struct{}{}
	}

	// 创建一个结果切片，用于存储差集
	var difference []string
	for _, item := range sliceA {
		if _, found := setB[item]; !found {
			difference = append(difference, item)
		}
	}

	return difference
}
