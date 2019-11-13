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

type PostsHandler interface {
	GetPostsHandler(http.ResponseWriter, *http.Request)
	PostPostsHandler(http.ResponseWriter, *http.Request)
}

type postsHandler struct {
	postsUsecase usecase.PostsUsecase
}

func NewPostsHandler(pu usecase.PostsUsecase) PostsHandler {
	return &postsHandler{
		postsUsecase: pu,
		//postsUsecase: usecase.NewPostsUsecase(),
	}
}

func (ph postsHandler) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/posts/" {
		http.NotFound(w, r)
		return
	}
	ctx := r.Context()
	posts, err := ph.postsUsecase.GetPosts(ctx)

	if err != nil {
		log.Fatal(err)
	}

	res := new(GetPostsResponse)

	for _, post := range posts {
		var bf PostsField
		bf = PostsField(*post)
		res.Posts = append(res.Posts, bf)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

type GetPostsResponse struct {
	Posts []PostsField `json:"posts"`
}

type PostsField struct {
	PostId int           `json:"post_id"`
	PostUserId int       `json:"post_user_id"`
	PostRestaurantId int `json:"post_restaurant_id"`
	PostImage string     `json:"post_image"`
	Good int             `json:"good"`
	Text string          `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

func (ph postsHandler) PostPostsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Lookin/posts/" {
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
	//var jsonBody map[string]interface{}
	var jsonBody = new(PostPostsRequest)
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := new(PostPostsResponse)
	res.PostId = 11
	res.PostUserId = jsonBody.PostUserId
	res.PostRestaurantId = jsonBody.PostRestaurantId
	res.PostImage = jsonBody.PostImage
	res.Text = jsonBody.Text
	res.Message = "post method success!! and send PostID"
	log.Println(res.PostUserId, res.PostRestaurantId, res.PostImage, res.Text)

	/*posts, err := ph.postsUsecase.PostPosts(ctx)

	if err != nil {
		log.Fatal(err)
	}

	res := new(GetPostsResponse)

	for _, post := range posts {
		var bf PostsField
		bf = PostsField(*post)
		res.Posts = append(res.Posts, bf)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}*/
	Success(w, res)
}

type PostPostsRequest struct {
	PostUserId int       `json:"post_user_id"`
	PostRestaurantId int `json:"post_restaurant_id"`
	PostImage string     `json:"post_image"`
	Text string          `json:"text"`
}

type PostPostsResponse struct {
	PostId int           `json:"post_id"`
	PostUserId int       `json:"post_user_id"`
	PostRestaurantId int `json:"post_restaurant_id"`
	PostImage string     `json:"post_image"`
	Text string          `json:"text"`
	Message string       `json:"message"`
}
