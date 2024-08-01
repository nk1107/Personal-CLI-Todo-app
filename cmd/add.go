package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Adding \"%s\" to your task list\n", task)
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
