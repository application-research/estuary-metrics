package statsapi

import (
	"database/sql"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sort"
	"time"
)

func ConfigStatsRouter(router gin.IRoutes) {
	router.GET("/stats/retrieval", api.ConvertHttpRouterToGin(GetRetrievalStats))
	router.GET("/stats/storage", api.ConvertHttpRouterToGin(GetStorageStats))
	router.GET("/stats/system", api.ConvertHttpRouterToGin(GetSystemStats))
	router.GET("/stats/users", api.ConvertHttpRouterToGin(GetUserStats))
	router.GET("/stats/info", api.ConvertHttpRouterToGin(GetPublicStats))
	router.GET("/stats/deal-metrics", api.ConvertHttpRouterToGin(GetDealMetrics))
}

type DealMetricsInfo struct {
	Time              time.Time `json:"time"`
	DealsOnChain      int       `json:"dealsOnChain"`
	DealsOnChainBytes int64     `json:"dealsOnChainBytes"`
	DealsAttempted    int       `json:"dealsAttempted"`
	DealsSealed       int       `json:"dealsSealed"`
	DealsSealedBytes  int64     `json:"dealsSealedBytes"`
	DealsFailed       int       `json:"dealsFailed"`
}

type MetricsDealJoin struct {
	CreatedAt        time.Time `json:"created_at"`
	Failed           bool      `json:"failed"`
	FailedAt         time.Time `json:"failed_at"`
	DealID           int64     `json:"deal_id"`
	Size             int64     `json:"size"`
	TransferStarted  time.Time `json:"transferStarted"`
	TransferFinished time.Time `json:"transferFinished"`
	OnChainAt        time.Time `json:"onChainAt"`
	SealedAt         time.Time `json:"sealedAt"`
}

type PublicStats struct {
	TotalStorage       sql.NullInt64 `json:"totalStorage"`
	TotalFilesStored   sql.NullInt64 `json:"totalFiles"`
	DealsOnChain       sql.NullInt64 `json:"dealsOnChain"`
	TotalObjectsRef    sql.NullInt64 `json:"totalObjectsRef"`
	TotalBytesUploaded sql.NullInt64 `json:"totalBytesUploaded"`
	TotalUsers         sql.NullInt64 `json:"totalUsers"`
	TotalStorageMiner  sql.NullInt64 `json:"totalStorageMiners"`
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
	ctx := api.InitializeContext(r)
	var stats PublicStats
	if err := dao.DB.Model(model.Content{}).Where("active and not aggregated_in > 0").Select("SUM(size) as total_storage").Scan(&stats).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := dao.DB.Model(model.Content{}).Where("active and not aggregate").Count(&stats.TotalFilesStored.Int64).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := dao.DB.Model(model.ContentDeal{}).Where("not failed and deal_id > 0").Count(&stats.DealsOnChain.Int64).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	//	this can be resource expensive but we are already caching it.
	if err := dao.DB.Table("obj_refs").Count(&stats.TotalObjectsRef.Int64).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := dao.DB.Table("objects").Select("SUM(size)").Find(&stats.TotalBytesUploaded.Int64).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := dao.DB.Model(model.User{}).Count(&stats.TotalUsers.Int64).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := dao.DB.Table("storage_miners").Count(&stats.TotalStorageMiner.Int64).Error; err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, stats)

}

func GetDealMetrics(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	metricsInfo, err := computeDealMetrics()
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, metricsInfo)
}

func computeDealMetrics() ([]*DealMetricsInfo, error) {

	var deals []*MetricsDealJoin
	if err := dao.DB.Model(model.ContentDeal{}).
		Joins("left join contents on content_deals.content = contents.id").
		Select("content_deals.failed as failed, failed_at, deal_id, size, transfer_started, transfer_finished, on_chain_at, sealed_at").
		Scan(&deals).Error; err != nil {
		return nil, err
	}

	coll := make(map[time.Time]*DealMetricsInfo)
	onchainbuckets := make(map[time.Time][]*MetricsDealJoin)
	attempts := make(map[time.Time][]*MetricsDealJoin)
	sealed := make(map[time.Time][]*MetricsDealJoin)
	beginning := time.Now().Add(time.Hour * -100000)
	failed := make(map[time.Time][]*MetricsDealJoin)

	for _, d := range deals {
		created := d.CreatedAt.Round(time.Hour * 24)
		attempts[created] = append(attempts[created], d)

		if !(d.DealID == 0 || d.Failed) {
			if d.OnChainAt.Before(beginning) {
				d.OnChainAt = time.Time{}
			}

			btime := d.OnChainAt.Round(time.Hour * 24)
			onchainbuckets[btime] = append(onchainbuckets[btime], d)
		}

		if d.SealedAt.After(beginning) {
			sbuck := d.SealedAt.Round(time.Hour * 24)
			sealed[sbuck] = append(sealed[sbuck], d)
		}

		if d.Failed {
			fbuck := d.FailedAt.Round(time.Hour * 24)
			failed[fbuck] = append(failed[fbuck], d)
		}
	}

	for bt, deals := range onchainbuckets {
		dmi := &DealMetricsInfo{
			Time:         bt,
			DealsOnChain: len(deals),
		}
		for _, d := range deals {
			dmi.DealsOnChainBytes += d.Size
		}

		coll[bt] = dmi
	}

	for bt, deals := range attempts {
		dmi, ok := coll[bt]
		if !ok {
			dmi = &DealMetricsInfo{
				Time: bt,
			}
			coll[bt] = dmi
		}

		dmi.DealsAttempted = len(deals)
	}

	for bt, deals := range sealed {
		dmi, ok := coll[bt]
		if !ok {
			dmi = &DealMetricsInfo{
				Time: bt,
			}
			coll[bt] = dmi
		}

		dmi.DealsSealed = len(deals)
		for _, d := range deals {
			dmi.DealsSealedBytes += d.Size
		}
	}

	for bt, deals := range failed {
		dmi, ok := coll[bt]
		if !ok {
			dmi = &DealMetricsInfo{
				Time: bt,
			}
			coll[bt] = dmi
		}

		dmi.DealsFailed = len(deals)
	}

	var out []*DealMetricsInfo
	for _, dmi := range coll {
		out = append(out, dmi)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Time.Before(out[j].Time)
	})

	return out, nil
}
