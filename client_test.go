package veil_api_client_go

import (
	"testing"
)

func Test_Client(t *testing.T) {
	client := NewClient("https://192.168.11.105", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo1NiwidXNlcm5hbWUiOiJidXIiLCJleHAiOjE5NTU0Mjc5OTEsInNzbyI6ZmFsc2UsIm9yaWdfaWF0IjoxNjQwOTMxOTkxfQ.BCPJi1hE_uvlv_sCjLYwGGq2qKJU8dbR9UUC5Cy79AA", true)
	_, _, err := client.DataPool.List()
	if err != nil {
		t.Error(err)
		return
	}
	//assert.Nil(t, err)

}
