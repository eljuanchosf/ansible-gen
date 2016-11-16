package ansibleGen

//AnsibleRole represents the structure of an Ansible Role
type AnsibleRole struct {
	name string
}

//NewAnsibleRole initializes the structure for a new Ansible role
func NewAnsibleRole(name string) *AnsibleRole {
	return &AnsibleRole{
		name: name,
	}
}
