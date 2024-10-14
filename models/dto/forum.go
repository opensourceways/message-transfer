package dto

import "time"

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
	UserName     int              `json:"user_name"`
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
}
