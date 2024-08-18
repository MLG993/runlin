package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getWinePrefix() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to determine home directory:", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, ".wine")
}

func listExeFiles(winePrefix string) []string {
	var exeFiles []string

	err := filepath.Walk(winePrefix, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(strings.ToLower(info.Name()), ".exe") {
			exeFiles = append(exeFiles, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error while searching for .exe files:", err)
		os.Exit(1)
	}

	return exeFiles
}

func grepExeFiles(exeFiles []string, searchTerm string) []string {
	var filteredFiles []string

	for _, exe := range exeFiles {
		if strings.Contains(strings.ToLower(filepath.Base(exe)), strings.ToLower(searchTerm)) {
			filteredFiles = append(filteredFiles, exe)
		}
	}

	return filteredFiles
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run finder.go [list | grep <search-term>]")
		return
	}

	command := os.Args[1]
	winePrefix := getWinePrefix()

	switch strings.ToLower(command) {
	case "list":
		exeFiles := listExeFiles(winePrefix)
		fmt.Println("Found .exe files:")
		for _, exe := range exeFiles {
			fmt.Println("- " + exe)
		}

	case "grep":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run finder.go grep <search-term>")
			return
		}
		searchTerm := os.Args[2]
		exeFiles := listExeFiles(winePrefix)
		filteredFiles := grepExeFiles(exeFiles, searchTerm)
		if len(filteredFiles) == 0 {
			fmt.Println("No matching .exe files found.")
		} else {
			fmt.Println("Matching .exe files:")
			for _, exe := range filteredFiles {
				fmt.Println("- " + exe)
			}
		}
	default:
		fmt.Println("Invalid command. Usage: go run finder.go [list | grep <search-term>]")
	}
}

