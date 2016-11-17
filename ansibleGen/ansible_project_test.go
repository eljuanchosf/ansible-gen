package ansibleGen

import (
	"reflect"
	"testing"
)

type testArgs struct {
	Name        string
	CustomRoles string
	GalaxyRoles string
}

var testProjectArgs = testArgs{}

func testProject() AnsibleProject {
	testProjectArgs.Name = "my_test_name"
	testProjectArgs.CustomRoles = "crole1,crole2"
	testProjectArgs.GalaxyRoles = "grole1,grole2,grole3"
	return *NewAnsibleProject(testProjectArgs.Name, testProjectArgs.CustomRoles, testProjectArgs.GalaxyRoles)
}

func Test_NewAnsibleProject(t *testing.T) {

	ap := testProject()
	apType := reflect.TypeOf(ap).Kind()
	if apType != reflect.Struct {
		t.Errorf("NewAnsibleProject didn't return an struct")
	}

}

func Test_ProjectHasAName(t *testing.T) {
	project := testProject()
	if project.Name != testProjectArgs.Name {
		t.Errorf("Project has a name, wanted %s, got %s", project.Name, testProjectArgs.Name)
	}
}

func Test_ProjectHasAnArrayOfRoles(t *testing.T) {
	project := testProject()
	want := len(splitRoles(testProjectArgs.CustomRoles))
	if got := len(project.CustomRoles); got != want {
		t.Errorf("Project has custom roles, wanted %d, got %d", want, got)
	}
	projectCustomRoleType := reflect.TypeOf(project.CustomRoles)
	roleType := reflect.TypeOf([]AnsibleRole{})
	if projectCustomRoleType != roleType {
		t.Errorf("Project has %d custom roles of type %s, expected of type %s", want, roleType, projectCustomRoleType)
	}
}

func Test_ProjectHasGalaxyRoles(t *testing.T) {
	project := testProject()
	want := len(splitRoles(testProjectArgs.GalaxyRoles))
	if got := len(project.GalaxyRoles); got != want {
		t.Errorf("Project has Galaxy roles, wanted %d, got %d", want, got)
	}
}

func Test_SplitRoles(t *testing.T) {
	type args struct {
		roles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Split two roles", args{"role1,role2"}, 2},
		{"Only one role", args{"role"}, 1},
		{"No role", args{""}, 0},
		{"Three roles", args{"role1, role2, role3"}, 3},
	}
	for _, tt := range tests {
		if got := len(splitRoles(tt.args.roles)); got != tt.want {
			t.Errorf("%q. splitRoles() = %d, want %d", tt.name, got, tt.want)
		}
	}
}

func Test_SplitCustomRoles(t *testing.T) {
	type args struct {
		roles string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Split two roles", args{"role1,role2"}, 2},
		{"Only one role", args{"role"}, 1},
		{"No role", args{""}, 0},
		{"Three roles", args{"role1, role2, role3"}, 3},
	}
	for _, tt := range tests {
		if got := len(splitCustomRoles(tt.args.roles)); got != tt.want {
			t.Errorf("%q. splitCustomRoles() = %d, want %d", tt.name, got, tt.want)
		}
	}
}

func Test_getProjectTreeTemplate(t *testing.T) {
	projectName := "a_project_name"
	tree := getProjectTreeTemplate(projectName)
	if tree.Name != projectName {
		t.Errorf("The tree structure does not have the project name")
	}
	if tree.Folders[0].Name != "group_vars" {
		t.Error("The tree structure is not correct")
	}
}

func Test_ProjectHasATreeStructure(t *testing.T) {
	p := testProject()
	if len(p.TreeStructure.Folders) == 0 {
		t.Error("Tree structure for the project is empty")
	}
}

func Test_ProjectAddsRoles(t *testing.T) {
	p := testProject()
	rolesIndex := p.rolesFolderIndex("roles")
	if len(p.TreeStructure.Folders[rolesIndex].Folders) != 2 {
		t.Error("Project does not have the roles in the tree structure")
	}
}
