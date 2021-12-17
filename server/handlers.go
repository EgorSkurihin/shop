package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func (srv *Server) getGoods(w http.ResponseWriter, r *http.Request) {
	albums, err := srv.store.GetAlbums()
	if err != nil {
		w.WriteHeader(400)
		return
	}
	res, err := json.Marshal(albums)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(res)
}

func (srv *Server) auth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("login")
	password := r.Form.Get("password")
	dbUser, err := srv.store.GetUser(name)
	if err != nil {
		w.Write([]byte("<h1>Error!</h1>"))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(password))
	if err != nil {
		w.Write([]byte("<h1>Wrong password or login!</h1>"))
		return
	}
	session, _ := srv.sessionStore.Get(r, "session")
	session.Values["userId"] = dbUser.Id
	session.Save(r, w)
	w.Write([]byte("<h1>You are logged in!</h1> <a href='/admin'>Amdin page</a>"))
}

func (srv *Server) sendMail(w http.ResponseWriter, r *http.Request) {
	// Get form values
	r.ParseForm()
	mailData := &MailData{}
	err := decoder.Decode(mailData, r.PostForm)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	//Parsing cart from string to object
	err = json.Unmarshal([]byte(mailData.CartString), &mailData.CartObj)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	//Reading goods.json
	goods, _ := srv.store.GetAlbums()

	//Forming the message to client
	message := "<!DOCTYPE html><html><body>"
	message += "<h1>Уведомление о покупке на сайте</h1>"
	message += "<p>Здравствуйте " + mailData.UName + ", вы сделали заказ на сайте vinylshop</p>"
	message += "<p>Для уточнения ваших данных мы позвоним вам по телефону: " + mailData.UTel + "</p>"
	message += "<p>Товары, которые вы заказали: </p>"
	totalCost := 0
	for key, count := range mailData.CartObj {
		id, _ := strconv.Atoi(key)
		message += fmt.Sprintf("%s (%s) - %d шт.<br>", goods[id].Title, goods[id].Performer, count)
		totalCost += count * goods[id].Cost
	}
	message += fmt.Sprintf("<strong>Общая стоимость: %d$</strong>", totalCost)

	// Gmail configs
	from := "vinyllshopp@gmail.com"
	password := "VH123123"
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, password, host)
	to := []string{"egorskurihin@gmail.com"}

	//Set headers for email message
	message = "MIME-Version: 1.0 \r\n" +
		"Content-type: text/html; charset=UTF-8 \r\n" +
		"From: " + from + "\r\n" +
		"To:" + mailData.UEmail + "\r\n" +
		"Subject: discount vinyl!\r\n" +
		"\r\n" + message + "</body></html>"

	err = smtp.SendMail(host+":587", auth, from, to, []byte(message))
	if err != nil {
		w.WriteHeader(400)
		return
	}
}

/*
	ADMIN functions
*/
func (srv *Server) getAdminPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "././build/templates/admin.html")
}

func (srv *Server) postGood(w http.ResponseWriter, r *http.Request) {
	alb := &Album{}
	r.ParseForm()
	err := decoder.Decode(alb, r.PostForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = srv.store.AddAlbum(alb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("New album has posted!"))
}

func (srv *Server) putGood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	alb := &Album{}
	r.ParseForm()
	err := decoder.Decode(alb, r.PostForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err = srv.store.UpdateAlbum(id, alb); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("The album with id " + id + " edited!"))
}

func (srv *Server) deleteGood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := srv.store.DeleteAlbum(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("The album with id " + id + " deleted!"))
}
