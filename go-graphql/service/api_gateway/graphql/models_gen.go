// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type House struct {
	Address string `json:"address"`
	OwnerID string `json:"ownerId"`
	Age     int    `json:"age"`
	ID      string `json:"id"`
}

type NewHouse struct {
	Address string `json:"address"`
	OwnerID string `json:"ownerId"`
	Age     int    `json:"age"`
}

type NewUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	ID        string `json:"id"`
}
