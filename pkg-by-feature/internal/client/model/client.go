package model

type Client struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

// Create Client DTOs
type CreateClientInput struct {
	Name string `json:"name"`
}

type CreateClientOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Get Client DTOs
type GetClientOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Update Client DTOs
type UpdateClientInput struct {
	Name string `json:"name"`
}

type UpdateClientOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List Clients DTOs
type ListClientsOutput struct {
	Clients []GetClientOutput `json:"clients"`
}

// Create Client methods
func (client Client) ToCreateClientOutput() CreateClientOutput {
	return CreateClientOutput{
		ID:   client.ID,
		Name: client.Name,
	}
}

func (input CreateClientInput) ToClient() Client {
	return Client{
		Name: input.Name,
	}
}

// Get Client methods
func (client Client) ToGetClientOutput() GetClientOutput {
	return GetClientOutput{
		ID:   client.ID,
		Name: client.Name,
	}
}

// Update Client methods
func (client Client) ToUpdateClientOutput() UpdateClientOutput {
	return UpdateClientOutput{
		ID:   client.ID,
		Name: client.Name,
	}
}

func (input UpdateClientInput) ToClient(id string) Client {
	return Client{
		ID:   id,
		Name: input.Name,
	}
}

// List Clients methods
func ClientsToListClientsOutput(clients []Client) ListClientsOutput {
	outputs := make([]GetClientOutput, len(clients))
	for i, client := range clients {
		outputs[i] = client.ToGetClientOutput()
	}
	return ListClientsOutput{
		Clients: outputs,
	}
}
