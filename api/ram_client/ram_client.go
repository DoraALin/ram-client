package ram_client

import (
	"encoding/json"
	"fmt"
	ram_utils "github.com/DoraALin/ram_api_go/api"
	"io"
	"io/ioutil"
	"log"
	http "net/http"
	"strings"
)

type AbstractClient interface {
	do_request(url_str string, uid string, pwd string, header map[string]string, method string) (*http.Response, error)
	Get(url string, uid string, pwd string, header map[string]string, format string) (interface{}, error)
	Delete(utl string, uid string, pwd string, header map[string]string) (int, error)
	Post(url string, uid string, pwd string, header map[string]string, format string, rc io.ReadCloser) (int, error)
	HandleResp(resp *http.Response, r interface{}, format string) error
}

type RAMClient struct {
}

func (ram *RAMClient) GetContent(url string, uid string, pwd string, header map[string]string) (io.ReadCloser, error) {
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_GET, nil)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body, err
}

func (ram *RAMClient) Head(url string, uid string, pwd string, header map[string]string) (*http.Response, error) {
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_HEAD, nil)
	if err != nil {
		log.Fatal(err)
	}

	//return code to indicate result
	return resp, err
}

func (ram *RAMClient) Delete(url string, uid string, pwd string, header map[string]string) (int, error) {
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_DELETE, nil)
	if err != nil {
		log.Fatal(err)
	}

	//return code to indicate result
	return resp.StatusCode, err
}

/**
Put method to update asset resource, with _map as request Body. for updating artifact using stream, refer to
func (ram *RAMClient) PutContent(url string, uid string, pwd string, header map[string]string, format string, rc io.ReadCloser) (int, error)
*/
func (ram *RAMClient) Put(url string, uid string, pwd string, header map[string]string, format string, _map map[string]string) (int, error) {
	//the url passed in is supposed to be RAM ws url
	//TODO: convert _map into io.Reader
	body, err := json.Marshal(_map)
	rc := ioutil.NopCloser(strings.NewReader(string(body)))
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_PUT, rc)
	if err != nil {
		log.Fatal(err)
	}

	//return code to indicate result
	return resp.StatusCode, err
}

func (ram *RAMClient) PutContent(url string, uid string, pwd string, header map[string]string, format string, rc io.ReadCloser) (int, error) {
	//the url passed in is supposed to be RAM ws url
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_PUT, rc)
	if err != nil {
		log.Fatal(err)
	}

	//return code to indicate result
	return resp.StatusCode, err
}

func (ram *RAMClient) Post(url string, uid string, pwd string, header map[string]string, format string, rc io.ReadCloser) (int, error) {
	//the url passed in is supposed to be RAM ws url
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_POST, rc)
	if err != nil {
		log.Fatal(err)
	}

	//return code to indicate result
	return resp.StatusCode, err
}

/**
perform GET request to url, given uid and pwd for authentication, function return interface for response, which means
the response could be parsed in to json(or otherstructured data, maybe), if one would like to fetch stream as response,
refer to func (ram * RAMClient) GetContent(url string, uid string, pwd string, header map[string]string) (io.ReadCloser, error)
*/
func (ram *RAMClient) Get(url string, uid string, pwd string, header map[string]string, format string) (interface{}, error) {
	//the url passed in is supposed to be RAM war url
	resp, err := ram.do_request(url, uid, pwd, header, ram_utils.M_GET, nil)
	if err != nil {
		log.Fatal(err)
	}
	var r interface{}
	ram.HandleResp(resp, &r, format)
	repo_map := r.(map[string]interface{})

	for k, v := range repo_map {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	return repo_map, nil
}

//internal get function to perform GET oslc servies
func (ram *RAMClient) do_request(url_str string, uid string, pwd string, header map[string]string, method string, body io.ReadCloser) (*http.Response, error) {
	//the url passed in is supposed to be RAM war url
	req, err := http.NewRequest(method, url_str, body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//add request header
	for k, v := range header {
		req.Header.Add(k, v)
	}
	//add authentication
	req.SetBasicAuth(uid, pwd)

	fmt.Printf("Request to : %s\n", url_str)
	for k, v := range req.Header {
		fmt.Printf("%s: %s\n", k, v)
	}

	http_client := &http.Client{}
	resp, err := http_client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resp, nil
}

func (ram *RAMClient) HandleResp(resp *http.Response, r interface{}, format string) error {
	fmt.Printf("Response returned with code: %i\n", resp.StatusCode)
	defer resp.Body.Close()
	con, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", con)
	if err != nil {
		return err
	}
	switch format {
	//default to handle json
	default:
		err = json.Unmarshal(con, r)
	}
	if err != nil {
		return err
	}

	return nil
}
