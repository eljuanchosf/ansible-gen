package ansibleGen

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/afero"

	yaml "gopkg.in/yaml.v2"
)

//AnsibleRole represents the structure of an Ansible Role
type AnsibleRole struct {
	Name          string
	TreeStructure Folder
}

//NewAnsibleRole initializes the structure for a new Ansible role
func NewAnsibleRole(name string) *AnsibleRole {
	return &AnsibleRole{
		Name:          name,
		TreeStructure: getRoleTreeTemplate(name),
	}
}

//Save run the tree structure creation for the project
func (role *AnsibleRole) Save(dryRun bool) {
	baseFs := afero.NewOsFs()
	rootDir, _ := os.Getwd()
	fmt.Println("Using root directory: ", rootDir)
	WriteTreeToDisk(rootDir, role.TreeStructure, &baseFs, dryRun, rootDir)
	fmt.Printf("Ansible role %s has been generated\n", role.Name)
}

func getRoleTreeTemplate(roleName string) Folder {
	roleTemplate := `
---
name: ` + roleName + `
folders:
- name: tasks
  files:
  - name: main.yml
  folders:
- name: handlers
  files:
  - name: main.yml
- name: vars
  files:
  - name: main.yml
- name: defaults
  files:
  - name: main.yml
- name: meta
  files:
  - name: main.yml
- name: templates
  files:
  - name: template.j2
- name: files
`

	dst := Folder{}
	err := yaml.Unmarshal([]byte(roleTemplate), &dst)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return dst
}
