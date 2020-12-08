package Seats

import (
	"EndTermArchitecture/Halls"
	"errors"
)

type HallSeats interface {
	CheckSeat(object *Seat)    (*Seat, error)
	AddSeat(object *Seat)    (*Seat, error)
	GetSeats()                  ([]*Seat,error)
	GetSeat(id int64)           (*Seat, error)
	UpdateSeat(object *Seat) (*Seat, error)
	DeleteSeat(object *Seat)            error
	GetSeatsFromHall (id int64)  ([]*Seat, error)
}

type HallSeatsClass struct {
	seat SeatsCollection
	hall Halls.HallsCollection
}

func NewHallSeats(hallscollection Halls.HallsCollection, seatscollection SeatsCollection )  HallSeats {
	return &HallSeatsClass{seat:seatscollection, hall: hallscollection}
}

func(obj *HallSeatsClass) CheckSeat (object *Seat)  (*Seat, error)  {

	if object.SeatID == 0 {
		return nil, errors.New("No Seat ID ")
	}
	if object.HallID == 0 {
		return nil, errors.New("No Hall ID ")
	}
	if object.SeatNumber == 0 {
		return nil, errors.New("No Seat number ")
	}
	_, err:=obj.hall.GetHall(object.HallID)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (obj *HallSeatsClass) GetSeats() ([]*Seat,error) {
	intern, err:=obj.seat.GetSeats()
	if err != nil {
		return nil, err
	}
	return intern, err
}

func (obj *HallSeatsClass) GetSeat(id int64)  (*Seat, error)  {

	intern, err := obj.seat.GetSeat(id)
	if err!= nil {
		return nil, err
	}
	return intern, nil
}

func (obj *HallSeatsClass)AddSeat (object *Seat)    (*Seat, error) {
	_, err := obj.CheckSeat(object)
	if err != nil {
		return nil, err
	}
	_, err = obj.seat.AddSeat(object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (obj *HallSeatsClass) UpdateSeat(object *Seat) (*Seat, error) {
	_, err := obj.CheckSeat(object)
	if err != nil {
		return nil, err
	}
	_, err = obj.seat.UpdateSeat(object)
	if err != nil {
		return nil, err
	}
	return object, nil

}
func (obj *HallSeatsClass) DeleteSeat(object *Seat)  error {
	if object.SeatID == 0 {
		return errors.New("NO ID")
	}
	err := obj.seat.DeleteSeat(object)
	if err != nil {
		return err
	}
	return err

}
func (obj *HallSeatsClass) GetSeatsFromHall (id int64)  ([]*Seat, error) {
	_, err := obj.hall.GetHall(id)
	if err != nil {
		return nil,err
	}
	interns, err:= obj.seat.GetSeatFromHall(id)
	if err != nil {
		return nil, err
	}
	return interns, nil

}









