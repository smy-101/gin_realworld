package security

import "testing"

func TestGenerateJWT(t *testing.T) {
	token, error := generateJWT("jack", "jack@gamil.com")
	if error != nil {
		t.Errorf("Error while generating")
	}
	t.Logf("Token: %s", token)
}
