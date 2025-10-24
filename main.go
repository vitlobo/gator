package main

import (
	"fmt"
	"log"

	"github.com/vitlobo/gator/internal/appcfg"
	"github.com/vitlobo/gator/internal/gatorsave"
)

func main() {
	path, err := gatorsave.DefaultPath()
	if err != nil { log.Fatal(err) }

	// Load previous config snapshot
	snap, err := gatorsave.Load(path)
	if err != nil { log.Fatalf("error reading config: %v", err) }
	fmt.Printf("Read config      : %+v\n", snap)

	cfg := &appcfg.Config{}
	appcfg.ApplySnapshot(cfg, snap)

	// Set current user and save snapshot
	err = cfg.SetUser("aawhite")
	if err != nil { log.Fatal(err) }

	// Load config again, apply, verify user is correct if changed from previous config
	snap, err = gatorsave.Load(path)
	if err != nil { log.Fatalf("error reading config: %v", err) }
	appcfg.ApplySnapshot(cfg, snap)

	fmt.Printf("Read config again: %+v\n", snap)
}