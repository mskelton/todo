package today

import (
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/todo/internal/printer"
	"github.com/mskelton/todo/internal/storage"
	"github.com/spf13/cobra"
)

var TodayCmd = &cobra.Command{
	Use:   "today",
	Short: "List tasks due today",
	Long: heredoc.Doc(`
    Bytes are stored as markdown using a unique id representing the date and time
    the byte was created. This command will generate a new id and print it to
    stdout. This is typically not needed as the id is automatically generated
    when creating a new byte.
  `),
	RunE: func(cmd *cobra.Command, args []string) error {
		today := time.Now().Format("2006-01-02")
		filters := []storage.Filter{
			{
				Key:      "due",
				Operator: storage.Neq,
				Value:    "",
			},
			{
				Key:      "date(json_extract(due, '$.date'))",
				Operator: storage.Lte,
				Value:    fmt.Sprintf("date('%s')", today),
				IsRaw:    true,
			},
		}

		tasks, err := storage.ListTasks(filters)
		if err != nil {
			return err
		}

		printer.PrintTasks(tasks)

		return nil
	},
}
