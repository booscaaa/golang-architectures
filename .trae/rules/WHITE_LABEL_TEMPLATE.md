# Template White Label - Clean Architecture Go

Este é um template white label baseado na arquitetura limpa (Clean Architecture) para projetos Go, seguindo o padrão de organização por features.

## Estrutura do Projeto

```
.
├── Dockerfile.dev
├── cmd
│   ├── cmd
│   │   ├── root.go
│   │   └── serve.go
│   └── main.go
├── database
│   └── migrations
│       └── 00001_init.up.sql
├── dev
│   ├── air.log
│   └── main
├── go.mod
├── go.sum
├── internal
│   ├── {module_name}
│   │   ├── contract
│   │   │   └── {module_name}.go
│   │   ├── controller
│   │   │   └── create_{module_name}.go
│   │   ├── model
│   │   │   └── {module_name}.go
│   │   ├── repository
│   │   │   └── create_{module_name}_database.go
│   │   └── usecase
│   │       └── create_{module_name}.go
└── pkg
    ├── di
    │   └── {module_name}.go
    └── rest
        └── rest.go
```

## Templates de Código

### 1. Model Template

**Arquivo:** `internal/{module_name}/model/{module_name}.go`

```go
package model

type {ModuleName} struct {
	ID   string `db:"id"`
	Name string `db:"name"`
	// Adicione outros campos conforme necessário
}

type Create{ModuleName}Input struct {
	Name string `json:"name"`
	// Adicione outros campos conforme necessário
}

type Create{ModuleName}Output struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Adicione outros campos conforme necessário
}

func (input *Create{ModuleName}Input) To{ModuleName}() {ModuleName} {
	return {ModuleName}{
		Name: input.Name,
		// Mapeie outros campos conforme necessário
	}
}

func ({module_name} {ModuleName}) ToCreate{ModuleName}Output() Create{ModuleName}Output {
	return Create{ModuleName}Output{
		ID:   {module_name}.ID,
		Name: {module_name}.Name,
		// Mapeie outros campos conforme necessário
	}
}
```

### 2. Contract Template

**Arquivo:** `internal/{module_name}/contract/{module_name}.go`

```go
package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/model"
)

type Create{ModuleName}DatabaseRepository interface {
	Execute(context.Context, model.{ModuleName}) (model.{ModuleName}, error)
}

type Create{ModuleName}UseCase interface {
	Execute(context.Context, model.Create{ModuleName}Input) (model.Create{ModuleName}Output, error)
}

type Create{ModuleName}Controller interface {
	Execute(http.ResponseWriter, *http.Request)
}
```

### 3. Controller Template

**Arquivo:** `internal/{module_name}/controller/create_{module_name}.go`

```go
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/model"
)

type create{ModuleName}Controller struct {
	create{ModuleName}UseCase contract.Create{ModuleName}UseCase
}

// Execute implements contract.Create{ModuleName}Controller.
func (controller *create{ModuleName}Controller) Execute(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	var input model.Create{ModuleName}Input
	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.create{ModuleName}UseCase.Execute(ctx, input)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(input)

	err = json.NewEncoder(response).Encode(output)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewCreate{ModuleName}Controller(create{ModuleName}UseCase contract.Create{ModuleName}UseCase) contract.Create{ModuleName}Controller {
	return &create{ModuleName}Controller{
		create{ModuleName}UseCase: create{ModuleName}UseCase,
	}
}
```

### 4. UseCase Template

**Arquivo:** `internal/{module_name}/usecase/create_{module_name}.go`

```go
package usecase

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/model"
)

type create{ModuleName}UseCase struct {
	repository contract.Create{ModuleName}DatabaseRepository
}

// Execute implements contract.Create{ModuleName}UseCase.
func (usecase *create{ModuleName}UseCase) Execute(ctx context.Context, input model.Create{ModuleName}Input) (model.Create{ModuleName}Output, error) {
	{module_name}, err := usecase.repository.Execute(ctx, input.To{ModuleName}())
	if err != nil {
		return model.Create{ModuleName}Output{}, err
	}
	return {module_name}.ToCreate{ModuleName}Output(), nil
}

func NewCreate{ModuleName}UseCase(repository contract.Create{ModuleName}DatabaseRepository) contract.Create{ModuleName}UseCase {
	return &create{ModuleName}UseCase{
		repository: repository,
	}
}
```

### 5. Repository Template

**Arquivo:** `internal/{module_name}/repository/create_{module_name}_database.go`

```go
package repository

import (
	"context"

	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/model"
	"github.com/booscaaa/initializers/postgres"
)

type create{ModuleName}DatabaseRepository struct {
	database postgres.Database
}

// Execute implements contract.Create{ModuleName}DatabaseRepository.
func (repository *create{ModuleName}DatabaseRepository) Execute(ctx context.Context, {module_name} model.{ModuleName}) (model.{ModuleName}, error) {
	var {module_name}Created model.{ModuleName}
	query := "INSERT INTO {module_name} (name) VALUES ($1) RETURNING *"
	err := repository.database.QueryRowxContext(
		ctx,
		query,
		{module_name}.Name,
	).StructScan(&{module_name}Created)
	if err != nil {
		return model.{ModuleName}{}, err
	}

	return {module_name}Created, nil
}

func NewCreate{ModuleName}DatabaseRepository(database postgres.Database) contract.Create{ModuleName}DatabaseRepository {
	return &create{ModuleName}DatabaseRepository{
		database: database,
	}
}
```

### 6. Dependency Injection Template

**Arquivo:** `pkg/di/{module_name}.go`

```go
package di

import (
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/contract"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/controller"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/repository"
	"github.com/booscaaa/clean-go/pkg-by-feature/internal/{module_name}/usecase"
	"github.com/booscaaa/initializers/postgres"
)

func NewCreate{ModuleName}Controller(database postgres.Database) contract.Create{ModuleName}Controller {
	return controller.NewCreate{ModuleName}Controller(
		usecase.NewCreate{ModuleName}UseCase(
			repository.NewCreate{ModuleName}DatabaseRepository(database),
		),
	)
}
```

### 7. Migration Template

**Arquivo:** `database/migrations/0000X_{module_name}.up.sql`

```sql
CREATE TABLE {module_name} (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

## Como Usar o Template

### Para criar um novo módulo:

1. **Substitua as variáveis:**
   - `{module_name}` → nome do módulo em snake_case (ex: `user`, `order`, `payment`)
   - `{ModuleName}` → nome do módulo em PascalCase (ex: `User`, `Order`, `Payment`)

2. **Crie a estrutura de diretórios:**
   ```bash
   mkdir -p internal/{module_name}/{contract,controller,model,repository,usecase}
   ```

3. **Crie os arquivos usando os templates acima**

4. **Adicione a rota no arquivo `pkg/rest/rest.go`:**
   ```go
   create{ModuleName}Controller := di.NewCreate{ModuleName}Controller(database)
   router.Post("/{module_name}s", create{ModuleName}Controller.Execute)
   ```

5. **Crie e execute a migração do banco de dados**

### Exemplo prático:

Para criar um módulo `user`:

- Substitua `{module_name}` por `user`
- Substitua `{ModuleName}` por `User`
- Crie os arquivos seguindo a estrutura
- Adicione a rota `/users` no rest.go
- Crie a migração `00002_user.up.sql`

## Arquitetura

Este template segue os princípios da Clean Architecture:

- **Model**: Entidades de domínio e DTOs
- **Contract**: Interfaces que definem os contratos entre as camadas
- **Controller**: Camada de apresentação (HTTP handlers)
- **UseCase**: Regras de negócio da aplicação
- **Repository**: Acesso a dados (banco de dados)
- **DI**: Injeção de dependências

## Dependências

- `github.com/go-chi/chi` - Router HTTP
- `github.com/spf13/cobra` - CLI
- `github.com/booscaaa/initializers/postgres` - Inicializador do PostgreSQL
- `github.com/jmoiron/sqlx` - Extensões SQL
- `github.com/lib/pq` - Driver PostgreSQL