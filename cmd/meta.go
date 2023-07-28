package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var errTag = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render("ERROR")

var urlStyle = lipgloss.NewStyle().Underline(true).Foreground(lipgloss.Color("#00FFFF"))

var longDesc = fmt.Sprintf(`Rug is an easy to use and simplified command runner.
	
No complicated setup like Make or CMake. Inspired by npm scripts. 
For details visit %v`, urlStyle.Render("https://github.com/Yakiyo/rug"))

var shortDesc = "Dead simple script runner"

var Version = "0.1.0"
