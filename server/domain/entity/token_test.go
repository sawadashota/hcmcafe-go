package entity

import (
	"testing"
	"time"
)

func TestGenerateSession(t *testing.T) {
	var newToken string
	var ts []string

	for i := 0; i < 10000; i++ {
		newToken = GenerateSession().Token
		for _, token := range ts {
			if token == newToken {
				t.Errorf("Session should be unique. generated %s", newToken)
			}
		}

		ts = append(ts, newToken)
	}
}

func TestSession_Refresh(t *testing.T) {
	token := &Session{
		Token:     "aaa",
		UpdatedAt: time.Now(),
	}

	token.Refresh()

	if token.Token == "aaa" {
		t.Errorf("Session value should be changed.")
	}
}

func TestSession_Flush(t *testing.T) {
	token := &Session{
		Token:     "aaa",
		UpdatedAt: time.Now(),
	}

	token.Flush()

	if token.Token != "" {
		t.Errorf("Session value should be empty.")
	}
}
