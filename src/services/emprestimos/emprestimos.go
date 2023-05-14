package emprestimos

import (
	"fmt"
	"os"
	"bufio" 
	"strings"
)

const emprestimosFile = "./src/data/emprestimos.txt" // Nome do arquivo de texto para armazenar os dados das disciplinas

type Emprestimo struct {
	codigo  int
	CPF_Professor  string
	Nome_Professor string
	Horario_inicio int
	Horario_fim int
}


func CadastrarEmprestimo(codigo int, CPF_Professor string, Nome_Professor string, Horario_inicio int, Horario_fim int) error {
	emprestimo := Emprestimo {
		codigo : codigo,
		CPF_Professor : CPF_Professor,
		Nome_Professor: Nome_Professor,
		Horario_inicio: Horario_inicio,
		Horario_fim: Horario_fim,
	}

    // Abrir o arquivo em modo de escrita, cria se não existir
    file, err := os.OpenFile(emprestimosFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados do professor no arquivo
    _, err = fmt.Fprintf(file, "%d,%s,%s,%d,%d\n", emprestimo.codigo, emprestimo.CPF_Professor, emprestimo.Nome_Professor, emprestimo.Horario_inicio, emprestimo.Horario_fim)
    if err != nil {
        return err
    }

    fmt.Println("Empréstimo cadastrado com sucesso!")
    return nil
}

func ListarEmprestimos() {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(emprestimosFile)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo de professores:", err)
        return
    }
    defer file.Close()

    // Ler os dados do arquivo e exibir na tela
    fmt.Println("Empréstimos cadastrados:")
    fmt.Println("codigo\tCPF_Professor\t\tNome_Professor\tHorario_inicio\tHorario_fim")
    fmt.Println("-----------------------------------")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        codigo := dados[0]
        CPF_Professor := dados[1]
        Nome_Professor := dados[2]
        Horario_inicio := dados[3]
        Horario_fim := dados[4]
        fmt.Printf("%s\t%s\t%s\t\t%s\t\t%s\n", codigo, CPF_Professor, Nome_Professor, Horario_inicio, Horario_fim)
    }
    fmt.Println("----------------------------------- \n")
}