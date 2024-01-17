package service

import (
	"github.com/eron97/login-authenticator.git/src/models"
	"github.com/eron97/login-authenticator.git/src/repository"
	"github.com/eron97/login-authenticator.git/src/service/crypto"
	"github.com/gin-gonic/gin"
)

type DomainService interface {
	CreateUser(request models.CreateUser) string
	ReadAllUsers(c *gin.Context) ([]models.GetUsers, error)
}

type useDomainService struct {
	userRepository repository.Database
}

func NewDomainService(
	db repository.Database,
) DomainService {
	return &useDomainService{db}
}

func (uds *useDomainService) ReadAllUsers(c *gin.Context) ([]models.GetUsers, error) {
	return uds.userRepository.ReadAllUsers()
}

func (service *useDomainService) CreateUser(request models.CreateUser) string {
	emailExists, err := service.userRepository.VerificExist(request.Email)
	if err != nil {
		return "Erro ao verificar a existência do e-mail"
	}

	if !emailExists {
		newPassword, err := crypto.HashPassword(request.Password)
		if err != nil {
			return "Erro ao criptografar a senha"
		}

		request.Password = newPassword

		err = service.userRepository.CreateUser([]models.CreateUser{request})
		if err == nil {
			return "Usuário criado com sucesso!"
		} else {
			return "Erro ao criar usuário no banco de dados"
		}
	}

	return "E-mail já existe e está associado a outra conta"
}
