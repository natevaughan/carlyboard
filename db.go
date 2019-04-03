package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type BoardDAO struct {
	host string
}

func (d *BoardDAO) GetBoard(id int64) (Board, error) {

	log.Printf("Retrieving %s from databse", id)
	var board Board

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return board, err
	}

	result, err := db.Query("SELECT id, name, description FROM board WHERE id = ?", id)

	if err != nil {
		return board, err
	}

	result.Next()
	result.Scan(&board.Id, &board.Name, &board.Description)

	return board, nil
}

func (d *BoardDAO) CreateBoard(board Board) (Board, error) {

	log.Printf("Saving %s to the databse", board.Id)

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return board, err
	}

	stmt, err := db.Prepare("INSERT INTO board SET id = ?, name = ?, description = ?")
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	res, err := stmt.Exec(board.Id, board.Name, board.Description)
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	id, _ := res.LastInsertId()

	if err != nil {
		return board, err
	}
	board.Id = id

	return board, nil
}

func (d *BoardDAO) GetSection(id int64) (Section, error) {

	log.Printf("Retrieving %s from databse", id)
	var section Section

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return section, err
	}

	result, err := db.Query("SELECT id, title, board_id FROM section WHERE id = ?", id)

	if err != nil {
		return section, err
	}

	result.Next()
	result.Scan(&section.Id, &section.Title, &section.BoardId)

	return section, nil
}

func (d *BoardDAO) CreateSection(section Section) (Section, error) {

	log.Printf("Saving %s to the databse", section.Id)

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return section, err
	}

	stmt, err := db.Prepare("INSERT INTO section SET title = ?, board_id = ?")
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	res, err := stmt.Exec(section.Title, section.BoardId)
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return section, err
	}

	section.Id = id

	return section, nil
}


func (d *BoardDAO) GetStickie(id int64) (Stickie, error) {

	log.Printf("Retrieving %s from databse", id)
	var stickie Stickie

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return stickie, err
	}

	result, err := db.Query("SELECT id, content, section_id FROM stickie WHERE id = ?", id)

	if err != nil {
		return stickie, err
	}

	result.Next()
	result.Scan(&stickie.Id, &stickie.Content, &stickie.SectionId)

	return stickie, nil
}

func (d *BoardDAO) CreateSticke(stickie Stickie) (Stickie, error) {

	log.Printf("Saving %s to the databse", stickie.Id)

	db, err := sql.Open("mysql", d.host)
	defer db.Close()

	if err != nil {
		return stickie, err
	}

	stmt, err := db.Prepare("INSERT INTO stickie SET content = ?, section_id = ?")
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	res, err := stmt.Exec(stickie.Content, stickie.SectionId)
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return stickie, err
	}

	stickie.Id = id

	return stickie, nil
}

type Board struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Section struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	BoardId int64  `json:"boardId"`
}


type Stickie struct {
	Id      int64  `json:"id"`
	Content   string `json:"content"`
	SectionId int64  `json:"sectionId"`
}
