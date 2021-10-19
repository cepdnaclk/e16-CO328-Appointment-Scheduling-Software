package database

import (
	"appoiment-backend/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClientServicses()  ([]*models.ClientServiceResponse,error){
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var returnVal []*models.ClientServiceResponse
	cur, err := ServicesCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
    	return nil,err
	}

	for cur.Next(context.TODO()) {
    
		var result models.Service
		err := cur.Decode(&result)
		if err != nil {
			return nil,err
		}
		var serviceOwner models.SignupUser

		filter := bson.D{{"email", result.Email}}

		if err := UserCollection.FindOne(context.TODO(), filter).Decode(&serviceOwner);err!=nil {
			return nil,err
		}
		returnVal = append(returnVal, &models.ClientServiceResponse{
			ServiceID: result.ServiceID,
			ServiceName: result.ServiceName,
			ServiceType: result.ServiceType,
			ServiceDiscription: result.ServiceDiscription,
			ServiceOwnerEmail: result.Email,
			OwnerFristName: serviceOwner.FristName,
			OwnerLastName: serviceOwner.LastName,
		})
	}

	return returnVal,nil
}




func ClientRequestingService(ownerEmail string,clientEmail string,serviceId string,slotId int,day string)(bool,error )  {
	filter := bson.M{"email": ownerEmail,"serviceId":serviceId}
	var result models.ServiceDayDetail
	if err := ServiceDayDetailsCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return false,err
	}

	for i := 0; i < len(result.DayDetails); i++ {
		if result.DayDetails[i].Date==day {
			for j := 0; j < len(result.DayDetails[i].SlotList); j++ {
				if result.DayDetails[i].SlotList[j].SlotId==slotId {
					if result.DayDetails[i].SlotList[j].ClientRequested==true {
						return false,errors.New("Already client Requested")
					}
					result.DayDetails[i].SlotList[j].ClientRequested=true
					result.DayDetails[i].SlotList[j].ClientEmail=clientEmail

					clientRequest:=models.ClientRequestedService{
						ServiceOwnerEmail: ownerEmail,
						ClientEmail: clientEmail,
						ServiceID: serviceId,
						SlotId: slotId,
						Approved: false,

					}

					if _, err := ClientRequestedCollection.InsertOne(context.TODO(), clientRequest);err!=nil{
						fmt.Println(err)
						return false,err
					}
					_, err := ServiceDayDetailsCollection.UpdateOne(
						context.TODO(),
						filter,
						bson.D{
							{"$set", bson.D{{"dayDetails", result.DayDetails}}},
						},
					)
					if err != nil {
						fmt.Println(err)
						return false,err
					}
					return true,nil
				}
			}
		}
	}
	return false,errors.New("Not found requested slot")
}

func CancelRequestedService(ownerEmail string,clientEmail string,serviceId string,slotId int,day string) error  {
	filter := bson.M{"email": ownerEmail,"serviceId":serviceId}
	var result models.ServiceDayDetail
	if err := ServiceDayDetailsCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return err
	}

	for i := 0; i < len(result.DayDetails); i++ {
		if result.DayDetails[i].Date==day {
			for j := 0; j < len(result.DayDetails[i].SlotList); j++ {
				if result.DayDetails[i].SlotList[j].SlotId==slotId {
					result.DayDetails[i].SlotList[j].ClientRequested=false
					result.DayDetails[i].SlotList[j].ClientEmail=""

					filter_D := bson.M{"serviceOwnerEmail": ownerEmail,"serviceId":serviceId,"clientEmail":clientEmail,"slotId":slotId}

					if _, err := ClientRequestedCollection.DeleteOne(context.TODO(),filter_D);err!=nil{
						return err
					}
					_, err := ServiceDayDetailsCollection.UpdateOne(
						context.TODO(),
						filter,
						bson.D{
							{"$set", bson.D{{"dayDetails", result.DayDetails}}},
						},
					)
					if err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return errors.New("Not found requested slot")
}

func GetAllRequestedServices(email string) ([]*models.ClientRequestedService,error)  {
	var returnVal []*models.ClientRequestedService
	filter := bson.D{{"email", email}}

	cur, err := ClientRequestedCollection.Find(context.TODO(), filter, nil)
	if err != nil {
    	return nil,err
	}

	for cur.Next(context.TODO()) {
    
		var result models.ClientRequestedService
		err := cur.Decode(&result)
		if err != nil {
			return nil,err
		}

		returnVal = append(returnVal, &result)
	}

	return returnVal,nil
}

func SearchByName(serviceName string) ([]*models.ClientServiceResponse,error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var returnVal []*models.ClientServiceResponse
	cur, err := ServicesCollection.Find(context.TODO(), bson.D{{"text", primitive.Regex{Pattern: serviceName, Options: ""}}}, findOptions)
	
	if err != nil {
    	return nil,err
	}

	for cur.Next(context.TODO()) {
    
		var result models.Service
		err := cur.Decode(&result)
		if err != nil {
			return nil,err
		}
		var serviceOwner models.SignupUser

		filter := bson.D{{"email", result.Email}}

		if err := UserCollection.FindOne(context.TODO(), filter).Decode(&serviceOwner);err!=nil {
			return nil,err
		}
		returnVal = append(returnVal, &models.ClientServiceResponse{
			ServiceID: result.ServiceID,
			ServiceName: result.ServiceName,
			ServiceType: result.ServiceType,
			ServiceDiscription: result.ServiceDiscription,
			ServiceOwnerEmail: result.Email,
			OwnerFristName: serviceOwner.FristName,
			OwnerLastName: serviceOwner.LastName,
		})
	}

	return returnVal,nil
}