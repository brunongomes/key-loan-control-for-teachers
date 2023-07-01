package handlers

import (
	"fmt"
	"net/http"
)

func DisciplinasHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de disciplinas
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Disciplina</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body class="container">
			<h1>Cadastrar Disciplina</h1>
			<div class="form-container">
				<form action="/disciplinas" method="POST">
					<label for="codigo">Código:</label><br>
					<input type="text" id="codigo" name="codigo"><br><br>
					<label for="nome">Nome:</label><br>
					<input type="text" id="nome" name="nome"><br><br>
					<label for="cargaHoraria">Carga Horária:</label><br>
					<input type="text" id="cargaHoraria" name="cargaHoraria"><br><br>
					<input class="button-form" type="submit" value="Cadastrar">
				</form>
			</div>
			<div class="btn"> 
				<a href="/visualizar-disciplinas" class="button-form" method="GET" >Visualizar</a>
				<a href="/" class="button-form" method="GET" >Voltar</a>
			</div>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}
