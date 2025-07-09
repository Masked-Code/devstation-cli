package installer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PackageManager interface for different installation methods
type PackageManager interface {
	Install(packageName string) error
	IsInstalled(packageName string) bool
	Update(packageName string) error
}

// ChocoManager implements PackageManager for Chocolatey
type ChocoManager struct{}

func (c *ChocoManager) Install(packageName string) error {
	fmt.Printf("Installing %s via Chocolatey...\n", packageName)
	cmd := exec.Command("choco", "install", packageName, "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *ChocoManager) IsInstalled(packageName string) bool {
	cmd := exec.Command("choco", "list", "--local-only", packageName)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), packageName)
}

func (c *ChocoManager) Update(packageName string) error {
	fmt.Printf("Updating %s via Chocolatey...\n", packageName)
	cmd := exec.Command("choco", "upgrade", packageName, "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// WingetManager implements PackageManager for Windows Package Manager
type WingetManager struct{}

func (w *WingetManager) Install(packageName string) error {
	fmt.Printf("Installing %s via winget...\n", packageName)
	cmd := exec.Command("winget", "install", packageName, "--accept-package-agreements", "--accept-source-agreements")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (w *WingetManager) IsInstalled(packageName string) bool {
	cmd := exec.Command("winget", "list", packageName)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), packageName)
}

func (w *WingetManager) Update(packageName string) error {
	fmt.Printf("Updating %s via winget...\n", packageName)
	cmd := exec.Command("winget", "upgrade", packageName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetAvailablePackageManager returns the first available package manager
func GetAvailablePackageManager() PackageManager {
	// Check for winget first (newer, built-in)
	if isCommandAvailable("winget") {
		return &WingetManager{}
	}
	
	// Fallback to chocolatey
	if isCommandAvailable("choco") {
		return &ChocoManager{}
	}
	
	return nil
}

func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// InstallPackageManager installs Chocolatey if no package manager is available
func InstallPackageManager() error {
	if GetAvailablePackageManager() != nil {
		return nil
	}
	
	fmt.Println("No package manager found. Installing Chocolatey...")
	cmd := exec.Command("powershell", "-Command", 
		"Set-ExecutionPolicy Bypass -Scope Process -Force; "+
		"[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; "+
		"iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
