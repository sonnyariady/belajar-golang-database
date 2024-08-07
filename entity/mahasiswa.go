package entity

//utk sqlserver

import "time"

type Mahasiswa struct {
	id            int
	Nama, Jurusan string
	TglLahir      time.Time
}
