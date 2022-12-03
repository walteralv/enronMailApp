package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	PATH_ENRON_USERS        = "./enron_mail_20110402/maildir"
	EMAIL_FIELDS_SEPARATOR  = "\r\n"     // carriage return + line feed
	EMAIL_CONTENT_SEPARATOR = "\r\n\r\n" // double "carriage return + line feed"
	INDEX                   = "enron"
	USER                    = "admin"
	PASSWORD                = "Complexpass#123"
	URL_ZINC_API            = "http://localhost:4080/api/"
)

type Email struct {
	MessageID               string `json:"Message-ID"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	To                      string `json:"To"`
	Subject                 string `json:"Subject"`
	MimeVersion             string `json:"Mime-Version"`
	ContentType             string `json:"Content-Type"`
	ContentTransferEncoding string `json:"Content-Transfer-Encoding"`
	XFrom                   string `json:"X-From"`
	XTo                     string `json:"X-To"`
	Xcc                     string `json:"X-cc"`
	Xbcc                    string `json:"X-bcc"`
	XFolder                 string `json:"X-Folder"`
	XOrigin                 string `json:"X-Origin"`
	XFileName               string `json:"X-FileName"`
	Content                 string `json:"Content"`
	Filepath                string `json:"Filepath"`
}

func ProcessEmailFile(filepathString string) (*Email, error) {
	file, err := os.ReadFile(filepath.Clean(filepathString))
	if err != nil {
		log.Println("Error in ProcessEmailFile - reading file: ", err)
		return nil, err
	}

	arr := strings.SplitN(string(file), EMAIL_CONTENT_SEPARATOR, 2)
	if len(arr) != 2 {
		log.Println("Error in ProcessEmailFile - Wrong email file found at: ", filepathString)
		return nil, err
	}

	allDetails, content := arr[0], arr[1]

	detailsArr := strings.Split(allDetails, EMAIL_FIELDS_SEPARATOR)

	email := mapEmailDetails(detailsArr, content)
	email.Filepath = filepathString

	return email, nil
}

func mapEmailDetails(details []string, content string) *Email {
	email := &Email{}

	for i := 0; i < len(details); i++ {
		keyValue := strings.SplitN(details[i], ": ", 2)
		switch keyValue[0] {
		case "Message-ID":
			email.MessageID = keyValue[1]
		case "Date":
			email.Date = keyValue[1]
		case "From":
			email.From = keyValue[1]
		case "To":
			email.To = keyValue[1]
		case "Subject":
			email.Subject = keyValue[1]
		case "Mime-Version":
			email.MimeVersion = keyValue[1]
		case "Content-Type":
			email.ContentType = keyValue[1]
		case "Content-Transfer-Encoding":
			email.ContentTransferEncoding = keyValue[1]
		case "X-From":
			email.XFrom = keyValue[1]
		case "X-To":
			email.XTo = keyValue[1]
		case "X-cc":
			email.Xcc = keyValue[1]
		case "X-bcc":
			email.Xbcc = keyValue[1]
		case "X-Folder":
			email.XFolder = keyValue[1]
		case "X-Origin":
			email.XOrigin = keyValue[1]
		case "X-FileName":
			email.XFileName = keyValue[1]
		default:
			continue
		}
	}
	email.Content = content
	return email
}

func ingestToZinc(data, index string) {
	req, err := http.NewRequest("POST", URL_ZINC_API+index+"/_doc", strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(USER, PASSWORD)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

func main() {

	err := filepath.Walk(PATH_ENRON_USERS, func(filepathString string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		spaces := strings.Count(filepathString, "/")
		buff := strings.Repeat(" ", (spaces-2)*4)

		if info.IsDir() {
			fmt.Println(buff + info.Name())
		} else {
			email, err := ProcessEmailFile(filepathString)
			if err != nil {
				fmt.Println(err)
				return err
			}
			emailJson, _ := json.Marshal(email)
			ingestToZinc(string(emailJson), INDEX)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
