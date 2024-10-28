package dto

type RawEvent interface {
	GetRelateUsers(CloudEvents)
	GetTodoUsers(CloudEvents)
	GetFollowUsers(CloudEvents)
	ToCloudEventsByConfig() CloudEvents
	//FUNC()
}
