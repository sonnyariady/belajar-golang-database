package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

type Karyawan struct {
	nama, jabatan string
	married       bool
	pasangan      sql.NullString
	tanggallahir  sql.NullTime
	gaji          sql.NullInt64
}

func TestExecInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	sqlcmd := "insert into Mahasiswa(id, nama, jurusan) VALUES('002','Andi','Ilmu Hukum')"
	_, err := db.ExecContext(ctx, sqlcmd)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert mahasiswa sukses!")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	sqlquery := "select id,nama,jurusan from mahasiswa"
	baris, err := db.QueryContext(ctx, sqlquery)

	if err != nil {
		panic(err)
	}
	i := 0
	fmt.Println("Cetak kolomnya:")
	fmt.Println(baris.Columns())

	for baris.Next() {
		var id, nama, jurusan string
		err := baris.Scan(&id, &nama, &jurusan)

		if err != nil {
			panic(err)
		}

		fmt.Println("Data Baris ke-", i)
		fmt.Println("Id : ", id)
		fmt.Println("Nama : ", nama)
		fmt.Println("Jurusan : ", jurusan)
		i++
	}
	defer baris.Close()
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	sqlquery := "select id,nama,married, pasangan, jabatan, tanggallahir, gaji from karyawan"
	baris, err := db.QueryContext(ctx, sqlquery)

	if err != nil {
		panic(err)
	}
	i := 0
	for baris.Next() {
		var id int
		var gaji sql.NullInt32
		var nama, jabatan string
		var pasangan sql.NullString
		var married []byte
		var isMarried bool = false
		var tanggallahir sql.NullTime

		err := baris.Scan(&id, &nama, &married, &pasangan, &jabatan, &tanggallahir, &gaji)

		if err != nil {
			panic(err)
		}

		if len(married) > 0 && married[0] == 1 {
			isMarried = true
		}

		fmt.Println("Data Baris ke-", i)
		fmt.Println("Id : ", id)
		fmt.Println("Nama : ", nama)
		fmt.Println("Married : ", isMarried)
		if pasangan.Valid {
			fmt.Println("Pasangan : ", pasangan.String)
		}

		fmt.Println("Jabatan : ", jabatan)
		if tanggallahir.Valid {
			fmt.Println("Tanggal Lahir : ", tanggallahir.Time)
		}

		if gaji.Valid {
			fmt.Println("Gaji: ", gaji.Int32)
		}

		i++
	}
	defer baris.Close()
}

func TestQueryUser(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "user1'; #"
	password := "rahasia2"

	sqlquery := "select username from useraccount where username='" + username + "' and password='" + password + "' limit 1"
	fmt.Println("Script: ", sqlquery)
	baris, err := db.QueryContext(ctx, sqlquery)

	if err != nil {
		panic(err)
	}

	if baris.Next() {
		var username string
		err := baris.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
	defer baris.Close()
}

func TestQueryWithParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	username := "user1'; #"
	password := "rahasia2"

	sqlquery := "select username from useraccount where username=? and password=? limit 1"
	fmt.Println("Script: ", sqlquery)
	baris, err := db.QueryContext(ctx, sqlquery, username, password)

	if err != nil {
		panic(err)
	}

	if baris.Next() {
		var username string
		err := baris.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses login", username)
	} else {
		fmt.Println("Gagal login")
	}
	defer baris.Close()
}

func TestExecInsertWithParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	kode := "003"
	nama := "Herman"
	jurusan := "Farmasi"

	sqlcmd := "insert into Mahasiswa(id, nama, jurusan) VALUES(?,?,?)"
	_, err := db.ExecContext(ctx, sqlcmd, kode, nama, jurusan)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert mahasiswa sukses!")
}

func TestExecInsertWithResultId(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	nama := "Herman"
	married := true
	pasangan := "Dewi"
	tanggallahir := time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local)
	jabatan := "Accounting"
	gaji := 20

	sqlcmd := "insert into Karyawan(nama, married, pasangan, jabatan, tanggallahir, gaji) VALUES(?,?,?,?,?,?)"
	hasil, err := db.ExecContext(ctx, sqlcmd, nama, married, pasangan, jabatan, tanggallahir, gaji)

	if err != nil {
		panic(err)
	}

	insertId, err := hasil.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Karyawan sukses!")
	fmt.Println("Id terinsert:", insertId)
}

func TestExecInsertPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	lstkaryawan := []Karyawan{}
	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Herman",
		jabatan:      "Accounting",
		married:      true,
		tanggallahir: sql.NullTime{Time: time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{String: "Henny", Valid: true},
		gaji:         sql.NullInt64{Int64: 15, Valid: true},
	})

	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Nanda",
		jabatan:      "Project Manager",
		married:      false,
		tanggallahir: sql.NullTime{Time: time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{},
		gaji:         sql.NullInt64{Int64: 15, Valid: true},
	})

	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Nanda",
		jabatan:      "Project Manager",
		married:      false,
		tanggallahir: sql.NullTime{},
		pasangan:     sql.NullString{},
		gaji:         sql.NullInt64{},
	})

	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Dimas",
		jabatan:      "Frontend",
		married:      false,
		tanggallahir: sql.NullTime{Time: time.Date(2000, 3, 14, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{},
		gaji:         sql.NullInt64{},
	})

	/*
		nama := "Herman"
		married := true
		pasangan := "Dewi"
		tanggallahir := time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local)
		jabatan := "Accounting"
		gaji := 20
	*/
	sqlcmd := "insert into Karyawan(nama, married, pasangan, jabatan, tanggallahir, gaji) VALUES(?,?,?,?,?,?)"
	stmt, err := db.PrepareContext(ctx, sqlcmd)

	if err != nil {
		panic(err)
	}

	for _, kar := range lstkaryawan {
		resinsert, err := stmt.ExecContext(ctx, kar.nama, kar.married, kar.pasangan, kar.jabatan, kar.tanggallahir, kar.gaji)
		if err != nil {
			panic(err)
		}

		lastid, err := resinsert.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Terinsert data ", kar.nama, "dengan id=", lastid)
	}

	defer stmt.Close()
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	lstkaryawan := []Karyawan{}
	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Anggi",
		jabatan:      "Business Analyst",
		married:      false,
		tanggallahir: sql.NullTime{Time: time.Date(1995, 12, 7, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{},
		gaji:         sql.NullInt64{Int64: 24, Valid: true},
	})

	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Septian",
		jabatan:      "QA",
		married:      false,
		tanggallahir: sql.NullTime{Time: time.Date(1994, 7, 20, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{},
		gaji:         sql.NullInt64{Int64: 19, Valid: true},
	})

	lstkaryawan = append(lstkaryawan, Karyawan{
		nama:         "Stella Winata",
		jabatan:      "PMO",
		married:      false,
		tanggallahir: sql.NullTime{Time: time.Date(1990, 5, 23, 0, 0, 0, 0, time.Local), Valid: true},
		pasangan:     sql.NullString{String: "Nindya", Valid: true},
		gaji:         sql.NullInt64{Int64: 17, Valid: true},
	})

	sqlcmd := "insert into Karyawan(nama, married, pasangan, jabatan, tanggallahir, gaji) VALUES(?,?,?,?,?,?)"
	stmt, err := tx.PrepareContext(ctx, sqlcmd)

	for _, kar := range lstkaryawan {
		resinsert, err := stmt.ExecContext(ctx, kar.nama, kar.married, kar.pasangan, kar.jabatan, kar.tanggallahir, kar.gaji)
		if err != nil {
			tx.Rollback()
			panic(err)
		}

		lastid, err := resinsert.LastInsertId()

		if err != nil {
			tx.Rollback()
			panic(err)
		}

		fmt.Println("Terinsert data ", kar.nama, "dengan id=", lastid)
	}

	defer stmt.Close()

	errcommit := tx.Rollback() // tx.Commit()

	if errcommit != nil {
		panic(errcommit)
	}

	fmt.Println("Sudah terinsert semuanya")
}
