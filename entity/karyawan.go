package entity

//utk mysql
import "database/sql"

type KaryawanTable struct {
	Id           int64
	Nama         string
	Jabatan      string
	Married      bool
	Pasangan     sql.NullString
	Tanggallahir sql.NullTime
	Gaji         sql.NullInt64
}
