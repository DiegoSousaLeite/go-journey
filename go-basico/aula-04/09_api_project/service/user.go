package service

//Camada de serviço (service) - Lógica de negócio
//Normalmente interage com a camada de modelo (model) e outras camadas (como repositório, controlador, etc.)

import (
	"errors"
	"github.com/diego/go-basico-01/model"
)

//Simulando um banco de dados com uma fatia (slice) de usuários
//Em um cenário real, isso seria substituído por chamadas a um banco de dados
//Seria o repositório (repository)
var users = []model.User{
	{ID: "1", Name: "Alice", Age: 30},
	{ID: "2", Name: "Bob", Age: 25},
	{ID: "3", Name: "Charlie", Age: 35},
}

func GetAllUsers() []model.User {
	return users
}

func GetUserByID(id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}