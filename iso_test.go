package veil_api_go

import (
	"testing"
)

func TestIsoService_List_Get(t *testing.T) {
	client := NewClient("", "")
	response, _, err := client.Iso.List()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range response.Results {
		entity, _, err := client.Iso.Get(v.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if entity.Id == "" {
			t.Error("Iso Id can not be empty")
			return
		}
		break
	}

	return

}

func TestIsoService_Upload(t *testing.T) {
	t.SkipNow()
	client := NewClient("", "")
	response, _, err := client.DataPool.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(response.Results) == 0 {
		t.SkipNow()
	}
	firstDp := response.Results[0]
	iso, _, err := client.Iso.Create(firstDp.Id, "test_live.iso")
	if err != nil {
		t.Error(err)
		return
	}
	if iso.Id == "" {
		t.Error("Iso Id can not be empty")
		return
	}
	return

}
