package reportingapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/filecoin-project/lassie/pkg/eventpublisher"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/libp2p/go-libp2p-core/peer"
)

func ConfigureRetrievalsPushRouter(router gin.IRoutes) {
	router.POST("/retrievals/push", api.ConvertHttpRouterToGin(AddRetrievalEvent))
}

// AddRetrievalEvent publishes a retrieval event into the database
// @Summary Add Retrival Event
// @Description Get device usages
// @Tags Environment
// @Accept  json
// @Body multiEventReport body multiEventReport true "multiEventReport"
// @Success 200 {object}
// @Failure 500 {object} api.HTTPError
// @Router /reporting/retrievals/log [post]
func AddRetrievalEvent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var multiRetrEvent multiEventReport
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(resBody, &multiRetrEvent)

	// @jcace TODO
	// evts := multiRetrEvent.Events
	// storeInMetricsDB(evts)
	// storeInPando(evts)
}

type multiEventReport struct {
	Events []eventReport `json:"events"`
}

type eventReport struct {
	RetrievalId       uuid.UUID            `json:"retrievalId"`
	InstanceId        string               `json:"instanceId"`
	Cid               string               `json:"cid"`
	StorageProviderId peer.ID              `json:"storageProviderId"`
	Phase             eventpublisher.Phase `json:"phase"`
	PhaseStartTime    time.Time            `json:"phaseStartTime"`
	EventName         eventpublisher.Code  `json:"eventName"`
	EventTime         time.Time            `json:"eventTime"`
	EventDetails      interface{}          `json:"eventDetails,omitempty"`
}

// eventDetailsSuccess is for the EventDetails in the case of a retrieval
// success
type eventDetailsSuccess struct {
	ReceivedSize uint64 `json:"receivedSize"`
	ReceivedCids uint64 `json:"receivedCids"`
	Confirmed    bool   `json:"confirmed"`
}

// eventDetailsError is for the EventDetails in the case of a query or retrieval
// failure
type eventDetailsError struct {
	Error string `json:"error"`
}
