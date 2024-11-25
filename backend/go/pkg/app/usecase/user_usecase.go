package usecase

import (
	"context"
	"path/filepath"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"github.com/auster-kaki/auster-mono/pkg/util/austerstorage"
)

type UserUseCase struct {
	repository repository.Repository
}

func NewUserUseCase(r repository.Repository) *UserUseCase {
	return &UserUseCase{repository: r}
}

func (u *UserUseCase) GetHobbies(ctx context.Context) (entity.Hobbies, error) {
	return u.repository.Hobby().GetAll(ctx)
}

type UsersGetOutput struct {
	Users entity.Users
}

func (u *UserUseCase) GetUsers(ctx context.Context) (*UsersGetOutput, error) {
	users, err := u.repository.User().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &UsersGetOutput{Users: users}, nil
}

type UserGetOutput struct {
	User    *entity.User
	Hobbies entity.Hobbies
	Photo   []byte
}

func (u *UserUseCase) GetUser(ctx context.Context, id entity.UserID) (*UserGetOutput, error) {
	user, err := u.repository.User().FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	photo, err := austerstorage.Get(user.ProfilePath)
	if err != nil {
		return nil, err
	}

	userHobbies, err := u.repository.UserHobby().GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	hobbies, err := u.repository.Hobby().GetByIDs(ctx, userHobbies.HobbyIDs())
	if err != nil {
		return nil, err
	}
	return &UserGetOutput{
		User:    user,
		Hobbies: hobbies,
		Photo:   photo,
	}, nil
}

type UserInput struct {
	ID      entity.UserID
	Name    string
	Gender  string
	Age     int
	Photo   Photo
	Hobbies entity.Hobbies
}

type Photo struct {
	Filename    string
	Body        []byte
	ContentType string
}

func (u *UserUseCase) CreateUser(ctx context.Context, input *UserInput) (entity.UserID, error) {
	userID := austerid.Generate[entity.UserID]()
	path, err := austerstorage.Save(
		austerstorage.ContentType(input.Photo.ContentType),
		filepath.Join("users", string(userID), input.Photo.Filename),
		input.Photo.Body,
	)
	if err != nil {
		return "", err
	}

	var (
		newHobbies  = make(entity.Hobbies, 0, len(input.Hobbies))
		userHobbies = make(entity.UserHobbies, len(input.Hobbies))
	)
	for i, hobby := range input.Hobbies {
		if hobby.ID == "" {
			hobby.ID = austerid.Generate[entity.HobbyID]()
			newHobbies = append(newHobbies, hobby)
		}
		userHobbies[i] = &entity.UserHobby{
			UserID:  userID,
			HobbyID: hobby.ID,
		}
	}

	if err := u.repository.UserHobby().Create(ctx, userHobbies...); err != nil {
		return "", err
	}
	if len(newHobbies) > 0 {
		if err := u.repository.Hobby().Create(ctx, newHobbies...); err != nil {
			return "", err
		}
	}

	if err := u.repository.User().Create(ctx, &entity.User{
		ID:          userID,
		Name:        input.Name,
		Gender:      input.Gender,
		Age:         input.Age,
		ProfilePath: path,
	}); err != nil {
		return "", err
	}

	return userID, nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, input *UserInput) error {
	path, err := austerstorage.Save(
		austerstorage.ContentType(input.Photo.ContentType),
		filepath.Join("users", string(input.ID), input.Photo.Filename),
		input.Photo.Body,
	)
	if err != nil {
		return err
	}
	var (
		newHobbies  = make(entity.Hobbies, 0, len(input.Hobbies))
		userHobbies = make(entity.UserHobbies, len(input.Hobbies))
	)
	for i, hobby := range input.Hobbies {
		if hobby.ID == "" {
			hobby.ID = austerid.Generate[entity.HobbyID]()
			newHobbies = append(newHobbies, hobby)
		}
		userHobbies[i] = &entity.UserHobby{
			UserID:  input.ID,
			HobbyID: hobby.ID,
		}
	}

	if len(newHobbies) > 0 {
		if err := u.repository.Hobby().Create(ctx, newHobbies...); err != nil {
			return err
		}
	}

	if err := u.repository.UserHobby().DeleteByUserID(ctx, input.ID); err != nil {
		return err
	}
	if err := u.repository.UserHobby().Create(ctx, userHobbies...); err != nil {
		return err
	}

	return u.repository.User().Update(ctx, &entity.User{
		ID:          input.ID,
		Name:        input.Name,
		Gender:      input.Gender,
		Age:         input.Age,
		ProfilePath: path,
	})
}
