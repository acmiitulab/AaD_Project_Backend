package Sessions

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


type MovieCollectionClass struct{
	dbcon *mongo.Database
}


func NewSessionCollection(config MongoConfig.MongoConfig) (SessionCollection, error){

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
	collection=db.Collection("Sessions")
	return &MovieCollectionClass{dbcon: db,},nil
}


func(scc *MovieCollectionClass) GetObjects() ([]*Session, error){

	findOptions:=options.Find()
	var objects []*Session
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var object Session
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


func (scc *MovieCollectionClass) AddObject(object *Session) (*Session,error){

	objects,err:=scc.GetObjects()
	n:=len(objects)
	if n!=0{
		lastSeat:=objects[n-1]
		object.MovieID = lastSeat.MovieID+1
	}else{
		object.MovieID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), object)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document", insertResult.InsertedID)
	return object,nil

}

func (scc *MovieCollectionClass) GetObject(id int64) (*Session,error){

	filter:=bson.D{{"sessionid",id}}
	object:=&Session{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&object)
	if err!=nil{
		return nil,err
	}
	return object,nil

}


func (scc *MovieCollectionClass) DeleteObject(object *Session) error{

	filter:=bson.D{{"sessionid",object.MovieID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (scc *MovieCollectionClass) UpdateObject (object *Session)  (*Session, error){

	filter:=bson.D{{"sessionid",object.MovieID}}
	update:=bson.D{{"$set",bson.D{
		{"movieid", object.MovieID},
		{"cinemaid", object.CinemaID},
		{"hallid", object.HallID},
		{"time",object.Time},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return object,nil
}

func (scc *MovieCollectionClass) GetSessionsByMovie(id int64) ([]*Session, error) {
	filter:=bson.D{{"movieid",id}}
	options:=options.Find()
	var seats []*Session
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var seat Session
		err:=cur.Decode(&seat)
		if err!=nil{
			return nil,err
		}
		seats = append(seats,&seat)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return seats,nil
}
func (scc *MovieCollectionClass) GetSessionsByCinema(id int64) ([]*Session, error) {
	filter:=bson.D{{"cinemaid",id}}
	options:=options.Find()
	var seats []*Session
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var seat Session
		err:=cur.Decode(&seat)
		if err!=nil{
			return nil,err
		}
		seats = append(seats,&seat)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return seats,nil
}
