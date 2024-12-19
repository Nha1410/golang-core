package pet

type Repository struct {
}

func NewRepository() RepositoryInterface {
	return &Repository{}
}
