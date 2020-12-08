package Halls

import (
	"EndTermArchitecture/MongoConfig"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	collection *mongo.Collection
)


type HallCollectionClass struct{
	dbcon *mongo.Database
}


func NewHallCollection(config MongoConfig.MongoConfig) (HallCollection, error){

	clientOptions:=options.Client().ApplyURI("mongodb://"+config.Host+":"+config.Port)
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		return nil,err
	}
	err = client.Ping(context.TODO(),nil)
	if err!=nil{
		return nil,err
	}

	db:=client.Database(config.Database)
	collection=db.Collection("Seats")
	return &HallCollectionClass{dbcon:db,},nil
}


func(scc *HallCollectionClass) GetObjects() ([]*Hall, error){

	findOptions:=options.Find()
	var objects []*Hall
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var object Hall
		err:=cur.Decode(&object)
		if err!=nil{
			return nil,err
		}
		objects = append(objects, &object)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return objects,nil
}


func (scc *HallCollectionClass) AddObject(object *Hall) (*Hall,error){

	objects,err:=scc.GetObjects()
	n:=len(objects)
	if n!=0{
		lastSeat:=objects[n-1]
		object.HallID = lastSeat.HallID+1
	}else{
		object.HallID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), object)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document", insertResult.InsertedID)
	return object,nil

}

func (scc *HallCollectionClass) GetObject(id int64) (*Hall,error){

	filter:=bson.D{{"hall_id",id}}
	object:=&Hall{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&object)
	if err!=nil{
		return nil,err
	}
	return object,nil

}


func (scc *HallCollectionClass) DeleteObject(object *Hall) error{

	filter:=bson.D{{"hall_id",object.HallID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (scc *HallCollectionClass) UpdateObject (object *Hall)  (*Hall, error){

	filter:=bson.D{{"seat_id",object.HallID}}
	update:=bson.D{{"$set",bson.D{
		{"hall_name",object.HallName},
		{"cinema_id",object.CinemaID},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return object,nil
}
func (scc *HallCollectionClass) GetObjectsFromParent (id int64)  ([]*Hall, error)  {
	filter:=bson.D{{"cinema_id",id}}
	options:=options.Find()
	var objects []*Hall
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var object Hall
		err:=cur.Decode(&object)
		if err!=nil{
			return nil,err
		}
		objects = append(objects,&object)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return objects,nil
}


