package ansibleGen

import (
	"testing"

	"github.com/spf13/afero"
)

func Test_WriteTreeToDiskOnProject(t *testing.T) {
	p := testProject()
	baseFs := afero.NewMemMapFs()
	WriteTreeToDisk("/tmp", p.TreeStructure, &baseFs, false)
	if _, err := baseFs.Stat("/tmp/my_test_name/roles/crole1"); err != nil {
		t.Errorf("Expected 'roles/crole1' to be present in filesystem structure")
	}
	if _, err := baseFs.Stat("/tmp/my_test_name/production"); err != nil {
		t.Errorf("Expected '/tmp/my_test_name/production' file to be present in filesystem structure")
	}
}

func Test_WriteTreeToDiskOnRoles(t *testing.T) {
	r := testRole()
	baseFs := afero.NewMemMapFs()
	WriteTreeToDisk("/tmp", r.TreeStructure, &baseFs, false)
	if _, err := baseFs.Stat("/tmp/my_role_name"); err != nil {
		t.Errorf("Expected 'my_role_name' to be present in filesystem structure")
	}
	if _, err := baseFs.Stat("/tmp/my_role_name/tasks/main.yml"); err != nil {
		t.Errorf("Expected '/tmp/my_role_name/tasks/main.yml' file to be present in filesystem structure")
	}
}

func Test_WriteTreeToDiskOnDryRun(t *testing.T) {
	r := testRole()
	baseFs := afero.NewMemMapFs()
	WriteTreeToDisk("/tmp", r.TreeStructure, &baseFs, true)
	if _, err := baseFs.Stat("/tmp/my_role_name"); err == nil {
		t.Errorf("Expected 'my_role_name' should not be present in the file structure while in Dry Run")
	}
}
