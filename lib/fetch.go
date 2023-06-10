package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"os/user"
	"strconv"

	"github.com/Siddheshk02/Securelee/model"
)

func Download(code int) string {

	os.Setenv("GODEBUG", "http2client=0")

	url := "https://cloud.appwrite.io/v1/storage/buckets/" + strconv.Itoa(code) + "/files/" + strconv.Itoa(code) + "/download"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatalln("Failed to get Data:", err)
	}
	req.Header.Add("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Add("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Add("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to make request:", err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Failed to read data:", err)
	}

	contentDisposition := res.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		fmt.Println("Failed to parse Content-Disposition:", err)
		return ""
	}

	filename := params["filename"]

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.WriteString(string(data))
	if err != nil {
		log.Fatalln(err)
	}
	Addusers(code)

	return "\nFile Downloaded!! - " + filename
}

func Addusers(code int) {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	fileName := currentUser.HomeDir + "/securelee/user.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	url := "https://cloud.appwrite.io/v1/databases/64849b4d30aa4d623fe6/collections/6484a5160346815cb7d9/documents"
	method := "POST"

	info := model.Download{
		DocID: fmt.Sprintf("%d", code),
		Data1: model.Data12{
			Code:         fmt.Sprintf("%d", code),
			Downloadedby: string(data),
		},
	}

	payload, err := json.Marshal(info)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Add("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Add("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
