// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// func Hnadle(w http.ResponseWriter, r *http.Request) {
// 	var body struct {
// 		ID int
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		fmt.Println("エラー1")
// 	}
// 	// b, err := GetBook(r.Context, body.ID)
// }

// func GetBook(ctx context.Context, id int) (*Book, error) {
// 	rows, err := db.QueryContext(ctx, "SELECT id, name, isdn, price FROM books WHERE id = ?;")
// }
