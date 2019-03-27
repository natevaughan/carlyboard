package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	log.Println("starting up...")
	log.Println("loading file appconfig.yml")

	var conf Config

	err := conf.Load("appconfig.yml")

	if err != nil {
		log.Fatal("Could not load appconfig.yml. Please make sure it is in the same directory as this executable: %s", err.Error())
	}

	dao := BoardDAO{conf.DbHost + "/" + conf.DbName}

	err = http.ListenAndServe(conf.HttpPort, carlyHandler{dao})
	log.Fatal(err)
}

type carlyHandler struct {
	Dao BoardDAO
}

func (c carlyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// todo: routing
	prefix := strings.Split(r.RequestURI, "?")[0]
	switch prefix {
	case "/board":
		println(r.RequestURI)
		switch r.Method {
		case "POST":
			println(r.Method)

			b, err := ioutil.ReadAll(r.Body)

			if err != nil {
				log.Printf("Error getting data from mysql: %s", err.Error())
			}

			var board Board

			err = json.Unmarshal(b, &board)

			if err != nil {
				log.Printf("Error getting data from mysql: %s", err.Error())
			}

			board, err = c.Dao.CreateBoard(board)

			if err != nil {
				log.Printf("Error marshalling json: %s", err.Error())
			}

			j, err := json.Marshal(board)

			if err != nil {
				log.Printf("Error marshalling json: %s", err.Error())
			}

			fmt.Fprint(w, string(j))
			break
		case "GET":
			params := r.URL.Query()["id"]
			if len(params) != 1 {
				log.Println("Must have only one ID")
				break
			}
			board, err := c.Dao.GetBoard(params[0])

			if err != nil {
				log.Printf("Error getting data from mysql: %s", err.Error())
			}

			log.Printf("Successfully retrieved board %s, %s, %s", board.Uuid, board.Name, board.Description)

			j, err := json.Marshal(board)

			if err != nil {
				log.Printf("Error marshalling json: %s", err.Error())
			}

			fmt.Fprint(w, string(j))

			break
		default:
			log.Println("could not map request" + r.RequestURI)
		}
	default:
		println("unknown: " + r.RequestURI)
	}

}
