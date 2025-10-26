package core

import (
	"fmt"

	"github.com/vitlobo/gator/internal/appcfg"
	"github.com/vitlobo/gator/internal/database"
)

type State struct {
	Cfg *appcfg.Config
	Db *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(*State, Command) error
}

var commandsInstance *Commands

func GetRegisteredCommands() *Commands {
	if commandsInstance == nil {
		commandsInstance = &Commands{
			RegisteredCommands: make(map[string]func(*State, Command) error),
		}
	}
	return commandsInstance
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.RegisteredCommands == nil {
		c.RegisteredCommands = make(map[string]func(*State, Command) error)
	}
	c.RegisteredCommands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf(
			"unknown command '%s'. Available commands: %v",
			cmd.Name,
			c.GetCommandNames(),
		)
	}
	return handler(s, cmd)
}

// GetCommandNames returns a slice of all registered command names
func (c *Commands) GetCommandNames() []string {
	names := make([]string, 0, len(c.RegisteredCommands))
	for name := range c.RegisteredCommands {
		names = append(names, name)
	}
	return names
}