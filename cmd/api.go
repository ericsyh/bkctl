package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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

type LedgerMeta struct {
	ID struct {
		LedgerID              int    `json:"ledgerId"`
		MetadataFormatVersion int    `json:"metadataFormatVersion"`
		EnsembleSize          int    `json:"ensembleSize"`
		WriteQuorumSize       int    `json:"writeQuorumSize"`
		AckQuorumSize         int    `json:"ackQuorumSize"`
		State                 string `json:"state"`
		Length                int    `json:"length"`
		LastEntryID           int    `json:"lastEntryId"`
		Ctime                 int64  `json:"ctime"`
		DigestType            string `json:"digestType"`
		Password              string `json:"password"`
		CustomMetadata        struct {
			Component           string `json:"component"`
			PulsarManagedLedger string `json:"pulsar/managed-ledger"`
			Application         string `json:"application"`
		} `json:"customMetadata"`
		Closed       bool `json:"closed"`
		AllEnsembles struct {
			EnsembleId []struct {
				ID string `json:"id"`
			} `json:"0"`
		} `json:"allEnsembles"`
		Ctoken int `json:"ctoken"`
	} `json:"2"`
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

func ledgermeta(url string, id int) []byte {
	resp, err := http.Get("http://" + url + "/api/v1/ledger/metadata/?ledger_id=" + strconv.Itoa(id))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Bookie endpoint doesn't own this Ledger Id")
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
