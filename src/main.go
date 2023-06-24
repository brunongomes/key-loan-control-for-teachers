package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"./services/disciplinas"
	"./services/emprestimos"
	"./services/professores"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o código da disciplina:")
	scanner.Scan()
	codigo := scanner.Text()

	fmt.Println("Digite o nome da disciplina:")
	scanner.Scan()
	nome := scanner.Text()

	fmt.Println("Digite a carga horária da disciplina:")
	scanner.Scan()
	cargaHorariaStr := scanner.Text()
	cargaHoraria, err := strconv.Atoi(cargaHorariaStr)
	if err != nil {
		log.Println("Erro ao converter carga horária:", err)
		return
	}

	err = disciplinas.CadastrarDisciplina(codigo, nome, cargaHoraria)
	if err != nil {
		log.Println("Erro ao cadastrar disciplina:", err)
	}
}

func excluirDisciplina() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o código da disciplina a ser excluída:")
	scanner.Scan()
	codigo := scanner.Text()

	err := disciplinas.ExcluirDisciplina(codigo)
	if err != nil {
		log.Println("Erro ao excluir disciplina:", err)
	}
}

func atualizarDisciplina() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o código da disciplina que será atualizada:")
	scanner.Scan()
	codigo := scanner.Text()

	err := disciplinas.AtualizarDisciplina(codigo)
	if err != nil {
		log.Println("Erro ao atualizar disciplina:", err)
	}
}

func cadastrarProfessor() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o CPF do professor:")
	scanner.Scan()
	cpfProfessor := scanner.Text()

	fmt.Println("Digite o nome do professor:")
	scanner.Scan()
	nomeProfessor := scanner.Text()

	err := professores.CadastrarProfessor(cpfProfessor, nomeProfessor)
	if err != nil {
		log.Println("Erro ao cadastrar professor:", err)
	}
}

func excluirProfessor() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o CPF do professor para exclusão:")
	scanner.Scan()
	cpfProfessor := scanner.Text()

	err := professores.ExcluirProfessor(cpfProfessor)
	if err != nil {
		log.Println("Erro ao excluir professor:", err)
	}
}

func atualizarProfessor() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o CPF do professor que será atualizado:")
	scanner.Scan()
	cpf := scanner.Text()

	err := professores.AtualizarProfessor(cpf)
	if err != nil {
		log.Println("Erro ao atualizar disciplina:", err)
	}
}

func cadastrarEmprestimo() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o código do empréstimo")
	scanner.Scan()
	codigoStr := scanner.Text()
	codigo, err := strconv.Atoi(codigoStr)
	if err != nil {
		log.Println("Erro ao converter código de empréstimo:", err)
		return
	}

	fmt.Println("Digite o CPF do professor:")
	scanner.Scan()
	cpfProfessor := scanner.Text()

	fmt.Println("Digite o nome do professor:")
	scanner.Scan()
	nomeProfessor := scanner.Text()

	fmt.Println("Digite o horário que o professor pegou a chave: (Utilize o formato HH:MM)")
	scanner.Scan()
	horarioInicio := scanner.Text()

	fmt.Println("Digite o horário que o professor devolveu a chave: (Utilize o formato HH:MM)")
	scanner.Scan()
	horarioFim := scanner.Text()

	err = emprestimos.CadastrarEmprestimo(codigo, cpfProfessor, nomeProfessor, horarioInicio, horarioFim)
	if err != nil {
		log.Println("Erro ao cadastrar empréstimo:", err)
	}
}

func excluirEmprestimo() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o código do empréstimo a ser excluído:")
	scanner.Scan()
	codigo := scanner.Text()

	err := emprestimos.ExcluirEmprestimos(codigo)
	if err != nil {
		log.Println("Erro ao excluir empréstimo:", err)
	}
}

func atualizarEmprestimo() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Digite o código do empréstimo que será atualizado:")
	scanner.Scan()
	codigoStr := scanner.Text()
	codigo, err := strconv.Atoi(codigoStr)
	if err != nil {
		log.Println("Erro ao converter código de empréstimo:", err)
		return
	}

	err = emprestimos.AtualizarEmprestimo(codigo)
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
	}

	// Cria os arquivos dentro do diretório
	files := []string{"professores.txt", "emprestimos.txt", "disciplinas.txt"}
	for _, file := range files {
		filePath := filepath.Join(dir, file)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			_, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf("erro ao criar o arquivo %s: %v", file, err)
			}
			fmt.Printf("Arquivo %s criado com sucesso!\n", file)
		}
	}

	return nil
}
