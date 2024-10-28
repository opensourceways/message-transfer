/*
Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved
*/

// Package dto models dto of eur build.
package dto

// EurBuildMessageRaw eur build message raw.
type EurBuildMessageRaw struct {
	Body struct {
		User    string      `json:"user"`
		Copr    string      `json:"copr"`
		Owner   string      `json:"owner"`
		Pkg     interface{} `json:"pkg"`
		Build   int         `json:"build"`
		Chroot  string      `json:"chroot"`
		Version interface{} `json:"version"`
		Status  int         `json:"status"`
		IP      string      `json:"ip"`
		Who     string      `json:"who"`
		Pid     int         `json:"pid"`
		What    string      `json:"what"`
	} `json:"body"`
	Headers struct {
		OpenEulerMessagingSchema string `json:"openEuler_messaging_schema"`
		SentAt                   string `json:"sent-at"`
	} `json:"headers"`
	ID          string `json:"id"`
	Topic       string `json:"topic"`
	SourceGroup string `json:"source_group"`
}

func (raw *EurBuildMessageRaw) ToCloudEventsByConfig() CloudEvents {
	rawMap := StructToMap(raw)
	return rawMap.ToCloudEventByConfig("gitee_issue_raw")
}

func (raw *EurBuildMessageRaw) GetRelateUsers(events CloudEvents) {
	releatedUsers := []string{}
	//TODO implement me

	//releatedUsers,_=utils.GetRepoSigInfo(raw.Repository.Name)
	events.SetExtension("releatedusers", releatedUsers)
}

func (raw *EurBuildMessageRaw) GetFollowUsers(events CloudEvents) {
	followUsers := []string{}
	//todo
	events.SetExtension("followUsers", followUsers)
}

func (raw *EurBuildMessageRaw) GetTodoUsers(events CloudEvents) {
	todoUsers := []string{}
	//todo
	events.SetExtension("todoUsers", todoUsers)
}
