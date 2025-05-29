package model

type Service struct {
	ID   string `db:"id"`
	Name string `db:"name"`
	// Adicione outros campos conforme necessário
}

type CreateServiceInput struct {
	Name string `json:"name"`
	// Adicione outros campos conforme necessário
}

type CreateServiceOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Adicione outros campos conforme necessário
}

func (input *CreateServiceInput) ToService() Service {
	return Service{
		Name: input.Name,
		// Mapeie outros campos conforme necessário
	}
}

func (service Service) ToCreateServiceOutput() CreateServiceOutput {
	return CreateServiceOutput{
		ID:   service.ID,
		Name: service.Name,
		// Mapeie outros campos conforme necessário
	}
}