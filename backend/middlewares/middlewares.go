package middlewares

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func extractToken(c *fiber.Ctx) (string,error){
	bearToken := string(c.Request().Header.Peek("Authorization"))
	
	if len(bearToken) == 0 {
	   return "",errors.New("No Authheader")
	}
	return bearToken,nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	   }
	   return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
  }
func AuthMiddleware(c *fiber.Ctx) error {
	bearToken ,err :=extractToken(c)
	if err!=nil {
		c.Status(403)
		return c.JSON(fiber.Map{"message": "Not found Auth Header"})
	}

	var token *jwt.Token

	token, err=verifyToken(bearToken)

  	if err != nil {
		c.Status(403)
		return c.JSON(fiber.Map{"message": err.Error()})
	}
	claims, ok := token.Claims.(jwt.MapClaims);  
  	if !ok && !token.Valid {
		c.Status(403)
		return c.JSON(fiber.Map{"message": "Authantication Failed"})
	}
	email, ok := claims["user_email"].(string)

    if !ok {
		fmt.Println(err)
		return err
    }
	  
	c.Set("Auth-Email",email)
	return c.Next()
}