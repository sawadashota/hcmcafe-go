package entity

import "testing"

func TestGenerateToken(t *testing.T) {
	var newToken string
	var ts []string

	for i := 0; i < 10000; i++ {
		newToken = GenerateToken().Token
		for _, token := range ts {
			if token == newToken {
				t.Errorf("Token should be unique. generated %s", newToken)
			}
		}

		ts = append(ts, newToken)
	}
}

func TestToken_Refresh(t *testing.T) {
	token := &Token{"aaa"}
	token.Refresh()

	if token.Token == "aaa" {
		t.Errorf("Token value should be changed.")
	}
}

func TestToken_Flush(t *testing.T) {
	token := &Token{"aaa"}
	token.Flush()

	if token.Token != "" {
		t.Errorf("Token value should be empty.")
	}
}
