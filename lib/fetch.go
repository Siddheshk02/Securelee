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
	"github.com/olekukonko/tablewriter"
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

	return "File Downloaded!! - " + filename
}

func Addusers(code int) {

	name, file := GetShares(code)

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
		DocID: "unique()",
		Data1: model.Data12{
			Code:         fmt.Sprintf("%d", code),
			Uploadedby:   name,
			Filename:     file,
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

func GetShares(code int) (string, string) {

	url := "https://cloud.appwrite.io/v1/databases/64849b4d30aa4d623fe6/collections/64849b5a260c48dbf759/documents"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Add("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Add("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")

	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	var response struct {
		Total     int            `json:"total"`
		Documents []model.Data11 `json:"documents"`
	}

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		log.Panic(err)
	}

	for _, doc := range response.Documents {
		if doc.Code == strconv.Itoa(code) {
			return doc.Uploadedby, doc.Filename
		}
	}
	return "", ""
}

func Getdownloads() {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	userPath := currentUser.HomeDir + "/securelee/user.txt"

	data, err := ioutil.ReadFile(userPath)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	url := "https://cloud.appwrite.io/v1/databases/64849b4d30aa4d623fe6/collections/6484a5160346815cb7d9/documents"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Add("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Add("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")

	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	var response struct {
		Total     int            `json:"total"`
		Documents []model.Data12 `json:"documents"`
	}

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		log.Panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Downloaded By", "Filename", "Code"})

	for _, doc := range response.Documents {
		if doc.Uploadedby == string(data) {
			table.Append([]string{doc.Downloadedby, doc.Filename, doc.Code})
		}
	}

	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold, tablewriter.BgBlackColor})

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor})

	table.Render()

}
