package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Siddheshk02/Securelee/model"
)

type Bucket struct {
	ID   string `json:"$id"`
	Name string `json:"name"`
}

type BucketsResponse struct {
	Total   int      `json:"total"`
	Buckets []Bucket `json:"buckets"`
}

func Upload(path string) string {

	os.Setenv("GODEBUG", "http2client=0")

	req, err := http.NewRequest("GET", "https://cloud.appwrite.io/v1/storage/buckets", nil)
	if err != nil {
		log.Fatalln("Failed to get Data:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Set("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Set("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")

	client2 := &http.Client{}
	resp, err := client2.Do(req)

	if err != nil {
		log.Fatalln("Failed to make request:", err)
	}

	defer resp.Body.Close()
	// fmt.Println(resp)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalln("Failed to create JWT. Status code:", resp.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Failed to read response body:", err)
	}

LOOP:
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(9000) + 1000
	// code := 9876
	res := IsValid(code, responseBody)

	if res == false {
		goto LOOP
	} else {
		CreateBucket(code)
		CreateFile(code, path)
		Addfiles(code, filepath.Base(path))
	}

	// const expirationTime = 20 * time.Minute

	return "Generated code " + strconv.Itoa(code)
}

func IsValid(targetInteger int, responseData []byte) bool {
	var bucketsResponse BucketsResponse
	err := json.Unmarshal([]byte(responseData), &bucketsResponse)
	if err != nil {
		log.Fatal("Failed to parse response data:", err)
	}

	for _, bucket := range bucketsResponse.Buckets {
		if targetInteger == parseIntFromBucketName(bucket.Name) {
			return false
		}
	}
	return true
}

func parseIntFromBucketName(name string) int {
	var targetInteger int
	_, err := fmt.Sscanf(name, "%d", &targetInteger)
	if err != nil {
		log.Println("Failed to parse integer from bucket name:", err)
	}
	return targetInteger
}

func CreateBucket(code int) {

	bucket := model.Bucket{
		BucketID: fmt.Sprintf("%d", code),
		Name:     fmt.Sprintf("%d", code),
	}

	payload, err := json.Marshal(bucket)
	if err != nil {
		log.Fatalln("Failed to marshal user data:", err)
	}

	os.Setenv("GODEBUG", "http2client=0")

	req, err := http.NewRequest("POST", "https://cloud.appwrite.io/v1/storage/buckets", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to get Data:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Set("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Set("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")

	client2 := &http.Client{}
	resp, err := client2.Do(req)

	if err != nil {
		log.Fatalln("Failed to make request:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalln("Failed to create Bucket. Status code:", resp.StatusCode)
	}

}

func CreateFile(code int, path string) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	FileID := fmt.Sprintf("%d", code)
	_ = writer.WriteField("fileId", FileID)

	file, errFile2 := os.Open(path)
	defer file.Close()
	part2, errFile2 := writer.CreateFormFile("file", filepath.Base(path))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return
	}

	err := writer.Close()
	if err != nil {
		log.Fatal("Failed to close writer:", err)
	}

	os.Setenv("GODEBUG", "http2client=0")

	url := "https://cloud.appwrite.io/v1/storage/buckets/" + strconv.Itoa(code) + "/files"

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalln("Failed to get Data:", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Set("X-Appwrite-Project", "64664ea036f955c9205f")
	req.Header.Set("X-Appwrite-Key", "00f7044ff0153cee6ec49c891db57cd41c2b16e80f1293291cfe145f9533ef83f6043a18fbe4729c04c77a78a693678c6fc37152e7a0f397fa946cde3c338c6a4db7b8857b9f01aebd44950ac0aabd456d2bebedcdf9ab3d7cd6eb20d77fd55ddf54b5977a7d96526c793cf76d0f0099950b077acb84edf94ee613263e6c4aac")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))

	client2 := &http.Client{}
	resp, err := client2.Do(req)

	if err != nil {
		log.Fatalln("Failed to make request:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalln("Failed to create File. Status code:", resp.StatusCode)
	}

}

func Addfiles(code int, filename string) {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	fileName := currentUser.HomeDir + "/securelee/user.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	url := "https://cloud.appwrite.io/v1/databases/64849b4d30aa4d623fe6/collections/64849b5a260c48dbf759/documents"
	method := "POST"

	info := model.Share{
		DocID: fmt.Sprintf("%d", code),
		Data1: model.Data11{
			Code:       fmt.Sprintf("%d", code),
			Uploadedby: string(data),
			Filename:   filename,
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
