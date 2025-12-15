package service

import (
	"context"
	"time"

	"github.com/surya/user-age-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (int32, string, time.Time, error) {
	user, err := s.repo.CreateUser(ctx, name, dob)
	if err != nil {
		return 0, "", time.Time{}, err
	}
	return user.ID, user.Name, user.Dob.Time, nil
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int32,
) (int32, string, time.Time, int, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return 0, "", time.Time{}, 0, err
	}

	dob := user.Dob.Time
	age := calculateAge(dob)

	return user.ID, user.Name, dob, age, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]struct {
	ID   int32
	Name string
	Dob  time.Time
	Age  int
}, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]struct {
		ID   int32
		Name string
		Dob  time.Time
		Age  int
	}, 0)

	for _, u := range users {
		dob := u.Dob.Time
		result = append(result, struct {
			ID   int32
			Name string
			Dob  time.Time
			Age  int
		}{
			ID:   u.ID,
			Name: u.Name,
			Dob:  dob,
			Age:  calculateAge(dob),
		})
	}

	return result, nil
}
