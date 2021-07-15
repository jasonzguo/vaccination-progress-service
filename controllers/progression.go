package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	service "github.com/jasonzguo/vaccination-progress-service/services"
	"github.com/julienschmidt/httprouter"
)

type progressionController struct {
}

var prograssionControllerInstance *progressionController = nil

func GetProgressionController() *progressionController {
	if prograssionControllerInstance == nil {
		prograssionControllerInstance = new(progressionController)
	}
	return prograssionControllerInstance
}

func (vc *progressionController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queries := r.URL.Query()

	payload, err := service.GetProgressionService().FindAll(r.Context(), queries.Get("lastId"))

	if err != nil {
		log.Fatal(fmt.Errorf("[Index] error in calling ProgressionController.GetAll  %v", err))
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(fmt.Errorf("[Index] error in calling json.Marshal  %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonPayload))
}
