package main

import (
	"EndTermArchitecture/Cinema"
	"EndTermArchitecture/Hall"
	"EndTermArchitecture/MongoConfig"
	"EndTermArchitecture/Movie"
	"EndTermArchitecture/Seats"
	"EndTermArchitecture/Sessions"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func run()  {





	fs := noDirListing(http.FileServer(http.Dir("./public/static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	router:=mux.NewRouter()
	conf:=MongoConfig.MongoConfig{
		Host:     "localhost",
		Database: "EndtermADIS",
		Port:     "27017",
	}

	////// Cinema
	cinemacol,err:=Cinema.NewCinemaCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	cinemaendpoints:=Cinema.NewEndpointsFactory(cinemacol)

	////// Movies
	moviescol,err:=Movie.NewMovieCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	movieendpoints:=Movie.NewEndpointsFactory(moviescol)

	//////Sessions
	sessioncol,err:=Sessions.NewSessionCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	sessionsendpoints:=Sessions.NewEndpointsFactory(sessioncol)

	//////Halls
	hallcol,err:=Hall.NewHallCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}

	hallsendpoints:=Hall.NewEndpointsFactory(hallcol)

	/////Seats
	seatscol,err:=Seats.NewSeatCollection(conf)
	if err!=nil{
		log.Fatal(err)
	}
	hallseats := Seats.NewHallSeats(hallcol, seatscol)
	seatsendpoints:=Seats.NewEndpointsFactory(hallseats)




	router.Methods("GET").Path("/cinema").HandlerFunc(cinemaendpoints.GetObjects())
	router.Methods("GET").Path("/cinema/{id}").HandlerFunc(cinemaendpoints.GetObject("id"))
	router.Methods("DELETE").Path("/cinema/{id}").HandlerFunc(cinemaendpoints.DeleteObject("id"))
	router.Methods("PUT").Path("/cinema/{id}").HandlerFunc(cinemaendpoints.UpdateObject("id"))
	router.Methods("POST").Path("/cinema/").HandlerFunc(cinemaendpoints.AddObject())


	router.Methods("GET").Path("/movie").HandlerFunc(movieendpoints.GetObjects())
	router.Methods("GET").Path("/movie/{id}").HandlerFunc(movieendpoints.GetObject("id"))
	router.Methods("DELETE").Path("/movie/{id}").HandlerFunc(movieendpoints.DeleteObject("id"))
	router.Methods("PUT").Path("/movie/{id}").HandlerFunc(movieendpoints.UpdateObject("id"))
	router.Methods("POST").Path("/movie/").HandlerFunc(movieendpoints.AddObject())


	router.Methods("GET").Path("/session").HandlerFunc(sessionsendpoints.GetObjects())
	router.Methods("GET").Path("/session/{id}").HandlerFunc(sessionsendpoints.GetObject("id"))
	router.Methods("DELETE").Path("/session/{id}").HandlerFunc(sessionsendpoints.DeleteObject("id"))
	router.Methods("PUT").Path("/session/{id}").HandlerFunc(sessionsendpoints.UpdateObject("id"))
	router.Methods("POST").Path("/session/").HandlerFunc(sessionsendpoints.AddObject())
	router.Methods("GET").Path("/session/{id}/movie").HandlerFunc(sessionsendpoints.GetSessionsByMovie("id"))
	router.Methods("GET").Path("/session/{id}/cinema").HandlerFunc(sessionsendpoints.GetSessionsByCinema("id"))


	router.Methods("GET").Path("/hall").HandlerFunc(hallsendpoints.GetObjects())
	router.Methods("GET").Path("/hall/{id}").HandlerFunc(hallsendpoints.GetObject("id"))
	router.Methods("DELETE").Path("/hall/{id}").HandlerFunc(hallsendpoints.DeleteObject("id"))
	router.Methods("PUT").Path("/hall/{id}").HandlerFunc(hallsendpoints.UpdateObject("id"))
	router.Methods("POST").Path("/hall/").HandlerFunc(hallsendpoints.AddObject())
	router.Methods("GET").Path("/hall/{id}/cinema/").HandlerFunc(hallsendpoints.GetObjectFromParent("id"))


	router.Methods("GET").Path("/seat").HandlerFunc(seatsendpoints.GetSeats())
	router.Methods("GET").Path("/seat/{id}").HandlerFunc(seatsendpoints.GetSeat("id"))
	router.Methods("DELETE").Path("/seat/{id}").HandlerFunc(seatsendpoints.DeleteSeat("id"))
	router.Methods("PUT").Path("/seat/{id}").HandlerFunc(seatsendpoints.UpdateSeat("id"))
	router.Methods("POST").Path("/seat/").HandlerFunc(seatsendpoints.AddSeat())
	router.Methods("GET").Path("/seat/{id}/cinema/").HandlerFunc(seatsendpoints.GetSeatsFromHall("id"))



	fmt.Println("Server is running")
	http.ListenAndServe(":8080",router)



}




func noDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}