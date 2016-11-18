package ansibleGen

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/afero"

	yaml "gopkg.in/yaml.v2"
)

//AnsibleProject represents the structure of an Ansible Project
type AnsibleProject struct {
	Name          string
	CustomRoles   []AnsibleRole
	GalaxyRoles   []string
	TreeStructure Folder
}

//NewAnsibleProject initializes the structure for a new Ansible project
func NewAnsibleProject(name string, customRoles string, galaxyRoles string) *AnsibleProject {
	ap := &AnsibleProject{
		Name:          name,
		CustomRoles:   splitCustomRoles(customRoles),
		GalaxyRoles:   splitRoles(galaxyRoles),
		TreeStructure: getProjectTreeTemplate(name),
	}
	ap.addRolesToTreeStructure()
	ap.addGalaxyRoles()
	return ap
}

//Save run the tree structure creation for the project
func (project *AnsibleProject) Save(dryRun bool) {
	baseFs := afero.NewOsFs()
	rootDir, _ := os.Getwd()
	fmt.Println("Using root directory: ", rootDir)
	WriteTreeToDisk(rootDir, project.TreeStructure, &baseFs, dryRun, rootDir)
	fmt.Printf("Ansible project %s has been generated\n", project.Name)
}

func (project *AnsibleProject) addGalaxyRoles() {
	if len(project.GalaxyRoles) > 0 {
		var galaxyRolesFile = File{Name: "galaxy-roles.yml"}
		fileContent := "---\n"
		for _, role := range project.GalaxyRoles {
			fileContent += fmt.Sprintln("-", role)
		}
		galaxyRolesFile.Content = fileContent
		project.TreeStructure.Files = append(project.TreeStructure.Files, galaxyRolesFile)
	}
}

func (project *AnsibleProject) addRolesToTreeStructure() {
	rolesIndex := project.rolesFolderIndex("roles")
	rolesFolder := &project.TreeStructure.Folders[rolesIndex]
	for _, role := range project.CustomRoles {
		rolesFolder.Folders = append(rolesFolder.Folders, role.TreeStructure)
	}
}

func (project *AnsibleProject) rolesFolderIndex(rolesFolderName string) int {
	for index, folder := range project.TreeStructure.Folders {
		if folder.Name == rolesFolderName {
			return index
		}
	}
	return 0
}

func splitCustomRoles(customRoles string) []AnsibleRole {
	roles := []AnsibleRole{}
	for _, roleName := range splitRoles(customRoles) {
		roles = append(roles, *NewAnsibleRole(roleName))
	}
	return roles
}

func splitRoles(roles string) []string {
	if len(roles) == 0 {
		return nil
	}
	return strings.Split(roles, ",")
}

func getProjectTreeTemplate(projectName string) Folder {
	projectTemplate := `
---
name: ` + projectName + `
files:
- name: production
- name: staging
- name: main.yml
folders:
- name: group_vars
  files:
  - name: group1
  - name: group2
  folders:
- name: hosts_vars 
  files:
  - name: hostname1
  - name: hostname2
- name: roles
`

	dst := Folder{}
	err := yaml.Unmarshal([]byte(projectTemplate), &dst)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return dst
}
