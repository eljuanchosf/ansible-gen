package ansibleGen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

//File represents a system file
type File struct {
	Name    string `yaml:"name"`
	Content string `yaml:"content"`
}

//Folder represents a system folder
type Folder struct {
	Name    string   `yaml:"name"`
	Files   []File   `yaml:"files"`
	Folders []Folder `yaml:"folders"`
}

//WriteTreeToDisk creates a directory structure based on the treeStructure parameter
func WriteTreeToDisk(rootDir string, treeStructure Folder, baseFs *afero.Fs, dryRun bool, baseDir string) {
	newRoot, basePath := createDirectory(rootDir, treeStructure, baseFs, dryRun, baseDir)
	for _, file := range treeStructure.Files {
		createFile(newRoot, file, &basePath, dryRun, baseDir)
	}
	for _, folder := range treeStructure.Folders {
		WriteTreeToDisk(newRoot, folder, baseFs, dryRun, baseDir)
	}
}

func createDirectory(rootDir string, folder Folder, baseFs *afero.Fs, dryRun bool, baseDir string) (string, afero.Fs) {
	fs := *baseFs
	basePath := afero.NewBasePathFs(fs, rootDir)
	newRoot := filepath.Join(rootDir, folder.Name)
	if dryRun {
		fmt.Printf("Create directory %s\n", strings.Replace(newRoot, baseDir, ".", 1))
	} else {
		basePath.Mkdir(folder.Name, 0755)
	}
	basePath = afero.NewBasePathFs(fs, newRoot)
	return newRoot, basePath
}

func createFile(rootDir string, file File, baseFs *afero.Fs, dryRun bool, baseDir string) {
	fs := *baseFs
	if dryRun {
		fmt.Printf("Create file      %s\n", strings.Replace(filepath.Join(rootDir, file.Name), baseDir, ".", 1))
	} else {

		if file.Content == "" {
			switch filepath.Ext(file.Name) {
			case ".yml":
				file.Content = yamlTemplate()
			case "":
				file.Content = variablesTemplate()
			}
		}
		afero.WriteFile(fs, file.Name, []byte(file.Content), 0644)
	}
}
