package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Mail(email string) int {

	key := Getkey()

	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(99999) + 10000

	from := mail.NewEmail("Securelee", "noreply@securelee.tech")
	subject := "Welcome to Securelee! [Verification]"
	to := mail.NewEmail("", email)
	plainTextContent := "Verify your email address. Hello, You have selected this email address as your Securelee ID. To verify that it's you, enter the code below on the email verification terminal : " + strconv.Itoa(code) + " Best Regards, Securelee Team"
	htmlContent := "<strong><p align=center>Verify your email address.</p></strong> <p>Hello,</p> <p>You have selected this email address as your Securelee ID. To verify that it's you, enter the code below on the email verification terminal : </p> <p><strong>" + strconv.Itoa(code) + "</strong></p> <br> <p>Best Regards,</p> Securelee Team"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(key)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
	}

	return code
}

func Getkeycipher() (string, string) {

	url := "https://cloud.appwrite.io/v1/databases/64849b4d30aa4d623fe6/collections/6486330b2a0d1a1a284b/documents"
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

	type Document struct {
		API_Key string `json:"API_Key"`
		ID      string `json:"$id"`
	}

	var response struct {
		Total     int        `json:"total"`
		Documents []Document `json:"documents"`
	}

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		log.Panic(err)
	}

	for _, doc := range response.Documents {
		return doc.ID, doc.API_Key
	}

	return "", ""
}

func Getkey() string {

	key, ciphertext := Getkeycipher()

	encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	if len(encrypted) < aes.BlockSize {
		return ""
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	cfb.XORKeyStream(decrypted, encrypted)

	return string(decrypted)

}
