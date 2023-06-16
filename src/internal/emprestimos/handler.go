package emprestimos

type Emprestimo struct {
	Codigo         int    `json:"codigo"`
	CPF_Professor  string `json:"cpfProfessor"`
	Nome_Professor string `json:"nomeProfessor"`
	Horario_inicio string `json:"horarioInicio"`
	Horario_fim    string `json:"horarioFim"`
}
