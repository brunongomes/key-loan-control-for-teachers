package main

import (
	"fmt"
	"log"
	"os"
	"./services/disciplinas"
	"./services/professores"
	"./services/emprestimos"
)

func main() {
	err := checkAndCreateDirectory()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Controle de empréstimo de chaves")
	for {
		fmt.Println("Escolha uma opção:")
		fmt.Println("1. Cadastrar disciplina")
		fmt.Println("2. Listar disciplinas")
		fmt.Println("3. Excluir disciplina")
		fmt.Println("4. Atualizar disciplina")
		fmt.Println("5. Cadastrar professor")
		fmt.Println("6. Listar professor")
		fmt.Println("7. Excluir professor")
		fmt.Println("8. Atualizar professor")
		fmt.Println("9. Cadastrar empréstimo")
		fmt.Println("10. Listar empréstimo")
		fmt.Println("11. Excluir empréstimo")
		fmt.Println("12. Atualizar empréstimo")
		fmt.Println("0. Sair")
		fmt.Println("----------------------------------- \n")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			cadastrarDisciplina()
		case 2:
			disciplinas.ListarDisciplinas()
		case 3:
			excluirDisciplina()
		case 4:
			atualizarDisciplina()
		case 5:
			cadastrarProfessor()
		case 6:
			professores.ListarProfessores()
		case 7:
			excluirProfessor()
		case 8:
			atualizarProfessor()
		case 9:
			cadastrarEmprestimo()
		case 10:
			emprestimos.ListarEmprestimos()
		case 11:
			excluirEmprestimo()
		case 12:
			atualizarEmprestimo()
		case 0:
			fmt.Println("Encerrando o programa...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente. \n")
		}
	}
}

func cadastrarDisciplina() {
	fmt.Println("Digite o código da disciplina:")
	var codigo string
	fmt.Scanln(&codigo)
	fmt.Println("Digite o nome da disciplina:")
	var nome string
	fmt.Scanln(&nome)
	fmt.Println("Digite a carga horária da disciplina:")
	var cargaHoraria int
	fmt.Scanln(&cargaHoraria)
	err := disciplinas.CadastrarDisciplina(codigo, nome, cargaHoraria)
	if err != nil {
		log.Println("Erro ao cadastrar disciplina:", err)
	}
}

func excluirDisciplina() {
	fmt.Println("Digite o código da disciplina a ser excluída:")
	var codigo string
	fmt.Scanln(&codigo)
	err := disciplinas.ExcluirDisciplina(codigo)
	if err != nil {
		log.Println("Erro ao excluir disciplina:", err)
	}
}

func atualizarDisciplina() {
	fmt.Println("Digite o código da disciplina que será atualizada:")
	var codigo string
	fmt.Scanln(&codigo)
	err := disciplinas.AtualizarDisciplina(codigo)
	if err != nil {
		log.Println("Erro ao atualizar disciplina:", err)
	}
}

func cadastrarProfessor() {
	fmt.Println("Digite o CPF do professor:")
	var cpfProfessor string
	fmt.Scanln(&cpfProfessor)
	fmt.Println("Digite o nome do professor:")
	var nomeProfessor string
	fmt.Scanln(&nomeProfessor)
	err := professores.CadastrarProfessor(cpfProfessor, nomeProfessor)
	if err != nil {
		log.Println("Erro ao cadastrar professor:", err)
	}
}

func excluirProfessor() {
	fmt.Println("Digite o cpf do professor para exclusão:")
	var cpfProfessor string
	fmt.Scanln(&cpfProfessor)
	err := professores.ExcluirProfessor(cpfProfessor)
	if err != nil {
		log.Println("Erro ao excluir professor:", err)
	}
}

func atualizarProfessor() {
	fmt.Println("Digite o cpf do professor que será atualizado:")
	var cpf string
	fmt.Scanln(&cpf)
	err := professores.AtualizarProfessor(cpf)
	if err != nil {
		log.Println("Erro ao atualizar disciplina:", err)
	}
}

func cadastrarEmprestimo() {
	fmt.Println("Digite o código do empréstimo")
	var codigo int
	fmt.Scanln(&codigo)
	fmt.Println("Digite o CPF do professor:")
	var cpfProfessor string
	fmt.Scanln(&cpfProfessor)
	fmt.Println("Digite o nome do professor:")
	var nomeProfessor string
	fmt.Scanln(&nomeProfessor)
	fmt.Println("Digite o horario que o professor pegou a chave: (Utilize o formato HH:MM)")
	var horarioInicio string
	fmt.Scanln(&horarioInicio)
	fmt.Println("Digite o horario que o professor devolveu a chave: (Utilize o formato HH:MM)")
	var horarioFim string
	fmt.Scanln(&horarioFim)
	err := emprestimos.CadastrarEmprestimo(codigo, cpfProfessor, nomeProfessor, horarioInicio, horarioFim)
	if err != nil {
		log.Println("Erro ao cadastrar professor:", err)
	}
}

func excluirEmprestimo() {
	fmt.Println("Digite o código do empréstimo a ser excluído:")
	var codigo string
	fmt.Scanln(&codigo)
	err := emprestimos.ExcluirEmprestimos(codigo)
	if err != nil {
		log.Println("Erro ao excluir empréstimo:", err)
	}
}

func atualizarEmprestimo() {
	fmt.Println("Digite o código do empréstimo que será atualizado:")
	var codigo int
	fmt.Scanln(&codigo)
	err := emprestimos.AtualizarEmprestimo(codigo)
	if err != nil {
		log.Println("Erro ao atualizar empréstimo:", err)
	}
}

func checkAndCreateDirectory() error {
	dir := "./data"

	// Verifica se o diretório já existe
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// Cria o diretório
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return fmt.Errorf("erro ao criar o diretório: %v", err)
		}
		fmt.Println("Diretório criado com sucesso!")
	} else if err != nil {
		return fmt.Errorf("erro ao verificar o diretório: %v", err)
	}

	return nil
}
