package handler

import (
	"encoding/json"
	"github.com/dsukesato/go13/pbl/app1-backend/usecase"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RestaurantsHandler interface {
	GetRestaurantsHandler(http.ResponseWriter, *http.Request)
	PostRestaurantsHandler(http.ResponseWriter, *http.Request)
}

type restaurantsHandler struct {
	restaurantsUsecase usecase.RestaurantsUsecase
}

func NewRestaurantsHandler(ru usecase.RestaurantsUsecase) RestaurantsHandler {
	return &restaurantsHandler{
		restaurantsUsecase: ru,
	}
}

func (rh restaurantsHandler) GetRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/restaurants/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()
	rests, err := rh.restaurantsUsecase.GetRestaurants(ctx)

	if err != nil {
		log.Fatal(err)
	}

	res := new(GetRestaurantsResponse)

	for _, rest := range rests {
		var rf RestaurantsField
		rf = RestaurantsField(*rest)
		res.Restaurants = append(res.Restaurants, rf)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type GetRestaurantsResponse struct {
	Restaurants []RestaurantsField `json:"restaurants"`
}

type RestaurantsField struct {
	RestaurantId int       `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	BusinessHours string   `json:"business_hours"`
	RestaurantImage string `json:"restaurant_image"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt time.Time    `json:"deleted_at"`
}

func (rh restaurantsHandler) PostRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/restaurants/" {
		http.NotFound(w, r)
		return
	}
	//ctx := r.Context()

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody = new(PostRestaurantRequest)
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := new(PostRestaurantResponse)
	res.RestaurantId = 7
	res.RestaurantName = jsonBody.RestaurantName
	res.BusinessHours = jsonBody.BusinessHours
	res.RestaurantImage = jsonBody.RestaurantImage
	res.Message = "post method success!! and send RestaurantID"
	log.Println(res.RestaurantName, res.BusinessHours, res.RestaurantImage)

	Success(w, res)
}

type PostRestaurantRequest struct {
	RestaurantName string  `json:"restaurant_name"`
	BusinessHours string   `json:"business_hours"`
	RestaurantImage string `json:"restaurant_image"`
}

type PostRestaurantResponse struct {
	RestaurantId int       `json:"restaurant_id"`
	RestaurantName string  `json:"restaurant_name"`
	BusinessHours string   `json:"business_hours"`
	RestaurantImage string `json:"restaurant_image"`
	Message string         `json:"massage"`
}