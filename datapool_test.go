package veil_api_go

import (
	"testing"
)

func TestDataPoolService_List_Get(t *testing.T) {

	client := NewClient("", "")
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
