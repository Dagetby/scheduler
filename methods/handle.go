package methods

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/scheduler/students/all", allStudents).Methods("GET")
	myRouter.HandleFunc("/scheduler/students/{id}", studentsById).Methods("GET")
	myRouter.HandleFunc("/schedulers/group/{name}", groupByName).Methods("GET")
	myRouter.HandleFunc("/schedulers/teacher/{id}", teacherById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8084", myRouter))
}
