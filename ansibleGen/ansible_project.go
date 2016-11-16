package ansibleGen

import "strings"

//AnsibleProject represents the structure of an Ansible Project
type AnsibleProject struct {
	name        string
	customRoles []AnsibleRole
	galaxyRoles []string
}

//NewAnsibleProject initializes the structure for a new Ansible project
func NewAnsibleProject(name string, customRoles string, galaxyRoles string) *AnsibleProject {
	return &AnsibleProject{
		name:        name,
		customRoles: splitCustomRoles(customRoles),
		galaxyRoles: splitRoles(galaxyRoles),
	}
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
