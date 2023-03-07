package models

type Student struct {
	Email   string `xorm:"varchar(256) not null unique(key) 'email'"`
	Teacher string `xorm:"varchar(256) not null unique(key) 'teacher'"`
}

type Suspend struct {
	Email string `xorm:"varchar(256) not null pk 'email'"`
}
