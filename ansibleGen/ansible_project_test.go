package ansibleGen

import (
	"reflect"
	"testing"
)

func testProject() AnsibleProject {
	return *NewAnsibleProject("my_test_name", "crole1,crole2", "grole1,grole2,grole3")
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
	want := "my_test_name"
	if project.name != want {
		t.Errorf("Project has a name, wanted %s, got %s", project.name, want)
	}
}

func Test_ProjectHasRoles(t *testing.T) {
	project := testProject()
	want := 2
	if got := len(project.customRoles); got != want {
		t.Errorf("Project has custom roles, wanted %d, got %d", want, got)
	}
	want = 3
	if got := len(project.galaxyRoles); got != want {
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
