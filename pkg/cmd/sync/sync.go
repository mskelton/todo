package sync

import (
	"github.com/mskelton/todo/internal/storage"
	"github.com/spf13/cobra"
)

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync with Todoist",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		db.Where("true").Delete(storage.Project{})
		db.Where("true").Delete(storage.Task{})

		db.Create(&res.Projects)
		db.Create(&res.Tasks)

		return nil
	},
}
