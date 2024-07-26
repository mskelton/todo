package cmd

import (
	"time"

	"github.com/fatih/color"
	"github.com/mskelton/todo/internal/arg_parser"
	"github.com/mskelton/todo/internal/printer"
	"github.com/mskelton/todo/internal/storage"
	"github.com/mskelton/todo/internal/utils"
)

func ListProjects(ctx arg_parser.ParseContext) error {
	db, err := storage.GetDB()
	if err != nil {
		return err
	}

	var projects []storage.Project
	tx := db.Where("is_archived == false and is_deleted == false").Find(&projects)
	if tx.Error != nil {
		return tx.Error
	}

	if len(projects) == 0 {
		printer.Message("No projects match filters")
		return nil
	}

	table := printer.Table{
		Columns: []string{"ID", "Name", "Favorite", "Created"},
		Rows:    []printer.Row{},
	}

	for _, project := range projects {
		var favorite string
		if project.IsFavorite && color.NoColor {
			favorite = "✔︎"
		}

		createdAt, err := time.Parse("2006-01-02T15:04:05Z", project.CreatedAt)
		if err != nil {
			return err
		}

		table.Rows = append(table.Rows, printer.Row{
			Cells: []string{
				// strconv.Itoa(int(project.ChildOrder)),
				project.ID,
				project.Name,
				favorite,
				utils.ShortDuration(createdAt),
			},
		})
	}

	table.Print()

	return nil
}
