package cmd

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

const RESOURCE_DIR="res"
const DIRECTORY_PERMISSIONS=0755

func onInit(cmd *cobra.Command, args []string) {
    log.Println("Initializing environment")
    if cwd, err := os.Getwd(); err == nil {
        os.MkdirAll(path.Join(cwd, RESOURCE_DIR), DIRECTORY_PERMISSIONS)
        os.MkdirAll(path.Join(cwd, RESOURCE_DIR, "views"), DIRECTORY_PERMISSIONS)
        os.MkdirAll(path.Join(cwd, RESOURCE_DIR, "static"), DIRECTORY_PERMISSIONS)
        log.Println("Environment is initialized")
        log.Println("Execute the migration command next!")
    } else {
        log.Fatal(err)
    }
}

var initCmd = &cobra.Command {
    Use: "init",
    Short: "Init application environment",
    Run: onInit,
}
