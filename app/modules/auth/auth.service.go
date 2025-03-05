package auth

import (
	"fmt"
	"time"
	"wordora/app/modules/auth/dto"
	"wordora/app/modules/otp"
	"wordora/app/modules/profiles"
	profileEnums "wordora/app/modules/profiles/enums"
	"wordora/app/modules/users"
	"wordora/app/utils/common"
	"wordora/app/utils/hash"
	"wordora/app/utils/mail"
	"wordora/app/utils/paseto"

	"wordora/app/utils/uuid"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	userRepo    *users.UserRepository
	profileRepo *profiles.ProfileRepository
	otpRepo     *otp.OTPRepository
	tokenHelper *paseto.TokenHelper
}

func NewAuthService(userRepo *users.UserRepository, profileRepo *profiles.ProfileRepository, otpRepo *otp.OTPRepository, tokenHelper *paseto.TokenHelper) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		otpRepo:     otpRepo,
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
		ID:              uuid.GenerateUUID(),
		Name:            req.Name,
		Email:           req.Email,
		Password:        hashedPassword,
		Role:            "viewer",
		IsEmailVerified: false,
	}

	if err := s.userRepo.CreateUser(&user); err != nil {
		common.GenerateErrorResponse(ctx, 500, "Failed to create user", nil)
		fmt.Println("Error creating user:", err)
		return
	}

	otpCode, expirationTime := mail.GenerateOTP()

	otp := otp.UserOTP{
		ID:        uuid.GenerateUUID(),
		UserID:    user.ID,
		OTPCode:   otpCode,
		ExpiredAt: expirationTime,
	}

	if err := s.otpRepo.CreateOTP(&otp); err != nil {
		fmt.Println("Error creating OTP:", err)
		common.GenerateErrorResponse(ctx, 500, "Failed to generate OTP", nil)
		return
	}

	if err := mail.SendOTPEmail(user.Email, otpCode); err != nil {
		common.GenerateErrorResponse(ctx, 500, "Failed to send OTP email", nil)
		return
	}

	response := dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Profiles: dto.ProfileResponse{
			FullName:    req.FullName,
			Gender:      req.Gender,
			PhoneNumber: req.PhoneNumber,
			Address:     req.Address,
		},
		CreatedAt:  user.CreatedAt,
	}

	common.GenerateSuccessResponseWithData(ctx, "User registered successfully. Check your email for OTP.", response)
}

func (s *AuthService) Login(ctx *gin.Context, req dto.LoginRequest) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		common.GenerateErrorResponse(ctx, 401, "Invalid credentials", nil)
		return
	}

	if !user.IsEmailVerified {
		common.GenerateErrorResponse(ctx, 403, "Email not verified", nil)
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

func (s *AuthService) VerifyOTP(ctx *gin.Context, req dto.VerifyOTPRequest) {
	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil || user == nil {
		common.GenerateErrorResponse(ctx, 400, "User not found", nil)
		return
	}

	otp, err := s.otpRepo.GetOTPByUserID(user.ID)
	if err != nil || otp == nil {
		common.GenerateErrorResponse(ctx, 400, "Invalid or expired OTP", nil)
		return
	}

	if time.Now().After(otp.ExpiredAt) {
		s.otpRepo.DeleteOTPByUserID(user.ID)
		common.GenerateErrorResponse(ctx, 400, "OTP has expired", nil)
		return
	}

	if otp.OTPCode != req.OTP {
		common.GenerateErrorResponse(ctx, 400, "Incorrect OTP", nil)
		return
	}

	user.IsEmailVerified = true
	if err := s.userRepo.UpdateUser(user); err != nil {
		common.GenerateErrorResponse(ctx, 500, "Failed to update user", nil)
		return
	}

	s.otpRepo.DeleteOTPByUserID(user.ID)

	common.GenerateSuccessResponse(ctx, "OTP verified successfully. Email is now verified.")
}