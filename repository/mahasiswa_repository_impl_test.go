package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMahasiswaInsert(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(belajar_golang_database.GetConnectionSQLServer())

	ctx := context.Background()
	karyawan := entity.Mahasiswa{
		Nama:     "Jaka",
		Jurusan:  "T.Arsitektur",
		TglLahir: time.Date(1987, 7, 20, 0, 0, 0, 0, time.Local),
	}
	hasil, err := mahasiswaRepository.Insert(ctx, karyawan)
	if err != nil {
		panic(err)
	}
	fmt.Println(hasil)
}

func TestFindByIdMahasiswa(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(belajar_golang_database.GetConnectionSQLServer())

	ctx := context.Background()
	hasil, err := mahasiswaRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(hasil)
}

func TestFindAllMahasiswa(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(belajar_golang_database.GetConnectionSQLServer())

	ctx := context.Background()
	hasil, err := mahasiswaRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, mhs := range hasil {
		fmt.Println(mhs)
	}
}
