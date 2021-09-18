package cli

import (
	"github.com/spf13/cobra"
)

// RootCmd main cobra command
var RootCmd = &cobra.Command{
	Use:   "CLI",
	Short: "Mine CLI",
	Long:  `A small command-line interface processes commands to a computer program in the form of lines of text`,
}