package emprestimos

import (
	"fmt"
	"os"
	"bufio" 
	"strings"
    "strconv"
)

const emprestimosFile = "./data/emprestimos.txt"

type Emprestimo struct {
	codigo         int
	CPF_Professor  string
	Nome_Professor string
	Horario_inicio string
	Horario_fim    string
}

func CadastrarEmprestimo(codigo int, CPF_Professor string, Nome_Professor string, Horario_inicio string, Horario_fim string) error {
	emprestimo := Emprestimo{
		codigo:          codigo,
		CPF_Professor:   CPF_Professor,
		Nome_Professor:  Nome_Professor,
	}

	// Solicitar e validar o horário de início
	for {
		if !validarHorario(Horario_inicio) {
			fmt.Println("Horário de início inválido. Utilize o formato HH:MM: ")
			fmt.Scanln(&Horario_inicio)
		} else {
			emprestimo.Horario_inicio = Horario_inicio
			break
		}
	}

	// Solicitar e validar o horário de fim
	for {
		if !validarHorario(Horario_fim) {
			fmt.Println("Horário de fim inválido. Utilize o formato HH:MM: ")
			fmt.Scanln(&Horario_fim)
		} else {
			emprestimo.Horario_fim = Horario_fim
			break
		}
	}

	// Abrir o arquivo em modo de escrita, cria se não existir
	file, err := os.OpenFile(emprestimosFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escrever os dados do professor no arquivo
	_, err = fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", emprestimo.codigo, emprestimo.CPF_Professor, emprestimo.Nome_Professor, emprestimo.Horario_inicio, emprestimo.Horario_fim)
	if err != nil {
		return err
	}

	fmt.Println("Empréstimo cadastrado com sucesso!")
	return nil
}

// Função auxiliar para validar o formato do horário (HH:MM)
func validarHorario(horario string) bool {
	parts := strings.Split(horario, ":")
	if len(parts) != 2 {
		return false
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}

	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return false
	}

	return true
}

func ListarEmprestimos() {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(emprestimosFile)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo de empréstimos:", err)
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
        fmt.Printf("%s\t%s\t\t%s\t\t%s\t\t%s\n", codigo, CPF_Professor, Nome_Professor, Horario_inicio, Horario_fim)
    }
    fmt.Println("----------------------------------- \n")
}

func ExcluirEmprestimos(codigo string) error {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(emprestimosFile)
    if err != nil {
        return err
    }
    defer file.Close()

    var emprestimos []Emprestimo
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        codigo := dados[0]
        CPF_Professor := dados[1]
        Nome_Professor := dados[2]
        Horario_inicio := dados[3]
        Horario_fim := dados[4]
        codigoInt, err := strconv.Atoi(codigo)
        if err != nil {
            return err
        }
        emprestimo := Emprestimo{
            codigoInt,
            CPF_Professor,
            Nome_Professor,
            Horario_inicio,
            Horario_fim,
        }
        emprestimos = append(emprestimos, emprestimo)
    }

    codigoInt, err := strconv.Atoi(codigo)
    if err != nil {
        return err
    }

    for i, emprestimo := range emprestimos {
        if emprestimo.codigo == codigoInt {
            emprestimos = append(emprestimos[:i], emprestimos[i+1:]...)
            break
        }
    }

    // Abrir o arquivo em modo de escrita para reescrever os dados
    file, err = os.OpenFile(emprestimosFile, os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados atualizados das disciplinas no arquivo
    for _, emprestimo := range emprestimos {
        _, err = fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", emprestimo.codigo, emprestimo.CPF_Professor, emprestimo.Nome_Professor, emprestimo.Horario_inicio, emprestimo.Horario_fim)
        if err != nil {
            return err
        }
    }

    fmt.Println("Empréstimo excluído com sucesso!")
    return nil
}

func AtualizarEmprestimo(codigo int) error {
    // Abrir o arquivo em modo leitura
    file, err := os.OpenFile(emprestimosFile, os.O_RDWR, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Ler o conteúdo do arquivo
    scanner := bufio.NewScanner(file)
    emprestimos := []Emprestimo{}
    for scanner.Scan() {
        emprestimoStr := scanner.Text()
        emprestimoArr := strings.Split(emprestimoStr, ",")
        emprestimoCodigo, _ := strconv.Atoi(emprestimoArr[0])
        if emprestimoCodigo == codigo {
            // Solicitar novos dados do empréstimo
            fmt.Println("Digite o CPF do professor:")
            var novoCPF string
            fmt.Scanln(&novoCPF)
            fmt.Println("Digite o nome do professor:")
            var novoNome string
            fmt.Scanln(&novoNome)
            fmt.Println("Digite o horário de início:")
            var novoInicio string
            fmt.Scanln(&novoInicio)
            fmt.Println("Digite o horário de fim:")
            var novoFim string
            fmt.Scanln(&novoFim)

            // Atualizar o empréstimo
            emprestimo := Emprestimo{
                codigo:         codigo,
                CPF_Professor:  novoCPF,
                Nome_Professor: novoNome,
                Horario_inicio: novoInicio,
                Horario_fim:    novoFim,
            }
            emprestimoStr := fmt.Sprintf("%d,%s,%s,%s,%s\n", emprestimo.codigo, emprestimo.CPF_Professor, emprestimo.Nome_Professor, emprestimo.Horario_inicio, emprestimo.Horario_fim)
            _, err = file.WriteString(emprestimoStr)
            if err != nil {
                return err
            }
            fmt.Println("Empréstimo atualizado com sucesso!")
            return nil
        }

        emprestimo := Emprestimo{
            codigo:         emprestimoCodigo,
            CPF_Professor:  emprestimoArr[1],
            Nome_Professor: emprestimoArr[2],
            Horario_inicio: emprestimoArr[3],
            Horario_fim:    emprestimoArr[4],
        }
        emprestimos = append(emprestimos, emprestimo)
    }

    // Caso não encontre o empréstimo, informar o usuário
    fmt.Println("Empréstimo não encontrado!")
    return nil
}
