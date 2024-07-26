package storage

type Project struct {
	CanAssignTasks bool    `json:"can_assign_tasks"`
	ChildOrder     uint32  `json:"child_order"`
	Collapsed      bool    `json:"collapsed"`
	Color          string  `json:"color"`
	CreatedAt      string  `json:"created_at"`
	ID             string  `json:"id" gorm:"primaryKey"`
	IsArchived     bool    `json:"is_archived"`
	IsDeleted      bool    `json:"is_deleted"`
	IsFavorite     bool    `json:"is_favorite"`
	Name           string  `json:"name"`
	ParentID       *string `json:"parent_id"`
	Shared         bool    `json:"shared"`
	SyncID         *string `json:"sync_id"`
	UpdatedAt      string  `json:"updated_at"`
	V2ID           string  `json:"v2_id"`
	V2ParentID     *string `json:"v2_parent_id"`
	ViewStyle      string  `json:"view_style"`
}
