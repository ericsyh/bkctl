package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var bklist map[string]string

var lglist map[string]string

type BookieInfo struct {
	FreeSpace  int64 `json:"freeSpace"`
	TotalSpace int64 `json:"totalSpace"`
}

type DiskFile struct {
	IndexFiles    string `json:"index files"`
	JournalFiles  string `json:"journal files"`
	EntrylogFiles string `json:"entrylog files"`
}

func bookieinfo(url string) []byte {
	resp, err := http.Get("http://" + url + "/api/v1/bookie/info")
	if err != nil {
		fmt.Println("Error:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func diskfile(url string) []byte {
	resp, err := http.Get("http://" + url + "/api/v1/bookie/list_disk_file")
	if err != nil {
		fmt.Println("Error:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func bookielist(url string) []byte {
	resp, err := http.Get("http://" + url + "/api/v1/bookie/list_bookies")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func ledgerlist(url string) []byte {
	resp, err := http.Get("http://" + url + "/api/v1/ledger/list")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
