package cmd

import (
	"fmt"
	"os"

	"github.com/eankeen/dotty/internal/util"
	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
)

// RootCmd is the root command
var RootCmd = &cobra.Command{
	Use:   "dotty",
	Short: "Dotfile CM Utility",
	Long:  "A CM (Configuration Management) utility for dotfiles. Used for managing local, user, or system-wide dotfiles",
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	unix.Umask(664)

	pf := RootCmd.PersistentFlags()
	pf.String("dot-dir", "", "The location of your dotfiles")

	err := cobra.MarkFlagRequired(pf, "dot-dir")
	util.P(err)

	err = cobra.MarkFlagDirname(pf, "dot-dir")
	util.P(err)
}
