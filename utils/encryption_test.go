package utils

import "testing"

func TestCompareHashAndPassword(t *testing.T) {
	pwd := "123456"
	hashedPassWord, err := GenerateFromPassword(pwd)
	if err != nil {
		return
	}
	err = CompareHashAndPassword(hashedPassWord, pwd)
	if nil != err {
		t.Fatalf("[CompareHashAndPassword] pwd is err: %v\n", err)
	} else {
		t.Fatal("[CompareHashAndPassword] success!")
	}
}
