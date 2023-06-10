package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/Siddheshk02/Securelee/model"
)

func Register(email string, pass []byte, name string) error {

	// password, err := bcrypt.GenerateFromPassword(pass, 14) //GenerateFromPassword returns the bcrypt hash of the password at the given cost i.e. (14 in our case).
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	user := model.User{
		UserID:   "unique()",
		Email:    fmt.Sprintf("%s", email),
		Password: fmt.Sprintf("%s", pass),
		Name:     fmt.Sprintf("%s", name),
	}

	payload, err := json.Marshal(user)
	// fmt.Println(bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to marshal user data:", err)
		return nil
	}

	os.Setenv("GODEBUG", "http2client=0")

	req, err := http.NewRequest("POST", "https://cloud.appwrite.io/v1/account", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to create request:", err)
		return nil
	}

	req.Header = http.Header{
		"Content-Type":               {"application/json"},
		"X-Appwrite-Response-Format": {"1.0.0"},
		"X-Appwrite-Project":         {"64664ea036f955c9205f"},
		"Accept-Encoding":            {"gzip,deflate,br"},
	}

	client1 := &http.Client{}
	resp, err := client1.Do(req)

	if err != nil {
		log.Fatalln("Failed to make request:", err)
		return nil
	}

	defer resp.Body.Close()
	// fmt.Println(resp)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalln("Failed to create user. Status code:", resp.StatusCode)
		return nil
	}

	// fmt.Println("User created successfully")
	println(resp.Proto == "HTTP/1.1")

	Login(email, string(pass))

	// JWT()

	return nil
}

func JWT() error {

	os.Setenv("GODEBUG", "http2client=0")

	req, err := http.NewRequest("POST", "https://cloud.appwrite.io/v1/account/jwt", nil)
	if err != nil {
		log.Fatalln("Failed to create JWT:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Appwrite-Response-Format", "1.0.0")
	req.Header.Set("X-Appwrite-Project", "64664ea036f955c9205f")

	client2 := &http.Client{}
	resp, err := client2.Do(req)

	if err != nil {
		log.Fatalln("Failed to make request:", err)
		return nil
	}

	defer resp.Body.Close()
	// fmt.Println(resp)

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Failed to create JWT. Status code:", resp.StatusCode)
		return nil
	}

	fmt.Println("JWT created successfully")

	println(resp.Proto == "HTTP/1.1")

	return nil

}

func Login(email string, password string) error {

	users := model.User{
		Email:    fmt.Sprintf("%s", email),
		Password: fmt.Sprintf("%s", password),
	}

	payload, err := json.Marshal(users)
	// fmt.Println(bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to marshal user data:", err)
		return nil
	}

	os.Setenv("GODEBUG", "http2client=0")

	req, err := http.NewRequest("POST", "https://cloud.appwrite.io/v1/account/sessions/email", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln("Failed to create request:", err)
		return nil
	}

	req.Header = http.Header{
		"Content-Type":               {"application/json"},
		"X-Appwrite-Response-Format": {"1.0.0"},
		"X-Appwrite-Project":         {"64664ea036f955c9205f"},
		"Accept-Encoding":            {"gzip,deflate,br"},
	}

	client1 := &http.Client{}
	resp, err := client1.Do(req)

	if err != nil {
		log.Fatalln("Failed to make request:", err)
		return nil
	}

	defer resp.Body.Close()
	// fmt.Println(resp)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Fatalln("Failed to create user. Status code:", resp.StatusCode)
		return nil
	}

	cookie := resp.Header.Get("Set-Cookie")

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	path := currentUser.HomeDir + "/securelee"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}
	// Construct the path to the token file in the user's home directory
	Session := currentUser.HomeDir + "/securelee/user_session.json"
	userfile := currentUser.HomeDir + "/securelee/user.txt"

	file, err := os.Create(Session)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file1, err := os.Create(userfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	// encode token into json format
	err = json.NewEncoder(file).Encode(cookie)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file1.WriteString(email)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("User created successfully")

	return nil

}
