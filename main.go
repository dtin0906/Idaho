package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo -implement DB
	restaurants = append(restaurants, Restaurant{NAME: "Krung Thep Thai Cuisine", Price: "Low", Location: "366 Shoup Ave, Idaho Falls, ID 83402", Hours: &Hours{Monday: "11am - 9:30 pm", Tuesday: "11am - 9:30 pm", Wednesday: "11am - 9:30 pm", Thursday: "11am - 9:30 pm", Friday: "11am - 9:30 pm", Saturday: "11am - 9:30 pm", Sunday: "4:30pm - 9:30pm"}})
	restaurants = append(restaurants, Restaurant{NAME: "The Healthier Place To Eat", Price: "Moderate", Location: "121 S Main St, Pocatello, ID 83204", Hours: &Hours{Monday: "11:30am - 8pm", Tuesday: "11:30am - 8pm", Wednesday: "11:30am - 8pm", Thursday: "11:30am - 8pm", Friday: "11:30am - 8pm", Saturday: "11:30am - 8pm", Sunday: "Closed On Sundays"}})
	//Route Handlers
	r.HandleFunc("/idahoapi/restaurants", getRestaurants).Methods("GET")
	r.HandleFunc("/idahoapi/restaurants{name}", getRestaurant).Methods("GET")
	r.HandleFunc("/idahoapi/restaurants", createRestaurant).Methods("POST")
	r.HandleFunc("/idahoapi/restaurant/{name}", updateRestaurant).Methods("PUT")
	r.HandleFunc("/idahoapi/restaurants{name}", deleteRestaurant).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Restaurant Struct
type Restaurant struct {
	NAME     string `json:"name"`
	Price    string `json:"price"`
	Location string `json:"location"`
	Hours    *Hours `json:"hours"`
}

//Hours Struct
type Hours struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

var restaurants []Restaurant

//Get all Restaurants
func getRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", " application/json")
	json.NewEncoder(w).Encode(restaurants)
}

//Get A Restaurant
func getRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find correct id
	for _, item := range restaurants {
		if item.NAME == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Restaurant{})
}

// func getRestaurant(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	//for _, item := range restaurants {
// 	switch item := range restaurants {
// 	case item.NAME == params["name"]:
// 		{
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}

// 	case item.price == params["Price"]:
// 		{
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	case item.location == params["Location"]:
// 		{
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// }

//Create A New Restaurant
func createRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var restaurant Restaurant
	_ = json.NewDecoder(r.Body).Decode(&restaurant)
	restaurants = append(restaurants, restaurant)
	json.NewEncoder(w).Encode(restaurant)
}

//Update Restaurant
func updateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range restaurants {
		if item.NAME == params["name"] {
			restaurants = append(restaurants[:index], restaurants[index+1:]...)
			var restaurant Restaurant
			_ = json.NewDecoder(r.Body).Decode(&restaurant)
			json.NewEncoder(w).Encode(restaurant)
			return
		}
	}
	json.NewEncoder(w).Encode(restaurants)
}

// Delete Restuarant
func deleteRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range restaurants {
		if item.NAME == params["name"] {
			restaurants = append(restaurants[:index], restaurants[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(restaurants)
}

//Route Handlers
// r.HandleFunc("/api/books", getBooks).Methods("GET")
// r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
// r.HandleFunc("/api/books", createBook).Methods("POST")
// r.HandleFunc("/api/book/{id}", updateBook).Methods("PUT")
// r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":8000", r))
// }

// // Book Struct (Model)
// type Book struct {
// 	ID     string  `json:"id"`
// 	Isbn   string  `json:"isbn"`
// 	Title  string  `json:"title"`
// 	Author *Author `json:"author"`
// }

// //Author Struct
// type Author struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// //Init Books Variable as a slice Book Struct
// var books []Book

// // Get All Books
// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(books)
// }

// //Get A Book
// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Get params
// 	// Loop through books and find correct id
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// //Create A New Book
// func createBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe ( don't use in production beacuase it could generate the same id potentially)
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)
// }

// //Update The Book
// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// // Delete book
// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }
