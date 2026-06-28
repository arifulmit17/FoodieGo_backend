package user

import (
	"errors"
	"foodiego/internal/dto"
	"foodiego/internal/models"
	repository "foodiego/internal/user"
	"foodiego/internal/utils"

	"gorm.io/gorm"
)

type UserService interface {
	Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	GetProfile(id string) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register User
func (s *userService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {

	// Check if email already exists
	_, err := s.userRepo.FindByEmail(req.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

// Login
func (s *userService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

// Get Logged-in User
func (s *userService) GetProfile(id string) (*dto.UserResponse, error) {

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
