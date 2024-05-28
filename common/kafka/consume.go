package kafka

type ConsumeConfig struct {
	Topic   string `json:"topic"  required:"true"`
	Publish string `json:"publish"  required:"true"`
	Group   string `json:"group"  required:"true"`
}
