package middleware

import (
	"ecomresto/initializers"
	"ecomresto/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode/validate it

	// Parse takes the token string and a function for looking up the key. The latter is especially

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Find the user with token sub
		var user models.Users
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Attach to req
		c.Set("user", user)
		//Continue

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

	c.Next()
}

// import (
// 	"ecomresto/models"
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type AdminController struct {
// 	DB *gorm.DB
// }

// // Secret key for JWT signing
// var jwtKey = []byte("secret_key")

// // Struct for login credentials
// type AdminLoginInput struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// // Struct for JWT claims
// type Claims struct {
// 	Email string `json:"email"`
// 	jwt.StandardClaims
// }

// // Admin login function
// func (controller *AdminController) AdminLogin(c *gin.Context) {
// 	var input AdminLoginInput
// 	var admin models.Admin // Assuming you have an admin model

// 	// Bind JSON input to struct
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	// Verify email and password from the database
// 	if err := controller.DB.Where("email = ?", input.Email).First(&admin).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	if !CheckPasswordHash(input.Password, admin.Password) {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
// 		return
// 	}

// 	// Create JWT token after successful login
// 	expirationTime := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		Email: input.Email,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
// 		return
// 	}

// 	// Respond with the JWT token
// 	c.JSON(http.StatusOK, gin.H{
// 		"token":   tokenString,
// 		"expires": expirationTime,
// 	})
// }

// // Function to validate password hash
// func CheckPasswordHash(password, hash string) bool {
// 	// Implement your password hashing logic here
// 	return password == hash // Simplified for example purposes
// }
