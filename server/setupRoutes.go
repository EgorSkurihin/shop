package main

import "net/http"

func (srv *Server) setupRoutes() {
	// Handle ADMIN routes
	admin := srv.router.PathPrefix("/admin").Subrouter()
	admin.Use(srv.authMiddleware)
	admin.HandleFunc("", srv.getAdminPage).Methods("GET")
	admin.HandleFunc("", srv.postGood).Methods("POST")
	admin.HandleFunc("/{id:[0-9]+}", srv.putGood).Methods("PUT")
	admin.HandleFunc("/{id:[0-9]+}", srv.deleteGood).Methods("DELETE")

	// Serve mail sending
	srv.router.HandleFunc("/mail", srv.sendMail).Methods("POST")

	//Server auth
	srv.router.HandleFunc("/auth", srv.auth).Methods("POST")

	// Serve images store and DB-get-goods
	srv.router.HandleFunc("/getGoods", srv.getGoods)
	srv.router.PathPrefix("/store/").Handler(http.StripPrefix("/store/", http.FileServer(http.Dir("././store"))))

	// Serve static files
	srv.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("././build/static/"))))

	/*
		Serve pages routes
	*/
	//Login page
	srv.router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "././build/templates/login.html")
	})

	// Cart page
	srv.router.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "././build/templates/cart.html")
	})
	// About page
	srv.router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "././build/templates/about.html")
	})
	// Home page
	srv.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "././build/templates/index.html")
	})
}
