package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	c := carlyHandler{dao}
	http.HandleFunc("/helloworld", helloWorldHandler)
	http.HandleFunc("/board", c.HandleBoardRequests)
	http.HandleFunc("/section", c.HandleSectionRequests)
	http.HandleFunc("/stickie", c.HandleStickieRequests)
	http.HandleFunc("/board/", c.HandleBoardRequests)
	http.HandleFunc("/section/", c.HandleSectionRequests)
	http.HandleFunc("/stickie/", c.HandleStickieRequests)

	err = http.ListenAndServe(conf.HttpPort, nil)
	log.Fatal(err)
}

type carlyHandler struct {
	Dao BoardDAO
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")

}

func (c carlyHandler) HandleBoardRequests(w http.ResponseWriter, r *http.Request) {

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

		respond(w, board)
		break
	case "GET":
		params := strings.Split(r.RequestURI, "/")
		if len(params) != 3 {
			log.Println("cannot find board: " + r.RequestURI)
			break
		}

		i, err := strconv.ParseInt(params[2], 10, 64)

		if err != nil {
			log.Printf(err.Error())
		}

		board, err := c.Dao.GetBoard(i)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		log.Printf("Successfully retrieved board %s, %s, %s", board.Id, board.Name, board.Description)

		respond(w, board)

		break
	default:
		fmt.Fprint(w, "Method not supported: "+r.Method)
	}
}

func (c carlyHandler) HandleStickieRequests(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		println(r.Method)

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		var stickie Stickie

		err = json.Unmarshal(b, &stickie)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		stickie, err = c.Dao.CreateSticke(stickie)

		if err != nil {
			log.Printf("Error marshalling json: %s", err.Error())
		}

		respond(w, stickie)
		break
	case "GET":
		params := strings.Split(r.RequestURI, "/")
		if len(params) != 3 {
			log.Println("cannot find stickie: " + r.RequestURI)
			break
		}

		i, err := strconv.ParseInt(params[2], 10, 64)

		if err != nil {
			log.Printf(err.Error())
		}

		stickie, err := c.Dao.GetStickie(i)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		log.Printf("Successfully retrieved stickie %s, %s, %s", stickie.Id, stickie.Content, stickie.SectionId)

		respond(w, stickie)

		break
	default:
		fmt.Fprint(w, "Method not supported: "+r.Method)
	}
}

func (c carlyHandler) HandleSectionRequests(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		println(r.Method)

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		var section Section

		err = json.Unmarshal(b, &section)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		section, err = c.Dao.CreateSection(section)

		if err != nil {
			log.Printf("Error marshalling json: %s", err.Error())
		}

		respond(w, section)
		break
	case "GET":
		params := strings.Split(r.RequestURI, "/")
		if len(params) != 3 {
			log.Println("cannot find section: " + r.RequestURI)
			break
		}

		i, err := strconv.ParseInt(params[2], 10, 64)

		if err != nil {
			log.Printf(err.Error())
		}

		section, err := c.Dao.GetSection(i)

		if err != nil {
			log.Printf("Error getting data from mysql: %s", err.Error())
		}

		respond(w, section)

		break
	default:
		fmt.Fprint(w, "Method not supported: "+r.Method)
	}
}

func respond(w http.ResponseWriter, i interface{}) {

	j, err := json.Marshal(i)

	if err != nil {
		log.Printf("Error marshalling json: %s", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(j))
}