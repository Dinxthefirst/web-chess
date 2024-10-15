package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	s.routes()

	s.PathPrefix("/static/styles/").Handler(http.StripPrefix("/static/styles/", http.FileServer(http.Dir("./static/styles/"))))
	s.PathPrefix("/static/dist/").Handler(http.StripPrefix("/static/dist/", http.FileServer(http.Dir("./static/dist/"))))
	s.PathPrefix("/static/images/").Handler(http.StripPrefix("/static/images/", http.FileServer(http.Dir("./static/images/"))))

	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", s.appHandler())

	gameHandler := &GameHandler{}
	s.HandleFunc("/new-game", gameHandler.NewGame)
	s.HandleFunc("/move", gameHandler.Move)
	s.HandleFunc("/current-state", gameHandler.CurrentState)
	s.HandleFunc("/legal-moves/{index}", gameHandler.LegalMoves)

	s.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
}

func (s *Server) appHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "static/index.html")
	}
}
