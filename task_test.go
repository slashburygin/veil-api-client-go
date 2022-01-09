package veil_api_client_go

import (
	"testing"
)

func TestTask(t *testing.T) {
	client := NewClient("", "", false)
	response, _, err := client.Task.List()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range response.Results {
		task, _, err := client.Task.Get(v.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if task.Id == "" {
			t.Error("Task Id can not be empty")
			return
		}
		break
	}

	return

}
