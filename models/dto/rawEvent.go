package dto

type RawEvent interface {
	GetRelateUsers(CloudEvents)
	GetTodoUsers(CloudEvents)
	GetFollowUsers(CloudEvents)
	GetApplyUsers(CloudEvents)
	ToCloudEventsByConfig(topic string) CloudEvents
	IsDone(events CloudEvents)
}
