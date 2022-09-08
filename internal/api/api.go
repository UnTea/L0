package api

import (
	"encoding/json"
	"fmt"
	"github.com/UnTea/L0/internal/model"
	"github.com/UnTea/L0/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
)

type Service interface {
	Get(id string) (*model.Data, error)
	GetAllIDs() ([]string, error)
}

type orderPage struct {
	Data model.Data
	Json string
}

type startPage struct {
	Exist bool
	IDs   []string
}

type handler struct {
	s Service
}

func NewHandler(db *pgxpool.Pool, ch chan model.Data) *handler {
	return &handler{
		s: service.NewService(db, ch),
	}
}

func (h *handler) getElemByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	data, err := h.s.Get(id)
	if err != nil {
		log.Printf(fmt.Sprintf("error while getting element by id: %s due to err: %v", id, err))
		fmt.Fprint(w, "There is no such id")
		return
	}

	jsonData, err := json.MarshalIndent(data, " ", " ")
	if err != nil {
		log.Printf(fmt.Sprintf("error while marshalling element by id: %s due to err: %v", id, err))
		fmt.Fprint(w, "There is no such id")
		return
	}

	t, err := template.ParseFiles("web/static/order.html")
	if err != nil {
		log.Printf("template parsing error: %v", err)
		return
	}

	p := orderPage{
		Data: *data,
		Json: string(jsonData),
	}

	err = t.Execute(w, p)
	if err != nil {
		log.Printf("error while parsing html template due to err: %v", err)
		return
	}
}

func (h *handler) startPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := template.ParseFiles("web/static/index.html")
	if err != nil {
		log.Printf("template parsing error: %v", err)
		return
	}

	ids, err := h.s.GetAllIDs()
	if err != nil {
		log.Printf("error: %v", err)
	}

	exist := false
	if len(ids) > 0 {
		exist = true
	}

	pg := startPage{
		Exist: exist,
		IDs:   ids,
	}

	err = t.Execute(w, pg)
	if err != nil {
		log.Printf("error while parsing html template due to err: %v", err)
		return
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/orders/:id", h.getElemByID)
	router.GET("/", h.startPage)
}
