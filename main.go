package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`;
	Isbn string `json: "isbn`;
	Title string `json:"title"`;
	Director *Director `json:"direcctor"`;
}

type Director struct {
	Firstname string `json:"firstname`;
	Lastname string `json:"lastname"`;
}

var movies []Movie;

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");

	json.NewEncoder(w).Encode(movies);
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");
	
	// Return route variable & params
	params := mux.Vars(r);

	for index, item := range movies {
		if(item.ID == params["id"]) {
			movies = append(movies[:index], movies[index+1:]...);
			break;
		}
	}

	json.NewEncoder(w).Encode(movies);
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");

	// Return route variable & params
	params := mux.Vars(r);

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return;
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");

	var movie Movie;

	// Decode request payload
	err := json.NewDecoder(r.Body).Decode(&movie);

	if err != nil {
		log.Fatal(err);
	}

	// Generate ID and convert to string
	movie.ID = strconv.Itoa(rand.Intn(100000000));

	movies = append(movies, movie);  

	json.NewEncoder(w).Encode(movie);
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json");

	params := mux.Vars(r);

	var movie Movie;

	err := json.NewDecoder(r.Body).Decode(&movie);

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range movies {
		if item.ID == params["id"] {
			fmt.Print(movie);
			item.ID = params["id"];
			item.Director = movie.Director;
			item.Isbn = movie.Isbn;
			item.Title = movie.Title;
			json.NewEncoder(w).Encode(movie);
		}
	}

}

func main() {
	r := mux.NewRouter();

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"} });

	movies = append(movies, Movie{ID: "2", Isbn: "45425", Title: "Movie Two", Director: &Director{Firstname: "Steven", Lastname: "Smith"}});

	r.HandleFunc("/movies", getMovies).Methods("GET");
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET");
	r.HandleFunc("/movies", createMovie).Methods("POST");
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT");
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE");

	fmt.Printf("Starting server at PORT 8000");

	err := http.ListenAndServe(":8000", r);

	if err != nil {
		log.Fatal("Error:", err);
	}
}