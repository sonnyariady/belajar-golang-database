package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type KaryawanRepository interface {
	Insert(ctx context.Context, karyawan entity.KaryawanTable) (entity.KaryawanTable, error)
	FindById(ctx context.Context, id int64) (entity.KaryawanTable, error)
	FindAll(ctx context.Context) ([]entity.KaryawanTable, error)
}
