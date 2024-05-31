package main

import (
	_ "embed"
	"os"

	"github.com/AaronSaikovski/go-azuresdk-sample-may2024/pkg/utils"

	"github.com/AaronSaikovski/go-azuresdk-sample-may2024/cmd/app"
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

// main - program main
func main() {

	if err := app.Run(version); err != nil {
		utils.HandleError(err)
		os.Exit(1)
	}
}
