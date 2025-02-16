package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/modestprophet/pirowflo_dbstore/internal/config"
	"github.com/modestprophet/pirowflo_dbstore/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config: %v", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
		return
	}
	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("start", startDataStorage)

	if len(os.Args) < 2 {
		fmt.Println("Error: command required")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := []string{}
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	err = cmds.run(s, command{
		name: cmdName,
		args: cmdArgs,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
