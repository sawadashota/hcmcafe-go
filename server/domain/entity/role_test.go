package entity

import "testing"

func TestNewRole(t *testing.T) {
	r := NewRole()

	if r.String() != MasterRole {
		t.Errorf("Expect %s, actual %s", MasterRole, r)
	}
}

func TestRole_SetMaster(t *testing.T) {
	r := NewRole()
	r.SetMaster()

	if r.String() != MasterRole {
		t.Errorf("Expect %s, actual %s", MasterRole, r)
	}
}

func TestRole_SetEditor(t *testing.T) {
	r := NewRole()
	r.SetEditor()

	if r.String() != EditorRole {
		t.Errorf("Expect %s, actual %s", EditorRole, r)
	}
}

func TestRole_SetWatcher(t *testing.T) {
	r := NewRole()
	r.SetWatcher()

	if r.String() != WatcherRole {
		t.Errorf("Expect %s, actual %s", WatcherRole, r)
	}
}
