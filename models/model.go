package models

type Teacher struct {
	Email string `db:"Email"`
}

type Student struct {
	Email   string `"db:"Email"`
	Teacher string `db:TeacherEmail`
}

type Suspended struct {
	Email string `"db:"Email"`
}
