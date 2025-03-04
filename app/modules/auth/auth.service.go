package auth

import (
	"wordora/app/modules/auth/dto"
	"wordora/app/modules/profiles"
	profileEnums "wordora/app/modules/profiles/enums"
	"wordora/app/modules/users"
	"wordora/app/utils/common"
	"wordora/app/utils/hash"
	"wordora/app/utils/paseto"

	"wordora/app/utils/uuid"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	userRepo    *users.UserRepository
	profileRepo *profiles.ProfileRepository
	tokenHelper *paseto.TokenHelper
}

func NewAuthService(userRepo *users.UserRepository, profileRepo *profiles.ProfileRepository, tokenHelper *paseto.TokenHelper) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		tokenHelper: tokenHelper,
	}
}

func (s *AuthService) Register(ctx *gin.Context, req dto.RegisterRequest) {
	if err := profileEnums.ValidateGender(req.Gender); err != nil {
		common.GenerateErrorResponse(ctx, 400, err.Error(), nil)
		return
	}

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		common.GenerateErrorResponse(ctx, 500, "Error hashing password", nil)
		return
	}

	user := users.User{
		ID:       uuid.GenerateUUID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "viewer",
	}

	if err := s.userRepo.CreateUser(&user); err != nil {
		common.GenerateErrorResponse(ctx, 500, "Failed to create user", nil)
		return
	}

	profile := profiles.Profile{
		ID:          uuid.GenerateUUID(),
		UserID:      user.ID,
		FullName:    req.FullName,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}

	if err := s.profileRepo.CreateProfile(&profile); err != nil {
		common.GenerateErrorResponse(ctx, 500, "Failed to create profile", nil)
		return
	}

	response := dto.RegisterResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		FullName:    profile.FullName,
		Gender:      profile.Gender,
		PhoneNumber: profile.PhoneNumber,
		Address:     profile.Address,
	}

	common.GenerateSuccessResponseWithData(ctx, "User registered successfully", response)
}



func (s *AuthService) Login(ctx *gin.Context, req dto.LoginRequest) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		common.GenerateErrorResponse(ctx, 401, "Invalid credentials", nil)
		return
	}

	if !hash.CheckPassword(req.Password, user.Password) {
		common.GenerateErrorResponse(ctx, 401, "Invalid credentials", nil)
		return
	}

	token, err := s.tokenHelper.GenerateToken(user)
	if err != nil {
		common.GenerateErrorResponse(ctx, 500, "Failed to generate token", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Login successful", dto.LoginResponse{Token: token})
}
