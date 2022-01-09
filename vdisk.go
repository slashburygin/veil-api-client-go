package veil_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const baseVdiskUrl string = "/api/vdisks/"

type VdiskService struct {
	client Client
}

type NameTypeDataPool struct {
	Id          string `json:"id,omitempty"`
	VerboseName string `json:"verbose_name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type VdiskSnapshot struct {
	Id     string `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
	Vdisk  string `json:"vdisk,omitempty"`
	Serial string `json:"serial,omitempty"`
}

type VdiskObjectsList struct {
	Id          string           `json:"id,omitempty"`
	Status      string           `json:"name,omitempty"`
	VerboseName string           `json:"verbose_name,omitempty"`
	Size        float64          `json:"size,omitempty"`
	DataPool    NameTypeDataPool `json:"datapool,omitempty"`
	Domain      NameDomain       `json:"vdisk,omitempty"`
	Hints       int              `json:"hints,omitempty"`
	VirtualSize int              `json:"virtual_size,omitempty"`
}

type VdiskObject struct {
	Id           string           `json:"id,omitempty"`
	VerboseName  string           `json:"verbose_name,omitempty"`
	Description  string           `json:"description,omitempty"`
	LockedBy     string           `json:"locked_by,omitempty"`
	EntityType   string           `json:"entity_type,omitempty"`
	Status       string           `json:"name,omitempty"`
	Created      string           `json:"created,omitempty"`
	Modified     string           `json:"modified,omitempty"`
	ReadOnly     bool             `json:"readonly,omitempty"`
	VirtualSize  int              `json:"virtual_size,omitempty"`
	DataPool     NameTypeDataPool `json:"datapool,omitempty"`
	Size         float64          `json:"size,omitempty"`
	Domain       NameDomain       `json:"vdisk,omitempty"`
	DiskType     string           `json:"disk_type,omitempty"`
	Device       string           `json:"device,omitempty"`
	DriverType   string           `json:"driver_type,omitempty"`
	DriverCache  string           `json:"driver_cache,omitempty"`
	Source       string           `json:"source,omitempty"`
	Shareable    bool             `json:"shareable,omitempty"`
	Ssd          bool             `json:"ssd,omitempty"`
	TargetBus    string           `json:"target_bus,omitempty"`
	ActualSource string           `json:"actual_source,omitempty"`
	TargetDev    string           `json:"target_dev,omitempty"`
	Snapshots    []VdiskSnapshot  `json:"snapshots,omitempty"`
	Consolidated bool             `json:"consolidated,omitempty"`
	Hints        int              `json:"hints,omitempty"`
	Permissions  []string         `json:"permissions,omitempty"`
}

type VdisksResponse struct {
	BaseListResponse
	Results []VdiskObjectsList `json:"results,omitempty"`
}

func (d *VdiskService) List(queryParams map[string]string) (*VdisksResponse, *http.Response, error) {
	listUrl := baseVdiskUrl
	if len(queryParams) != 0 {
		params := url.Values{}
		for k, v := range queryParams {
			params.Add(k, v)
		}
		listUrl += "?"
		listUrl += params.Encode()
	}
	response := new(VdisksResponse)
	res, err := d.client.ExecuteRequest("GET", listUrl, []byte{}, response)
	return response, res, err
}

func (d *VdiskService) Get(Id string) (*VdiskObject, *http.Response, error) {

	vdisk := new(VdiskObject)
	res, err := d.client.ExecuteRequest("GET", fmt.Sprint(baseVdiskUrl, Id, "/"), []byte{}, vdisk)
	return vdisk, res, err
}

func (d *VdiskService) Create(verboseName string, preallocation bool,
	datapool string, size float64, asynced bool) (*VdiskObject, *http.Response, error) {

	vdisk := new(VdiskObject)

	body := struct {
		VerboseName   string  `json:"verbose_name,omitempty"`
		Datapool      string  `json:"datapool,omitempty"`
		Size          float64 `json:"size,omitempty"`
		Preallocation bool    `json:"preallocation,omitempty"`
	}{verboseName, datapool, size, preallocation}

	b, _ := json.Marshal(body)
	if !asynced {
		res, err := d.client.ExecuteRequest("POST", baseVdiskUrl, b, vdisk)
		return vdisk, res, err
	}
	asyncResp := new(AsyncResponse)
	res, err := d.client.ExecuteRequest("POST", baseVdiskUrl+"?async=1", b, asyncResp)
	WaitTaskReady(asyncResp.Task.Id, true, 0, true)
	res, err = d.client.ExecuteRequest("GET", fmt.Sprint(baseVdiskUrl, asyncResp.Entity, "/"), []byte{}, vdisk)
	return vdisk, res, err
}

func (d *VdiskService) Update(Id string, description string) (*VdiskObject, *http.Response, error) {

	vdisk := new(VdiskObject)

	body := struct {
		Description string `json:"description,omitempty"`
	}{description}

	b, _ := json.Marshal(body)

	res, err := d.client.ExecuteRequest("PUT", fmt.Sprint(baseVdiskUrl, Id, "/"), b, vdisk)

	return vdisk, res, err
}

func (d *VdiskService) Remove(Id string) (bool, *http.Response, error) {

	res, err := d.client.ExecuteRequest("POST", fmt.Sprint(baseVdiskUrl, Id, "/remove/"), []byte{}, nil)

	if err != nil {
		return false, res, err
	}

	return true, res, err
}
