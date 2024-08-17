package security

import "testing"

func TestHashPassword(t *testing.T) {
	hashPassword, err := HashPassword("secret123")
	if err != nil {
		t.Error("Error while hashing password")
		return
	}
	t.Logf("Hashed password: %v\n", hashPassword)
	check := CheckPassword("secret123", hashPassword)
	if !check {
		t.Error("Password does not match")
	}
	t.Logf("Password match: %v\n", check)
}
