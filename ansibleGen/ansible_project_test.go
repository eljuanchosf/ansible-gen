package ansibleGen

import "testing"

func Test_ProjectHasTheRightName(t *testing.T) {
	type args struct {
		name        string
		customRoles string
		galaxyRoles string
	}
	test := struct {
		name string
		args args
		want string
	}{
		"Project name is test_project", args{"test_project", "role1,role2", "grole1,grole2"}, "test_project",
	}
	ap := *NewAnsibleProject(test.args.name, test.args.customRoles, test.args.galaxyRoles)
	if ap.name != test.want {
		t.Errorf("%q, wanted %s, got %s", test.name, ap.name, test.want)
	}
}

func Test_ProjectHasTheCustomRoles(t *testing.T) {
	type args struct {
		name        string
		customRoles string
		galaxyRoles string
	}
	test := struct {
		name string
		args args
		want int
	}{
		"Project has two custom roles", args{"test_project", "role1,role2", "grole1,grole2"}, 2,
	}
	ap := *NewAnsibleProject(test.args.name, test.args.customRoles, test.args.galaxyRoles)
	if len(ap.customRoles) != test.want {
		t.Errorf("%q, wanted %d, got %d", test.name, len(ap.customRoles), test.want)
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
