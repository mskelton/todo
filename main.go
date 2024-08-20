package main

import (
	"github.com/mskelton/todo/internal/storage"
	"github.com/mskelton/todo/pkg/cmd/root"
)

func main() {
	// TODO: Improve syncing
	go storage.Sync("*")

	root.Execute()
}
