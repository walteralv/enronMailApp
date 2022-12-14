package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"testing"
)

const (
	PATH_ENRON_USERS_TEST = "./enron_mail_test"
	INDEX_TEST            = "emails_test"
)

func createEnronDirTest(num_users int) {
	_, err := os.Stat(PATH_ENRON_USERS_TEST)
	if os.IsNotExist(err) {
		log.Println("Creating Test Directory enron_mail_test/ with", num_users, "users.")
		if err := os.MkdirAll(PATH_ENRON_USERS_TEST, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		users, _ := os.ReadDir(DEFAULT_ENRON_USERS_PATH)
		for _, user := range users[:num_users] {
			cmd := exec.Command("cp", "-r", DEFAULT_ENRON_USERS_PATH+"/"+user.Name(), PATH_ENRON_USERS_TEST)
			cmd.Run()
		}
	}
}




func TestIndexer(t *testing.T) {
	// createEnronDirTest(2)
	listEmailFilePaths, err := getListEmailFilePaths(PATH_ENRON_USERS_TEST)
	if err != nil {
		log.Fatal(err)
	}
	numEmailsFound := len(listEmailFilePaths)
	log.Println(numEmailsFound, "email files were found")
	var emails = make([]Email, numEmailsFound)

	for _, EmailFilePath := range listEmailFilePaths {
		log.Println("Indexing file ", EmailFilePath)
		email, err := ProcessEmailFile(EmailFilePath)
		if err != nil {
			log.Fatal(err)
		}
		emails = append(emails, *email)
		// emailJson, _ := json.Marshal(email)
		// err = ingestToZinc(string(emailJson), INDEX_TEST)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

	bulkv2Request := Bulkv2Request{
		Index:   INDEX,
		Records: emails,
	}
	bulkv2RequestJson, _ := json.Marshal(bulkv2Request)
	log.Println("Ingeting emails found to ZincSearch!")
	bulkv2IngestToZinc(string(bulkv2RequestJson))

}
