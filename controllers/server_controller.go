package controllers

import (
	"github.com/dsukesato/go13/pbl/app1-backend/controllers/handler"
	"github.com/dsukesato/go13/pbl/app1-backend/infrastructure/persistence"
	"github.com/dsukesato/go13/pbl/app1-backend/usecase"

	"github.com/gorilla/mux"

	"log"
	"net/http"
	"os"
)

func Serve() {
	//依存関係注入
	//投稿
	postsPersistence := persistence.NewPostsPersistence()
	postsUsecase := usecase.NewPostsUsecase(postsPersistence)
	postsHandler := handler.NewPostsHandler(postsUsecase)
	//お店
	restaurantsPersistence := persistence.NewRestaurantsPersistence()
	restaurantsUsecase := usecase.NewRestaurantsUsecase(restaurantsPersistence)
	restaurantsHandler := handler.NewRestaurantsHandler(restaurantsUsecase)

	//mux := http.NewServeMux() //マルチプレクサを用いてhandlerを管理
	r := mux.NewRouter()

	hindex := handler.IndexHandler{}
	hr := restaurantsHandler
	hp := postsHandler

	r.HandleFunc("/Lookin/", hindex.LookinHandler).Methods("GET")
	r.HandleFunc("/Lookin/restaurants/", hr.GetRestaurantsHandler).Methods("GET")
	r.HandleFunc("/Lookin/restaurants/", hr.PostRestaurantsHandler).Methods("POST")
	r.HandleFunc("/Lookin/posts/", hp.GetPostsHandler).Methods("GET")
	r.HandleFunc("/Lookin/posts/", hp.PostPostsHandler).Methods("POST")

	/*
	mux.Handle("/Lookin/", handler.MethodHandler{"GET": http.HandlerFunc(hindex.LookinHandler)})
	mux.Handle("/Lookin/restaurants/", handler.MethodHandler{"GET": http.HandlerFunc(hr.GetRestaurantsHandler)})
	mux.Handle("/Lookin/posts/", handler.MethodHandler{"GET": http.HandlerFunc(hp.GetPostsHandler)})
	mux.Handle("/Lookin/posts/", handler.MethodHandler{"POST": http.HandlerFunc(hp.PostPostsHandler)})
	*/

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

	//http.ListenAndServe(":8080", mux)
}

func PostsDI () {

}
