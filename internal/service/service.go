package service

import "github.com/tommytan/garen/internal/dao"

var Dao *dao.Dao

// New new a service and return.
func New() {
	Dao = dao.New()
}

// Close close the resource.
func Close() {
	Dao.Close()
}
