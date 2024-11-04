package models

type User struct {
	ID        int64
	Firstname string
	Lastname  string
	Username  string
	Password  string
	PhoneID   int64
}

type Product struct {
	ID          int64
	Title       string
	Description string
	Price       int64
	Seller      int64
}

type Phone struct {
	ID         int64
	Number     string
	IsValidate bool
}
