package main

type DatabaseProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]DatabaseProfile{
	"user1": {
		Email: "george_bush@email.com",
		Id:    "user1",
		Name:  "George Bush",
		Token: "123",
	},
	"user2": {
		Email: "bill_clienton@email.com",
		Id:    "user2",
		Name:  "Bill Clinton",
		Token: "456",
	},
}

// Response uchun struct
type ClientProfile struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	Name  string `json:"name"`
}
