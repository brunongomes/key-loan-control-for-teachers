package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Home</title>
				<link rel="stylesheet" href="/static/style.css">
			</head>
			<body class="container">
				<h1>Home</h1>
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
