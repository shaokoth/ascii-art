package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/asciiArt"
)

func main() {
	fileName := "standard.txt"

	// Check if an argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string as an argument.")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Many arguments passed")
		return
	}
	if os.Args[1] == "" {
		return
	}
	// print a new line and exit in case argument is a new line character only
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	// Load the banner map from the file
	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}

	// Process the provided argument
	args := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	args = strings.ReplaceAll(args, "\\t", "    ")
	lines := strings.Split(args, "\n")

	// Generate the ASCII art for each line
	for _, line := range lines {
		asciiArt.PrintLineBanner(line, bannerMap)
	}
}
