package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string   `json:"firstName"`
	Age       int      `json:"age"`
	Interests []string `json:"interests"`
	Bio       string   `json:"bio"`
	Drink     string   `json:"drink"`
	Picture   string   `json:"picture"`
}

//func homePage(w http.ResponseWriter, _ *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]interface{}{
//		"message": "Welcome to MinkUp!",
//	})
//}

func createUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    w.Header().Set("Content-Type", "application/json")

    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    db.Create(&user)

    json.NewEncoder(w).Encode(user)
}

func getUsers(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		query := r.URL.Query().Get("q")
		var users []User
		if query != "" {
			if err := db.Where("first_name LIKE ?", "%"+query+"%").Find(&users).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			if err := db.Find(&users).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		json.NewEncoder(w).Encode(users)
	}
}

func getUserByID(id int, db *gorm.DB) (User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func getUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := getUserByID(id, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    w.Header().Set("Content-Type", "application/json")
    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var updatedUser User
    err = json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user User
    result := db.First(&user, id)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusNotFound)
        return
    }

    user.FirstName = updatedUser.FirstName
    user.Age = updatedUser.Age
    user.Interests = updatedUser.Interests
    user.Bio = updatedUser.Bio
    user.Drink = updatedUser.Drink
    user.Picture = updatedUser.Picture

    db.Save(&user)

    json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}

	db.Delete(&user)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func getMatches(userID int, db *gorm.DB) ([]User, error) {
    var user User
    if err := db.First(&user, userID).Error; err != nil {
        return nil, err
    }

    var matches []User
    if err := db.Where("id != ?", user.ID).Where("array_to_string(interests, ',') ILIKE ANY (array_to_string(?, ',')::text[])", user.Interests).Find(&matches).Error; err != nil {
        return nil, err
    }

    return matches, nil
}

func handleRequests(db *gorm.DB) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		createUser(w, r, db)
	}).Methods("POST")
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		updateUser(w, r, db)
	}).Methods("PUT")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteUser(w, r, db)
	}).Methods("DELETE")
	router.HandleFunc("/users/{id}/matches", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
	
		matches, err := getMatches(id, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(matches)
	}).Methods("GET")


	// Serve the HTML page for creating a new user
	router.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "create_user.html")
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
    dsn := "host=localhost user=postgres password=postgres dbname=minkup port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&User{})
    seed(db) 

    handleRequests(db)
}
