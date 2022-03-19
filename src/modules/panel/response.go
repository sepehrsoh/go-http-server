package panel

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gorilla/csrf"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
)

const (
	PublicPath = "/server/panel/html/"
)

func response(writer http.ResponseWriter, r *http.Request, file string, PageResponse *PageResponse) {
	pwd, _ := os.Getwd()
	bareAddr := fmt.Sprintf("%s%s", pwd, PublicPath)
	path := fmt.Sprintf("%s%s", bareAddr, file)

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Errorln(err)
	}
	if PageResponse.Code == 0 {
		PageResponse.Code = http.StatusOK
	}
	writer.WriteHeader(PageResponse.Code)
	responses := structs.Map(PageResponse)

	responses[csrf.TemplateTag] = csrf.TemplateField(r)
	err = tmpl.Execute(writer, responses)
	if err != nil {
		log.Errorln(err)
	}
}

type PageResponse struct {
	Title      string
	Parameters interface{}
	Errors     interface{}
	Code       int
	Action     string
	Messages   []string
}
