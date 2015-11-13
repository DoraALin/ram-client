package data_model

import ()

type RAMAsset struct {
	id        *AssetIdentification
	repo_info *Repository
	asset_com *RAMCommunity
	//asset_type *RAMAssetType
	s_desc string

	//lifecycle properties
	//state *AssetState
}

func NewRAMAsset() *RAMAsset {
	return &RAMAsset{}
}
