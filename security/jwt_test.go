package security

import "testing"

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("jack", "jack@gmail.com")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("token: %v\n", token)
}

func TestVerifyJWT(t *testing.T) {
	claim, valid, err := VerifyJWT("eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMwMjAzNTcsImlhdCI6MTcyMjkzMzk1NywidXNlciI6eyJlbWFpbCI6ImphY2tAZ21haWwuY29tIiwidXNlcm5hbWUiOiJqYWNrIn19.SjpT38zMQDkSDD2Xib-QBYUaDV3i5HtMJ47IQ_GorC8DUnoQSfUIoq1B7g-6AYZqXPoAUideDQeFhveLtK8BX8cB1TySYkhQHn0StwQxsycwZFE280mJ7ZTXU8SvzvoABFAVwq9YVVIyPBULu52yRxLqu4RfAvRoCZH5EMvpBaR1mpvYovZ8dUd1fT99K_T6P-AF5o3ESMhDzHirjBCg34rfBSxAO5IYn5tcWH8inHp1CUZo5Gn6mfX_7GWlMucyDSEtGIws4lLAzyenHf6yvPwNQc-pnjqVhVRr-U6mifSg9V8IsXDxuW8AB-1sMdFSuXpAMrXxXinWSS_CGHP_vA")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("verify jwt: %v\n", valid)
	t.Logf("claim: %v\n", claim)
}
