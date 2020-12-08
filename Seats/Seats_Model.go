package Seats

type SeatsCollection interface {
	AddSeat(intern *Seat) (*Seat, error)
	GetSeats() ([]*Seat,error)
	GetSeat(id int64) (*Seat, error)
	UpdateSeat(intern *Seat) (*Seat, error)
	DeleteSeat(intern *Seat) error
	GetSeatFromHall (id int64)  ([]*Seat, error)
}

type Seat struct {
	SeatID        int64      `json:"seat_id"`
	SeatNumber    int   	 `json:"seat_number"`
	HallID 		  int64      `json:"hall_id"`
	IsFree		  bool		 `json:"is_free"`



}

