package controllers

import (
	"appoiment-backend/database"
	"appoiment-backend/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func createToken(email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 3600).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
  }

func Login(c *fiber.Ctx)  error{
	user:=new(models.LoginUser)
	if err := c.BodyParser(user); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}
	hashedPassword,err:=database.Login(user)

	if  err != nil{
		c.Status(409)
		return c.JSON(fiber.Map{"message": err.Error()})
	}
	if checkPasswordHash(user.Password,hashedPassword) {
		token,_:=createToken(user.Email)
		c.Set("Authorization",token)
		return c.SendString("Logged")
	}else{
		c.Status(409)
		return c.JSON(fiber.Map{"message":"Wrong password"})
	}
	
}

func SignUp(c *fiber.Ctx)  error{
	user:=new(models.SignupUser)
	if err := c.BodyParser(user); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message":"Required arguments not found"})
	}

	user.Password,_=hashPassword(user.Password)

	if err := database.Signup(user); err!=nil {
		c.Status(409)
		return c.JSON(fiber.Map{"message": err.Error()})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{"message": "Done"})
}