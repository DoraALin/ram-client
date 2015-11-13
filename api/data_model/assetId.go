package data_model

import (
	"fmt"
)

type AssetIdentification struct {
	guid    string
	version string
}

func NewAssetId(_guid, _version string) *AssetIdentification {
	return &AssetIdentification{
		guid:    _guid,
		version: _version,
	}
}

func (id *AssetIdentification) GetVersion() string {
	return id.version
}

func (id *AssetIdentification) GetGUID() string {
	return id.guid
}

func (id *AssetIdentification) String() string {
	return fmt.Sprintf("/%s/%s", id.guid, id.version)
}
