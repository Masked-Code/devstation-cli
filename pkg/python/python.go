package python

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"devstation-cli/pkg/installer"
)

// PythonSetup handles Python development environment setup
type PythonSetup struct {
	PackageManager installer.PackageManager
}

// NewPythonSetup creates a new Python setup instance
func NewPythonSetup(pm installer.PackageManager) *PythonSetup {
	return &PythonSetup{PackageManager: pm}
}

// InstallPython installs Python and essential tools
func (p *PythonSetup) InstallPython() error {
	fmt.Println("=== Setting up Python Development Environment ===")
	
	// Install Python
	if err := p.PackageManager.Install("python"); err != nil {
		return fmt.Errorf("failed to install Python: %v", err)
	}
	
	// Install pip (usually comes with Python, but ensure it's available)
	if err := p.ensurePipInstalled(); err != nil {
		return fmt.Errorf("failed to ensure pip is installed: %v", err)
	}
	
	// Install essential Python packages
	if err := p.installEssentialPackages(); err != nil {
		return fmt.Errorf("failed to install essential packages: %v", err)
	}
	
	fmt.Println("✓ Python development environment setup complete!")
	return nil
}

// InstallPythonTools installs common Python development tools
func (p *PythonSetup) InstallPythonTools() error {
	fmt.Println("=== Installing Python Development Tools ===")
	
	tools := []string{
		"git",
		"vscode", // Visual Studio Code
	}
	
	for _, tool := range tools {
		if err := p.PackageManager.Install(tool); err != nil {
			fmt.Printf("Warning: Failed to install %s: %v\n", tool, err)
		}
	}
	
	return nil
}

// ensurePipInstalled checks if pip is available and installs it if needed
func (p *PythonSetup) ensurePipInstalled() error {
	// Check if pip is available
	cmd := exec.Command("python", "-m", "pip", "--version")
	if err := cmd.Run(); err != nil {
		fmt.Println("Installing pip...")
		// Download and install pip
		cmd := exec.Command("python", "-m", "ensurepip", "--upgrade")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	
	fmt.Println("✓ pip is available")
	return nil
}

// installEssentialPackages installs essential Python packages
func (p *PythonSetup) installEssentialPackages() error {
	fmt.Println("Installing essential Python packages...")
	
	packages := []string{
		"virtualenv",     // Virtual environment management
		"pip-tools",      // Package dependency management
		"black",          // Code formatter
		"flake8",         // Linting
		"pytest",         // Testing framework
		"requests",       // HTTP library
		"numpy",          // Scientific computing
		"pandas",         // Data analysis
		"jupyter",        // Interactive notebooks
	}
	
	for _, pkg := range packages {
		fmt.Printf("Installing %s...\n", pkg)
		cmd := exec.Command("python", "-m", "pip", "install", pkg)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Warning: Failed to install %s: %v\n", pkg, err)
		}
	}
	
	return nil
}

// CreateProjectStructure creates a basic Python project structure
func (p *PythonSetup) CreateProjectStructure(projectName string) error {
	fmt.Printf("Creating Python project structure for '%s'...\n", projectName)
	
	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %v", err)
	}
	
	// Create subdirectories
	dirs := []string{
		filepath.Join(projectName, "src"),
		filepath.Join(projectName, "tests"),
		filepath.Join(projectName, "docs"),
	}
	
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	
	// Create essential files
	files := map[string]string{
		filepath.Join(projectName, "requirements.txt"):     "# Add your dependencies here\n",
		filepath.Join(projectName, "README.md"):            fmt.Sprintf("# %s\n\nDescription of your project.\n", projectName),
		filepath.Join(projectName, ".gitignore"):           pythonGitignore,
		filepath.Join(projectName, "src", "__init__.py"):   "",
		filepath.Join(projectName, "tests", "__init__.py"): "",
		filepath.Join(projectName, "setup.py"):             generateSetupPy(projectName),
	}
	
	for filename, content := range files {
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filename, err)
		}
	}
	
	// Create virtual environment
	fmt.Println("Creating virtual environment...")
	cmd := exec.Command("python", "-m", "venv", filepath.Join(projectName, "venv"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to create virtual environment: %v\n", err)
	}
	
	fmt.Printf("✓ Python project '%s' created successfully!\n", projectName)
	fmt.Printf("To activate the virtual environment, run:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  venv\\Scripts\\activate\n")
	
	return nil
}

// generateSetupPy generates a basic setup.py file
func generateSetupPy(projectName string) string {
	return fmt.Sprintf(`from setuptools import setup, find_packages

setup(
    name="%s",
    version="0.1.0",
    description="A Python project",
    packages=find_packages(where="src"),
    package_dir={"": "src"},
    python_requires=">=3.6",
    install_requires=[
        # Add your dependencies here
    ],
    extras_require={
        "dev": [
            "pytest",
            "black",
            "flake8",
        ],
    },
)
`, projectName)
}

// pythonGitignore contains a comprehensive .gitignore for Python projects
const pythonGitignore = `# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
pip-wheel-metadata/
share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# PyInstaller
*.manifest
*.spec

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
htmlcov/
.tox/
.nox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.py,cover
.hypothesis/
.pytest_cache/

# Virtual environments
venv/
env/
ENV/
env.bak/
venv.bak/

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`
