package veil_api_client_go

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseDomainUrl string = "/api/domains/"

type DomainService struct {
	client Client
}

type NameNode struct {
	Id          string `json:"id,omitempty"`
	VerboseName string `json:"verbose_name,omitempty"`
}

type GuestUtils struct {
	Id          string `json:"id,omitempty"`
	VerboseName string `json:"verbose_name,omitempty"`
}

type DomainObjectsList struct {
	Id                 string           `json:"id,omitempty"`
	VerboseName        string           `json:"verbose_name,omitempty"`
	MemoryCount        int              `json:"memory_count,omitempty"`
	Status             string           `json:"name,omitempty"`
	Parent             NameDomain       `json:"parent,omitempty"`
	CpuCount           int              `json:"cpu_count,omitempty"`
	MemoryPool         string           `json:"memory_pool,omitempty"`
	VmachineInfsCount  int              `json:"vmachine_infs_count,omitempty"`
	VdisksCount        int              `json:"vdisks_count,omitempty"`
	VfunctionsCount    int              `json:"vfunctions_count,omitempty"`
	LunsCount          int              `json:"luns_count,omitempty"`
	UserPowerState     int              `json:"user_power_state,omitempty"`
	Node               NameNode         `json:"node,omitempty"`
	Template           bool             `json:"template,omitempty"`
	MdevsCount         int              `json:"mdevs_count,omitempty"`
	Tags               []Tags           `json:"tags,omitempty"`
	Hints              int              `json:"hints,omitempty"`
	ResourcePool       NameResourcePool `json:"resource_pool,omitempty"`
	GuestUtils         GuestUtils       `json:"guest_utils,omitempty"`
	Thin               bool             `json:"thin,omitempty"`
	Replication        bool             `json:"replication,omitempty"`
	CpuUsedPercentUser string           `json:"cpu_used_percent_user,omitempty"`
	MemUsedPercentUser string           `json:"mem_used_percent_user,omitempty"`
	Priority           int              `json:"priority,omitempty"`
}

type DomainObject struct {
	Id                 string           `json:"id,omitempty"`
	VerboseName        string           `json:"verbose_name,omitempty"`
	Description        string           `json:"description,omitempty"`
	LockedBy           string           `json:"locked_by,omitempty"`
	Permissions        []string         `json:"permissions,omitempty"`
	Created            string           `json:"created,omitempty"`
	Modified           string           `json:"modified,omitempty"`
	MemoryCount        int              `json:"memory_count,omitempty"`
	Status             string           `json:"name,omitempty"`
	Parent             NameDomain       `json:"parent,omitempty"`
	CpuCount           int              `json:"cpu_count,omitempty"`
	MemoryPool         string           `json:"memory_pool,omitempty"`
	VmachineInfsCount  int              `json:"vmachine_infs_count,omitempty"`
	VdisksCount        int              `json:"vdisks_count,omitempty"`
	VfunctionsCount    int              `json:"vfunctions_count,omitempty"`
	LunsCount          int              `json:"luns_count,omitempty"`
	UserPowerState     int              `json:"user_power_state,omitempty"`
	Node               NameNode         `json:"node,omitempty"`
	Template           bool             `json:"template,omitempty"`
	MdevsCount         int              `json:"mdevs_count,omitempty"`
	Tags               []Tags           `json:"tags,omitempty"`
	Hints              int              `json:"hints,omitempty"`
	ResourcePool       NameResourcePool `json:"resource_pool,omitempty"`
	GuestUtils         GuestUtils       `json:"guest_utils,omitempty"`
	Thin               bool             `json:"thin,omitempty"`
	Replication        bool             `json:"replication,omitempty"`
	CpuUsedPercentUser string           `json:"cpu_used_percent_user,omitempty"`
	MemUsedPercentUser string           `json:"mem_used_percent_user,omitempty"`
	Priority           int              `json:"priority,omitempty"`
}

type DomainsResponse struct {
	BaseListResponse
	Results []DomainObjectsList `json:"results,omitempty"`
}

type DomainCreate struct {
	VerboseName string `json:"verbose_name,omitempty"`
	Description string `json:"description,omitempty"`
	Node        string `json:"node,omitempty"`
	CpuCount    string `json:"cpu_count,omitempty"`
	MemoryCount int    `json:"memory_count,omitempty"`
}

func (d *DomainService) List() (*DomainsResponse, *http.Response, error) {

	response := new(DomainsResponse)

	res, err := d.client.ExecuteRequest("GET", baseDomainUrl, []byte{}, response)

	return response, res, err
}

func (d *DomainService) Create(verboseName string, domainID string) (*DomainObject, *http.Response, error) {

	domain := new(DomainObject)

	body := struct {
		VerboseName string `json:"verbose_name,omitempty"`
		DomainId    string `json:"domain_id,omitempty"`
		MemoryCount int    `json:"memory_count,omitempty"`
	}{verboseName, domainID, 50}

	b, _ := json.Marshal(body)
	res, err := d.client.ExecuteRequest("POST", baseDomainUrl, b, domain)

	return domain, res, err
}

func (d *DomainService) Get(domainID string) (*DomainObject, *http.Response, error) {

	domain := new(DomainObject)

	res, err := d.client.ExecuteRequest("GET", fmt.Sprint(baseDomainUrl, domainID, "/"), []byte{}, domain)

	return domain, res, err
}

func (d *DomainService) Update(domainID string, description string) (*DomainObject, *http.Response, error) {

	domain := new(DomainObject)

	body := struct {
		Description string `json:"description,omitempty"`
	}{description}

	b, _ := json.Marshal(body)

	res, err := d.client.ExecuteRequest("PUT", fmt.Sprint(baseDomainUrl, domainID, "/"), b, domain)

	return domain, res, err
}

func (d *DomainService) Start(domain *DomainObject) (*DomainObject, *http.Response, error) {
	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseDomainUrl, domain.Id, "/start/"), []byte{}, domain)
	return domain, res, err
}

func (d *DomainService) Suspend(domain *DomainObject) (*DomainObject, *http.Response, error) {
	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseDomainUrl, domain.Id, "/suspend/"), []byte{}, domain)
	return domain, res, err
}

func (d *DomainService) Resume(domain *DomainObject) (*DomainObject, *http.Response, error) {
	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseDomainUrl, domain.Id, "/resume/"), []byte{}, domain)
	return domain, res, err
}

func (d *DomainService) Shutdown(domain *DomainObject, force bool) (*DomainObject, *http.Response, error) {
	body := struct {
		Force bool `json:"force,omitempty"`
	}{force}
	b, _ := json.Marshal(body)
	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseDomainUrl, domain.Id, "/shutdown/"), b, domain)
	return domain, res, err
}

func (d *DomainService) Reboot(domain *DomainObject, force bool) (*DomainObject, *http.Response, error) {
	body := struct {
		Force bool `json:"force,omitempty"`
	}{force}
	b, _ := json.Marshal(body)
	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseDomainUrl, domain.Id, "/reboot/"), b, domain)
	return domain, res, err
}

func (d *DomainService) Template(domain *DomainObject, template bool) (*DomainObject, *http.Response, error) {
	body := struct {
		Template bool `json:"template"`
	}{template}
	b, _ := json.Marshal(body)
	res, err := d.client.ExecuteRequest("PUT", fmt.Sprint(baseDomainUrl, domain.Id, "/template/"), b, domain)
	return domain, res, err
}

func (d *DomainService) Remove(domainID string) (bool, *http.Response, error) {

	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseDomainUrl, domainID, "/remove/"), []byte{}, nil)

	if err != nil {
		return false, res, err
	}

	return true, res, err
}
