package veil_api_client_go

import (
	"github.com/google/uuid"
	"testing"
)

var TestDomainName = NameGenerator("domain")
var TestDomainID = uuid.NewString()

func Test_DomainList(t *testing.T) {
	client := NewClient("", "", false)
	_, _, err := client.Domain.List()
	if err != nil {
		t.Error(err)
		return
	}

	return

}

func Test_DomainCreate(t *testing.T) {
	client := NewClient("", "", false)
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

func Test_DomainGet(t *testing.T) {
	client := NewClient("", "", false)
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

func Test_DomainUpdate(t *testing.T) {
	client := NewClient("", "", false)
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

func Test_DomainRemove(t *testing.T) {
	client := NewClient("", "", false)
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
