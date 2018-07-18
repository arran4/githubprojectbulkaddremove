package main

import (
	"flag"
	"os"
	"log"
	"encoding/csv"
	"net/http"
	"strings"
	"io/ioutil"
	"io"
		)

var (
	incsv = flag.String("csv", "userchange.csv", "CSV of user changes")
	githubapikey = flag.String("ghtoken", "", "Github api token")
)

func main() {
	flag.Parse()
	f, err := os.Open(*incsv)
	if err != nil {
		log.Panic(err)
		return
	}
	defer f.Close()
	csvr := csv.NewReader(f)
	i := 0
	for {
		line, err := csvr.Read()
		i++
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicf("Error on line %d with err: %v", i, err)
			return
		}
		if len(line) < 3 {
			log.Printf("Line %s doesn't have the right columns\n", i)
			continue
		}
		switch line[0] {
		case "remove user from repo":
			project := line[1]
			user := line[2]
			u := "https://api.github.com/repos/"+project+"/collaborators/"+user
			req, err := http.NewRequest("DELETE", u, nil)
			if err != nil {
				log.Printf("Error line %d: Put failed to construct: %v\n", i, err)
				continue
			}
			req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")
			req.Header.Set("Authorization", "token " + *githubapikey)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Error line %d: Put failed to do request: %v\n", i, err)
				continue
			}
			if resp.StatusCode == http.StatusNoContent {
				log.Printf("Remove for %s to %s success\n", user, project)
			} else {
				b, _ := ioutil.ReadAll(resp.Body)
				log.Printf("Remove for %s to %s failed: %s\n", user, project, string(b))
			}
		case "add user to repo":
			project := line[1]
			user := line[2]
			u := "https://api.github.com/repos/"+project+"/collaborators/"+user
			req, err := http.NewRequest("PUT", u, strings.NewReader(`{"permission":"push"}`))
			if err != nil {
				log.Printf("Error line %d: Put failed to construct: %v\n", i, err)
				continue
			}
			req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")
			req.Header.Set("Authorization", "token " + *githubapikey)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("Error line %d: Put failed to do request: %v\n", i, err)
				continue
			}
			if resp.StatusCode == http.StatusCreated {
				log.Printf("Add for %s to %s success\n", user, project)
			} else {
				b, _ := ioutil.ReadAll(resp.Body)
				log.Printf("%d Add for %s to %s failed: %s\n", resp.StatusCode, user, project, string(b))
			}
		default:
			log.Printf("Line %d Unknown command %s\n", i, line[0])
		}
	}
}
