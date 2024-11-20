package dto

import (
	"slices"
	"time"
)

// Forum 是包含所有通知信息的顶级结构体
type Forum struct {
	Notifications         []Notification `json:"notifications"`
	TotalRows             int            `json:"total_rows_notifications"`
	SeenNotificationID    int            `json:"seen_notification_id"`
	LoadMoreNotifications string         `json:"load_more_notifications"`
}

// Notification 表示单个通知信息
type Notification struct {
	ID           int              `json:"id"`
	UserID       int              `json:"user_id"`
	UserName     string           `json:"user_name"`
	Type         int              `json:"notification_type"`
	Read         bool             `json:"read"`
	HighPriority bool             `json:"high_priority"`
	CreatedAt    time.Time        `json:"created_at"`
	PostNumber   *int             `json:"post_number,omitempty"`
	TopicID      *int             `json:"topic_id,omitempty"`
	Slug         *string          `json:"slug,omitempty"`
	FancyTitle   *string          `json:"fancy_title,omitempty"`
	Data         NotificationData `json:"data"`
}

// NotificationData 包含通知的具体数据
type NotificationData struct {
	BadgeID          *int    `json:"badge_id,omitempty"`
	BadgeName        *string `json:"badge_name,omitempty"`
	BadgeSlug        *string `json:"badge_slug,omitempty"`
	BadgeTitle       *bool   `json:"badge_title,omitempty"`
	Username         *string `json:"username,omitempty"`
	TopicTitle       *string `json:"topic_title,omitempty"`
	OriginalPostID   *int    `json:"original_post_id,omitempty"`
	OriginalPostType *int    `json:"original_post_type,omitempty"`
	OriginalUsername *string `json:"original_username,omitempty"`
	RevisionNumber   *int    `json:"revision_number,omitempty"`
	DisplayUsername  *string `json:"display_username,omitempty"`
	GroupID          *int    `json:"group_id,omitempty"`
	GroupName        *string `json:"group_name,omitempty"`
	InBoxCount       *int    `json:"inbox_count,omitempty"`
}

func (raw *Notification) isSystemMessage() bool {
	return slices.Contains([]int{12, 24, 37}, raw.Type)
}

func (raw *Notification) isSystemAboutMessage() bool {
	return raw.Type == 6 && *raw.Data.OriginalUsername == "system"
}

func (raw *Notification) GetRelateUsers(events CloudEvents) {
	if raw.isSystemMessage() {
		return
	}
	if raw.isSystemAboutMessage() {
		return
	}

	events.SetExtension("releatedusers", []string{raw.UserName})
}

func (raw *Notification) GetTodoUsers(events CloudEvents) {
	events.SetExtension("todoUsers", []string{})
}

func (raw *Notification) GetFollowUsers(events CloudEvents) {
	if raw.isSystemMessage() || raw.isSystemAboutMessage() {
		events.SetExtension("followusers", []string{raw.UserName})
	}
}

func (raw *Notification) ToCloudEventsByConfig() CloudEvents {
	rawMap := StructToMap(raw)
	return rawMap.ToCloudEventByConfig("forum_raw")
}

func (raw *Notification) IsDone(events CloudEvents) {
	events.SetExtension("isdone", false)
}
