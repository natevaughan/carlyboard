package main

import (
	"log"
	"encoding/json"
	"net/http"
	"fmt"
)


func main() {
	log.Println("starting up...")
	log.Println("loading file appconfig.yml")

	var conf Config
	
	err := conf.Load("appconfig.yml")
	
	if err != nil {
		log.Fatal("Could not load appconfig.yml. Please make sure it is in the same directory as this executable: %s", err.Error())
	}

	dao := BoardDAO{conf.Dbconn}

	err = http.ListenAndServe(conf.Port, carlyHandler{dao})
	log.Fatal(err)
}

type carlyHandler struct {
	Dao BoardDAO
}

func (c carlyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

        board, err := c.Dao.GetBoard("123abc")
        
        if err != nil {
             log.Printf("Error getting data from mysql: %s", err.Error())
        }

	log.Printf("Successfully retrieved board %s, %s, %s", board.Uuid, board.Name, board.Description)
	
	j, err := json.Marshal(board)

        if err != nil {
             log.Printf("Error marshalling json: %s", err.Error())
        }

	fmt.Fprintf(w, string(j))
}
