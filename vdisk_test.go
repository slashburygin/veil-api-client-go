package veil_api_go

import (
	"testing"
)

func Test_VdiskCreate(t *testing.T) {
	client := NewClient("", "")
	dpResponse, _, err := client.DataPool.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(dpResponse.Results) == 0 {
		t.SkipNow()
	}
	firstDp := dpResponse.Results[0]
	vdisk, _, err := client.Vdisk.Create(NameGenerator("vdisk"), false, firstDp.Id, 0.1, false)
	if err != nil {
		t.Error(err)
		return
	}
	vdisk, _, err = client.Vdisk.Get(vdisk.Id)
	if err != nil {
		t.Error(err)
		return
	}

	vdiskResponse, _, err := client.Vdisk.List(map[string]string{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(vdiskResponse.Results) == 0 {
		t.Errorf("no vdisks")
	}
	vdisk, _, err = client.Vdisk.Update(vdisk.Id, "test")
	if err != nil {
		t.Error(err)
		return
	}

	status, _, err := client.Vdisk.Remove(vdisk.Id)
	if err != nil {
		t.Error(err)
		return
	}

	if !status {
		t.Error(err)
		return
	}
	return
}

func Test_VdiskCreateAsync(t *testing.T) {
	client := NewClient("", "")
	dpResponse, _, err := client.DataPool.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(dpResponse.Results) == 0 {
		t.SkipNow()
	}
	firstDp := dpResponse.Results[0]
	vdisk, _, err := client.Vdisk.Create(NameGenerator("vdisk"), false, firstDp.Id, 0.1, true)
	if err != nil {
		t.Error(err)
		return
	}
	vdisk, _, err = client.Vdisk.Get(vdisk.Id)
	if err != nil {
		t.Error(err)
		return
	}

	vdiskResponse, _, err := client.Vdisk.List(map[string]string{})
	if err != nil {
		t.Error(err)
		return
	}
	if len(vdiskResponse.Results) == 0 {
		t.Errorf("no vdisks")
	}
	vdisk, _, err = client.Vdisk.Update(vdisk.Id, "test")
	if err != nil {
		t.Error(err)
		return
	}

	status, _, err := client.Vdisk.Remove(vdisk.Id)
	if err != nil {
		t.Error(err)
		return
	}

	if !status {
		t.Error(err)
		return
	}
	return
}
