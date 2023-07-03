package home

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Controle de chaves</title>
				<link rel="stylesheet" href="/static/style.css">
				<link rel="icon" href="/static/favicon.ico" type="image/ico">
			</head>
			<body class="container">
				<h1>Controle de empréstimo de chaves</h1>
				<img src="/static/logo.png" alt="Logo">
				<div class="container">
  				<a href="/disciplinas" class="button">Disciplinas</a>
  				<a href="/emprestimos" class="button">Empréstimos</a>
  				<a href="/professores" class="button">Professores</a>
				</div>
			</body>
			</html>
			`
		fmt.Fprintf(w, html)
	}
}
