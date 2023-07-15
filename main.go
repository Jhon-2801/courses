package main

import (
	"log"
	"net/http"
	"time"
	"v/internal/course"
	"v/internal/enrollment"
	"v/internal/user"
	"v/pkg/boostrap"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db, err := boostrap.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	userRepo := user.NewRepo(db)
	userSrv := user.NewService(userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	courseRepo := course.NewRepo(db)
	courseSrv := course.NewService(courseRepo)
	courseEnd := course.MakeEndpoints(courseSrv)

	enrollRepo := enrollment.NewRepo(db)
	enrollSrv := enrollment.NewService(enrollRepo, userSrv, courseSrv)
	enrollEnd := enrollment.MakeEndpoints(enrollSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	router.HandleFunc("/courses", courseEnd.Create).Methods("POST")
	router.HandleFunc("/courses", courseEnd.GetAll).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Get).Methods("GET")
	router.HandleFunc("/courses/{id}", courseEnd.Update).Methods("PATCH")

	router.HandleFunc("/enrollments", enrollEnd.Create).Methods("POST")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
