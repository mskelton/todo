package cmd

import (
	"github.com/mskelton/todo/internal/arg_parser"
	"github.com/mskelton/todo/internal/printer"
	"github.com/mskelton/todo/internal/storage"
)

func Today(ctx arg_parser.ParseContext) error {
	tasks, err := storage.ListTasks()
	if err != nil {
		return err
	}

	printer.PrintTasks(tasks)

	return nil
}
