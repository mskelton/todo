package project

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/fatih/color"
	"github.com/mskelton/todo/internal/printer"
	"github.com/mskelton/todo/internal/storage"
	"github.com/mskelton/todo/internal/utils"
	"github.com/spf13/cobra"
)

func listProjects() ([]storage.Project, error) {
	db, err := storage.GetDB()
	if err != nil {
		return nil, err
	}

	var projects []storage.Project
	tx := db.Where("is_archived == false and is_deleted == false").Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return projects, nil
}

func printProjects(projects []storage.Project) error {
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

		table.Rows = append(table.Rows, printer.Row{
			Cells: []string{
				// strconv.Itoa(int(project.ChildOrder)),
				project.ID,
				project.Name,
				favorite,
				utils.ShortDuration(project.CreatedAt, "-"),
			},
		})
	}

	table.Print()

	return nil
}

var ProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "List tasks due today",
	Long: heredoc.Doc(`
    Bytes are stored as markdown using a unique id representing the date and time
    the byte was created. This command will generate a new id and print it to
    stdout. This is typically not needed as the id is automatically generated
    when creating a new byte.
  `),
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := listProjects()
		if err != nil {
			return err
		}

		return printProjects(projects)
	},
}
