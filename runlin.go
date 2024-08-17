package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isInstalled(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}

func installDependencies(deps []string) {
	fmt.Println("Checking and installing Dependencies...")

	for _, dep := range deps {
		if !isInstalled(dep) {
			fmt.Printf("Installing %s...\n", dep)
			installCmd := exec.Command("sudo", "yay", "-S", "--needed", "--noconfirm", dep)
			installCmd.Stdout = os.Stdout
			installCmd.Stderr = os.Stderr
			err := installCmd.Run()
			if err != nil {
				fmt.Printf("Error while installing %s: %v\n", dep, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("%s is already installed.\n", dep)
		}
	}
	fmt.Println("Every Dependencies are installed.")
}

func installPerformanceEnhancers() {
	fmt.Println("Checking optional Performance opportunities...")

	performanceDeps := []string{"dxvk-bin", "vkd3d-bin"} 
	for _, dep := range performanceDeps {
		if !isInstalled(dep) {
			fmt.Printf("%s is not installed? Do you want to install it? (y/n): ", dep)
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) == "j" {
				installDependencies([]string{dep})
			}
		} else {
			fmt.Printf("%s is already installed.\n", dep)
		}
	}
	fmt.Println("Every Perfomance increasing opportunity is installed.")
}

func runWithWine(appPath string) {
	fmt.Printf("Starting %s with Wine...\n", appPath)
	cmd := exec.Command("wine", appPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error while starting: %v\n", err)
		os.Exit(1)
	}
}

func runWithProton(appPath string) {
	fmt.Printf("Starting %s with Proton...\n", appPath)
	protonPath := "/usr/bin/proton" 

	cmd := exec.Command(protonPath, "run", appPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
    fmt.Printf("Error while starting: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run runlin.go [proton | wine | install | performance] <Path to .exe-file>")
		return
	}

	command := os.Args[1]
	appPath := filepath.Clean(os.Args[2])

	switch strings.ToLower(command) {
	case "install":
		installDependencies([]string{"wine", "proton-ge-custom-bin"})
	case "wine":
		if isInstalled("wine") {
			runWithWine(appPath)
		} else {
			fmt.Println("Wine is not installed pls type go run runlin.go install")
		}
	case "proton":
		if isInstalled("proton") {
			runWithProton(appPath)
		} else {
			fmt.Println("Proton is not installed pls type go run runlin.go install.")
		}
	case "performance":
		installPerformanceEnhancers()
		if isInstalled("proton") {
			runWithProton(appPath)
		} else if isInstalled("wine") {
			runWithWine(appPath)
		} else {
			fmt.Println("No supported Executor.")
		}
	default:
		fmt.Println("Invalid command. Usage: go run runlin.go [proton | wine | install | performance] <Path to .exe-file>")
	}
}

