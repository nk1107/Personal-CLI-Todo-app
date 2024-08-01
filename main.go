package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"task/cmd"
	"task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")

	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
