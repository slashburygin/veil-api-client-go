package veil_api_client_go

import (
	"testing"
)

func Test_IsoListGet(t *testing.T) {
	client := NewClient("", "", false)
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

func Test_IsoUpload(t *testing.T) {
	t.SkipNow()
	client := NewClient("", "", false)
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
