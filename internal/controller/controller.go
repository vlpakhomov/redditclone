package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/handler"
	"gitlab.com/vk-golang/lectures/06_databases/99_hw/redditclone/internal/controller/middleware"
)

type controller struct {
	Handler *handler.Handler
	Router  *mux.Router
}

func NewController(handl *handler.Handler, router *mux.Router) *controller {
	return &controller{
		Handler: handl,
		Router:  router,
	}
}

func (c *controller) RegisterRoutes() {
	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("../../static/")),
	)
	c.Router.PathPrefix("/static/").Handler(staticHandler).Methods("GET")

	c.Router.HandleFunc("/api/register", c.Handler.SignUp).Methods("POST")
	c.Router.HandleFunc("/api/login", c.Handler.Login).Methods("POST")

	c.Router.HandleFunc("/api/posts/", c.Handler.GetPosts).Methods("GET")
	c.Router.HandleFunc("/api/posts", c.Handler.AddPost).Methods("POST")
	c.Router.HandleFunc("/api/posts/{categoryName}", c.Handler.GetPostsWithCategory).Methods("GET")
	c.Router.HandleFunc("/api/post/{postID}", c.Handler.GetPost).Methods("GET")
	c.Router.HandleFunc("/api/post/{postID}/upvote", c.Handler.Upvote).Methods("GET")
	c.Router.HandleFunc("/api/post/{postID}/downvote", c.Handler.Downvote).Methods("GET")
	c.Router.HandleFunc("/api/post/{postID}/unvote", c.Handler.Unvote).Methods("GET")
	c.Router.HandleFunc("/api/post/{postID}", c.Handler.DeletePost).Methods("DELETE")
	c.Router.HandleFunc("/api/user/{username}", c.Handler.GetPostsWithUser).Methods("GET")

	c.Router.HandleFunc("/api/post/{postID}", c.Handler.AddComment).Methods("POST")
	c.Router.HandleFunc("/api/post/{postID}/{commentID}", c.Handler.DeleteComment).Methods("DELETE")

	c.Router.PathPrefix("/").HandlerFunc(c.Handler.Index).Methods("GET")
}

func (c *controller) UseMiddleware() http.Handler {
	mux := middleware.Auth(c.Handler.Auth, c.Router)
	mux = middleware.AccessLog(mux)
	mux = middleware.Panic(mux)
	return mux
}
