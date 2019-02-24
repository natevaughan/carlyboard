package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type BoardDAO struct {
	host string;
}

func (d *BoardDAO) GetBoard(uuid string) (Board, error) {
	log.Printf("Retrieving %s from databse", uuid)
	var board Board

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return board, err
	}

	result, err := db.Query("SELECT uuid, name, description FROM board WHERE uuid = ?", uuid)

	if err != nil {
		return board, err
	}

	result.Next()
	result.Scan(&board.Uuid, &board.Name, &board.Description)

	return board, nil
}

type Board struct {
	Uuid        string  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
