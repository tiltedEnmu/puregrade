package entities

type MQRecord struct {
	UserIds []string `json:"userIds,omitempty"`
	PostIds []string `json:"postIds,omitempty"`
}
