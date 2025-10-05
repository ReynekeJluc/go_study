-- SQLite

DROP TABLE IF EXISTS Book;
DROP TABLE IF EXISTS Author;
DROP TABLE IF EXISTS Genre;
DROP TABLE IF EXISTS BookAuthor;
DROP TABLE IF EXISTS BookGenre;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Tokens;


PRAGMA foreign_keys = ON;


CREATE TABLE Book (
	book_id INTEGER PRIMARY KEY AUTOINCREMENT,
	book_name TEXT NOT NULL,
	book_total_quantity INTEGER NOT NULL,
	book_pages_amount INTEGER NOT NULL,
	book_desc TEXT,
	book_price REAL NOT NULL,
	book_cover INTEGER NOT NULL,          -- 0/1
	book_super_cover INTEGER NOT NULL,    -- 0/1
	book_publisher TEXT,
	book_year_release INTEGER NOT NULL,
	book_isbn TEXT NOT NULL
);


CREATE TABLE Author (
	author_id INTEGER PRIMARY KEY AUTOINCREMENT,
	author_name TEXT NOT NULL,
	author_desc TEXT,
	author_birth INTEGER,
	author_country TEXT
);


CREATE TABLE Genre (
	genre_id INTEGER PRIMARY KEY AUTOINCREMENT,
	genre_name TEXT NOT NULL,
	genre_desc TEXT
);


CREATE TABLE BookAuthor (
	book_id INTEGER NOT NULL,
	author_id INTEGER NOT NULL,

	FOREIGN KEY (book_id) REFERENCES Book(book_id) 
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	FOREIGN KEY (author_id) REFERENCES Author(author_id) 
		ON DELETE CASCADE
		ON UPDATE CASCADE
);


CREATE TABLE BookGenre (
	book_id INTEGER NOT NULL,
	genre_id INTEGER NOT NULL,

	FOREIGN KEY (book_id) REFERENCES Book(book_id) 
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	FOREIGN KEY (genre_id) REFERENCES Genre(genre_id) 
		ON DELETE CASCADE
		ON UPDATE CASCADE
);


CREATE TABLE Users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE Tokens (
    token_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token_hash TEXT NOT NULL,
    revoked BOOLEAN DEFAULT 0,
    expires_at DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE
);
