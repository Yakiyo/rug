package cmd

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rug",
	Short: "Dead simple script runner",
	Long: `Rug is an easy to use and simplified command runner.
	
No complicated setup like Make or Cmake. Rug is inspired by npm scripts. 
For details visit https://github.com/Yakiyo/rug`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	// Set up colored cobra
	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "rug version")
}
