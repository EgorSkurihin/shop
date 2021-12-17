package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	dsn := "root:232323@tcp(localhost:3306)/shop"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() {
	s.db.Close()
}

/* func (s *Store) AddUser() error {
	u := &User{Name: "Admin", OpenPassword: "Admin"}
	stmt := `INSERT INTO users (name, password) VALUES(?, ?)`
	u.BeforeCreate()
	_, err := s.db.Exec(stmt, u.Name, u.HashedPassword)
	if err != nil {
		return err
	}
	return nil
} */

func (s *Store) GetUser(name string) (*User, error) {
	u := &User{}
	err := s.db.QueryRow("SELECT * FROM users WHERE name=?", name).Scan(
		&u.Id,
		&u.Name,
		&u.HashedPassword,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Store) GetAlbums() (map[int]*Album, error) {
	rows, err := s.db.Query("SELECT * FROM `albums`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albums := make(map[int]*Album)
	for rows.Next() {
		// Создаем указатель на новую структуру Snippet
		a := &Album{}
		var id int
		err = rows.Scan(&id, &a.Title, &a.Performer, &a.Cost, &a.Image)
		if err != nil {
			return nil, err
		}
		// Добавляем структуру в вывод.
		albums[id] = a
	}
	return albums, nil
}

func (s *Store) AddAlbum(a *Album) error {
	// Check Image exists
	if a.Image == "" || !isImageExists(a.Image) {
		a.Image = "store/albums_img/image-not-found.jpg"
	} else {
		a.Image = "store/albums_img/image/" + a.Image
	}

	stmt := `INSERT INTO albums (title, performer, cost, image) VALUES(?, ?, ?, ?)`

	_, err := s.db.Exec(stmt, a.Title, a.Performer, a.Cost, a.Image)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteAlbum(id string) error {
	_, err := s.db.Exec(`DELETE FROM albums WHERE id=?;`, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateAlbum(id string, newAlb *Album) error {
	//Get old album to compare with new
	oldAlb := &Album{}
	err := s.db.QueryRow("SELECT * FROM albums WHERE id=?", id).Scan(
		&id,
		&oldAlb.Title,
		&oldAlb.Performer,
		&oldAlb.Cost,
		&oldAlb.Image,
	)
	if err != nil {
		return err
	}
	// Compare ald album with new to find fileds? that need to update
	if newAlb.Title == "" {
		newAlb.Title = oldAlb.Title
	}
	if newAlb.Performer == "" {
		newAlb.Performer = oldAlb.Performer
	}
	if newAlb.Cost == 0 {
		newAlb.Cost = oldAlb.Cost
	}
	if newAlb.Image != "" && !isImageExists(newAlb.Image) {
		newAlb.Image = "store/albums_img/image-not-found.jpg"
	}
	if newAlb.Image == "" {
		newAlb.Image = oldAlb.Image
	}
	//Update albums in DB
	stmt := `UPDATE albums SET title = ?, performer = ?, cost = ?, image = ? where id = ?`
	_, err = s.db.Exec(stmt, newAlb.Title, newAlb.Performer, newAlb.Cost, newAlb.Image, id)
	if err != nil {
		return err
	}
	return nil
}

func isImageExists(string) bool {
	if _, err := os.Stat("../store/albums_img"); err != nil {
		return false
	}
	return true
}
