package main

import (
	"context"
	"bufio"
	"os"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"./pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type Disciplina struct {
	Codigo       string
	Nome         string
	CargaHoraria int
}

type Emprestimo struct {
	Codigo         int
	CPF_Professor  string
	Nome_Professor string
	Horario_inicio string
	Horario_fim    string
}

type Professor struct {
	CPF  string
	Nome string
}

func main() {
	db, err := mongodb.ConnectToMongoDB()
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
	}
	defer db.Client.Disconnect(context.Background())

	var opcao int
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar disciplina")
		fmt.Println("2. Listar disciplinas")
		fmt.Println("3. Excluir disciplina")
		fmt.Println("4. Atualizar disciplina")
		fmt.Println("5. Cadastrar professor")
		fmt.Println("6. Listar professores")
		fmt.Println("7. Excluir professor")
		fmt.Println("8. Atualizar professor")
		fmt.Println("9. Cadastrar empréstimo")
		fmt.Println("10. Listar empréstimos")
		fmt.Println("11. Excluir empréstimo")
		fmt.Println("12. Atualizar empréstimo")
		fmt.Println("0. Sair")
		fmt.Println("----------------------------------- \n")
		fmt.Scanln(&opcao)

		switch opcao {
		case 0:
			fmt.Println("Saindo do programa...")
			return
		case 1:
			err = cadastrarDisciplina(db)
		case 2:
			err = listarColecao(db, "disciplinas")
		case 3:
			err = excluirDocumento(db, "disciplinas")
		case 4:
			err = atualizarDocumento(db, "disciplinas")
		case 5:
			err = cadastrarProfessor(db)
		case 6:
			err = listarColecao(db, "professores")
		case 7:
			err = excluirDocumento(db, "professores")
		case 8:
			err = atualizarDocumento(db, "professores")
		case 9:
			err = cadastrarEmprestimo(db)
		case 10:
			err = listarColecao(db, "emprestimos")
		case 11:
			err = excluirDocumento(db, "emprestimos")
		case 12:
			err = atualizarDocumento(db, "emprestimos")
		default:
			fmt.Println("Opção inválida")
		}

		if err != nil {
			log.Println("Erro:", err)
		}
		fmt.Println("----------------------------------- \n")
	}
}

func cadastrarDisciplina(db *mongodb.MongoDB) error {
	var disciplina Disciplina
	fmt.Println("Digite o código da disciplina:")
	fmt.Scanln(&disciplina.Codigo)
	fmt.Println("Digite o nome da disciplina:")
	disciplina.Nome = readLine()
	fmt.Println("Digite a carga horária da disciplina:")
	fmt.Scanln(&disciplina.CargaHoraria)

	err := db.Insert("disciplinas", disciplina)
	if err != nil {
		return fmt.Errorf("Erro ao cadastrar disciplina: %s", err)
	}
	fmt.Println("Disciplina cadastrada com sucesso!")
	return nil
}

func cadastrarProfessor(db *mongodb.MongoDB) error {
	var professor Professor
	fmt.Println("Digite o CPF do professor:")
	fmt.Scanln(&professor.CPF)
	fmt.Println("Digite o nome do professor:")
	professor.Nome = readLine()

	err := db.Insert("professores", professor)
	if err != nil {
		return fmt.Errorf("Erro ao cadastrar professor: %s", err)
	}
	fmt.Println("Professor cadastrado com sucesso!")
	return nil
}

func cadastrarEmprestimo(db *mongodb.MongoDB) error {
	var emprestimo Emprestimo
	fmt.Println("Digite o código do empréstimo:")
	fmt.Scanln(&emprestimo.Codigo)
	fmt.Println("Digite o CPF do professor:")
	fmt.Scanln(&emprestimo.CPF_Professor)
	fmt.Println("Digite o nome do professor:")
	emprestimo.Nome_Professor = readLine()
	fmt.Println("Digite o horário de início do empréstimo:")
	fmt.Scanln(&emprestimo.Horario_inicio)
	fmt.Println("Digite o horário de fim do empréstimo:")
	fmt.Scanln(&emprestimo.Horario_fim)

	err := db.Insert("emprestimos", emprestimo)
	if err != nil {
		return fmt.Errorf("Erro ao cadastrar empréstimo: %s", err)
	}
	fmt.Println("Empréstimo cadastrado com sucesso!")
	return nil
}

func listarColecao(db *mongodb.MongoDB, collection string) error {
	filter := bson.M{}
	results, err := db.Read(collection, filter)
	if err != nil {
		return fmt.Errorf("Erro ao listar %s: %s", collection, err)
	}
	for _, result := range results {
		data, err := json.Marshal(result)
		if err != nil {
			return fmt.Errorf("Erro ao serializar resultado: %s", err)
		}
		fmt.Println(string(data))
	}
	return nil
}

func excluirDocumento(db *mongodb.MongoDB, collection string) error {
	var campo, valor string

	switch collection {
	case "professores":
		fmt.Println("Digite o CPF do professor a ser excluído:")
		fmt.Scanln(&valor)
		campo = "CPF"
	case "disciplinas":
		fmt.Println("Digite o código da disciplina a ser excluída:")
		fmt.Scanln(&valor)
		campo = "codigo"
	case "emprestimos":
		fmt.Println("Digite o código do empréstimo a ser excluído:")
		fmt.Scanln(&valor)
		campo = "Codigo"
	default:
		return fmt.Errorf("Coleção inválida")
	}

	filter := bson.M{campo: valor}
	err := db.Delete(collection, filter)
	if err != nil {
		return fmt.Errorf("Erro ao excluir documento: %s", err)
	}
	fmt.Println("Documento excluído com sucesso!")
	return nil
}

func atualizarDocumento(db *mongodb.MongoDB, collection string) error {
	var codigo string
	fmt.Println("Digite o código do documento a ser atualizado:")
	fmt.Scanln(&codigo)

	filter := bson.M{"codigo": codigo}
	update := bson.M{}

	switch collection {
	case "disciplinas":
		var disciplina Disciplina
		fmt.Println("Digite o novo nome da disciplina:")
		disciplina.Nome = readLine()
		update = bson.M{"$set": bson.M{"nome": disciplina.Nome}}
	case "professores":
		var professor Professor
		fmt.Println("Digite o novo nome do professor:")
		professor.Nome = readLine()
		update = bson.M{"$set": bson.M{"nome": professor.Nome}}
	case "emprestimos":
		var emprestimo Emprestimo
		fmt.Println("Digite o novo horário de início do empréstimo:")
		fmt.Scanln(&emprestimo.Horario_inicio)
		update = bson.M{"$set": bson.M{"horario_inicio": emprestimo.Horario_inicio}}
	default:
		return fmt.Errorf("Coleção inválida")
	}

	err := db.Update(collection, filter, update)
	if err != nil {
		return fmt.Errorf("Erro ao atualizar documento: %s", err)
	}
	fmt.Println("Documento atualizado com sucesso!")
	return nil
}

// Função auxiliar para ler uma linha inteira do console, incluindo espaços em branco.
func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
