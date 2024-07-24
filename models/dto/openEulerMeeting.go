package dto

type OpenEulerMeetingRaw struct {
	Action string `json:"action"`
	Msg    struct {
		Topic     string      `json:"topic"`
		Platform  interface{} `json:"platform"`
		Sponsor   int         `json:"sponsor"`
		GroupName string      `json:"group_name"`
		GroupId   interface{} `json:"group_id"`
		Date      int         `json:"date"`
		Start     string      `json:"start"`
		End       string      `json:"end"`
		Etherpad  int         `json:"etherpad"`
		Agenda    string      `json:"agenda"`
		EmailList string      `json:"emailist"`
		Record    string      `json:"record"`
		JoinUrl   string      `json:"join_url"`
	} `json:"msg"`
}
