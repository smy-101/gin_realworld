package security

import (
	"gin_realworld/utils"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	token, error := GenerateJWT("jack", "jack@gamil.com")
	if error != nil {
		t.Errorf("Error while generating")
	}
	t.Logf("Token: %s", token)
}

func TestVerifyJWT(t *testing.T) {
	claim, valid, err := VerifyJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMwMjEyMDYsImlhdCI6MTcyMjkzNDgwNiwidXNlciI6eyJlbWFpbCI6ImphY2tAZ2FtaWwuY29tIiwidXNlcm5hbWUiOiJqYWNrIn19.TkpqnYBuEe-YoJm4qshARIpsReNH3d9RsBM1Evas56g")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("verify jwt: %v, claim: %v\n", valid, utils.JsonMarshal(claim))
}
