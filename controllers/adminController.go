package controllers

import (
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

func Login(c *gin.Context) {
	//Get the email and pass off req body
	var body struct {
		Email    string
		Password string
	}

	if err := c.ShouldBind(&body); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Failed to read body"})
		return
	}
	// Look up requested user
	var user models.Users
	initializers.DB.Where(&user, "email = ?", body.Email)
	// if user.ID == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalidf email or password",
	// 	})
	// 	return
	// }

	// // Compare sent in password with saved user pass hash
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "invalid password or email",
	// 	})
	// 	return
	// }

	// Gen a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed, can't create token",
		})
		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("userAuthorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I am loggged in.",
	})
}
