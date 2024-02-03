package main

import (
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
	// Meta struct {
	// 	Visits int
	// }
	Mate UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("./hello.gohtml")
	if err != nil {
		panic(err)
	}

	// 匿名 结构体 可能更加有效
	// user := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Tom",
	// 	Age:  20,
	// }
	user := User{
		Name: "Jon Tom",
		Age:  25,
		Mate: UserMeta{
			Visits: 111,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
