package Movie

type MovieCollection interface {
	AddObject(object *Movie) (*Movie, error)
	GetObjects() ([]*Movie,error)
	GetObject(id int64) (*Movie, error)
	UpdateObject(object *Movie) (*Movie, error)
	DeleteObject(object *Movie) error
}

type Movie struct {
	MovieID        int64      `json:"hall_id"`
	MovieName      string     `json:"movie_name"`
	ReleaseDate    string     `json:"release_date"`
	Director	   string	  `json:"director"`
	Actors 		  []string	  `json:"actors"`

}

