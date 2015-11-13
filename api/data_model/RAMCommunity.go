package data_model

import (
	"fmt"
	ram_utils "github.com/DoraALin/ram_api_go/api"
)

type RAMCommunity struct {
	session *RAMSession
	url     string
	name    string
	desc    string
	id      int
	//lazyfetch to decide if community should be retrieved when community is initialized
	lazy bool
}

func NewRAMCommunity(_session *RAMSession, _url string, _lazy bool) *RAMCommunity {
	com := &RAMCommunity{
		session: _session,
		url:     _url,
		lazy:    _lazy,
	}
	if !com.lazy {
		com.fetchCommunityInfo()
	}

	return com
}

//internal function to fetch the community information from url
func (com *RAMCommunity) fetchCommunityInfo() {
	if com.session == nil {
		//error
	}

	resp, err := com.session.Get(com.url, ram_utils.OSLC_JSON_HEADER, ram_utils.RESPONSE_FORMAT["json"])
	if err != nil {

	}
	r_map := resp.(map[string]interface{})
	for k, v := range r_map {
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
	//	com.name = r_map["name"]
	//	com.desc = r_map["description"]
	//	com.id = r_map["dbid"]
}
