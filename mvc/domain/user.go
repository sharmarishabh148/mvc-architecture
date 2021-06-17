package domain

type User struct {
	Id        uint64 `josn:"id"`
	FirstName string `josn:"first_name"`
	LastName  string `josn:"last_name"`
	Email     string `josn:"email"`
}
