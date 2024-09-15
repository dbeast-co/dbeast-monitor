//go:build mage
// +build mage

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	// mage:import
	build "github.com/grafana/grafana-plugin-sdk-go/build"
)

// Default sets the default mage target to CopyResources
var Default = Build

//var Default = build.BuildAll

func Build() error {
	// Call CopyResources and check for errors
	if err := CopyResources(); err != nil {
		return err
	}

	// Call build.BuildAll, assuming it does not return an error
	build.BuildAll() // just call the function without handling return since it doesn't return any value

	return nil
}

// CopyResources recursively copies resource files to the build directory.
func CopyResources() error {
	sourceDir := "./pkg/resources" // Source directory containing resources
	outputDir := "./dist/data"     // Target directory for resources

	// Create the output directory if it doesn't already exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create resource directory: %v", err)
	}

	// Walk through the source directory and copy each file
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root source directory
		if path == sourceDir {
			return nil
		}

		// Create target path based on the relative path to the source directory
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(outputDir, relPath)

		if info.IsDir() {
			// Create directory if it's a directory
			return os.MkdirAll(targetPath, 0755)
		} else {
			// Copy file if it's a file
			return copyFile(path, targetPath)
		}
	})

	if err != nil {
		return fmt.Errorf("failed to copy resources: %v", err)
	}

	fmt.Println("Resources copied successfully.")
	return nil
}

// copyFile copies a single file from src to dst.
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
