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
	fs := *baseFs
	bp := afero.NewBasePathFs(fs, rootDir)
	newRoot := filepath.Join(rootDir, treeStructure.Name)
	if dryRun {
		fmt.Printf("D: %s\n", newRoot)
	} else {
		bp.Mkdir(treeStructure.Name, 0755)
	}
	bp = afero.NewBasePathFs(fs, newRoot)
	for _, file := range treeStructure.Files {
		if dryRun {
			fmt.Printf("F: %s\n", filepath.Join(newRoot, file.Name))
		} else {
			bp.Create(file.Name)
		}
	}
	for _, folder := range treeStructure.Folders {
		WriteTreeToDisk(newRoot, folder, baseFs, dryRun)
	}
}
