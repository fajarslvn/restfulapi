package repository

import (
	"context"
	"database/sql"

	"github.com/fajarslvn/restfulapi/model/domain"
)

// Buat kontrak dgn interface.
// context berguna untuk mempermudah mengatur sebuah fungsi.
type CategoryRepository interface {
	// Save Category: Fungsi ini untuk membuat category baru. Dibutuhkan context, query database dan Name dari category. Untuk Id akan di generate otomatis. Fungsi ini mengembalikan Category struct (Id dan Name). */
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category

	/* Update Category: Fungsi ini untuk meng-update detail dari category. Dibutuhkan context, query database, category Id dan Name. Dimana category Name harus sama dgn Id. Fungsi ini mengembalikan Category struct. */
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category

	/* Delete Category: Fungsi ini untuk menghapus category, yg dihapus adalah category Id otomatis semua yg berhubungan dgnnya akan terhapus. Fungsi ini tdk perlu mengembalikan apa-apa. Dibutuhkan context, query database dan category Id. */
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)

	/* Find Category by Id: Fungsi ini sesuai namanya akan mencari category berdasarkan Id. Dibutuhkan context, query database, dan nomor Id dari category. Fungsi ini akan mengembalikan info dari Id category yg dicari, dan akan mengembalikan error message jika category Id tidak ditemukan. */
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)

	/* Find Category by All: Fungsi ini akan mencari category berdasarkan Id dan Name. Dibutuhkan hanya context dan query database. Fungsi ini akan mengembalikan info apa saja dari category yg dicari. */
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
