package entity

//utk mysql
import "database/sql"

type KaryawanTable struct {
	id           int
	nama         string
	jabatan      string
	married      bool
	pasangan     sql.NullString
	tanggallahir sql.NullTime
	gaji         sql.NullInt64
}
