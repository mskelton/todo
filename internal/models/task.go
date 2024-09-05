package models

import (
	"gorm.io/datatypes"
)

type DueDate struct {
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Lang        string `json:"lang"`
	String      string `json:"string"`
	Timezone    string `json:"timezone"`
}

type Task struct {
	AddedAt        string                       `json:"added_at"`
	AddedByUID     string                       `json:"added_by_uid"`
	AssignedByUID  *string                      `json:"assigned_by_uid"`
	Checked        bool                         `json:"checked"`
	ChildOrder     int32                        `json:"child_order"`
	Collapsed      bool                         `json:"collapsed"`
	CompletedAt    *string                      `json:"completed_at"`
	Content        string                       `json:"content"`
	DayOrder       int                          `json:"day_order"`
	Description    string                       `json:"description"`
	Due            *datatypes.JSONType[DueDate] `json:"due"`
	Duration       *string                      `json:"duration"`
	ID             string                       `json:"id"`
	IsDeleted      bool                         `json:"is_deleted"`
	Labels         datatypes.JSONSlice[string]  `json:"labels"`
	ParentID       *string                      `json:"parent_id"`
	Priority       int                          `json:"priority"`
	ProjectID      string                       `json:"project_id"`
	ResponsibleUID *string                      `json:"responsible_uid"`
	SectionID      *string                      `json:"section_id"`
	SyncID         *string                      `json:"sync_id"`
	UpdatedAt      string                       `json:"updated_at"`
	UserID         string                       `json:"user_id"`
	V2ID           string                       `json:"v2_id"`
	V2ParentID     *string                      `json:"v2_parent_id"`
	V2ProjectID    string                       `json:"v2_project_id"`
	V2SectionID    *string                      `json:"v2_section_id"`
}
