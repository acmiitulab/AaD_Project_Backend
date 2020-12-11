package Movie

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


type MovieCollectionClass struct{
	dbcon *mongo.Database
}


func NewMovieCollection(config MongoConfig.MongoConfig) (MovieCollection, error){

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
	collection=db.Collection("Movies")
	return &MovieCollectionClass{dbcon: db,},nil
}


func(scc *MovieCollectionClass) GetObjects() ([]*Movie, error){

	findOptions:=options.Find()
	var objects []*Movie
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var object Movie
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


func (scc *MovieCollectionClass) AddObject(object *Movie) (*Movie,error){

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

func (scc *MovieCollectionClass) GetObject(id int64) (*Movie,error){

	filter:=bson.D{{"movieid",id}}
	object:=&Movie{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&object)
	if err!=nil{
		return nil,err
	}
	return object,nil

}


func (scc *MovieCollectionClass) DeleteObject(object *Movie) error{

	filter:=bson.D{{"movieid",object.MovieID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (scc *MovieCollectionClass) UpdateObject (object *Movie)  (*Movie, error){

	filter:=bson.D{{"movieid",object.MovieID}}
	update:=bson.D{{"$set",bson.D{
		{"moviename",object.MovieName},
		{"releasedate",object.ReleaseDate},
		{"director",object.Director},
		{"actors",object.Actors},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return object,nil
}
