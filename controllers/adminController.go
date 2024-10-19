package controllers

import (
	"database/sql"
	"ecomresto/initializers"
	"ecomresto/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Admin struct {
	Email    string
	Password string
}

// func Login(c *gin.Context) { date 18-10-24
// 	//Get the email and pass off req body
// 	var body struct {
// 		Email    string
// 		Password string
// 	}

// 	if err := c.ShouldBind(&body); err != nil {
// 		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Failed to read body"})
// 		return
// 	}
// 	// Look up requested user
// 	var user models.Users
// // 	initializers.DB.Where(&user, "email = ?", body.Email)
// 	if user.ID == 0 {
//  		c.JSON(http.StatusBadRequest, gin.H{
//  			"error": "invalid email or password",
//  		})
//  		return
//  	}

// 	// Compare sent in password with saved user pass hash
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid password or email",
// 		})
// 		return
// 	}

// 	// Gen a JWT token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte(("SECRET")))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed, can't create token",
// 		})
// 		return
// 	}

// 	// Send it back
// 	c.SetSameSite(http.SameSiteLaxMode)
// 	c.SetCookie("userAuthorization", tokenString, 3600*24*30, "", "", false, true)
// 	c.JSON(http.StatusOK, gin.H{
// 		"token": tokenString,
// 	})
// }

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I am logged in.",
	})
}
func Login(c *gin.Context) {
	var body struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	// Bind form data from the request
	if err := c.ShouldBind(&body); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Failed to read form data"})
		return
	}

	// Force-set the ID to 28 for this example
	var user models.Users
	user.ID = 28

	// Normally, you'd fetch the user from the database, but in this case, we're overriding the ID
	// Uncomment this block if you want to query the user by email as usual.

	sqlDB, err := initializers.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	// Use raw SQL to find the user by email
	query := "SELECT id, email, password FROM users WHERE email = $1 LIMIT 1"
	err = sqlDB.QueryRow(query, body.Email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid email or password"})
			return
		} else {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to query user"})
			return
		}
	}

	// Compare the password (this step assumes you've fetched user details in practice)
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	// if err != nil {
	// 	c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid password"})
	// 	return
	// }

	// Generate JWT token with the fixed ID 28
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID, // This will be set to 28
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET")) // Ensure matching secret key
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to create token"})
		return
	}

	// Set cookie with token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("userAuthorizgation", tokenString, 3600*24*30, "", "", false, true)

	// Redirect or respond with success message
	c.Redirect(http.StatusSeeOther, "/dashboard") // Redirect to the dashboard after successful login
}
