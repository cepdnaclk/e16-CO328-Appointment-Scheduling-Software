package database

import (
	"appoiment-backend/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddNewService(service *models.Service,serDayDetails *models.ServiceDayDetail)bool{
	_, err := ServicesCollection.InsertOne(context.TODO(), service)
	if err!=nil {
		return false
	}
	_, err = ServiceDayDetailsCollection.InsertOne(context.TODO(), serDayDetails)
	if err!=nil {
		return false
	}
	return true
}

func GetAllServicesCreated(email string) ([]*models.ServiceResponse,error)  {
	
	var returnVal []*models.ServiceResponse
	filter := bson.D{{"email", email}}

	cur, err := ServicesCollection.Find(context.TODO(), filter, nil)
	if err != nil {
    	return nil,err
	}

	for cur.Next(context.TODO()) {
    
		var result models.Service
		err := cur.Decode(&result)
		if err != nil {
			return nil,err
		}

		returnVal = append(returnVal, &models.ServiceResponse{
			ServiceID: result.ServiceID,
			ServiceName: result.ServiceName,
			ServiceType: result.ServiceType,
			ServiceDiscription: result.ServiceDiscription,
		})
	}

	return returnVal,nil
}


func CheckExpired(email string,id string) (bool,error){
	var result models.Service
	filter := bson.M{"email": email,"serviceid":id}
	if err := ServicesCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return false,err
	}
	today := time.Now()

	return today.After(result.ExpiredDay),nil
}


func GetAllSlotsOfService(email string,id string)(*[]models.DayDetail,error)  {
	var result models.ServiceDayDetail
	filter := bson.M{"email": email,"serviceid":id}
	if err := ServiceDayDetailsCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return nil,err
	}
	return &result.DayDetails,nil
}

func DeleteService(email string,id string) bool{
	filter := bson.M{"email": email,"serviceid":id}
	
	if _, err := ServicesCollection.DeleteOne(context.TODO(), filter);err != nil {
		return false
	}
	
	if _, err := ServiceDayDetailsCollection.DeleteOne(context.TODO(), filter);err != nil {
		return false
	}
	filter_d := bson.M{"serviceowneremail": email,"serviceid":id}
	if _, err := ClientRequestedCollection.DeleteMany(context.TODO(), filter_d);err != nil {
		return false
	}
	return true

}


func UpdateApproved(email string,serviceId string,slotId int,day string) bool{
	filter_s := bson.M{"email": email,"serviceid":serviceId}
	var result models.ServiceDayDetail
	if err := ServiceDayDetailsCollection.FindOne(context.TODO(), filter_s).Decode(&result);err!=nil {
		return false
	}

	for i := 0; i < len(result.DayDetails); i++ {
		if result.DayDetails[i].Date==day {
			for j := 0; j < len(result.DayDetails[i].SlotList); j++ {
				if result.DayDetails[i].SlotList[j].SlotId==slotId {
					result.DayDetails[i].SlotList[j].Approved=true
					filter_c := bson.M{"clientemail": result.DayDetails[i].SlotList[j].ClientEmail,"serviceid":serviceId,"slotid":result.DayDetails[i].SlotList[j].SlotId,"serviceowneremail":email}

					if _, err := ClientRequestedCollection.UpdateOne(
						context.TODO(),
						filter_c,
						bson.D{
							{"$set", bson.D{{"approved", true}}},
						},
					) ;err!=nil{
						fmt.Println(err)
						return false
					}

					_, err := ServiceDayDetailsCollection.UpdateOne(
						context.TODO(),
						filter_s,
						bson.D{
							{"$set", bson.D{{"daydetails", result.DayDetails}}},
						},
					)
					if err != nil {
						fmt.Println(err)
						return false
					}
					return true
				}
			}
		}
	}
	return false
}

func UpdateRemovedClient(email string,serviceId string,slotId int,day string) bool  {
	filter := bson.M{"email": email,"serviceid":serviceId}
	var result models.ServiceDayDetail
	if err := ServiceDayDetailsCollection.FindOne(context.TODO(), filter).Decode(&result);err!=nil {
		return false
	}

	for i := 0; i < len(result.DayDetails); i++ {
		if result.DayDetails[i].Date==day {
			for j := 0; j < len(result.DayDetails[i].SlotList); j++ {
				if result.DayDetails[i].SlotList[j].SlotId==slotId {
					result.DayDetails[i].SlotList[j].Approved=false
					result.DayDetails[i].SlotList[j].ClientRequested=false
					result.DayDetails[i].SlotList[j].ClintName=""
					result.DayDetails[i].SlotList[j].ClientEmail=""
					_, err := ServiceDayDetailsCollection.UpdateOne(
						context.TODO(),
						filter,
						bson.D{
							{"$set", bson.D{{"daydetails", result.DayDetails}}},
						},
					)
					if err != nil {
						fmt.Println(err)
						return false
					}
					filter_d := bson.M{"serviceowneremail": email,"serviceid":serviceId,"slotid":slotId}
					if _, err := ClientRequestedCollection.DeleteOne(context.TODO(), filter_d);err != nil {
						return false
					}
					return true
				}
			}
		}
	}
	return false
}