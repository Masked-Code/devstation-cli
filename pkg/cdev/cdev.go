package cdev

import (
	"fmt"
	"os"
	"path/filepath"

	"devstation-cli/pkg/installer"
)

// CDevSetup handles C development environment setup
type CDevSetup struct {
	PackageManager installer.PackageManager
}

// NewCDevSetup creates a new C development setup instance
func NewCDevSetup(pm installer.PackageManager) *CDevSetup {
	return &CDevSetup{PackageManager: pm}
}

// InstallCDevelopmentTools installs C compiler and development tools
func (c *CDevSetup) InstallCDevelopmentTools() error {
	fmt.Println("=== Setting up C Development Environment ===")
	
	// Install Microsoft Visual Studio Build Tools or MinGW
	if err := c.installCompiler(); err != nil {
		return fmt.Errorf("failed to install C compiler: %v", err)
	}
	
	// Install additional development tools
	if err := c.installDevelopmentTools(); err != nil {
		return fmt.Errorf("failed to install development tools: %v", err)
	}
	
	fmt.Println("✓ C development environment setup complete!")
	return nil
}

// installCompiler installs a C compiler (MinGW-w64 or MSVC)
func (c *CDevSetup) installCompiler() error {
	fmt.Println("Installing C compiler...")
	
	// Try to install MinGW-w64 first (more portable)
	if err := c.PackageManager.Install("mingw"); err != nil {
		fmt.Printf("MinGW installation failed: %v\n", err)
		
		// Fallback to Visual Studio Build Tools
		fmt.Println("Trying Visual Studio Build Tools...")
		if err := c.PackageManager.Install("visualstudio2022buildtools"); err != nil {
			return fmt.Errorf("failed to install both MinGW and VS Build Tools: %v", err)
		}
	}
	
	return nil
}

// installDevelopmentTools installs additional C development tools
func (c *CDevSetup) installDevelopmentTools() error {
	fmt.Println("Installing C development tools...")
	
	tools := []string{
		"cmake",           // Build system
		"make",            // Make utility
		"git",             // Version control
		"vscode",          // IDE
		"clang-format",    // Code formatter
		"gdb",             // Debugger
	}
	
	for _, tool := range tools {
		if err := c.PackageManager.Install(tool); err != nil {
			fmt.Printf("Warning: Failed to install %s: %v\n", tool, err)
		}
	}
	
	return nil
}

// CreateCProject creates a basic C project structure
func (c *CDevSetup) CreateCProject(projectName string) error {
	fmt.Printf("Creating C project structure for '%s'...\n", projectName)
	
	// Create project directory
	if err := os.MkdirAll(projectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %v", err)
	}
	
	// Create subdirectories
	dirs := []string{
		filepath.Join(projectName, "src"),
		filepath.Join(projectName, "include"),
		filepath.Join(projectName, "tests"),
		filepath.Join(projectName, "build"),
		filepath.Join(projectName, "docs"),
	}
	
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	
	// Create essential files
	files := map[string]string{
		filepath.Join(projectName, "README.md"):              generateReadme(projectName),
		filepath.Join(projectName, ".gitignore"):             cGitignore,
		filepath.Join(projectName, "CMakeLists.txt"):         generateCMakeLists(projectName),
		filepath.Join(projectName, "Makefile"):               generateMakefile(projectName),
		filepath.Join(projectName, "src", "main.c"):          generateMainC(projectName),
		filepath.Join(projectName, "src", projectName+".c"):  generateSourceC(projectName),
		filepath.Join(projectName, "include", projectName+".h"): generateHeader(projectName),
		filepath.Join(projectName, "tests", "test_main.c"):   generateTestC(projectName),
	}
	
	for filename, content := range files {
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filename, err)
		}
	}
	
	fmt.Printf("✓ C project '%s' created successfully!\n", projectName)
	fmt.Printf("To build the project, run:\n")
	fmt.Printf("  cd %s\n", projectName)
	fmt.Printf("  mkdir build && cd build\n")
	fmt.Printf("  cmake ..\n")
	fmt.Printf("  make\n")
	
	return nil
}

// generateReadme generates a README.md file for the C project
func generateReadme(projectName string) string {
	return fmt.Sprintf("# %s\n\nA C project created with DevStation CLI.\n\n## Building\n\n### Using CMake\n```bash\nmkdir build\ncd build\ncmake ..\nmake\n```\n\n### Using Make directly\n```bash\nmake\n```\n\n## Running\n```bash\n./build/%s\n```\n\n## Testing\n```bash\nmake test\n```\n\n## Project Structure\n- `src/` - Source files\n- `include/` - Header files\n- `tests/` - Test files\n- `build/` - Build artifacts\n- `docs/` - Documentation\n", projectName, projectName)
}

// generateCMakeLists generates a CMakeLists.txt file
func generateCMakeLists(projectName string) string {
	return fmt.Sprintf(`cmake_minimum_required(VERSION 3.10)
project(%s)

# Set C standard
set(CMAKE_C_STANDARD 11)
set(CMAKE_C_STANDARD_REQUIRED ON)

# Add include directory
include_directories(include)

# Add executable
add_executable(%s src/main.c src/%s.c)

# Add test executable
add_executable(test_main tests/test_main.c src/%s.c)

# Enable testing
enable_testing()
add_test(NAME test_main COMMAND test_main)
`, projectName, projectName, projectName, projectName)
}

// generateMakefile generates a basic Makefile
func generateMakefile(projectName string) string {
	return fmt.Sprintf(`CC=gcc
CFLAGS=-Wall -Wextra -std=c11 -Iinclude
SRCDIR=src
BUILDDIR=build
TESTDIR=tests
SOURCES=$(SRCDIR)/main.c $(SRCDIR)/%s.c

# Default target
all: $(BUILDDIR)/%s

# Create build directory
$(BUILDDIR):
	mkdir -p $(BUILDDIR)

# Build main executable
$(BUILDDIR)/%s: $(SOURCES) | $(BUILDDIR)
	$(CC) $(CFLAGS) -o $@ $(SOURCES)

# Build and run tests
test: $(BUILDDIR)/test_main
	$(BUILDDIR)/test_main

$(BUILDDIR)/test_main: $(TESTDIR)/test_main.c $(SRCDIR)/%s.c | $(BUILDDIR)
	$(CC) $(CFLAGS) -o $@ $^

# Clean build artifacts
clean:
	rm -rf $(BUILDDIR)

# Run the program
run: $(BUILDDIR)/%s
	$(BUILDDIR)/%s

.PHONY: all test clean run
`, projectName, projectName, projectName, projectName, projectName, projectName)
}

// generateMainC generates a main.c file
func generateMainC(projectName string) string {
	return fmt.Sprintf(`#include <stdio.h>
#include "%s.h"

int main(void) {
    printf("Hello from %s!\\n");
    
    // Call a function from your header
    greet();
    
    return 0;
}
`, projectName, projectName)
}

// generateSourceC generates a source file with function implementations
func generateSourceC(projectName string) string {
	return fmt.Sprintf(`#include <stdio.h>
#include "%s.h"

void greet(void) {
    printf("Hello from %s library!\\n");
}
`, projectName, projectName)
}

// generateHeader generates a header file
func generateHeader(projectName string) string {
	return fmt.Sprintf(`#ifndef %s_H
#define %s_H

// Function prototypes
void greet(void);

#endif // %s_H
`, projectName, projectName, projectName)
}

// generateTestC generates a basic test file
func generateTestC(projectName string) string {
	return fmt.Sprintf(`#include <stdio.h>
#include <assert.h>
#include "%s.h"

void test_basic_functionality(void) {
    // Add your tests here
    printf("Running basic functionality test...\\n");
    // Example: assert(some_function() == expected_value);
    printf("✓ Basic functionality test passed\\n");
}

int main(void) {
    printf("Running tests for %s...\\n");
    
    test_basic_functionality();
    
    printf("All tests passed!\\n");
    return 0;
}
`, projectName, projectName)
}

// cGitignore contains a comprehensive .gitignore for C projects
const cGitignore = `# Object files
*.o
*.ko
*.obj
*.elf

# Linker output
*.ilk
*.map
*.exp

# Precompiled Headers
*.gch
*.pch

# Libraries
*.lib
*.a
*.la
*.lo

# Shared objects (inc. Windows DLLs)
*.dll
*.so
*.so.*
*.dylib

# Executables
*.exe
*.out
*.app
*.i*86
*.x86_64
*.hex

# Debug files
*.dSYM/
*.su
*.idb
*.pdb

# Build directories
build/
Build/
BUILD/
out/
bin/
obj/

# CMake
CMakeCache.txt
CMakeFiles/
cmake_install.cmake
Makefile
*.cmake
!CMakeLists.txt

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS files
.DS_Store
Thumbs.db

# Core dumps
core
*.core

# Temporary files
*.tmp
*.temp
*.bak
*.backup
`
