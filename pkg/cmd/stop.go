package cmd

import (
	"errors"
	"fmt"

	"github.com/mskelton/todo/internal/arg_parser"
	"github.com/mskelton/todo/internal/printer"
	"github.com/mskelton/todo/internal/storage"
	"github.com/mskelton/todo/internal/utils"
)

func Stop(ctx arg_parser.ParseContext) {
	requireFilters(ctx, "stop")

	filters := buildFilters(ctx)
	count, err := storage.Count(filters)
	if err != nil {
		printer.Error(err)
		return
	}

	if count == 0 {
		printer.Error(errors.New("No tasks match filters"))
		return
	}

	fmt.Printf(
		"This command will stop %d %s\n",
		count,
		utils.Pluralize(count, "task", "tasks"),
	)

	if utils.IsBulk(ctx, count) && !printer.Confirm("Are you sure you want to continue?") {
		return
	}

	edits := []storage.QueryEdit{{
		Path:  "status",
		Value: string(storage.TaskStatusPending),
	}}

	ids, err := storage.Edit(filters, edits)
	if err != nil {
		printer.Error(err)
		return
	}

	for _, id := range ids {
		fmt.Printf("Stoped task %d\n", id)
	}
}
