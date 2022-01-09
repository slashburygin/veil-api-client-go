package veil_api_client_go

import (
	"testing"
)

func Test_DataPoolListGet(t *testing.T) {

	client := NewClient("", "", false)
	response, _, err := client.DataPool.List()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range response.Results {
		entity, _, err := client.DataPool.Get(v.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if entity.Id == "" {
			t.Error("DataPool Id can not be empty")
			return
		}
		break
	}

	return

}
