# scheduler

Приложение написанно с целью отработки и треннировки навыков 
  - RESTful API 
  - PostgreSQL
  - SQL
  
  Использованные библеотеки:
  - "github.com/gorilla/mux"
  - "github.com/jinzhu/gorm/dialects/postgres"
  
  
  Методы реализованные в проекте: 
   
  ("/scheduler/students/all", allStudents).Methods("GET") - получение всех студентов из базы данны
	("/scheduler/students/{id}", studentsById).Methods("GET") - поиск определенного студента по id 
	("/schedulers/group/{name}", groupByName).Methods("GET") - поиск группы по имени
	("/schedulers/teacher/{id}", teacherById).Methods("GET") - поиск  учителя по id 
  
 
 В файле script.txt - лежит скрипт для создания и заполнения БД.
 
