package cli

import (
	"github.com/spf13/cobra"
)

// SearchByLink command
var RandomJoke = &cobra.Command{
	Use:   "Random",
	Short: "Random a dad joke",
	Long:  `Using a public API to random a dad joke`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		RandowJoke()
	},
}

// SearchByFile Command
var SearchByFile = &cobra.Command{
	Use:   "File",
	Short: "Search for the anime scene by existing image file",
	Long:  `CLI-anime file <PATH_TO_IMAGE>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SearchByImageFile(args[0])
	},
}

// SearchByLink command
var SearchByLink = &cobra.Command{
	Use:   "Link",
	Short: "Search for the anime scene by existing image url",
	Long:  `CLI-anime Link <IMAGE_URL>`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		SearchByImageLink(args[0])
	},
}

// Checking IP command
var IPAdress = &cobra.Command{
	Use:   "Ip",
	Short: "Show your network IP",
	Long:  `Lookup the IP address for a particular host`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		getIPAdress()
	},
}

var MembersProfile = &cobra.Command{
	Use:   "Group",
	Short: "Show members profile of KPOP groups",
	Long:  `Lookup the members profile by group name`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Idols(args[0])
	},
}

var DetailProfile = &cobra.Command{
	Use:   "Detail",
	Short: "Show detail member profile of KPOP groups",
	Long:  `Lookup the detail member profile by group and idol name`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		DetailIdol(args[0] , args[1])
	},
}

// AddCommands launches all commands
func AddCommands() {
	RootCmd.AddCommand(SearchByFile)
	RootCmd.AddCommand(SearchByLink)
	RootCmd.AddCommand(RandomJoke)
	RootCmd.AddCommand(IPAdress)
	RootCmd.AddCommand(MembersProfile)
	RootCmd.AddCommand(DetailProfile)
}