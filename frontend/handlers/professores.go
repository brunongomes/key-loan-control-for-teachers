package handlers

import (
	"fmt"
	"net/http"
)

func ProfessoresHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de professores
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Professor</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body class="container" >
			<h1>Cadastrar Professor</h1>
			<div class="form-container">
				<form action="/professores" method="POST">
					<label for="cpf">CPF:</label><br>
					<input type="text" id="cpf" name="cpf"><br><br>
					<label for="nome">Nome:</label><br>
					<input type="text" id="nome" name="nome"><br><br>
					<input class="button-form" type="submit" value="Cadastrar">
				</form>
				</div>
				<div class="btn"> 
					<a href="/visualizar-professores" class="button-form" method="GET" >Visualizar</a>
					<a href="/" class="button-form" method="GET" >Voltar</a>
				</div>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}
