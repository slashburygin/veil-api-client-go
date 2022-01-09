package veil_api_go

import (
	"testing"
)

func TestNodeService_List_Get(t *testing.T) {
	client := NewClient("", "")
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
