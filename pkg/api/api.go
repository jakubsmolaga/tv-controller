package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/smolagakuba/tv-controller/pkg/tv"
)

type Api struct {
	tv tv.TV
	r  chi.Router
}

func Init(tv tv.TV) Api {
	r := chi.NewRouter()
	r.Post("/turnoff", tv.TurnOff)
	r.Post("/turnon", tv.TurnOn)
	r.Post("/reboot", tv.Reboot)
	r.Post("/select-hdmi1", tv.SelectHDMI1)
	r.Post("/select-displayport", tv.SelectDisplayPortPC)
	return Api{tv, r}
}

func (api Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.r.ServeHTTP(w, r)
}
