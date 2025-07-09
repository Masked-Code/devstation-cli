package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"devstation-cli/pkg/installer"
	"devstation-cli/pkg/python"
	"devstation-cli/pkg/cdev"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up development environment",
	Long:  `Set up Python and C development environments with all necessary tools and dependencies.`,
}

// pythonCmd represents the python setup command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Set up Python development environment",
	Long:  `Install Python, pip, and essential packages for Python development.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := setupPythonEnvironment(); err != nil {
			fmt.Printf("Error setting up Python environment: %v\n", err)
			os.Exit(1)
		}
	},
}

// cCmd represents the C setup command
var cCmd = &cobra.Command{
	Use:   "c",
	Short: "Set up C development environment",
	Long:  `Install C compiler (MinGW or MSVC), build tools, and development utilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := setupCEnvironment(); err != nil {
			fmt.Printf("Error setting up C environment: %v\n", err)
			os.Exit(1)
		}
	},
}

// allCmd represents the command to set up both environments
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Set up both Python and C development environments",
	Long:  `Install and configure both Python and C development environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setting up complete development environment...")
		
		if err := setupPythonEnvironment(); err != nil {
			fmt.Printf("Error setting up Python environment: %v\n", err)
			os.Exit(1)
		}
		
		if err := setupCEnvironment(); err != nil {
			fmt.Printf("Error setting up C environment: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Println("🎉 Complete development environment setup finished!")
	},
}

// newCmd represents the new project command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project",
	Long:  `Create a new project structure for Python or C development.`,
}

// newPythonCmd creates a new Python project
var newPythonCmd = &cobra.Command{
	Use:   "python [project-name]",
	Short: "Create a new Python project",
	Long:  `Create a new Python project with proper structure, virtual environment, and configuration files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		pm := installer.GetAvailablePackageManager()
		if pm == nil {
			fmt.Println("No package manager available. Please install Chocolatey or winget first.")
			os.Exit(1)
		}
		
		pythonSetup := python.NewPythonSetup(pm)
		if err := pythonSetup.CreateProjectStructure(projectName); err != nil {
			fmt.Printf("Error creating Python project: %v\n", err)
			os.Exit(1)
		}
	},
}

// newCCmd creates a new C project
var newCCmd = &cobra.Command{
	Use:   "c [project-name]",
	Short: "Create a new C project",
	Long:  `Create a new C project with proper structure, build files, and configuration.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		pm := installer.GetAvailablePackageManager()
		if pm == nil {
			fmt.Println("No package manager available. Please install Chocolatey or winget first.")
			os.Exit(1)
		}
		
		cSetup := cdev.NewCDevSetup(pm)
		if err := cSetup.CreateCProject(projectName); err != nil {
			fmt.Printf("Error creating C project: %v\n", err)
			os.Exit(1)
		}
	},
}

// statusCmd shows the current development environment status
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check development environment status",
	Long:  `Check which development tools are installed and their versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkEnvironmentStatus()
	},
}

// setupPythonEnvironment sets up the Python development environment
func setupPythonEnvironment() error {
	fmt.Println("🐍 Setting up Python development environment...")
	
	// Ensure package manager is available
	if err := installer.InstallPackageManager(); err != nil {
		return fmt.Errorf("failed to install package manager: %v", err)
	}
	
	pm := installer.GetAvailablePackageManager()
	if pm == nil {
		return fmt.Errorf("no package manager available")
	}
	
	pythonSetup := python.NewPythonSetup(pm)
	
	// Install Python and essential packages
	if err := pythonSetup.InstallPython(); err != nil {
		return err
	}
	
	// Install development tools
	if err := pythonSetup.InstallPythonTools(); err != nil {
		return err
	}
	
	return nil
}

// setupCEnvironment sets up the C development environment
func setupCEnvironment() error {
	fmt.Println("⚙️  Setting up C development environment...")
	
	// Ensure package manager is available
	if err := installer.InstallPackageManager(); err != nil {
		return fmt.Errorf("failed to install package manager: %v", err)
	}
	
	pm := installer.GetAvailablePackageManager()
	if pm == nil {
		return fmt.Errorf("no package manager available")
	}
	
	cSetup := cdev.NewCDevSetup(pm)
	
	// Install C development tools
	if err := cSetup.InstallCDevelopmentTools(); err != nil {
		return err
	}
	
	return nil
}

// checkEnvironmentStatus checks and displays the current environment status
func checkEnvironmentStatus() {
	fmt.Println("=== Development Environment Status ===")
	
	// Check package managers
	fmt.Println("\nPackage Managers:")
	checkCommand("winget", "Windows Package Manager")
	checkCommand("choco", "Chocolatey")
	
	// Check Python environment
	fmt.Println("\nPython Environment:")
	checkCommand("python", "Python")
	checkCommand("pip", "pip")
	
	// Check C environment
	fmt.Println("\nC Development Environment:")
	checkCommand("gcc", "GCC (MinGW)")
	checkCommand("cl", "Microsoft C Compiler")
	checkCommand("cmake", "CMake")
	checkCommand("make", "Make")
	
	// Check common tools
	fmt.Println("\nCommon Development Tools:")
	checkCommand("git", "Git")
	checkCommand("code", "Visual Studio Code")
}

// checkCommand checks if a command is available and shows its version
func checkCommand(command, displayName string) {
	if isCommandAvailable(command) {
		fmt.Printf("✓ %s: Available\n", displayName)
	} else {
		fmt.Printf("✗ %s: Not found\n", displayName)
	}
}

// isCommandAvailable checks if a command is available in PATH
func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// InitCommands initializes and adds all commands to the root command
func InitCommands(rootCmd *cobra.Command) {
	// Add setup subcommands
	setupCmd.AddCommand(pythonCmd)
	setupCmd.AddCommand(cCmd)
	setupCmd.AddCommand(allCmd)
	
	// Add new project subcommands
	newCmd.AddCommand(newPythonCmd)
	newCmd.AddCommand(newCCmd)
	
	// Add all commands to root
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(statusCmd)
}
