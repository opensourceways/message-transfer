/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

package kafka

// ConsumeConfig kafka consume config.
type ConsumeConfig struct {
	Topic   string `json:"topic"  required:"true"`
	Publish string `json:"publish"  required:"true"`
	Group   string `json:"group"  required:"true"`
}
