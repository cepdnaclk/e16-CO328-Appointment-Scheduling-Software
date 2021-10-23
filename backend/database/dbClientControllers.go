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
	filter := bson.M{"email": ownerEmail,"serviceid":serviceId}
	var result models.ServiceDayDetail
	if err := ServiceDayDetailsCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return false,err
	}

	for i := 0; i < len(result.DayDetails); i++ {
		if result.DayDetails[i].Date==day {
			for j := 0; j < len(result.DayDetails[i].SlotList); j++ {
				if result.DayDetails[i].SlotList[j].SlotId==slotId {
					if result.DayDetails[i].SlotList[j].ClientRequested==true {
						return false,nil
					}
					user,_:=GetUser(clientEmail)
					result.DayDetails[i].SlotList[j].ClientRequested=true
					result.DayDetails[i].SlotList[j].ClientEmail=clientEmail
					result.DayDetails[i].SlotList[j].ClintName=user.FristName+" "+user.LastName

					clientRequest:=models.ClientRequestedService{
						Date: day,
						ServiceOwnerEmail: ownerEmail,
						ClientEmail: clientEmail,
						ServiceID: serviceId,
						SlotId: slotId,
						Approved: false,
						Time: result.DayDetails[i].SlotList[j].Time,
					}

					if _, err := ClientRequestedCollection.InsertOne(context.TODO(), clientRequest);err!=nil{
						fmt.Println(err)
						return false,err
					}
					_, err := ServiceDayDetailsCollection.UpdateOne(
						context.TODO(),
						filter,
						bson.D{
							{"$set", bson.D{{"daydetails", result.DayDetails}}},
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
	filter := bson.M{"email": ownerEmail,"serviceid":serviceId}
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
					result.DayDetails[i].SlotList[j].ClintName=""
					filter_D := bson.M{"serviceowneremail": ownerEmail,"serviceid":serviceId,"clientemail":clientEmail,"slotid":slotId}

					if _, err := ClientRequestedCollection.DeleteOne(context.TODO(),filter_D);err!=nil{
						return err
					}
					_, err := ServiceDayDetailsCollection.UpdateOne(
						context.TODO(),
						filter,
						bson.D{
							{"$set", bson.D{{"daydetails", result.DayDetails}}},
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

func GetAllRequestedServices(email string) ([]*models.ClientRequestedServicesResponse,error)  {
	var returnVal []*models.ClientRequestedServicesResponse
	filter := bson.D{{"clientemail", email}}

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
		filter := bson.M{"email": result.ServiceOwnerEmail,"serviceid":result.ServiceID}
		var service models.Service
		if err := ServicesCollection.FindOne(context.TODO(), filter).Decode(&service);err!=nil {
			return nil,err
		}
		
		returnVal = append(returnVal, &models.ClientRequestedServicesResponse{
			ServiceOwnerEmail: result.ServiceOwnerEmail,
			ServiceID: result.ServiceID,
			SlotId: result.SlotId,
			Approved: result.Approved,
			Time: result.Time,
			ServiceName: service.ServiceName,
			ServiceDiscription: service.ServiceDiscription,
			Date: result.Date,
		})
	}

	return returnVal,nil
}

func SearchByName(serviceName string) ([]*models.ClientServiceResponse,error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var returnVal []*models.ClientServiceResponse
	cur, err := ServicesCollection.Find(context.TODO(), bson.D{{"servicename", primitive.Regex{Pattern: serviceName}}}, findOptions)
	
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