package ansibleGen

import (
	"reflect"
	"testing"
)

func testRole() AnsibleRole {
	return *NewAnsibleRole("my_role_name")
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
	want := "my_role_name"
	if got := role.name; got != want {
		t.Errorf("Role doesn't have expected name, wanted %s, got %s", want, got)
	}
}
