package professores

import (
    "bufio"    // Pacote para leitura e escrita de arquivos de texto
    "fmt"      // Pacote para formatação e exibição de texto na tela
    "os"       // Pacote para manipulação de arquivos e variáveis de ambiente
    // "strconv"  // Pacote para conversão de tipos de dados, como string para int
    "strings"  // Pacote para manipulação de strings, como separação em substrings por um separador
)

const professoresFile = "./data/professores.txt" // Nome do arquivo de texto para armazenar os dados dos professores

type Professor struct {
    CPF  string
    Nome string
}

func CadastrarProfessor(cpf, nome string) error {
    professor := Professor{
        CPF:  cpf,
        Nome: nome,
    }

    // Abrir o arquivo em modo de escrita, cria se não existir
    file, err := os.OpenFile(professoresFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados do professor no arquivo
    _, err = fmt.Fprintf(file, "%s,%s\n", professor.CPF, professor.Nome)
    if err != nil {
        return err
    }

    fmt.Println("Professor cadastrado com sucesso!")
    return nil
}

func ListarProfessores() {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(professoresFile)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo de professores:", err)
        return
    }
    defer file.Close()

    // Ler os dados do arquivo e exibir na tela
    fmt.Println("Professores cadastrados:")
    fmt.Println("CPF\tNome")
    fmt.Println("-----------------------------------")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        cpf := dados[0]
        nome := dados[1]
        fmt.Printf("%s\t%s\n", cpf, nome)
    }
    fmt.Println("----------------------------------- \n")
}

func ExcluirProfessor(cpf string) error {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(professoresFile)
    if err != nil {
        return err
    }
    defer file.Close()

    // Ler os dados do arquivo e armazenar em um slice de Professores
    var professores []Professor
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        cpfAtual := dados[1]
        nomeAtual := dados[0]
        professor := Professor{
            CPF:  cpfAtual,
            Nome: nomeAtual,
        }
        professores = append(professores, professor)
    }

    // Procurar o professor pelo CPF e removê-lo do slice
    for i, professor := range professores {
        if professor.CPF == cpf {
            professores = append(professores[:i], professores[i+1:]...)
            break
        }
    }

    // Abrir o arquivo em modo de escrita para reescrever os dados
    file, err = os.OpenFile(professoresFile, os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados atualizados dos professores no arquivo
    for _, professor := range professores {
        _, err = fmt.Fprintf(file, "%s,%s\n", professor.CPF, professor.Nome)
        if err != nil {
            return err
        }
    }

    fmt.Println("Professor excluído com sucesso!")
    return nil
}
