package app

import (
	"int-gateway/handler"
	"os"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type app struct {
	router *mux.Router
	logger *logrus.Logger
}

const (
	get     = "GET"
	post    = "POST"
	put     = "PUT"
	delete  = "DELETE"
	options = "OPTIONS"

	port = "2001"
)

// InitializeApp Initializes the router of the gateway
func InitializeApp() *app {
	a := app{}
	a.router = mux.NewRouter()
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.JSONFormatter{})

	a.logger = logger
	a.registerRoutes()
	return &a
}

// Run runs the http server
func (a *app) Run() error {
	// configure CORS
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// Serve http
	err := http.ListenAndServe(":"+port, handlers.CORS(headers, origins, methods)(a.router))
	if err != nil {
		a.logger.Error("CONNECTION ERROR")
		return errors.Wrap(err, "connection error")
	}
	return nil
}

func (a *app) registerRoutes() {
	h := handler.InitiateHandler(a.logger)
	a.registerJournalistRoutes(h)
	a.registerSeasonRoutes(h)
	a.registerCelebrityRoutes(h)
	a.registerGenreRoutes(h)
	a.registerArticleRoutes(h)
	a.registerShowRoutes(h)
	a.registerEpisodeRoutes(h)
}

func (a *app) registerShowRoutes(h *handler.Handler) {
	a.router.Path("/show").
		Methods(post).
		HandlerFunc(h.CreateShow)

	a.router.Path("/show").
		Methods(get).
		HandlerFunc(h.GetShow)

	a.router.Path("/show").
		Methods(put).
		HandlerFunc(h.UpdateShow)

	a.router.Path("/show/list").
		Methods(get).
		HandlerFunc(h.ListShows)

	a.router.Path("/movie/posters").
		Methods(post).
		HandlerFunc(h.UploadMoviePosters)

	a.router.Path("/movie/poster").
		Methods(delete).
		HandlerFunc(h.DeleteMoviePoster)

	a.router.Path("/series/posters").
		Methods(post).
		HandlerFunc(h.UploadSeriesPosters)

	a.router.Path("/series/poster").
		Methods(delete).
		HandlerFunc(h.DeleteSeriesPoster)
}

func (a *app) registerJournalistRoutes(h *handler.Handler) {
	a.router.Path("/journalist").
		Methods(post).
		HandlerFunc(h.CreateJournalist)

	a.router.Path("/journalist").
		Methods(get).
		HandlerFunc(h.GetJournalist)

	a.router.Path("/journalist/list").
		Methods(get).
		HandlerFunc(h.ListJournalists)

	a.router.Path("/journalist").
		Methods(put).
		HandlerFunc(h.UpdateJournalist)
}

func (a *app) registerCelebrityRoutes(h *handler.Handler) {
	a.router.Path("/celebrity").
		Methods(post).
		HandlerFunc(h.CreateCelebrity)

	a.router.Path("/celebrity").
		Methods(get).
		HandlerFunc(h.GetCelebrity)

	a.router.Path("/celebrity").
		Methods(put).
		HandlerFunc(h.UpdateCelebrity)

	a.router.Path("/celebrity/posters").
		Methods(post).
		HandlerFunc(h.UploadCelebrityPosters)

	a.router.Path("/celebrity/poster").
		Methods(delete).
		HandlerFunc(h.DeleteCelebrityPoster)

	a.router.Path("/celebrity/list").
		Methods(get).
		HandlerFunc(h.ListCelebrities)
}

func (a *app) registerSeasonRoutes(h *handler.Handler) {
	a.router.Path("/season").
		Methods(post).
		HandlerFunc(h.CreateSeason)

	a.router.Path("/season").
		Methods(get).
		HandlerFunc(h.GetSeason)

	a.router.Path("/season/posters").
		Methods(post).
		HandlerFunc(h.UploadSeasonPosters)

	a.router.Path("/season/poster").
		Methods(delete).
		HandlerFunc(h.DeleteSeasonPoster)

	a.router.Path("/season/list").
		Methods(get).
		HandlerFunc(h.ListSeasons)

	a.router.Path("/season").
		Methods(put).
		HandlerFunc(h.UpdateSeason)
}
func (a *app) registerGenreRoutes(h *handler.Handler) {
	a.router.Path("/genre").
		Methods(post).
		HandlerFunc(h.CreateGenre)

	a.router.Path("/genre").
		Methods(get).
		HandlerFunc(h.GetGenre)

	a.router.Path("/genre").
		Methods(put).
		HandlerFunc(h.UpdateGenre)

	a.router.Path("/genre/list").
		Methods(get).
		HandlerFunc(h.ListGenres)
}
 
func (a *app) registerArticleRoutes(h *handler.Handler) {
	a.router.Path("/article").
		Methods(post).
		HandlerFunc(h.CreateArticle)

	a.router.Path("/article").
		Methods(get).
		HandlerFunc(h.GetArticle)

	a.router.Path("/article").
		Methods(put).
		HandlerFunc(h.UpdateArticle)

	a.router.Path("/article/list").
		Methods(get).
		HandlerFunc(h.ListArticles)

	a.router.Path("/article/posters").
		Methods(post).
		HandlerFunc(h.UploadArticlePosters)

	a.router.Path("/article/poster").
		Methods(delete).
		HandlerFunc(h.DeleteArticlePoster)
}

func (a *app) registerEpisodeRoutes(h *handler.Handler) {
	a.router.Path("/episode").
		Methods(post).
		HandlerFunc(h.CreateEpisode)

	a.router.Path("/episode").
		Methods(get).
		HandlerFunc(h.GetEpisode)

	a.router.Path("/episode").
		Methods(put).
		HandlerFunc(h.UpdateEpisode)

	a.router.Path("/episode/posters").
		Methods(post).
		HandlerFunc(h.UploadEpisodePosters)

	a.router.Path("/episode/poster").
		Methods(delete).
		HandlerFunc(h.DeleteEpisodePoster)

	a.router.Path("/episode/list").
		Methods(get).
		HandlerFunc(h.ListEpisodes)
}
