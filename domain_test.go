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
	config := new(DomainCreateConfig)
	config.DomainId = TestDomainID
	config.VerboseName = TestDomainName
	domain, _, err := client.Domain.Create(*config)
	assert.Nil(t, err)
	assert.NotEqual(t, domain.Id, "", "Domain Id can not be empty")

	return
}

func Test_DomainGet(t *testing.T) {
	client := NewClient("", "", false)

	domain, _, err := client.Domain.Get(TestDomainID)
	assert.Nil(t, err)
	assert.NotEqual(t, domain.Id, "", "Domain Id can not be empty")

	return
}

func Test_DomainPower(t *testing.T) {
	client := NewClient("", "", false)

	domain, _, err := client.Domain.Get(TestDomainID)
	assert.Nil(t, err)
	assert.NotEqual(t, domain.Id, "", "Domain Id can not be empty")
	domain, _, err = client.Domain.Start(domain)
	assert.Nil(t, err)
	domain, _, err = client.Domain.Suspend(domain)
	assert.Nil(t, err)
	domain, _, err = client.Domain.Resume(domain)
	assert.Nil(t, err)
	domain, _, err = client.Domain.Reboot(domain, true)
	assert.Nil(t, err)
	domain, _, err = client.Domain.Shutdown(domain, true)
	assert.Nil(t, err)
	domain, _, err = client.Domain.Template(domain, true)
	assert.Nil(t, err)
	assert.True(t, domain.Template)
	domain, _, err = client.Domain.Template(domain, false)
	assert.Nil(t, err)
	assert.False(t, domain.Template)

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
	assert.True(t, status)

	return
}
