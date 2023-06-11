package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
	"time"
)

func Check() (bool, string) {

	var cookieData string
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	sessionPath := currentUser.HomeDir + "/securelee/user_session.json"

	file, err := ioutil.ReadFile(sessionPath)
	if err != nil {
		return false, "No User logged in. Log in to Securelee"
	}
	err = json.Unmarshal(file, &cookieData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	cookieParts := strings.Split(cookieData, ";")
	var expires string
	for _, part := range cookieParts {
		if strings.Contains(part, "expires=") {
			expires = strings.TrimPrefix(part, "expires=")
			break
		}
	}

	expires = strings.TrimSpace(expires)
	expires = strings.TrimPrefix(expires, "expires=")

	expirationTime, err := time.Parse("Mon, 02-Jan-2006 15:04:05 MST", expires)
	if err != nil {
		log.Fatal(err)
	}

	isExpired := time.Now().After(expirationTime)

	if isExpired {
		return false, "Session is expired. Log in to Securelee."
	}

	fileName := currentUser.HomeDir + "/securelee/user.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return true, "Logged In as " + string(data)
}

func Logout() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error occured!, try again.")
	}

	sessionPath := currentUser.HomeDir + "/securelee/user_session.json"
	userPath := currentUser.HomeDir + "/securelee/user.txt"

	e := os.Remove(sessionPath)
	e = os.Remove(userPath)
	if e != nil {
		return "No User logged in. Log in to Securelee"
	}
	return "User Logged Out."
}
