package models

type Student struct {
	Email       string `xorm:"varchar(256) not null pk 'email'"`
	Teacher     string `xorm:"varchar(256) 'teacher'"`
	IsSuspended bool   `xorm:"bool 'is_suspended'"`
}
