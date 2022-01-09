package veil_api_go

import (
	"github.com/google/uuid"
	"testing"
)

var TestDomainName = NameGenerator("domain")
var TestDomainID = uuid.NewString()

func TestDomainService_List(t *testing.T) {
	client := NewClient("", "")
	_, _, err := client.Domain.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func TestDomainService_Create(t *testing.T) {
	client := NewClient("", "")
	domain, _, err := client.Domain.Create(TestDomainName, TestDomainID)
	if err != nil {
		t.Error(err)
		return
	}

	if domain.Id == "" {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func TestDomainService_Get(t *testing.T) {
	client := NewClient("", "")
	domain, _, err := client.Domain.Get(TestDomainID)
	if err != nil {
		t.Error(err)
		return
	}

	if domain.Id == "" {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func TestDomainService_Update(t *testing.T) {
	client := NewClient("", "")
	domain, _, err := client.Domain.Update(TestDomainID, "test")
	if err != nil {
		t.Error(err)
		return
	}

	if domain.Id == "" {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func TestDomainService_Remove(t *testing.T) {
	client := NewClient("", "")
	status, _, err := client.Domain.Remove(TestDomainID)
	if err != nil {
		t.Error(err)
		return
	}

	if !status {
		t.Fail()
		return
	}

	return

}
