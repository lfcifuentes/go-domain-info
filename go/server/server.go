package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
	"testTruora/search"
)

func StartServer() {
	// iniciar el enrutador
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// habilitar las conexiones desde cualquier sitio
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
	})
	r.Use(cors.Handler)

	// ruta para consultar informacion de un dominio
	r.Post("/search",SearchHandler)
	// ruta para consultar la lista de dominios que se han buscado
	r.Post("/servers",ServersHandler)
	http.ListenAndServe(":3333", r)
}
// funcion para validar los parametros de la peticion e ir a buscar la informacion
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	data := &search.SearchRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	searchurl := data.Search
	string := search.SearchInformation(*searchurl)

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(string))
}
// funcion parra mostrar la lista de servidores consultados
func ServersHandler(w http.ResponseWriter, r *http.Request){
	string := search.GetData()
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(string))
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}
