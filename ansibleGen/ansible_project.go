package ansibleGen

import "strings"

//AnsibleProject represents the structure of an Ansible Project
type AnsibleProject struct {
	name        string
	customRoles []string
	galaxyRoles []string
}

//NewAnsibleProject initializes the structure for a new Ansible project
func NewAnsibleProject(name string, customRoles string, galaxyRoles string) *AnsibleProject {
	projectCustomRoles := splitRoles(customRoles)
	projectGalaxyRoles := splitRoles(galaxyRoles)
	return &AnsibleProject{
		name:        name,
		customRoles: projectCustomRoles,
		galaxyRoles: projectGalaxyRoles,
	}
}

func splitRoles(roles string) []string {
	if len(roles) == 0 {
		return nil
	}
	return strings.Split(roles, ",")
}
