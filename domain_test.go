package veil_api_client_go

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestDomainName = NameGenerator("domain")
var TestDomainID = uuid.NewString()

func Test_DomainList(t *testing.T) {
	client := NewClient("", "", false)
	_, _, err := client.Domain.List()
	assert.Nil(t, err)

	return

}

func Test_DomainCreate(t *testing.T) {
	client := NewClient("", "", false)
	domain, _, err := client.Domain.Create(TestDomainName, TestDomainID)
	assert.Nil(t, err)

	if domain.Id == "" {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func Test_DomainGet(t *testing.T) {
	client := NewClient("", "", false)
	domain, _, err := client.Domain.Get(TestDomainID)
	assert.Nil(t, err)

	if domain.Id == "" {
		t.Error("Domain ID can not be empty")
		return
	}

	return

}

func Test_DomainUpdate(t *testing.T) {
	client := NewClient("", "", false)
	domain, _, err := client.Domain.Update(TestDomainID, "test")
	assert.Nil(t, err)
	assert.NotEqual(t, domain.Id, "", "Domain Id can not be empty")

	return

}

func Test_DomainRemove(t *testing.T) {
	client := NewClient("", "", false)
	status, _, err := client.Domain.Remove(TestDomainID)
	assert.Nil(t, err)

	if !status {
		t.Fail()
		return
	}

	return

}
