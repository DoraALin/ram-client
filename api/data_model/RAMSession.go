package data_model

import (
	"fmt"
	api "github.com/DoraALin/ram_api_go/api"
	ram_client "github.com/DoraALin/ram_api_go/api/ram_client"
	"io"
	"log"
	"net/http"
)

type RAMSession struct {
	repo  *Repository
	url   string
	uid   string
	pwd   string
	valid string
}

// FromParameters constructs a new Driver with a given parameters map.
func NewRAMSession(_url, _uid, _pwd string) (*RAMSession, error) {
	fmt.Printf("NewRAMSession %s:%s\n", _uid, _pwd)
	ram := &RAMSession{
		url:  _url,
		repo: NewRepoInfo(),
		uid:  _uid,
		pwd:  _pwd,
	}
	ram.validate()
	return ram, nil
}

// New constructs a new Driver with the given Azure Storage Account credentials
func (ram *RAMSession) validate() error {
	//the url passed in is supposed to be RAM war url
	client := &ram_client.RAMClient{}
	repo_info := ram.url + api.REPO_URL
	_, err := client.Get(repo_info, ram.uid, ram.pwd, map[string]string{
		"accept": "application/json",
	}, api.M_GET)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//get content represented by url, any function call this function need to close Response body
func (ram *RAMSession) GetContent(url string, header map[string]string) (io.ReadCloser, error) {
	client := &ram_client.RAMClient{}
	resp, err := client.GetContent(url, ram.uid, ram.pwd, header)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return resp.Body, err
}

const repository_prefix = "repositories"
const artifactContentUrl_pattern = "%s/oslc/assets/%s/%s/artifactContents"
const artifactUrl_pattern = "%s/oslc/assets/%s/%s/artifacts"

// locate path into asset artifact url
func (ram *RAMSession) locateArtifactPath(id *AssetIdentification, subPath string) string {
	//	idx := strings.Index(subPath, repository_prefix) + len(repository_prefix)
	//	artifactPath := subPath[idx:]
	basePath := fmt.Sprintf(artifactContentUrl_pattern, ram.url, id.GetGUID(), id.GetVersion())
	//	fmt.Printf("basepath: %s\n", basePath)
	url := basePath + subPath
	//	fmt.Printf("url: %s\n", url)
	return url
}

func (ram *RAMSession) createArtifactPath(id *AssetIdentification, subPath string) string {
	basePath := fmt.Sprintf(artifactUrl_pattern, ram.url, id.GetGUID(), id.GetVersion())
	//fmt.Printf("basepath: %s\n", basePath)
	if len(subPath) > 0 {
		basePath = basePath + subPath
	}
	return basePath
}

/**
do PUT to update a new artifact in specified subPath for asset with id as identification
*/
func (ram *RAMSession) PutArtifactContent(id *AssetIdentification, subPath string, header map[string]string, rc io.ReadCloser) (int, error) {
	client := &ram_client.RAMClient{}
	header["name"] = subPath
	fmt.Printf("PUT: %s\n", subPath)
	return client.PutContent(ram.createArtifactPath(id, ""), ram.uid, ram.pwd, header, "" /*there is no meta data*/, rc)
}

func (ram *RAMSession) HeadArtifactContent(id *AssetIdentification, subPath string, header map[string]string) (int, *http.Header, error) {
	client := &ram_client.RAMClient{}
	resp, err := client.Head(ram.locateArtifactPath(id, subPath), ram.uid, ram.pwd, header)
	defer resp.Body.Close()
	return resp.StatusCode, &resp.Header, err
}

/**
return size(int64) and name(string) of artifact stored in header of a head request
*/
func (ram *RAMSession) GetArtifactInfo(header *http.Header) *ArtifactInfo {
	return NewArtifactInfo(header)
}

func (ram *RAMSession) GetArtifactContent(id *AssetIdentification, subPath string, header map[string]string) (io.ReadCloser, error) {
	client := &ram_client.RAMClient{}
	resp, err := client.GetContent(ram.locateArtifactPath(id, subPath), ram.uid, ram.pwd, header)
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return resp.Body, err
}

/**
do POST to create a new artifact in specified subPath for asset with id as identification
*/
func (ram *RAMSession) PostArtifactContent(id *AssetIdentification, subPath string, header map[string]string, rc io.ReadCloser) (int, error) {
	client := &ram_client.RAMClient{}
	header["name"] = subPath
	fmt.Printf("POST: %s\n", subPath)
	return client.Post(ram.createArtifactPath(id, ""), ram.uid, ram.pwd, header, "" /*there is no meta data*/, rc)
}

func (ram *RAMSession) DeleteArtifactContent(id *AssetIdentification, subPath string, header map[string]string) (int, error) {
	return ram.Delete(ram.createArtifactPath(id, subPath), header, "")
}

func (ram *RAMSession) Delete(url string, header map[string]string, format string) (int, error) {
	client := &ram_client.RAMClient{}
	return client.Delete(url, ram.uid, ram.pwd, header)
}

func (ram *RAMSession) Get(url string, header map[string]string, format string) (interface{}, error) {
	client := &ram_client.RAMClient{}
	resp, err := client.Get(url, ram.uid, ram.pwd, header, api.M_GET)
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

func (ram *RAMSession) GetAsset(asset_id *AssetIdentification) (interface{}, error) {
	//the url passed in is supposed to be RAM war url
	oslc_asset_service_url := ram.url + api.OSLC_BASE + api.OSLC_ASSETS + asset_id.String()
	client := &ram_client.RAMClient{}
	resp, err := client.Get(oslc_asset_service_url, ram.uid, ram.pwd, api.OSLC_JSON_HEADER, api.M_GET)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}
