package controllers

import (
	"ecomresto/initializers"
	"ecomresto/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// func Signup(c *gin.Context) {
// 	var body struct {
// 		Username  string `form:"username" json:"username" binding:"required"`
// 		Email     string `form:"email" json:"email" binding:"required,email"`
// 		Password  string `form:"password" json:"password" binding:"required"`
// 		CPassword string `form:"c_password" json:"c_password" binding:"required"`
// 		Phone     int    `form:"phone" json:"phone" binding:"required"`
// 	}

// 	if err := c.ShouldBind(&body); err != nil {
// 		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Failed to read body"})
// 		return
// 	}

// 	user := models.Users{
// 		Username: body.Username,
// 		Email:    body.Email,
// 		Password: body.Password, // Hash the password before saving
// 		Phone:    body.Phone,
// 	}

// 	if err := initializers.DB.Create(&user).Error; err != nil {
// 		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.HTML(http.StatusOK, "signup.html", gin.H{"message": "User created successfully"})
// }

func Signup(c *gin.Context) {
	var body struct {
		Username  string `form:"username" json:"username" binding:"required"`
		Email     string `form:"email" json:"email" binding:"required,email"`
		Password  string `form:"password" json:"password" binding:"required"`
		CPassword string `form:"c_password" json:"c_password" binding:"required"`
		Phone     int    `form:"phone" json:"phone" binding:"required"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Failed to read body"})
		return
	}

	// Check if passwords match
	if body.Password != body.CPassword {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Passwords do not match"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user model with hashed password
	user := models.Users{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hashedPassword), // Store the hashed password
		Phone:    body.Phone,
	}

	// Save user to the database
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to create user"})
		return
	}

	// Success
	c.HTML(http.StatusOK, "signup.html", gin.H{"message": "User created successfully"})
}
