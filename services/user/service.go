package user

import (
	"context"
	"fmt"
	"krown/common/types"
	"krown/config"
	"krown/db"
	protouser "krown/services/genproto/user"
	"krown/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userStore *Store
}

func NewUserService(userStore *Store) *UserService {
	return &UserService{userStore}
}

func (s *UserService) Register(c *fiber.Ctx, payload db.CreateUserParams) *types.ServiceResponse {
	count, err := s.userStore.CheckUserByEmail(payload.Email)
	if err != nil {
		return utils.NewServiceResponse(fiber.StatusInternalServerError, fmt.Sprintf("Error getting the user by email, %s", err))
	}
	if count != 0 {
		return utils.NewServiceResponse(fiber.StatusConflict, "User with the same email already exists")
	}

	if err := s.userStore.CreateUser(c, payload); err != nil {
		return utils.NewServiceResponse(fiber.StatusConflict, "Error creating user")
	}
	return nil
}

func (s *UserService) Login(c *fiber.Ctx, payload types.LoginUserPayload) (string, *types.ServiceResponse) {
	user, err := s.userStore.userQueries.GetUserByEmail(c.Context(), payload.Email)
	if err != nil {
		return "", utils.NewServiceResponse(fiber.StatusConflict, "Credentials not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", utils.NewServiceResponse(fiber.StatusConflict, "Credentials not found")
	}

	token := jwt.New(jwt.SigningMethodHS512)
	secretKey := []byte(config.Envs.SecretKey)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", utils.NewServiceResponse(fiber.StatusInternalServerError, "Error generating token")
	}
	return signedToken, nil
}

//TODO: review
func (u *UserService) ValidateAuth(c context.Context, req *protouser.AuthRequest) error {
	tokenString := req.Token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Envs.SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
