package Cinema

import (
	"EndTermArchitecture/MongoConfig"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
)


type CinemaCollectionClass struct{
	dbcon *mongo.Database
}


func NewCinemaCollection(config MongoConfig.MongoConfig) (CinemaCollection, error){

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
	collection=db.Collection("Cinema")
	return &CinemaCollectionClass{dbcon: db,},nil
}


func(scc *CinemaCollectionClass) GetObjects() ([]*Cinema, error){

	findOptions:=options.Find()
	var objects []*Cinema
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var object Cinema
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


func (scc *CinemaCollectionClass) AddObject(object *Cinema) (*Cinema,error){

	objects,err:=scc.GetObjects()
	n:=len(objects)
	if n!=0{
		lastSeat:=objects[n-1]
		object.CinemaID = lastSeat.CinemaID+1
	}else{
		object.CinemaID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), object)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document", insertResult.InsertedID)
	return object,nil

}

func (scc *CinemaCollectionClass) GetObject(id int64) (*Cinema,error){

	filter:=bson.D{{"cinemaid",id}}
	object:=&Cinema{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&object)
	if err!=nil{
		return nil,err
	}
	return object,nil

}


func (scc *CinemaCollectionClass) DeleteObject(object *Cinema) error{

	filter:=bson.D{{"cinemaid",object.CinemaID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (scc *CinemaCollectionClass) UpdateObject (object *Cinema)  (*Cinema, error){

	filter:=bson.D{{"cinemaid",object.CinemaID}}
	update:=bson.D{{"$set",bson.D{
		{"cinemaname",object.CinemaName},
		{"cinemalocation",object.CinemaLocation},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return object,nil
}
