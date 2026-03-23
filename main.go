package main

import (
	"github.com/opencode-ai/teamcode/cmd"
	"github.com/opencode-ai/teamcode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
