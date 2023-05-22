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
        Horario_inicio, _ := strconv.Atoi(dados[3])
        Horario_fim, _ := strconv.Atoi(dados[4])
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
        _, err = fmt.Fprintf(file, "%s,%s,%s,%d,%d\n", emprestimo.codigo, emprestimo.CPF_Professor, emprestimo.Nome_Professor, emprestimo.Horario_inicio, emprestimo.Horario_fim)
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
            var novoInicio int
 			fmt.Scanln(&novoInicio)
            fmt.Println("Digite o horário de fim:")
            var novoFim int
 			fmt.Scanln(&novoFim)

            // Atualizar o empréstimo
            emprestimo := Emprestimo{
                codigo:         codigo,
                CPF_Professor:  novoCPF,
                Nome_Professor: novoNome,
                Horario_inicio: novoInicio,
                Horario_fim:    novoFim,
            }
            emprestimoStr := fmt.Sprintf("%d,%s,%s,%d,%d\n", emprestimo.codigo, emprestimo.CPF_Professor, emprestimo.Nome_Professor, emprestimo.Horario_inicio, emprestimo.Horario_fim)
            _, err = file.WriteString(emprestimoStr)
            if err != nil {
                return err
            }
            fmt.Println("Empréstimo atualizado com sucesso!")
            return nil
        }

        Horario_inicio, err := strconv.Atoi(emprestimoArr[3])
        if err != nil {
            return err
        }

        Horario_fim, err := strconv.Atoi(emprestimoArr[3])
        if err != nil {
            return err
        }

        emprestimo := Emprestimo{
            codigo:         emprestimoCodigo,
            CPF_Professor:  emprestimoArr[1],
            Nome_Professor: emprestimoArr[2],
            Horario_inicio: Horario_inicio,
            Horario_fim:    Horario_fim,
        }
        emprestimos = append(emprestimos, emprestimo)
    }

    // Caso não encontre o empréstimo, informar o usuário
    fmt.Println("Empréstimo não encontrado!")
    return nil
}
