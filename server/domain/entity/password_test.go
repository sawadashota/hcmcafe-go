package entity

import "testing"

func TestPassword(t *testing.T) {
	p, err := HashPassword("password")

	if err != nil {
		t.Errorf("%s", err)
	}

	if p.Verify("hoge") {
		t.Errorf("Invalid password should never verify")
	}

	if !p.Verify("password") {
		t.Errorf("Valid password should verify")
	}
}
