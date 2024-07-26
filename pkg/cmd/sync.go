package cmd

import (
	"github.com/mskelton/todo/internal/arg_parser"
	"github.com/mskelton/todo/internal/storage"
)

func Sync(ctx arg_parser.ParseContext) error {
	res, err := storage.Sync("*")
	if err != nil {
		return err
	}

	err = storage.Migrate()
	if err != nil {
		return err
	}

	db, err := storage.GetDB()
	if err != nil {
		return err
	}

	tx := db.Create(&res.Projects)
	return tx.Error
}
