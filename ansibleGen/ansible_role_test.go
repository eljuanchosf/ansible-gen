package ansibleGen

import (
	"reflect"
	"testing"
)

var roleName string

func testRole() AnsibleRole {
	roleName = "my_role_name"
	return *NewAnsibleRole(roleName)
}

func Test_NewAnsibleRole(t *testing.T) {
	ar := testRole()
	arType := reflect.TypeOf(ar).Kind()
	if arType != reflect.Struct {
		t.Errorf("NewAnsibleRole didn't return an struct")
	}
}

func Test_AnsibleRoleHasName(t *testing.T) {
	role := testRole()
	if got := role.name; got != roleName {
		t.Errorf("Role doesn't have expected name, wanted %s, got %s", roleName, got)
	}
}

func Test_getRoleTreeTemplate(t *testing.T) {
	role := getRoleTreeTemplate(roleName)
	if role.Name != roleName {
		t.Errorf("The tree structure does not have the role name")
	}
}
