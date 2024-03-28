package service

import (
	"net/http"
	"strings"

	"github.com/bagasjs/go-backend-template/app/core"
	"github.com/bagasjs/go-backend-template/app/entity"
	"github.com/bagasjs/go-backend-template/app/model"
	"github.com/bagasjs/go-backend-template/app/repository"
)

type UserService struct {
    UserRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
    return &UserService{
        UserRepository: repository,
    }
}

func (service *UserService) List() ([]model.GeneralUserResponse, *core.Error) {
    users, err:= service.UserRepository.ListAll()
    if err != nil {
        return []model.GeneralUserResponse{}, err
    }

    usersModels := []model.GeneralUserResponse{}
    for _, user := range users {
        userModel := model.GeneralUserResponse {
            ID: user.ID,
            Name: user.Name,
            Email: user.Email,
            Created: user.Created,
            Updated: user.Updated,
        }
        usersModels = append(usersModels, userModel)
    }
    return usersModels, nil
}

func (service *UserService) Find(id int) (model.GeneralUserResponse, *core.Error) {
    response := model.GeneralUserResponse{}
    query := core.NewQueryBuilder()
    users, err:=  service.UserRepository.Query(query.Where("id", "=", id).Limit(1))
    if err != nil {
        return response, err
    }

    if len(users) == 0 {
        return response, core.NewError(http.StatusNotFound, "Failed to find user")
    }

    response.Name = users[0].Name
    response.Email = users[0].Email
    response.ID = users[0].ID
    response.Created = users[0].Created
    response.Updated = users[0].Updated
    return response, nil
}

func (service *UserService) Create(request model.CreateUpdateUserRequest) *core.Error {
    if strings.Compare(request.Password, request.PasswordConfirmation) != 0 {
        return core.NewError(http.StatusForbidden, "Password and it's confirmation should be equal")
    }

    entt := entity.User {
        Name: request.Name,
        Password: request.Password,
        Email: request.Email,
    }

    if err := service.UserRepository.Insert(entt); err != nil {
        return err
    }

    return nil
}

func (service *UserService) Update(id int, request model.CreateUpdateUserRequest) *core.Error {
    if strings.Compare(request.Password, request.PasswordConfirmation) != 0 {
        return core.NewError(http.StatusForbidden, "Password and it's confirmation should be equal")
    }

    entt := entity.User {
        Name: request.Name,
        Email: request.Email,
        Password: request.Password,
    }
    if err := service.UserRepository.Update(id, entt); err != nil {
        return err
    }

    return nil
}

func (service *UserService) Destroy(id int) *core.Error {
    if err := service.UserRepository.Destroy(id); err != nil {
        return err
    }
    return nil
}

