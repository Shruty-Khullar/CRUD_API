package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" //GitHub - gorilla/mux: A powerful HTTP router and URL matcher for building Go web. The gorilla/mux package is perhaps the most famous Go router,
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isnb"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Fistname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//HTTP headers let the client and the server pass additional information with an HTTP request or response. An HTTP header consists of its case-insensitive name followed by a colon ( : ), then by its value.29-May-2022
	w.Header().Set("Content-Type", "application/json") //So the header is a key value pair with. Here for key "C-T" we have set the value "appliaction/json"
	json.NewEncoder(w).Encode(movies)                  //An object that encodes instances of a data type as JSON objects.
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	//in the request we will get id of the movie we want to delete along with movies slice
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range movies {
		if item.ID == parameters["id"] {
			movies = append(movies[:index], movies[index+1:]...) //left the index ele
			break
		}
		json.NewEncoder(w).Encode(movies)
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for _, item := range movies {
		if item.ID == parameters["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie                            //var of type Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //now as we need to create a movie so we need to decode the info sent in request and we will save decoded info in movie var
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

/*Steps for updating a movie is:
1. ste content type to application/json
2.loop over the movies range
3. delete the mpvie with id that you've sent
4. create newmovie-that we have sent in body of postman
*/

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index, item := range movies {
		if item.ID == parameters["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = parameters["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	//The name mux stands for "HTTP request multiplexer". Like the standard http. ServeMux , mux. Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "23456", Title: "Movie one", Director: &Director{Fistname: "Saurav", Lastname: "Sharma"}})
	movies = append(movies, Movie{ID: "2", Isbn: "23457", Title: "Movie Two", Director: &Director{Fistname: "Shruty", Lastname: "Khullar"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at Port 8080")
	//Port number 8080 is usually used for web servers. When a port number is added to the end of the domain name, it drives traffic to the web server. However, users can not reserve port 8080 for secondary web servers.
	log.Fatal(http.ListenAndServe(":8000", r)) //starting the router
}
