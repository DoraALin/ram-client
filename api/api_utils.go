package ram_utils

const RAM_API_VER = "7.5.2.4"
const REPO_URL = "/internal/repository"
const OSLC_BASE = "/oslc"
const OSLC_ASSETS = "/assets"

const (
	M_GET    = "GET"
	M_PUT    = "PUT"
	M_DELETE = "DELETE"
	M_POST   = "POST"
	M_HEAD   = "HEAD"
)

var OSLC_JSON_HEADER = map[string]string{
	"OSLC-Core-Version": "2.0",
	"accept":            "application/json",
}

var OSLC_JSON_HEADER_V1 = map[string]string{
	"OSLC-Core-Version": "1.0",
	"accept":            "application/json",
}

var RESPONSE_FORMAT = map[string]string{
	"json": "application/json",
}
