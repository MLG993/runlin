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
			installCmd := exec.Command("yay", "-S", "--needed", "--noconfirm", dep)
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
	fmt.Println("All Dependencies are installed.")
}

func installPerformanceEnhancers() {
	fmt.Println("Checking optional Performance opportunities...")

	performanceDeps := []string{"dxvk-bin", "vkd3d-git"}
	for _, dep := range performanceDeps {
		if !isInstalled(dep) {
			fmt.Printf("%s is not installed? Do you want to install it? (y/n): ", dep)
			var response string
			fmt.Scanln(&response)
			if strings.ToLower(response) == "y" {
				installDependencies([]string{dep})
			}
		} else {
			fmt.Printf("%s is already installed.\n", dep)
		}
	}
	fmt.Println("All Performance increasing opportunities are installed.")
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

func runWinetricksCommand(command string) error {
	cmd := exec.Command("winetricks", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error while running winetricks command: %v", err)
	}
	return nil
}

func optimizeApp(appPath string) error {
	var winetricksCommands []string

	if strings.Contains(appPath, "d3d9") || strings.Contains(appPath, "dx9") || strings.Contains(appPath, "directx9") {
		winetricksCommands = append(winetricksCommands, "d3dx9")
	}
	if strings.Contains(appPath, "d3d10") || strings.Contains(appPath, "dx10") || strings.Contains(appPath, "directx10") {
		winetricksCommands = append(winetricksCommands, "d3dx10")
	}
	if strings.Contains(appPath, "d3d11") || strings.Contains(appPath, "dx11") || strings.Contains(appPath, "directx11") {
		winetricksCommands = append(winetricksCommands, "d3dx11")
	}
	if strings.Contains(appPath, "vcrun2019") || strings.Contains(appPath, "vcrun") {
		winetricksCommands = append(winetricksCommands, "vcrun2019")
	}
	if strings.Contains(appPath, ".net") || strings.Contains(appPath, "dotnet") {
		winetricksCommands = append(winetricksCommands, "dotnet48")
	}

	if len(winetricksCommands) == 0 {
		fmt.Println("No specific dependencies identified, installing default components.")
		winetricksCommands = append(winetricksCommands, "corefonts")
	}

	fmt.Println("Optimizing the application with the following Winetricks components:")
	for _, cmd := range winetricksCommands {
		fmt.Printf("- %s\n", cmd)
		err := runWinetricksCommand(cmd)
		if err != nil {
			return fmt.Errorf("Error installing %s: %v", cmd, err)
		}
	}

	fmt.Println("Optimization completed.")
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run runlin.go [proton | wine | install | compatability | optimize] <Path to .exe-file>")
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
			fmt.Println("Wine is not installed, please run: go run runlin.go install")
		}
	case "proton":
		if isInstalled("proton") {
			runWithProton(appPath)
		} else {
			fmt.Println("Proton is not installed, please run: go run runlin.go install")
		}
	case "compatability":
		installPerformanceEnhancers()
		if isInstalled("proton") {
			runWithProton(appPath)
		} else if isInstalled("wine") {
			runWithWine(appPath)
		} else {
			fmt.Println("No supported executor found.")
		}
	case "optimize":
		err := optimizeApp(appPath)
		if err != nil {
			fmt.Printf("Error during optimization: %v\n", err)
		}
	default:
		fmt.Println("Invalid command. Usage: go run runlin.go [proton | wine | install | compatability | optimize] <Path to .exe-file>")
	}
}

