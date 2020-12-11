package Sessions

type SessionCollection interface {
	AddObject(object *Session) (*Session, error)
	GetObjects() ([]*Session,error)
	GetObject(id int64) (*Session, error)
	UpdateObject(object *Session) (*Session, error)
	DeleteObject(object *Session) error
	GetSessionsByMovie(id int64) ([]*Session, error)
	GetSessionsByCinema(id int64) ([]*Session, error)
}

type Session struct {
	SessionID      int64      `json:"session_id"`
	MovieID        int64      `json:"movie_id"`
	CinemaID       int64      `json:"cinema_id"`
	HallID		   int64	  `json:"hall_id"`
	Time	       string     `json:"time"`


}

