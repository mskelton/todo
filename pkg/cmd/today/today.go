package today

import (
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
		tasks, err := storage.ListTasks()
		if err != nil {
			return err
		}

		printer.PrintTasks(tasks)

		return nil
	},
}
