package handlers

import (
	"fmt"
	"net/http"
)

func EmprestimosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Exibir a página de empréstimos
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Cadastrar Disciplina</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body class="container" >
			<h1>Cadastrar Empréstimos</h1>
			<div class="form-container">
				<form action="/emprestimos" method="POST">
					<label for="codigo">Código:</label><br>
					<input type="text" id="codigo" name="codigo"><br><br>
					<label for="cpfProfessor">CPF do professor:</label><br>
					<input type="text" id="cpfProfessor" name="cpfProfessor"><br><br>
					<label for="nomeProfessor">Nome do professor:</label><br>
					<input type="text" id="nomeProfessor" name="nomeProfessor"><br><br>
					<label for="horarioInicio">Horário início:</label><br>
					<input type="text" id="horarioInicio" name="horarioInicio"><br><br>
					<label for="horarioFim">Horário início:</label><br>
					<input type="text" id="horarioFim" name="horarioFim"><br><br>
					<input class="button-form" type="submit" value="Cadastrar">
				</form>
				</div>
				<div class="btn"> 
					<a href="/visualizar-emprestimos" class="button-form" method="GET" >Visualizar</a>
					<a href="/" class="button-form" method="GET" >Voltar</a>
				</div>
		</body>
		</html>
			`
		fmt.Fprintf(w, html)
	}
}
