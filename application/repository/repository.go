package repository

type ExampleRepository interface {
	FetchData() string
}

type exampleRepositoryImpl struct{}

func NewExampleRepository() ExampleRepository {
	return &exampleRepositoryImpl{}
}

func (r *exampleRepositoryImpl) FetchData() string {
	return "Data from Repository"
}
