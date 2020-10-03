// /1) I dont know how get dynamic data so created ownself(need your help for this).
// Done and checked manually
// compared data if they are same, show message:file already exists(in ternminal)
// check validation.(required field),if not error is shown in terminal.
// if everyrhing is okay..it create file in json and save it.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	First_name string `json:"First_name" validate:"required"`
	Last_name  string `json:"Lirst_name"  validate:"required"`
	Email      string `json:"Email" validate:"required,email"`
}

func main() {
	router := httprouter.New()
	router.POST("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := User{
			First_name: "Mark",
			Last_name:  "jghad",
			Email:      "mark@gmail.com",
		}

		validate := validator.New()
		err := validate.Struct(data)
		if err != nil {
			fmt.Println(err.Error())
		}
		file, _ := json.MarshalIndent(data, "", " ")

		_ = ioutil.WriteFile("data.json", file, 0644)

		data2 := User{
			First_name: "Mark",
			Last_name:  "Jones",
			Email:      "mrk@gmail.com",
		}
		if data == data2 {
			fmt.Println("file already exists")
		} else {
			file, _ := json.MarshalIndent(data, "", " ")

			_ = ioutil.WriteFile("new.json", file, 0644)
		}
	})

	http.ListenAndServe(":8080", router)

}
