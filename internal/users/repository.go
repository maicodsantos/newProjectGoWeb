package users

import "fmt"

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Create(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
	LastID() (int, error)
	Update(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error)
	UpdateNome(id int, nome string) (User, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Create(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	u := User{id, nome, sobrenome, email, idade, altura, ativo, dataDeCriacao}
	users = append(users, u)
	lastID = u.Id
	return u, nil
}

func (repository) Update(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCriacao string) (User, error) {
	u := User{Nome: nome, Sobrenome: sobrenome, Email: email, Idade: idade, Altura: altura, Ativo: ativo, DataDeCriacao: dataDeCriacao}
	updated := false
	for i := range users {
		if users[i].Id == id {
			u.Id = id
			users[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("usuario %d não encontrado", id)
	}
	return u, nil
}

func (repository) UpdateNome(id int, nome string) (User, error) {
	var u User
	updated := false
	for i := range users {
		if users[i].Id == id {
			users[i].Nome = nome
			updated = true
			u = users[i]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return u, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d não encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}
