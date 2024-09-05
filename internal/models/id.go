package models

type StorageType string

const (
	StorageTypeProject StorageType = "project"
	StorageTypeTask    StorageType = "task"
)

type IdMapping struct {
	Type    StorageType
	ShortId int
	RealId  string
}
