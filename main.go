package main

import (
	"database/sql"
	"os"
	"time"

	"github.com/fatih/color"
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
		color.Red("Error:", err)
		os.Exit(1)
	}
	// Load previous config snapshot
	snap, err := gatorsave.Load(path)
	if err != nil {
		color.Yellow("Warning: could not read config:", err)
	}

	cfg := &appcfg.Config{}
	appcfg.ApplySnapshot(cfg, snap)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		color.Red("error connecting to db:", err)
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
		color.Yellow("Usage: gator <command> [args...]")
		color.Yellow("Available commands:", commands.GetCommandNames())
		os.Exit(1)
	}

	command := core.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(state, command); err != nil {
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}