package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

type setVar struct{}

var reqBody []byte
var write setVar

func (s setVar) Write(bs []byte) (int, error) {
	reqBody = bs
	return len(bs), nil
}

func setRoutes() {
	//get all
	http.HandleFunc("/api/get", func(res http.ResponseWriter, req *http.Request) {
		//checks for get method
		if req.Method != "GET" {
			return
		}

		file, _ := ioutil.ReadFile("data.json")

		res.Header().Set("content-type", "application/json")

		res.Write(file)
	})

	//add new
	http.HandleFunc("/api/add", func(res http.ResponseWriter, req *http.Request) {
		//checks for post method
		if req.Method != "POST" {
			return
		}

		io.Copy(write, req.Body)

		data, msg := addURL(string(reqBody))

		if msg != "" {
			res.Write([]byte(msg))
			return
		}

		res.Header().Set("content-type", "appliction/json")
		res.Write(data)
	})

	//delete one
	http.HandleFunc("/api/delete", func(res http.ResponseWriter, req *http.Request) {
		//checks fort delete method
		if req.Method != "DELETE" {
			return
		}

		io.Copy(write, req.Body)

		deleteURL(string(reqBody))

		res.Write([]byte{})
	})
}
