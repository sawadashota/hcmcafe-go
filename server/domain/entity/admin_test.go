package entity

import (
	"testing"
	"time"
)

func TestAdmin(t *testing.T) {
	now := time.Now()
	emptyTime := time.Time{}

	admin := NewAdmin("澤田",
		"翔太",
		"sawada@exmaple.com",
		"123345678",
		"Hello",
		"aaa",
		now,
		now,
		emptyTime)

	admin.Delete()

	if !admin.IsDeleted() {
		t.Errorf("Admin should be deleted")
	}
}
