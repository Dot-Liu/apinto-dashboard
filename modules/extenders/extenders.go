package extenders

import (
	"fmt"
	apinto_dashboard "github.com/eolinker/apinto-dashboard"
	"github.com/eolinker/apinto-dashboard/internal/apinto"
	"github.com/eolinker/apinto-dashboard/modules/professions"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Extenders struct {
	header *professions.ListHeader
	*httprouter.Router
	name string
}

func NewExtenders(name string) *Extenders {
	e := &Extenders{
		name: name,
		header: &professions.ListHeader{
			Title: map[apinto_dashboard.ZoneName][]string{
				apinto_dashboard.ZhCn: {"名称", "扩展ID", "项目", "分组", "版本"},
				apinto_dashboard.EnUs: {"Name", "ID", "Project", "Group", "Version"},
			},
			Fields: []string{"name", "id", "project", "group", "version"},
		},
	}
	e.createRouter()
	return e
}

func (e *Extenders) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.Router.ServeHTTP(w, r)
}

func (e *Extenders) Lookup(r *http.Request) (view string, data interface{}, has bool) {
	return e.name, e.header, true
}

func (e *Extenders) createRouter() {
	r := httprouter.New()
	r.GET(fmt.Sprintf("/api/%s/", e.name), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		data, code, err := apinto.Client().Extenders()
		if err != nil {
			apinto.WriteResult(w, 500, []byte(err.Error()))
			return
		}
		apinto.WriteResult(w, code, data)
	})
	r.GET(fmt.Sprintf("/api/%s/:id", e.name), func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		data, code, err := apinto.Client().Extender(id)
		if err != nil {
			apinto.WriteResult(w, 500, []byte(err.Error()))
			return
		}
		apinto.WriteResult(w, code, data)
	})

	e.Router = r
}
