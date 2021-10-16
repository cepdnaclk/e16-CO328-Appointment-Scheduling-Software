package database

import (
	"appoiment-backend/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(user *models.LoginUser) (string,error){
	var result models.SignupUser
	filter := bson.D{{"email", user.Email}}
	if err := UserCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return "",err
	}
	return result.Password,nil
}


func Signup(user *models.SignupUser) error {
	var result models.SignupUser
	filter := bson.D{{"email", user.Email}}
	err := UserCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil || result != (models.SignupUser{}){
		
		return errors.New("Error Occured")
	}
	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}