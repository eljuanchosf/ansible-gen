package ansibleGen

import (
	"fmt"
	"path/filepath"

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
func WriteTreeToDisk(rootDir string, treeStructure Folder, baseFs *afero.Fs, dryRun bool) {
	newRoot, basePath := createDirectory(rootDir, treeStructure, baseFs, dryRun)
	for _, file := range treeStructure.Files {
		createFile(newRoot, file, &basePath, dryRun)
	}
	for _, folder := range treeStructure.Folders {
		WriteTreeToDisk(newRoot, folder, baseFs, dryRun)
	}
}

func createDirectory(rootDir string, folder Folder, baseFs *afero.Fs, dryRun bool) (string, afero.Fs) {
	fs := *baseFs
	bp := afero.NewBasePathFs(fs, rootDir)
	newRoot := filepath.Join(rootDir, folder.Name)
	if dryRun {
		fmt.Printf("D: %s\n", newRoot)
	} else {
		bp.Mkdir(folder.Name, 0755)
	}
	bp = afero.NewBasePathFs(fs, newRoot)
	return newRoot, bp
}

func createFile(rootDir string, file File, baseFs *afero.Fs, dryRun bool) {
	fs := *baseFs
	if dryRun {
		fmt.Printf("F: %s\n", filepath.Join(rootDir, file.Name))
	} else {
		fs.Create(file.Name)
	}
}
