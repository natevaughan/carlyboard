DROP DATABASE IF EXISTS carlyboard;
CREATE DATABASE carlyboard; 
USE carlyboard; 

CREATE TABLE board (
    uuid VARCHAR(64) NOT NULL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description TEXT
);

CREATE TABLE section (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    board_uuid VARCHAR(64) NOT NULL, 
    title VARCHAR(128) NOT NULL,
    FOREIGN KEY fk_section_board(board_uuid)
    REFERENCES board(uuid)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE stickie (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    section_id BIGINT NOT NULL,
    content TEXT,
    FOREIGN KEY fk_stickie_section(section_id)
    REFERENCES section(id)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);
