package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/rug/parser"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rug [args]",
	Short: shortDesc,
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// Check for flags first

		if lo.Must(cmd.Flags().GetBool("version")) {
			fmt.Printf("Rug v%v\n", Version)
			return
		}

		if lo.Must(cmd.Flags().GetBool("list")) {
			// TODO
			return
		}

		if lo.Must(cmd.Flags().GetBool("init")) {
			file, err := cmd.Flags().GetString("config")
			// this should usually never happen
			if err != nil {
				// panic in the off-case it does
				panic(err)
			}
			if fileExists(file) {
				fmt.Fprintf(os.Stderr, "%v %v\n", errTag, "Rug File already exists")
				os.Exit(1)
			}
			if err := parser.WriteDefault(file); err != nil {
				fmt.Fprintf(os.Stderr, "%v %v\n", errTag, err)
				os.Exit(1)
			}
			fmt.Println("Initialized rug file in", file)
			return
		}

		// Now check for scripts

		// If theres no args, show help and exit
		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}
		rugPath, _ := cmd.Flags().GetString("config")
		rug, err := parser.ReadRugFile(rugPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v %v\n", errTag, err)
			os.Exit(1)
		}
		script := args[0]
		passThrough := args[1:]

		if !rug.HasScript(script) {
			fmt.Fprintf(os.Stderr, "%v Missing script by the name `%v`.\nRun `rug --list` to see a list of all commands\n", errTag, script)
			os.Exit(1)
		}

		exitCode, err := Run(rug, script, passThrough)

		if exitCode != 0 || err != nil {
			fmt.Fprintf(os.Stderr, "%v Process exited with status code %v. (%v)\n", errTag, exitCode, err)
			os.Exit(1)
		}
	},
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

	if len(os.Args[1:]) < 1 {
		rootCmd.Help()
		os.Exit(1)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "rug version")
	rootCmd.Flags().BoolP("list", "l", false, "list available scripts")
	rootCmd.Flags().BoolP("init", "i", false, "init rug file")

	rootCmd.Flags().StringP("config", "c", "rug.json", "path to rug file")
}
