package models

type Group struct {
	Id             int32  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	NumberOfPeople int    `json:"number_of_people,omitempty"`
}

type Students struct {
	Id                  int32  `json:"id,omitempty"`
	FirstName           string `json:"first_name,omitempty"`
	LastName            string `json:"last_name,omitempty"`
	Age                 int32  `json:"age,omitempty"`
	Course              int32  `json:"course,omitempty"`
	Academic_prefomance int    `json:"academic_perfomance,omitempty"`
	GroupID             Group  `json:"group_id,omitempty"`
}

type Subject struct {
	Id   int32  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Teacher struct {
	Id        int32   `json:"id,omitempty"`
	FirstName string  `json:"firstName,omitempty"`
	LastName  string  `json:"lastName,omitempty"`
	SubjectId Subject `json:"subjects,omitempty"`
	Post      string  `json:"post,omitempty"`
}
