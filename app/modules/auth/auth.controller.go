package auth

import (
	"net/http"
	"wordora/app/modules/auth/dto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user by providing necessary details
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register Request"
// @Success 201 {object} dto.RegisterResponse
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.authService.Register(ctx, req)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.authService.Login(ctx, req)
}

// VerifyOTP godoc
// @Summary Verify an OTP code
// @Description Verify a user’s OTP code for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.VerifyOTPRequest true "Verify OTP Request"
// @Router /auth/verify-otp [post]
func (c *AuthController) VerifyOTP(ctx *gin.Context) {
	var req dto.VerifyOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.authService.VerifyOTP(ctx, req)
}

// ResendOTP godoc
// @Summary Resend OTP code
// @Description Request a new OTP code for verification
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.ResendOTPRequest true "Resend OTP Request"
// @Router /auth/resend-otp [post]
func (c *AuthController) ResendOTP(ctx *gin.Context) {
	var req dto.ResendOTPRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.authService.ResendOTP(ctx, req)
}
