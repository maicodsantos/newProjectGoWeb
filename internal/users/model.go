package users

type User struct {
	Id            int    `json:"id"`
	Nome          string `json:"nome"`
	Sobrenome     string `json:"sobrenome"`
	Email         string `json:"email"`
	Idade         int    `json:"idade"`
	Altura        int    `json:"altura"`
	Ativo         bool   `json:"ativo"`
	DataDeCriacao string `json:"data_de_criacao"`
}
