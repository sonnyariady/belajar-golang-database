package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

type Mahasiswa struct {
	id            int
	Nama, Jurusan string
	TglLahir      time.Time
}

func TestTransactionSqlServer(t *testing.T) {
	db := GetConnectionSQLServer()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	sqlcmd := "insert into Mahasiswa(nama, jurusan, tgllahir) OUTPUT INSERTED.id VALUES(@nama, @jurusan, @tgllahir)"
	stmt, err := tx.PrepareContext(ctx, sqlcmd)

	listmhs := []Mahasiswa{}

	listmhs = append(listmhs, Mahasiswa{
		Nama:     "Andi",
		Jurusan:  "T.Kimia",
		TglLahir: time.Date(1985, 5, 20, 0, 0, 0, 0, time.Local),
	})

	listmhs = append(listmhs, Mahasiswa{
		Nama:     "Jenny",
		Jurusan:  "T.Industri",
		TglLahir: time.Date(1986, 3, 14, 0, 0, 0, 0, time.Local),
	})

	listmhs = append(listmhs, Mahasiswa{
		Nama:     "Andre",
		Jurusan:  "T.Informatika",
		TglLahir: time.Date(1984, 8, 27, 0, 0, 0, 0, time.Local),
	})

	for _, mhs := range listmhs {
		var lastid int
		err := stmt.QueryRow(sql.Named("nama", mhs.Nama), sql.Named("jurusan", mhs.Jurusan), sql.Named("tgllahir", mhs.TglLahir)).Scan(&lastid)
		if err != nil {
			tx.Rollback()
			panic(err)
		}

		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println("Terinsert data ", mhs.Nama, "dengan id=", lastid)
	}
	defer stmt.Close()
	errcommit := tx.Commit()

	if errcommit != nil {
		panic(errcommit)
	}

	fmt.Println("Sudah terinsert semuanya")
}

func TestQuerySQLServer(t *testing.T) {
	db := GetConnectionSQLServer()
	defer db.Close()
	ctx := context.Background()

	sqlquery := "select id,nama, jurusan, tgllahir from Mahasiswa"
	lstmhs := []Mahasiswa{}
	baris, err := db.QueryContext(ctx, sqlquery)
	if err != nil {
		panic(err)
	}
	//i := 0
	for baris.Next() {
		var nama, jurusan string
		var id int
		var tgllahir time.Time

		err := baris.Scan(&id, &nama, &jurusan, &tgllahir)
		if err != nil {
			panic(err)
		}
		lstmhs = append(lstmhs, Mahasiswa{
			id:       id,
			Nama:     nama,
			Jurusan:  jurusan,
			TglLahir: tgllahir,
		})
		/*
			fmt.Println("Data Baris ke-", i)
			fmt.Println("Id : ", id)
			fmt.Println("Nama : ", nama)
			fmt.Println("Jurusan : ", jurusan)
			fmt.Println("Tgl Lahir : ", tgllahir)
			i++
		*/
	}
	defer baris.Close()
	fmt.Println(lstmhs)
}

func TestQueryWithParamSQLServer(t *testing.T) {
	db := GetConnectionSQLServer()
	defer db.Close()
	ctx := context.Background()

	paramjurusan := "T.Informatika"
	sqlquery := "select id,nama, jurusan, tgllahir from Mahasiswa where jurusan = @paramjurusan"
	lstmhs := []Mahasiswa{}
	baris, err := db.QueryContext(ctx, sqlquery, sql.Named("paramjurusan", paramjurusan))
	if err != nil {
		panic(err)
	}
	//i := 0
	for baris.Next() {
		var nama, jurusan string
		var id int
		var tgllahir time.Time

		err := baris.Scan(&id, &nama, &jurusan, &tgllahir)
		if err != nil {
			panic(err)
		}
		lstmhs = append(lstmhs, Mahasiswa{
			id:       id,
			Nama:     nama,
			Jurusan:  jurusan,
			TglLahir: tgllahir,
		})
		/*
			fmt.Println("Data Baris ke-", i)
			fmt.Println("Id : ", id)
			fmt.Println("Nama : ", nama)
			fmt.Println("Jurusan : ", jurusan)
			fmt.Println("Tgl Lahir : ", tgllahir)
			i++
		*/
	}
	defer baris.Close()
	fmt.Println(lstmhs)
}
