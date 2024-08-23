package storage

import (
	"gorm.io/gorm"
)

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

func SaveMapping(storageType StorageType, idMap map[int]string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("type = ?", storageType).Delete(&IdMapping{}).Error; err != nil {
			return err
		}

		mappings := make([]IdMapping, 0, len(idMap))
		for shortId, realId := range idMap {
			mapping := IdMapping{
				ShortId: shortId,
				RealId:  realId,
				Type:    storageType,
			}

			mappings = append(mappings, mapping)
		}

		return tx.Create(&mappings).Error
	})
}
