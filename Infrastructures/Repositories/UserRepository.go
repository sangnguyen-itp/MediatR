package Repositories

type UserRepository struct {
	*Repository
}

func NewUserRepository(base *Repository) *UserRepository {
	return &UserRepository{base}
}

