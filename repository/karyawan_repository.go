package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type KaryawanRepository interface {
	Insert(ctx context.Context, karyawan entity.Karyawan) (entity.Karyawan, error)
	FindById(ctx context.Context, id int32) (entity.Karyawan, error)
	FindAll(ctx context.Context) ([]entity.Karyawan, error)
}
