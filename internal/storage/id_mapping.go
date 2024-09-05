package storage

import (
	"github.com/mskelton/todo/internal/models"
	"gorm.io/gorm"
)

func SaveIdMapping(storageType models.StorageType, idMap map[int]string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("type = ?", storageType).Delete(&models.IdMapping{}).Error; err != nil {
			return err
		}

		mappings := make([]models.IdMapping, 0, len(idMap))
		for shortId, realId := range idMap {
			mapping := models.IdMapping{
				ShortId: shortId,
				RealId:  realId,
				Type:    storageType,
			}

			mappings = append(mappings, mapping)
		}

		return tx.Create(&mappings).Error
	})
}
