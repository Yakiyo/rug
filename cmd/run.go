package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Yakiyo/rug/parser"
)

func Run(rug parser.RugFile, script string, args []string) (int, error) {
	rugScript := rug.GetScript(script)

	fullCommand := strings.Join([]string{rugScript, strings.Join(args, " ")}, " ")

	c := exec.Command("sh", "-c", fullCommand)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Run()
	if err == nil {
		return 0, nil
	}
	if exiterr, ok := err.(*exec.ExitError); ok {
		return exiterr.ProcessState.ExitCode(), err
	}
	return 1, err
}
