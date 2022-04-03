package main

import (
	"github.com/brunohprada/go-api-gin/database"
	"github.com/brunohprada/go-api-gin/models"
	"github.com/brunohprada/go-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Bruno Prada", CPF: "12345678912", RG: "123456789"},
		{Nome: "João das Neves", CPF: "98765432123", RG: "1238762344"},
		{Nome: "Hermengarda das Planicies", CPF: "87612365434", RG: "983462374"},
		{Nome: "Guimarães Cravo", CPF: "76523456712", RG: "874569823"},
	}
	routes.HandleRequests()
}
