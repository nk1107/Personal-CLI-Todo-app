package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("All tasks completed !!")
			return
		}
		fmt.Println()
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
		fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
