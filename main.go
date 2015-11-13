package main

import (
	ram_data "github.com/DoraALin/ram_api_go/api/data_model"
	"github.com/docker/docker/vendor/src/github.com/Sirupsen/logrus"

	"fmt"
	"github.com/DoraALin/ram_api_go/api"
	"io/ioutil"
	"os"
)

func main() {
	ram_session, err := ram_data.NewRAMSession("http://9.110.74.108:9080/ram.ws7.5.2.3", "admin", "admin")
	if err != nil {
		logrus.Fatal(err)
	}
	asset_example := ram_data.NewAssetId("F5656860-9D7D-D8AA-2494-765DAD160E68", "1.0")
	ram_session.GetAsset(asset_example)
	header := ram_utils.OSLC_JSON_HEADER
	//argh... no use
	//header["Range"]="bytes=2"
	is, err := ram_session.GetArtifactContent(asset_example, "/test.txt", header)
	defer is.Close()
	if err != nil {

	}
	ouput, err := ioutil.ReadAll(is)

	ioutil.WriteFile("C:/test.txt", ouput, os.FileMode(0666))
	fmt.Printf("%s", string(ouput))

	//try post a file
	file, _ := os.OpenFile("C:/test.txt", os.O_RDONLY, 0644)
	code, _ := ram_session.PostArtifactContent(asset_example, "/base/test1.txt", header, file)
	fmt.Printf("POST return: %d", code)
	code, resp_header, _ := ram_session.HeadArtifactContent(asset_example, "/base/test1.txt", ram_utils.OSLC_JSON_HEADER_V1)
	file_name, file_size := ram_session.GetArtifactInfo(resp_header)
	fmt.Printf("HEAD return: %d", code)
	fmt.Printf("ArtifactInfo name: %s, size: %d", file_name, file_size)
	file, _ = os.OpenFile("C:/test.txt", os.O_RDONLY, 0644)
	code, _ = ram_session.PutArtifactContent(asset_example, "/base/test1.txt", header, file)
	fmt.Printf("Put return: %d", code)
	code, _ = ram_session.DeleteArtifactContent(asset_example, "/base/test1.txt", ram_utils.OSLC_JSON_HEADER_V1)
	fmt.Printf("DELETE return: %d", code)
}
