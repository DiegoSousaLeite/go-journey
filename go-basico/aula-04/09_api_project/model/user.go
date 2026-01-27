package model

//Camada de modelo (model) - Define a estrutura de dados
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}