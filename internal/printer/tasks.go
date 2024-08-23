package printer

import (
	"strconv"
	"strings"

	"github.com/mskelton/todo/internal/storage"
	"github.com/mskelton/todo/internal/utils"
)

func trunc(s string, length int) string {
	if len(s) > length {
		return s[:length]
	}

	return s
}

func PrintTasks(tasks []storage.Task) error {
	table := Table{
		Columns: []string{
			"ID",
			"P",
			"Due",
			"Name",
			"Project",
			"Labels",
			"Age",
		},
		Rows: []Row{},
	}

	// TODO: Sections
	// TODO: Description

	for _, task := range tasks {
		due := ""
		if task.Due != nil {
			due = utils.ShortDuration(task.Due.Data().Date, "")
		}

		priority := ""
		if task.Priority > 1 {
			priority = strconv.Itoa(task.Priority)
		}

		table.Rows = append(table.Rows, Row{
			Cells: []string{
				task.ID,
				priority,
				due,
				trunc(task.Content, 50),
				task.ProjectID,
				strings.Join(task.Labels, ", "),
				// If the duration is less than 1 second, we just return "-". This is
				// primarily to make the tests more stable.
				utils.ShortDuration(task.AddedAt, "-"),
			},
		})
	}

	return table.Print(storage.StorageTypeTask)
}
