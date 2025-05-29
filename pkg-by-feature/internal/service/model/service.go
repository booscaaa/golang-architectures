package model

type Service struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int    `db:"price"`
}

// Create Service types
type CreateServiceInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type CreateServiceOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Get Service types
type GetServiceOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Update Service types
type UpdateServiceInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type UpdateServiceOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// List Services types
type ListServicesOutput struct {
	Services []ServiceItem `json:"services"`
}

type ServiceItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// Conversion methods
func (service Service) ToCreateServiceOutput() CreateServiceOutput {
	return CreateServiceOutput{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Price:       service.Price,
	}
}

func (service Service) ToGetServiceOutput() GetServiceOutput {
	return GetServiceOutput{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Price:       service.Price,
	}
}

func (service Service) ToUpdateServiceOutput() UpdateServiceOutput {
	return UpdateServiceOutput{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Price:       service.Price,
	}
}

func (service Service) ToServiceItem() ServiceItem {
	return ServiceItem{
		ID:          service.ID,
		Name:        service.Name,
		Description: service.Description,
		Price:       service.Price,
	}
}

func (input CreateServiceInput) ToService() Service {
	return Service{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
}

func (input UpdateServiceInput) ToService(id string) Service {
	return Service{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
}

func ServicesToListServicesOutput(services []Service) ListServicesOutput {
	serviceItems := make([]ServiceItem, len(services))
	for i, service := range services {
		serviceItems[i] = service.ToServiceItem()
	}
	return ListServicesOutput{
		Services: serviceItems,
	}
}
