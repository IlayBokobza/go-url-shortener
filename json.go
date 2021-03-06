package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"regexp"
	"time"
)

type jsonData []struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

type jsonDataS struct {
	URL string `json:"url"`
	ID  string `json:"id"`
}

//delete a url
func deleteURL(url string) {
	regX := regexp.MustCompile(`(https?:\/\/)?www.`)
	url = regX.ReplaceAllString(url, "")

	data := getJSONData()
	toDeleteIndex := -1

	for i, v := range data {
		if v.URL == url {
			toDeleteIndex = i
			break
		}
	}

	if toDeleteIndex == -1 {
		return
	}

	data = append(data[:toDeleteIndex], data[toDeleteIndex+1:]...)

	newJSON, _ := json.Marshal(&data)

	ioutil.WriteFile("data.json", newJSON, 0666)
}

//add a url
func addURL(url string) ([]byte, string) {
	//parses url
	regX := regexp.MustCompile(`(https?:\/\/)?www.`)
	url = regX.ReplaceAllString(url, "")

	data := getJSONData()

	//check if url is in file
	for _, storedURL := range data {
		if storedURL.URL == url {
			return []byte{}, "Url already stored"
		}
	}

	newID := generateID()

	urlObj := jsonDataS{
		URL: url,
		ID:  newID,
	}

	newData := append(data, urlObj)

	newJSON, _ := json.Marshal(&newData)

	ioutil.WriteFile("data.json", newJSON, 0666)

	byteData, _ := json.Marshal(&urlObj)

	return byteData, ""
}

//get and parse json file
func getJSONData() jsonData {
	file, _ := ioutil.ReadFile("data.json")
	data := jsonData{}

	json.Unmarshal([]byte(file), &data)

	return data
}

//generate id for url
func generateID() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var output []byte

	index := 0
	for index < 10 {
		output = append(output, byte(r.Intn(122-97)+97))
		index++
	}

	return string(output)
}
