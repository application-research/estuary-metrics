package statsapi

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func ConfigStatsRouter(router gin.IRoutes) {
	//	 retrieval
	router.GET("/stats/retrieval-rates", api.ConvertHttpRouterToGin(GetRetrievalRateStats))
	router.GET("/stats/total-retrievals", api.ConvertHttpRouterToGin(GetTotalRetrievals))

	//	 storage
	router.GET("/stats/total-content-deals-attempted", api.ConvertHttpRouterToGin(GetContentDealsAttempted))
	router.GET("/stats/total-storage", api.ConvertHttpRouterToGin(GetTotalStorageInTib))
	router.GET("/stats/total-files", api.ConvertHttpRouterToGin(GetTotalFiles))
	router.GET("/stats/storage-rates", api.ConvertHttpRouterToGin(GetStorageRateStats))

	// deals
	router.GET("/stats/deal-metrics", api.ConvertHttpRouterToGin(GetDealMetrics))
	router.GET("/stats/info", api.ConvertHttpRouterToGin(GetInfo))

	// social media
	router.GET("/stats/to-twitter", api.ConvertHttpRouterToGin(GetStatsForTwitter))
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
	RetrievalRateStats
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

type RetrievalRateStats struct {
	DealSuccessRate string `json:"dealSuccessRate"`
	DealFailureRate string `json:"dealFailureRate"`
}
type StorageRateStats struct {
	DealSuccessRate string `json:"dealSuccessRate"`
	DealFailureRate string `json:"dealFailureRate"`
}
type StorageStats struct {
	StorageRateStats
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
	TotalObjectsPinned        int `json:"totalObjectsPinned"`
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

// GetTotalRetrievals returns the total number of retrieval deals attempted
// @Summary Get total number of retrieval deals attempted
// @Description Get total number of retrieval deals attempted
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} int64
// @Router /stats/total-retrievals [get]
func GetTotalRetrievals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	var totalRetrievals int64
	totalRetrievals, err := dao.Metrics.GetTotalRetrievals()
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, totalRetrievals)
}

// GetRetrievalRateStats returns the retrieval stats
// @Summary Returns the retrieval stats
// @Description Returns the retrieval stats
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} RetrievalStats
// @Router /stats/retrieval-rates [get]
func GetRetrievalRateStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	//select ((t.success  * 1.0 /t.total  * 1.0) * 100) as "Success Rate", ((t.failed * 1.0 / t.total * 1.0) * 100) as "Failure Rate" from (select (select (select count(*) from retrieval_success_records) + (select count(*) from retrieval_failure_records)) as total, (select count(*) from retrieval_success_records)               as success,  (select count(*) from retrieval_failure_records) as failed) as t;
	ctx := api.InitializeContext(r)
	var retrievalStats RetrievalStats
	successFailRate := dao.DB.Raw("select ((t.success  * 1.0 /t.total  * 1.0) * 100), ((t.failed * 1.0 / t.total * 1.0) * 100) from (select (select (select count(*) from retrieval_success_records) + (select count(*) from retrieval_failure_records)) as total, (select count(*) from retrieval_success_records)               as success,  (select count(*) from retrieval_failure_records) as failed) as t").Scan(&retrievalStats).Error
	if successFailRate != nil {
		api.ReturnError(ctx, w, r, successFailRate)
		return
	}

	api.WriteJSON(ctx, w, retrievalStats)

}

// GetContentDealsAttempted returns the total number of content deals attempted
// @Summary Returns the total number of content deals attempted
// @Description Returns the total number of content deals attempted
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} int
// @Router /stats/total-content-deals-attempted [get]
func GetContentDealsAttempted(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//select count(*) from content_deals as c1, contents as c2 where c1.id = c2.id;
	ctx := api.InitializeContext(r)
	var totalDealsAttempted int
	err := dao.DB.Raw("select count(*) from content_deals as c1, contents as c2 where c1.id = c2.id").Scan(&totalDealsAttempted).Error
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, totalDealsAttempted)

}

// GetTotalFiles returns the total number of files stored
// @Summary Returns the total number of files stored
// @Description Returns the total number of files stored
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} int
// @Router /stats/total-files [get]
func GetTotalFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	var totalFiles int64
	// select count(*) as "Total Files" from contents where active and not aggregate
	err := dao.DB.Raw("select count(*) from contents where active and not aggregate").Scan(&totalFiles).Error
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, totalFiles)
}

type TwitterStats struct {
	TotalContentDealsSize int64 `json:"totalContentDealsSize"`
	TotalContentDeals     int64 `json:"totalContentDeals"`
	TotalSealedDeals      int64 `json:"totalSealedDeals"`
	TotalUsers            int64 `json:"totalUsers"`
	TotalStorageProviders int64 `json:"totalStorageProviders"`
}

// GetStatsForTwitter returns the total number of storage deals attempted
// @Summary Returns the total number of storage deals attempted
// @Description Returns the total number of storage deals attempted
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} int
// @Router /stats/to-twitter [get]
func GetStatsForTwitter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if from == "" || to == "" {
		api.ReturnError(ctx, w, r, errors.New("from and to dates are required"))
		return
	}

	// validate date format
	_, err := time.Parse("2006-01-02", from)
	if err != nil {
		api.ReturnError(ctx, w, r, errors.New("from date is invalid"))
		return
	}

	_, err = time.Parse("2006-01-02", to)
	if err != nil {
		api.ReturnError(ctx, w, r, errors.New("to date is invalid"))
		return
	}

	// from date must be before to
	fromDate, _ := time.Parse("2006-01-02", from)
	toDate, _ := time.Parse("2006-01-02", to)

	if fromDate.After(toDate) {
		api.ReturnError(ctx, w, r, errors.New("from date must be before to date"))
		return
	}

	twitterStats, err := dao.Cacher.Get("/stats/to-twitter?from="+from+"&to="+to, time.Second*1, func() (interface{}, error) {
		var twitterStats TwitterStats
		var totalContentDealsSize sql.NullInt64
		err := dao.DB.Raw("select sum(c.size) as total from content_deals as cd, contents as c where (cd.created_at between ? and ?) and cd.deal_id > 0 and c.id = cd.content", from, to).Scan(&totalContentDealsSize).Error
		if err != nil {
			api.ReturnError(ctx, w, r, err)
			return nil, err
		}

		value, err := totalContentDealsSize.Value()
		if err != nil {
			return nil, err
		}

		if value == nil {
			twitterStats.TotalContentDealsSize = 0
		} else {
			twitterStats.TotalContentDealsSize = value.(int64)
		}

		err = dao.DB.Raw("select count(*) from content_deals where (created_at between ? and ?)", from, to).Scan(&twitterStats.TotalContentDeals).Error
		if err != nil {
			api.ReturnError(ctx, w, r, err)
			return nil, err
		}

		err = dao.DB.Raw("select count(*) from content_deals where deal_id > 0 and deleted_at is null and sealed_at between ? and ?", from, to).Scan(&twitterStats.TotalSealedDeals).Error
		if err != nil {
			api.ReturnError(ctx, w, r, err)
			return nil, err
		}

		err = dao.DB.Raw("select count(*) from users where (created_at between ? and ?)", from, to).Scan(&twitterStats.TotalUsers).Error
		if err != nil {
			api.ReturnError(ctx, w, r, err)
			return nil, err
		}

		err = dao.DB.Raw("select count(*) from storage_miners where (created_at between ? and ?)", from, to).Scan(&twitterStats.TotalStorageProviders).Error
		if err != nil {
			api.ReturnError(ctx, w, r, err)
			return nil, err

		}
		return twitterStats, nil
	})

	api.WriteJSON(ctx, w, twitterStats)
}

// GetTotalStorageInTib GetTotalStorage returns the total storage
// @Summary Returns the total storage
// @Description Returns the total storage
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} int
// @Router /stats/total-storage [get]
func GetTotalStorageInTib(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//select (sum(size)/1000000000000) as "Total Storage" from contents where active and not aggregated_in > 0
	ctx := api.InitializeContext(r)
	var totalStorage int64
	err := dao.DB.Raw("select (sum(size)/1000000000000) from contents where active and not aggregated_in > 0").Scan(&totalStorage).Error
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, totalStorage)
}

// GetStorageRateStats returns the storage rate stats
// @Summary Returns the storage rate stats
// @Description Returns the storage rate stats
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} StorageRateStats
// @Router /stats/storage-rates [get]
func GetStorageRateStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var storageRateStats *StorageRateStats

	//select ((t.success  * 1.0 /t.total  * 1.0) * 100) as "Success", ((t.failed * 1.0 / t.total * 1.0) * 100) as "Failure" from (select
	//	(select count(*) from content_deals as c1 where failed = false and deleted_at is null) as total,
	//	(select count(*) from content_deals as c1 where failed = false and deal_id > 0 and deleted_at is null) as success,
	//	(select count(*) from content_deals as c1 where failed = false and deal_id = 0 and deleted_at is null) as failed) as t;

	ctx := api.InitializeContext(r)
	err := dao.DB.Raw("" +
		"select " +
		" ((t.success  * 1.0 / t.total  * 1.0) * 100) as \"DealSuccessRate\"," +
		" ((t.failed * 1.0 / t.total * 1.0) * 100) as \"DealFailureRate\" " +
		"from (select (select count(*) from content_deals as c1 where failed = false and deleted_at is null) as total, " +
		"	(select count(*) from content_deals as c1 where failed = false and deal_id > 0 and deleted_at is null) as success, " +
		"	(select count(*) from content_deals as c1 where failed = false and deal_id = 0 and deleted_at is null) as failed" +
		") as t").Scan(&storageRateStats).Error

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	success, err := strconv.ParseFloat(storageRateStats.DealSuccessRate, 64)
	failure, err := strconv.ParseFloat(storageRateStats.DealFailureRate, 64)

	storageRateStats.DealSuccessRate = fmt.Sprintf("%.2f", success)
	storageRateStats.DealFailureRate = fmt.Sprintf("%.2f", failure)
	api.WriteJSON(ctx, w, storageRateStats)
}

//	GetPublicStats returns the public stats
//	@Summary Returns the public stats
//	@Description Returns the public stats
//	@Tags Stats
//	@Accept  json
//	@Produce  json
//	@Success 200 {object} PublicStats
//	@Router /stats/info [get]
func GetInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	// 	cache for 30mins
	stats, err := dao.Cacher.Get("/stats/info", time.Minute*60, func() (interface{}, error) {
		var stats PublicStats
		if err := dao.DB.Model(model.Content{}).Where("active and not aggregated_in > 0").Select("SUM(size) as total_storage").Scan(&stats).Error; err != nil {
			api.ReturnError(ctx, w, r, err)
			return stats, err
		}

		if err := dao.DB.Model(model.Content{}).Where("active and not aggregate").Count(&stats.TotalFilesStored.Int64).Error; err != nil {
			api.ReturnError(ctx, w, r, err)
			return stats, err
		}

		if err := dao.DB.Model(model.ContentDeal{}).Where("not failed and deal_id > 0").Count(&stats.DealsOnChain.Int64).Error; err != nil {
			api.ReturnError(ctx, w, r, err)
			return stats, err
		}

		//	this can be resource expensive but we are already caching it.
		if err := dao.DB.Table("obj_refs").Select("id").Order("id desc").Limit(1).Scan(&stats.TotalObjectsRef.Int64).Error; err != nil {
			api.ReturnError(ctx, w, r, err)
			return stats, err
		}

		//var objects []model.Object
		//var totalBytesUploadsize int64
		//dao.DB.Model(&model.Object{}).FindInBatches(&objects, 10000000,
		//	func(tx *gorm.DB, batch int) error {
		//		//tx.Select("SUM(size) as size").Scan(&totalBytesUploadsize)
		//		for _, object := range objects {
		//			totalBytesUploadsize += object.Size
		//		}
		//		return nil
		//	})

		//result.
		//stats.TotalBytesUploaded = sql.NullInt64{Int64: totalBytesUploadsize, Valid: true}
		stats.TotalBytesUploaded = stats.TotalFilesStored

		//if err := dao.DB.Table("objects").Select("SUM(size)").Find(&stats.TotalBytesUploaded.Int64).Error; err != nil {
		//	api.ReturnError(ctx, w, r, err)
		//	return stats, err
		//}

		if err := dao.DB.Model(model.User{}).Count(&stats.TotalUsers.Int64).Error; err != nil {
			api.ReturnError(ctx, w, r, err)
			return stats, err
		}

		if err := dao.DB.Table("storage_miners").Count(&stats.TotalStorageMiner.Int64).Error; err != nil {
			api.ReturnError(ctx, w, r, err)
			return stats, err
		}
		return stats, nil
	})
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	jsonResponse := map[string]interface{}{
		"totalStorage":       stats.(PublicStats).TotalStorage.Int64,
		"totalFilesStored":   stats.(PublicStats).TotalFilesStored.Int64,
		"dealsOnChain":       stats.(PublicStats).DealsOnChain.Int64,
		"totalObjectsRef":    stats.(PublicStats).TotalObjectsRef.Int64,
		"totalBytesUploaded": stats.(PublicStats).TotalBytesUploaded.Int64,
		"totalUsers":         stats.(PublicStats).TotalUsers.Int64,
		"totalStorageMiner":  stats.(PublicStats).TotalStorageMiner.Int64,
	}

	api.WriteJSON(ctx, w, jsonResponse)

}

// GetDealMetrics returns the deal metrics
// @Summary Returns the deal metrics
// @Description Returns the deal metrics
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} []DealMetricsInfo
// @Router /stats/deal-metrics [get]
func GetDealMetrics(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	dealMetrics, err := dao.Cacher.Get("/stats/deal-metrics", time.Minute*30, func() (interface{}, error) {
		return computeDealMetrics()
	})

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, dealMetrics)
}

// computeDealMetrics computes the deal metrics
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
