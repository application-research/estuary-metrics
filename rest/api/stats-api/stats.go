package statsapi

import (
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigStatsRouter(router gin.IRoutes) {
	router.GET("/stats/retrieval", api.ConvertHttpRouterToGin(GetRetrievalStats))
	router.GET("/stats/storage", api.ConvertHttpRouterToGin(GetStorageStats))
	router.GET("/stats/system", api.ConvertHttpRouterToGin(GetSystemStats))
	router.GET("/stats/users", api.ConvertHttpRouterToGin(GetUserStats))
	router.GET("/stats/info", api.ConvertHttpRouterToGin(GetUserStats))
}

type PublicStats struct {
	DealsOnChain       int   `json:"dealsOnChain"`
	TotalBytesUploaded int64 `json:"totalBytesUploaded"`
	TotalFilesStored   int   `json:"totalFilesStored"`
	TotalObjectsRef    int   `json:"totalObjectsRef"`
	TotalStorage       int64 `json:"totalStorage"`
	TotalStorageMiner  int   `json:"totalStorageMiner"`
	TotalUsers         int   `json:"totalUsers"`
}

type RetrievalStats struct {
	DealSuccessRate                    string `json:"dealSuccessRate"`
	DealAcceptanceRate                 string `json:"dealAcceptanceRate"`
	TotalRetrievalDealsProposed        string `json:"totalRetrievalDealsProposed"`
	TotalRetrievalDealProposalAccepted string `json:"totalRetrievalDealProposalAccepted"`
	TotalRetrievalDealProposalRejected string `json:"totalRetrievalDealProposalRejected"`
	TotalNumberOfSuccessfulRetrieval   string `json:"totalNumberOfSuccessfulRetrieval"`
	TotalNumberOfFailedRetrieval       string `json:"totalNumberOfFailedRetrieval"`
	FailedDealsDueToUndialableMiners   string `json:"failedDealsDueToUndialableMiners"`
	TimeToFirstByte                    string `json:"timeToFirstByte"`
	//Total number of retrieval deals attempted (per day and per week breakdown)
}

type StorageStats struct {
	DealSuccessRate                    string `json:"dealSuccessRate"`
	DealFailureRate                    string `json:"dealFailureRate"`
	DealAcceptanceRate                 string `json:"dealAcceptanceRate"`
	TotalStorageDealsProposed          string `json:"totalStorageDealsProposed"`
	TotalStorageDealProposalAccepted   string `json:"totalStorageDealProposalAccepted"`
	TotalStorageDealProposalRejected   string `json:"totalStorageDealProposalRejected"`
	TotalNumberOfStorageDealsAttempted string `json:"totalNumberOfStorageDealsAttempted"`
	TotalNumberOfSuccessfulStorage     string `json:"totalNumberOfSuccessfulStorage"`
	TotalNumberOfFailedStorage         string `json:"totalNumberOfFailedStorage"`

	// Distribution of data size uploaded per user
	// Performance metrics
	//Time to a successful deal
	//how does that scale with data size?
}

type SystemStats struct {
	TotalObjecsPinned         int `json:"totalObjecsPinned"`
	TotalSizeUploaded         int `json:"totalSizeUploaded"`
	totalSizeSealedOnFilecoin int `json:"totalSizeSealedOnFilecoin"`
	AvailableFreeSpace        int `json:"availableFreeSpace"`
	TotalSpaceCapacity        int `json:"totalSpaceCapacity"`

	//Downtime
	//Performance
}

type UserStats struct {
	TotalNumberOfStorageProviders int `json:"totalNumberOfStorageProviders"`
	TotalNumberOfUsers            int `json:"totalNumberOfUsers"`

	//Ongoing user activity — DAUs, WAUs, MAUs etc. Are users coming back? (custom Grafana plugin)
	//For Storage/Retrieval deal metrics, in addition to aggregate, we also want the following breakdowns
	//per day breakdown.
	//per week breakdown.
	//per provider breakdown.

}

func GetRetrievalStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetStorageStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetSystemStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetUserStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetPublicStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
