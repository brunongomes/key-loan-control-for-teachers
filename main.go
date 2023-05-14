package main

import (
	"fmt"
	"log"
	"./services/disciplinas"
	"./services/professores"
	"./services/emprestimos"
)

func main() {
	fmt.Println("Controle de empréstimo de chaves")
 	for {
 		// Exibir menu de opções
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
 			// Solicitar dados da disciplina
 			fmt.Println("Digite o código da disciplina:")
 			var codigo string
 			fmt.Scanln(&codigo)
 			fmt.Println("Digite o nome da disciplina:")
 			var nome string
 			fmt.Scanln(&nome)
 			fmt.Println("Digite a carga horária da disciplina:")
 			var cargaHoraria int
 			fmt.Scanln(&cargaHoraria)
 
 			// Chamar a função cadastrarDisciplina para cadastrar a disciplina
 			err := disciplinas.CadastrarDisciplina(codigo, nome, cargaHoraria)
 			if err != nil {
 				log.Println("Erro ao cadastrar disciplina:", err)
 			}
 
 		case 2:
 			// Chamar a função listarDisciplinas para listar as disciplinas cadastradas
 			disciplinas.ListarDisciplinas()
 
 		case 3:
 			// Solicitar código da disciplina a ser excluída
 			fmt.Println("Digite o código da disciplina a ser excluída:")
 			var codigo string
 			fmt.Scanln(&codigo)
 
 			// Chamar a função excluirDisciplina para excluir a disciplina
 			err := disciplinas.ExcluirDisciplina(codigo)
 			if err != nil {
 				log.Println("Erro ao excluir disciplina:", err)
 			}
		
		case 4:

		case 5:
			// Solicitar dados do professor
			fmt.Println("Digite o nome do professor:") 
			var cpfProfessor string
			fmt.Scanln(&cpfProfessor)
			fmt.Println("Digite o CPF do professor:")
			var nomeProfessor string
			fmt.Scanln(&nomeProfessor)
			
			// Chamar a função cadastrarDisciplina para cadastrar a disciplina
			err := professores.CadastrarProfessor(cpfProfessor, nomeProfessor)
			if err != nil {
				log.Println("Erro ao cadastrar professor:", err)
			}

		case 6:
			professores.ListarProfessores()

		case 7:
			// Solicitar código da disciplina a ser excluída
			fmt.Println("Digite o cpf do professor para exclusão:") 
			var cpfProfessor string
			fmt.Scanln(&cpfProfessor)

			// Chamar a função excluirDisciplina para excluir o professor
			err := professores.ExcluirProfessor(cpfProfessor) 
			if err != nil {
			 log.Println("Erro ao excluir professor:", err)
			}

		case 8:

		case 9:
			// Solicitar dados do emprestimo
			fmt.Println("Digite o codigo do empréstimo") 
			var codigo int
			fmt.Scanln(&codigo)
			fmt.Println("Digite o CPF do professor:")
			var cpfProfessor string
			fmt.Scanln(&cpfProfessor)
			fmt.Println("Digite o nome do professor:")
			var nomeProfessor string
			fmt.Scanln(&nomeProfessor)
			fmt.Println("Digite o horario que o professor pegou a chave:")
			var horarioInicio int
			fmt.Scanln(&horarioInicio)
			fmt.Println("Digite o horario que o professor devolveu a chave:")
			var horarioFim int
			fmt.Scanln(&horarioFim)
			
			// Chamar a função cadastrarDisciplina para cadastrar a disciplina
			err := emprestimos.CadastrarEmprestimo(codigo, cpfProfessor, nomeProfessor, horarioInicio, horarioFim)
			if err != nil {
				log.Println("Erro ao cadastrar professor:", err)
			}

		case 10:
			emprestimos.ListarEmprestimos()

		case 11:

		case 12:

 		case 0:
 			// Encerrar o programa
 			fmt.Println("Encerrando o programa...")
 			return
 
 		default:
 			fmt.Println("Opção inválida, tente novamente. \n")
 		}
 	}
}
