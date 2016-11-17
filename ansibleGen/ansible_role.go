package ansibleGen

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

//AnsibleRole represents the structure of an Ansible Role
type AnsibleRole struct {
	name          string
	TreeStructure Folder
}

//NewAnsibleRole initializes the structure for a new Ansible role
func NewAnsibleRole(name string) *AnsibleRole {
	return &AnsibleRole{
		name:          name,
		TreeStructure: getRoleTreeTemplate(name),
	}
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
