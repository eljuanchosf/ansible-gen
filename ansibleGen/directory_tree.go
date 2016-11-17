package ansibleGen

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
