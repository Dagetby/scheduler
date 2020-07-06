package methods

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"scheduler/models"
	"scheduler/routers"

	"github.com/gorilla/mux"
)

func check(text string, err error) {
	if err != nil {
		log.Fatalf(text, err.Error())
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func allStudents(w http.ResponseWriter, r *http.Request) {
	qu, err := routers.GetDB().Query("select name, first_name , last_name, age, " +
		"course, academic_perfomance from students right join group_ ON(group_.id = students.group_id )")
	check("Error on get able students", err)
	defer qu.Close()
	students := []models.Students{}
	for qu.Next() {
		student := models.Students{}
		err := qu.Scan(&student.GroupID.Name, &student.FirstName, &student.LastName, &student.Age,
			&student.Course, &student.Academic_prefomance)
		check("[SCAN]: ", err)
		students = append(students, student)
	}
	json.NewEncoder(w).Encode(students)
}

func studentsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	qu, err := routers.GetDB().Query("select students.id, name, first_name , last_name, course, "+
		"academic_perfomance, age from students right join group_ ON(group_.id = students.group_id )"+
		" where students.id = $1", id)
	check("[ById]: ", err)
	defer qu.Close()
	student := models.Students{}
	for qu.Next() {
		s := models.Students{}
		err := qu.Scan(&s.Id, &s.GroupID.Name, &s.FirstName, &s.LastName,
			&s.Age, &s.Course, &s.Academic_prefomance)
		check("[SCAN]: ", err)
		student = s
	}

	json.NewEncoder(w).Encode(student)

}

func groupByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	qu, err := routers.GetDB().Query("select * from group_ Where name=$1", name)
	check("[Query Group]: ", err)
	defer qu.Close()
	group := models.Group{}
	for qu.Next() {
		g := models.Group{}
		err := qu.Scan(&g.Id, &g.Name, &g.NumberOfPeople)
		check("[Scan Group]: ", err)
		group = g
	}
	json.NewEncoder(w).Encode(group)
}

func teacherById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	qu, err := routers.GetDB().Query("select teacher.id, first_name, "+
		"last_name, name, post from teacher join subject"+
		" on(subject.id = teacher.subject_id) Where teacher.id=$1", id)
	check("[ById teacher]: ", err)
	defer qu.Close()
	teacher := models.Teacher{}
	for qu.Next() {
		t := models.Teacher{}
		err := qu.Scan(&t.Id, &t.FirstName, &t.LastName, &t.SubjectId.Name, &t.Post)
		check("[SCAN teacher]: ", err)
		teacher = t
	}

	json.NewEncoder(w).Encode(teacher)

}
