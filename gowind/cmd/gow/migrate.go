package main

import (
	"github.com/Hzn8Hub/teamtalk-tools/gowind/internal/migrate"
)

func init() {
	rootCmd.AddCommand(migrate.CmdMigrate)
}
