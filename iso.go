package veil_api_client_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const baseIsoUrl string = "/api/iso/"

type IsoService struct {
	client Client
}

type IsoObjectsList struct {
	Id       string           `json:"id,omitempty"`
	Status   string           `json:"name,omitempty"`
	FileName string           `json:"filename,omitempty"`
	Size     float64          `json:"size,omitempty"`
	DataPool NameTypeDataPool `json:"datapool,omitempty"`
	Domains  []NameDomain     `json:"domains,omitempty"`
	Created  string           `json:"created,omitempty"`
}

type IsoObject struct {
	Id          string           `json:"id,omitempty"`
	FileName    string           `json:"filename,omitempty"`
	Description string           `json:"description,omitempty"`
	LockedBy    string           `json:"locked_by,omitempty"`
	EntityType  string           `json:"entity_type,omitempty"`
	Status      string           `json:"name,omitempty"`
	Created     string           `json:"created,omitempty"`
	Modified    string           `json:"modified,omitempty"`
	DataPool    NameTypeDataPool `json:"datapool,omitempty"`
	Domains     []NameDomain     `json:"domains,omitempty"`
	Size        float64          `json:"size,omitempty"`
	Path        string           `json:"path,omitempty"`
	Permissions []string         `json:"permissions,omitempty"`
	UploadUrl   string           `json:"upload_url,omitempty"`
}

type IsosResponse struct {
	BaseListResponse
	Results []IsoObjectsList `json:"results,omitempty"`
}

func (d *IsoService) List() (*IsosResponse, *http.Response, error) {

	response := new(IsosResponse)

	res, err := d.client.ExecuteRequest("GET", baseIsoUrl, []byte{}, response)

	return response, res, err
}

func (d *IsoService) Get(Id string) (*IsoObject, *http.Response, error) {

	node := new(IsoObject)

	res, err := d.client.ExecuteRequest("GET", fmt.Sprint(baseIsoUrl, Id, "/"), []byte{}, node)

	return node, res, err
}

func (d *IsoService) Create(DataPoolId string, FileName string) (*IsoObject, *http.Response, error) {
	// Part 1
	iso := new(IsoObject)

	body := struct {
		DataPoolId string `json:"datapool,omitempty"`
		FileName   string `json:"filename,omitempty"`
	}{DataPoolId, FileName}

	b, _ := json.Marshal(body)
	res, err := d.client.ExecuteRequest("PUT", baseIsoUrl, b, iso)
	if err != nil {
		return nil, res, err
	}

	// Part 2
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/file_data/" + FileName)
	if err != nil {
		return iso, res, err
	}
	defer file.Close()

	fileBody := &bytes.Buffer{}
	writer := multipart.NewWriter(fileBody)
	part, err := writer.CreateFormFile("file", filepath.Base(FileName))
	if err != nil {
		return nil, res, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, res, err
	}
	request, err := http.NewRequest("POST", fmt.Sprint("http://192.168.11.105", iso.UploadUrl), fileBody)
	err = writer.Close()
	if err != nil {
		return nil, res, err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	return iso, response, err
}
