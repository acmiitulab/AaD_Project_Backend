package Seats

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


type SeatsCollectionClass struct{
	dbcon *mongo.Database
}


func NewInternCollection(config MongoConfig.MongoConfig) (SeatsCollection, error){

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
	return &SeatsCollectionClass{dbcon:db,},nil
}


func(scc *SeatsCollectionClass) GetSeats() ([]*Seat, error){

	findOptions:=options.Find()
	var seats []*Seat
	cur,err :=collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err!=nil{
		return nil,err
	}
	for cur.Next(context.TODO()){
		var seat Seat
		err:=cur.Decode(&seat)
		if err!=nil{
			return nil,err
		}
		seats = append(seats, &seat)
	}
	if err:=cur.Err();err!=nil{
		return nil,err
	}
	cur.Close(context.TODO())
	return seats,nil
}


func (scc *SeatsCollectionClass) AddSeat(seat *Seat) (*Seat,error){

	seats,err:=scc.GetSeats()
	n:=len(seats)
	if n!=0{
		lastSeat:=seats[n-1]
		seat.SeatID = lastSeat.SeatID+1
	}else{
		seat.SeatID = 1
	}
	insertResult,err:=collection.InsertOne(context.TODO(), seat)
	if err!=nil{
		return nil,err
	}
	fmt.Println("Inserted document", insertResult.InsertedID)
	return seat,nil

}

func (scc *SeatsCollectionClass) GetSeat(id int64) (*Seat,error){

	filter:=bson.D{{"seat_id",id}}
	seat:=&Seat{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&seat)
	if err!=nil{
		return nil,err
	}
	return seat,nil

}


func (scc *SeatsCollectionClass) DeleteSeat(seat *Seat) error{

	filter:=bson.D{{"seat_id",seat.SeatID}}
	_,err:=collection.DeleteOne(context.TODO(),filter)
	if err!=nil{
		return err
	}
	return nil
}

func (scc *SeatsCollectionClass) UpdateSeat (seat *Seat)  (*Seat, error){

	filter:=bson.D{{"seat_id",seat.SeatID}}
	update:=bson.D{{"$set",bson.D{
		{"seat_number",seat.SeatNumber},
		{"hall_id",seat.HallID},
		{"isFree",seat.IsFree},
	}}}
	_,err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil{
		return nil,err
	}
	return seat,nil
}
func (scc *SeatsCollectionClass) GetSeatFromHall (id int64)  ([]*Seat, error)  {
	filter:=bson.D{{"hall_id",id}}
	options:=options.Find()
	var seats []*Seat
	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()){
		var seat Seat
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


