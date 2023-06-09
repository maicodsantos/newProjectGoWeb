package users

type Service interface {
	GetAll() ([]User, error)
	Create(nome, sobrenome, email string, idade, altura int, ativo bool, DataDeCriacao string) (User, error)
	GetById() ([]User, error)
	Update(id int, nome, sobrenome, email string, idade, altura int, ativo bool, dataDeCraicao string) (User, error)
	UpdateNome(id int, nome string) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetById() ([]User, error) {
	return nil, nil
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) Create(nome, sobrenome, email string, idade, altura int, ativo bool, DataDeCriacao string) (User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return User{}, err
	}

	lastID++

	user, err := s.repository.Create(lastID, nome, sobrenome, email, idade, altura, ativo, DataDeCriacao)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s service) Update(id int, nome, sobrenome, email string, idade, altura int, ativo bool, DataDeCriacao string) (User, error) {
	user, err := s.repository.Update(id, nome, sobrenome, email, idade, altura, ativo, DataDeCriacao)
	if err != nil {
		return User{}, err
	}
	return user, err
}

func (s service) UpdateNome(id int, nome string) (User, error) {
	user, err := s.repository.UpdateNome(id, nome)
	if err != nil {
		return User{}, err
	}
	return user, err
}

func (s service) Delete(id int) error {
	return s.repository.Delete(id)
}
