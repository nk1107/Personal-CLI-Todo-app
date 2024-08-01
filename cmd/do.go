package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid id: %s\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		err := db.DeleteTasks(ids)
		if err != nil {
			fmt.Println("something went wrong", err)
			return
		}
		fmt.Println("tasks deleted")

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
