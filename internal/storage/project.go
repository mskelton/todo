package storage

import "github.com/mskelton/todo/internal/models"

func ListProjects() ([]models.Project, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}

	var projects []models.Project
	tx := db.Where("is_archived == false and is_deleted == false").Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return projects, nil
}
