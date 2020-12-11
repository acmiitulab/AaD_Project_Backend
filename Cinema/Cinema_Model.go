package Cinema

type CinemaCollection interface {
	AddObject(object *Cinema) (*Cinema, error)
	GetObjects() ([]*Cinema,error)
	GetObject(id int64) (*Cinema, error)
	UpdateObject(object *Cinema) (*Cinema, error)
	DeleteObject(object *Cinema) error
}

type Cinema struct {
	CinemaID        int64      `json:"cinema_id"`
	CinemaName      string   	`json:"cinema_name"`
	CinemaLocation 	string      `json:"cinema_location"`


}

