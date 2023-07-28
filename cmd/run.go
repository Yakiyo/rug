package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/Yakiyo/rug/parser"
)

func Run(rug *parser.RugFile, script string, args []string) (int, error) {
	rugScript := rug.GetScript(script)

	if rug.Hooks && rug.HasScript("pre"+script) {
		pre := command(rug, rug.GetScript("pre"+script))
		err := pre.Run()
		if err != nil {
			return handleErr(err)
		}
	}

	fullCommand := strings.Join([]string{rugScript, strings.Join(args, " ")}, " ")

	c := command(rug, fullCommand)

	err := c.Run()

	if err != nil {
		return handleErr(err)
	}

	if rug.Hooks && rug.HasScript("post"+script) {
		post := command(rug, rug.GetScript("post"+script))
		err := post.Run()
		if err != nil {
			return handleErr(err)
		}
	}
	return 0, nil
}

func command(rug *parser.RugFile, args ...string) *exec.Cmd {
	c := exec.Command("sh", "-c", strings.Join(args, " "))
	c.Env = append(os.Environ(), rug.GetEnv()...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c
}

func handleErr(err error) (int, error) {
	if err == nil {
		return 0, nil
	}
	if exiterr, ok := err.(*exec.ExitError); ok {
		return exiterr.ProcessState.ExitCode(), err
	}
	return 1, err
}
