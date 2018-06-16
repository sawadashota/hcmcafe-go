package entity

import (
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	now := time.Now()
	emptyTime := time.Time{}
	ts := NewTimestamp(now, now, emptyTime)

	if ts.CreatedAt != now {
		t.Errorf("CreatedAt should be %s, but %s", now, ts.CreatedAt)
	}

	if ts.UpdatedAt != now {
		t.Errorf("UpdatedAt should be %s, but %s", now, ts.UpdatedAt)
	}

	if ts.IsDeleted() {
		t.Errorf("DeletedAt should be %s, but %s", emptyTime, ts.DeletedAt)
	}

	ts.Delete()

	if !ts.IsDeleted() {
		t.Errorf("DeletedAt should be empty")
	}

	ts.StampTime()

	if ts.UpdatedAt.Before(now) {
		t.Errorf("UpdateAt should be future campare with first value")
	}
}

func TestCreateTimestamp(t *testing.T) {
	emptyTime := time.Time{}

	ts := CreateTimestamp()

	if ts.DeletedAt != emptyTime {
		t.Errorf("DeletedAt should be empty")
	}
}
