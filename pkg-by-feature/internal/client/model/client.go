package model

type Client struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type CreateClientInput struct {
	Name string `json:"name"`
}

type CreateClientOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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
