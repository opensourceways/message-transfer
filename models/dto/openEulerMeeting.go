package dto

type OpenEulerMeetingRaw struct {
	Action string `json:"action"`
	Msg    struct {
		Id        int    `json:"id"`
		Topic     string `json:"topic"`
		Community string `json:"community"`
		GroupName string `json:"group_name"`
		Sponsor   string `json:"sponsor"`
		Date      string `json:"date"`
		Start     string `json:"start"`
		End       string `json:"end"`
		Duration  string `json:"duration"`
		Agenda    string `json:"agenda"`
		Etherpad  string `json:"etherpad"`
		Emaillist string `json:"emaillist"`
		HostId    string `json:"host_id"`
		Mid       string `json:"mid"`
		Mmid      string `json:"mmid"`
		JoinUrl   string `json:"join_url"`
		IsDelete  int    `json:"is_delete"`
		StartUrl  string `json:"start_url"`
		Timezone  string `json:"timezone"`
		User      int    `json:"user"`
		Group     int    `json:"group"`
		Mplatform string `json:"mplatform"`
	} `json:"msg"`
}
