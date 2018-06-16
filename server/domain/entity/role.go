package entity

import "fmt"

const (
	MasterRole  = "master"
	EditorRole  = "editor"
	WatcherRole = "watcher"
)

type Role string

// NewRole return WatcherRole
func NewRole() *Role {
	var r Role = WatcherRole
	return &r
}

func (r *Role) String() string {
	return fmt.Sprint(*r)
}

func (r *Role) SetMaster() {
	*r = MasterRole
}

func (r *Role) SetEditor() {
	*r = EditorRole
}

func (r *Role) SetWatcher() {
	*r = WatcherRole
}
