package disciplinas

import (
    "bufio"    // Pacote para leitura e escrita de arquivos de texto
    "fmt"      // Pacote para formatação e exibição de texto na tela
    "os"       // Pacote para manipulação de arquivos e variáveis de ambiente
    "strconv"  // Pacote para conversão de tipos de dados, como string para int
    "strings"  // Pacote para manipulação de strings, como separação em substrings por um separador
)

const disciplinasFile = "./src/data/disciplinas.txt" // Nome do arquivo de texto para armazenar os dados das disciplinas

type Disciplina struct {
    Codigo       string
    Nome         string
    CargaHoraria int
}

func CadastrarDisciplina(codigo, nome string, cargaHoraria int) error {
    disciplina := Disciplina{
        Codigo:       codigo,
        Nome:         nome,
        CargaHoraria: cargaHoraria,
    }

    // Abrir o arquivo em modo de escrita, cria se não existir
    file, err := os.OpenFile(disciplinasFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados da disciplina no arquivo
    _, err = fmt.Fprintf(file, "%s,%s,%d\n", disciplina.Codigo, disciplina.Nome, disciplina.CargaHoraria)
    if err != nil {
        return err
    }

    fmt.Println("Disciplina cadastrada com sucesso!")
    return nil
}

func ListarDisciplinas() {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(disciplinasFile)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo de disciplinas:", err)
        return
    }
    defer file.Close()

    // Ler os dados do arquivo e exibir na tela
    fmt.Println("Disciplinas cadastradas:")
    fmt.Println("Código\tNome\tCarga Horária")
    fmt.Println("-----------------------------------")
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        codigo := dados[0]
        nome := dados[1]
        cargaHoraria, _ := strconv.Atoi(dados[2])
        fmt.Printf("%s\t%s\t%d\n", codigo, nome, cargaHoraria)
    }
    fmt.Println("----------------------------------- \n")
}

func ExcluirDisciplina(codigo string) error {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(disciplinasFile)
    if err != nil {
        return err
    }
    defer file.Close()

    // Ler os dados do arquivo e armazenar em um slice de Disciplinas
    var disciplinas []Disciplina
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        codigoAtual := dados[0]
        nomeAtual := dados[1]
        cargaHorariaAtual, _ := strconv.Atoi(dados[2])
        disciplina := Disciplina{
            Codigo:       codigoAtual,
            Nome:         nomeAtual,
            CargaHoraria: cargaHorariaAtual,
        }
        disciplinas = append(disciplinas, disciplina)
    }

    // Procurar a disciplina pelo código e removê-la do slice
    for i, disciplina := range disciplinas {
        if disciplina.Codigo == codigo {
            disciplinas = append(disciplinas[:i], disciplinas[i+1:]...)
            break
        }
    }

    // Abrir o arquivo em modo de escrita para reescrever os dados
    file, err = os.OpenFile(disciplinasFile, os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados atualizados das disciplinas no arquivo
    for _, disciplina := range disciplinas {
        _, err = fmt.Fprintf(file, "%s,%s,%d\n", disciplina.Codigo, disciplina.Nome, disciplina.CargaHoraria)
        if err != nil {
            return err
        }
    }

    fmt.Println("Disciplina excluída com sucesso!")
    return nil
}

func AtualizarDisciplina(codigo string) error {
    // Abrir o arquivo em modo de leitura
    file, err := os.Open(disciplinasFile)
    if err != nil {
        return err
    }
    defer file.Close()

    var disciplinas []Disciplina
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        linha := scanner.Text()
        dados := strings.Split(linha, ",")
        codigo := dados[0]
        nome := dados[1]
        cargaHoraria, _ := strconv.Atoi(dados[2])
        disciplina := Disciplina{
            Codigo:       codigo,
            Nome:         nome,
            CargaHoraria: cargaHoraria,
        }
        disciplinas = append(disciplinas, disciplina)
    }

    // Procurar a disciplina pelo código
    index := -1
    for i, disciplina := range disciplinas {
        if disciplina.Codigo == codigo {
            index = i
            break
        }
    }

    if index == -1 {
        return fmt.Errorf("Código %s não encontrado. Digite um código válido.", codigo)
    }

    // Pedir os novos dados da disciplina
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite o novo nome da disciplina: ")
    nome, _ := reader.ReadString('\n')
    nome = strings.TrimSpace(nome)
    fmt.Print("Digite a nova carga horária da disciplina: ")
    cargaHorariaStr, _ := reader.ReadString('\n')
    cargaHorariaStr = strings.TrimSpace(cargaHorariaStr)
    cargaHoraria, _ := strconv.Atoi(cargaHorariaStr)

    // Atualizar os dados da disciplina na lista
    disciplinas[index].Nome = nome
    disciplinas[index].CargaHoraria = cargaHoraria

    // Abrir o arquivo em modo de escrita para reescrever os dados
    file, err = os.OpenFile(disciplinasFile, os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    // Escrever os dados atualizados das disciplinas no arquivo
    for _, disciplina := range disciplinas {
        _, err = fmt.Fprintf(file, "%s,%s,%d\n", disciplina.Codigo, disciplina.Nome, disciplina.CargaHoraria)
        if err != nil {
            return err
        }
    }

    fmt.Println("Disciplina atualizada com sucesso!")
    return nil
}
