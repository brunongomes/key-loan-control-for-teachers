# Sistema de Controle de Empréstimo de Chaves para Professores 📝

O sistema de Controle de Empréstimo de Chaves para Professores é uma aplicação que permite cadastrar, visualizar, atualizar e excluir informações relacionadas a professores, disciplinas e empréstimos.

## Requisitos ✅

Antes de começar, certifique-se de ter instalado o seguinte:

- [GO](https://go.dev/dl/)
- [MongoDB](https://www.mongodb.com/cloud/atlas/lp/try4?utm_source=google&utm_campaign=search_gs_pl_evergreen_atlas_core_prosp-brand_gic-null_amers-br_ps-all_desktop_eng_lead&utm_term=mongodb&utm_medium=cpc_paid_search&utm_ad=e&utm_ad_campaign_id=12212624308&adgroup=115749706023&cq_cmp=12212624308&gad=1&gclid=CjwKCAjwkLCkBhA9EiwAka9QRl846vPE0kXkCtmekDxAserqSfHRGRIJsxsZv90fypc8tm658DIQVhoCsM0QAvD_BwE)

## Dependências

Execute o seguinte comando para baixar as dependências necessárias:

```
go mod download
```

## Iniciar 🚀 

### Backend 

Para executar o backend, utilize o seguinte comando:

```
go run src/main.go
```

### Frontend

Para executar o frontend, utilize o seguinte comando:

```
go run frontend/main.go
```

## Rotas disponíveis 🛣️

A aplicação oferece as seguintes rotas:

### Disciplinas

- **Cadastrar uma disciplina**: `POST /disciplinas`
- **Listar disciplinas**: `GET /disciplinas`
- **Excluir uma disciplina**: `DELETE /disciplinas/{codigo}`
- **Atualizar uma disciplina**: `PUT /disciplinas/{codigo}`

### Professores

- **Cadastrar um professor**: `POST /professores`
- **Listar professores**: `GET /professores`
- **Excluir um professor**: `DELETE /professores/{cpf}`
- **Atualizar um professor**: `PUT /professores/{cpf}`

### Empréstimos

- **Cadastrar um empréstimo**: `POST /emprestimos`
- **Listar empréstimos**: `GET /emprestimos`
- **Excluir um empréstimo**: `DELETE /emprestimos/{codigo}`
- **Atualizar um empréstimo**: `PUT /emprestimos/{codigo}`

Sinta-se à vontade para utilizar essas rotas para interagir com o sistema e gerenciar as informações de professores, disciplinas e empréstimos.

# Key Loan Control System for Teachers 📝

The Key Loan Control System for Teachers is an application that allows you to register, view, update, and delete information related to teachers, subjects, and loans.

## Requirements ✅

Before getting started, make sure you have the following installed:

- [GO](https://go.dev/dl/)
- [MongoDB](https://www.mongodb.com/cloud/atlas/lp/try4?utm_source=google&utm_campaign=search_gs_pl_evergreen_atlas_core_prosp-brand_gic-null_amers-br_ps-all_desktop_eng_lead&utm_term=mongodb&utm_medium=cpc_paid_search&utm_ad=e&utm_ad_campaign_id=12212624308&adgroup=115749706023&cq_cmp=12212624308&gad=1&gclid=CjwKCAjwkLCkBhA9EiwAka9QRl846vPE0kXkCtmekDxAserqSfHRGRIJsxsZv90fypc8tm658DIQVhoCsM0QAvD_BwE)

## Dependencies

Run the following command to download the necessary dependencies:

```
go mod download
```

## Start 🚀

### Backend

To run the backend, use the following command:

```
go run src/main.go
```

### Frontend

To run the frontend, use the following command:

```
go run frontend/main.go
```

## Available Routes 🛣️

The application offers the following routes:

### Subjects

- **Create a subject**: `POST /subjects`
- **List subjects**: `GET /subjects`
- **Delete a subject**: `DELETE /subjects/{code}`
- **Update a subject**: `PUT /subjects/{code}`

### Teachers

- **Create a teacher**: `POST /teachers`
- **List teachers**: `GET /teachers`
- **Delete a teacher**: `DELETE /teachers/{cpf}`
- **Update a teacher**: `PUT /teachers/{cpf}`

### Loans

- **Create a loan**: `POST /loans`
- **List loans**: `GET /loans`
- **Delete a loan**: `DELETE /loans/{code}`
- **Update a loan**: `PUT /loans/{code}`

Feel free to use these routes to interact with the system and manage the information of teachers, subjects, and loans.
