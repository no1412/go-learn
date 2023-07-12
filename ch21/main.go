package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
网络编程：Go 语言如何玩转 RESTful API 服务？
*/
func main() {
	//http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users", handleUsers2)
	http.ListenAndServe(":8080", nil)
}
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ID:1,Name:张三")
		fmt.Fprintln(w, "ID:2,Name:李四")
		fmt.Fprintln(w, "ID:3,Name:王五")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "not found")
	}
}

// 数据源，类似MySQL中的数据
var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

// json返回数据
func handleUsers2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")
	}
}

// 用户
type User struct {
	ID   int
	Name string
}
