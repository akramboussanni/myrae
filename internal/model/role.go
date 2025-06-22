package model

type Role struct {
	ID   int64  // primary key
	Name string // unique role name, e.g. "admin", "uploader"
}
