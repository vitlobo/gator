package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/vitlobo/gator/cmd"

	"github.com/vitlobo/gator/internal/appcfg"
	"github.com/vitlobo/gator/internal/core"
	"github.com/vitlobo/gator/internal/database"
	"github.com/vitlobo/gator/internal/gatorapi"
	"github.com/vitlobo/gator/internal/gatorsave"
)

func main() {
	path, err := gatorsave.DefaultPath()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Load previous config snapshot
	snap, err := gatorsave.Load(path)
	if err != nil {
		fmt.Println("Warning: could not read config:", err)
	}

	cfg := &appcfg.Config{}
	appcfg.ApplySnapshot(cfg, snap)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("error connecting to db:", err)
		os.Exit(1)
	}
	defer db.Close()
	dbQueries := database.New(db)

	gatorClient := gatorapi.NewClient(10*time.Second)

	state := &core.State{
		Cfg: cfg,
		Db: dbQueries,
		GatorClient: &gatorClient,
	}

	commands := core.GetRegisteredCommands()

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		fmt.Println("Available commands:", commands.GetCommandNames())
		os.Exit(1)
	}

	command := core.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(state, command); err != nil {
		fmt.Println("Command failed:", err)
		os.Exit(1)
	}
}