package data_model

type Repository struct {
	webServerPath   string
	webServicesPath string
	version         string
	identifier      string
}

func NewRepoInfo() *Repository {
	return &Repository{}
}
