package main

import (
	"database/sql"
	"os"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Album struct {
	ID int64
	Title string
	Artist string
	Price float32
}

func main() {
	dbConnectionString := "postgres://mugambi:mugambi@localhost:5432/recordings"
	var err error
	db, err = sql.Open("postgres", dbConnectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to DB: %v\n",err)
		os.Exit(1)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database connected")

	albId, err := addAlbum(Album{
		Title: "The modern sound of Betty Carter",
		Artist: "John Coltrane",
		Price: 49.99,
	})
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album: %v\n", albId)

	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("albums  found: %v \n", albums)

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close() // close handle just before the function returns

	for rows.Next() {
		var alb Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}

		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func albumByID (id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("select * from album where id = $1", id)

	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return alb, nil
}

func addAlbum (alb Album) (int64, error) {
	var id int64
	err := db.QueryRow("insert into album (title, artist, price) values ($1, $2, $3) returning id", alb.Title, alb.Artist, alb.Price).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}