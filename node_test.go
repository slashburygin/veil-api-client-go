package veil_api_client_go

import (
	"testing"
)

func Test_NodeListGet(t *testing.T) {
	client := NewClient("", "", false)
	response, _, err := client.Node.List()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range response.Results {
		node, _, err := client.Node.Get(v.Id)
		if err != nil {
			t.Error(err)
			return
		}

		if node.Id == "" {
			t.Error("Node Id can not be empty")
			return
		}
		break
	}

	return

}
