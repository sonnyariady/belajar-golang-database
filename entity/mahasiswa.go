package entity

//utk sqlserver

import "time"

type Mahasiswa struct {
	Id       int
	Nama     string
	Jurusan  string
	TglLahir time.Time
}
