package Halls

type HallCollection interface {
	AddObject(object *Hall) (*Hall, error)
	GetObjects() ([]*Hall,error)
	GetObject(id int64) (*Hall, error)
	UpdateObject(object *Hall) (*Hall, error)
	DeleteObject(object *Hall) error
	GetObjectsFromParent (id int64)  ([]*Hall, error)
}

type Hall struct {
	HallID        int64      `json:"hall_id"`
	HallName      string   	 `json:"hall_name"`
	CinemaID 	  int64      `json:"cinema_id"`

}

