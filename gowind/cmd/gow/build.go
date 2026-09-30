package main

import (
	"github.com/Hzn8Hub/teamtalk-tools/gowind/internal/build"
)

func init() {
	rootCmd.AddCommand(build.CmdBuild)
}
