package users

import (
	"net/http"
	"wordora/app/modules/users/dto"
	"wordora/app/utils/common"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

// GetMe godoc
// @Summary Get current authenticated user
// @Description Retrieve details of the authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Router /users/me [get]
// @Security BearerAuth
func (c *UserController) GetMe(ctx *gin.Context) {
	userID, exists := ctx.Get("sub")
	if !exists {
		common.GenerateErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	user, err := c.service.GetMe(userID.(string))
	if err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve user", nil)
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "User retrieved successfully", user)
}

// UpdateUserRole godoc
// @Summary Update user role
// @Description Update the role of a specific user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body dto.UpdateUserRoleRequest true "Role"
// @Router /users/{id}/role [put]
// @Security BearerAuth
func (c *UserController) UpdateUserRole(ctx *gin.Context) {
	userID := ctx.Param("id")
	var req dto.UpdateUserRoleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	if err := c.service.UpdateUserRole(userID, req.Role); err != nil {
		common.GenerateErrorResponse(ctx, http.StatusInternalServerError, "Failed to update role", nil)
		return
	}


	common.GenerateSuccessResponse(ctx, "User role updated successfully")
}
