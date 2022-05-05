package goalexa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aivahealth/goalexa/alexaapi"
	"go.uber.org/zap"
)

type Skill struct {
	applicationId string
	handlers      []Handler
}

func NewSkill(applicationId string) *Skill {
	return &Skill{
		applicationId: applicationId,
		handlers:      []Handler{},
	}
}

func (s *Skill) RegisterHandlers(handler ...Handler) {
	if s.handlers == nil {
		s.handlers = []Handler{}
	}
	s.handlers = append(s.handlers, handler...)
}

func (s *Skill) HandleRequest(r *alexaapi.RequestRoot) (*alexaapi.ResponseRoot, error) {
	for _, h := range s.handlers {
		if h.CanHandle(r) {
			return h.Handle(r)
		}
	}
	return nil, fmt.Errorf("No handler found for request (%q)", r.Request.Type)
}

func (s *Skill) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := validateAlexaRequest(w, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestJson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var root alexaapi.RequestRoot
	err = json.Unmarshal(requestJson, &root)
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if root.Context.System.Application.ApplicationId == "" || root.Context.System.Application.ApplicationId != s.applicationId {
		err := fmt.Errorf("Unable to verify applicationId")
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := s.HandleRequest(&root)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, bytes.NewReader(responseJson))
	if err != nil {
		Logger.Error("ServeHTTP failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
