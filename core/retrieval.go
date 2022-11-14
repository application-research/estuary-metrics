package core

import (
	"github.com/application-research/estuary-metrics/core/generated/query"
	"log"
)

type RetrievalDealRate struct {
	AcceptanceRate int64
	SuccessRate    float64
	FailureRate    float64
}
type RetrievalAcceptanceDealRate struct {
	AcceptanceRate int64
}

func (m Metrics) GetRetrievalDealSuccessRate() (RetrievalDealRate, error) {
	retreivalSuccessU := query.Use(DB).RetrievalSuccessRecord
	retrievalFailureU := query.Use(DB).RetrievalFailureRecord
	retrievalSuccess, errRetSucc := retreivalSuccessU.WithContext(m.Context).Count()
	retrievalFail, errRetFail := retrievalFailureU.WithContext(m.Context).Count()

	if errRetSucc != nil {
		log.Println("Error getting retrieval success rate: ", errRetSucc)
	}

	if errRetFail != nil {
		log.Println("Error getting retrieval failure rate: ", errRetFail)
	}

	retrievalTotal := retrievalSuccess + retrievalFail
	retrievalSuccessRate := (float64(retrievalSuccess) / float64(retrievalTotal)) * 100
	retrievalFailureRate := (float64(retrievalFail) / float64(retrievalTotal)) * 100

	return RetrievalDealRate{
		AcceptanceRate: retrievalTotal,
		SuccessRate:    retrievalSuccessRate,
		FailureRate:    retrievalFailureRate,
	}, nil
}

func (m Metrics) GetRetrievalDealAcceptanceRate() (RetrievalAcceptanceDealRate, error) {
	retreivalSuccessU := query.Use(DB).RetrievalSuccessRecord
	retrievalFailureU := query.Use(DB).RetrievalFailureRecord
	retrievalSuccess, errRetSucc := retreivalSuccessU.WithContext(m.Context).Count()
	retrievalFail, errRetFail := retrievalFailureU.WithContext(m.Context).Count()

	if errRetSucc != nil {
		log.Println("Error getting retrieval success rate: ", errRetSucc)
	}

	if errRetFail != nil {
		log.Println("Error getting retrieval failure rate: ", errRetFail)
	}

	retrievalTotal := retrievalSuccess + retrievalFail
	return RetrievalAcceptanceDealRate{
		AcceptanceRate: retrievalTotal,
	}, nil

}

func (m Metrics) GetTotalRetrievals() (int64, error) {
	retrievalSuccessU := query.Use(DB).RetrievalSuccessRecord
	retrievalFailureU := query.Use(DB).RetrievalFailureRecord
	retrievalSuccess, errRetSucc := retrievalSuccessU.WithContext(m.Context).Count()
	retrievalFail, errRetFail := retrievalFailureU.WithContext(m.Context).Count()

	if errRetSucc != nil {
		log.Println("Error getting retrieval success rate: ", errRetSucc)
	}

	if errRetFail != nil {
		log.Println("Error getting retrieval failure rate: ", errRetFail)
	}

	retrievalTotal := retrievalSuccess + retrievalFail
	return retrievalTotal, nil
}

func (m Metrics) GetRetrievalSuccess() (int64, error) {
	retrievalSuccessU := query.Use(DB).RetrievalSuccessRecord
	retrievalSuccess, errRetSucc := retrievalSuccessU.WithContext(m.Context).Count()

	if errRetSucc != nil {
		log.Println("Error getting retrieval success rate: ", errRetSucc)
	}

	return retrievalSuccess, nil
}

func (m Metrics) GetRetrievalFail() (int64, error) {
	retrievalFailureU := query.Use(DB).RetrievalFailureRecord
	retrievalFail, errRetFail := retrievalFailureU.WithContext(m.Context).Count()

	if errRetFail != nil {
		log.Println("Error getting retrieval failure rate: ", errRetFail)
		return 0, errRetFail
	}

	return retrievalFail, nil
}

type MinerRetrievalSpeed struct {
	Miner string
	Speed int64 // seconds
}

func (m Metrics) GetRetrievalAverageSpeedTop(top int) ([]MinerRetrievalSpeed, error) {
	//select a.miner, (avg(a.average_speed)/1000) as "Average Speed" from retrieval_success_records a where a.average_speed > 0 group by a.miner order by avg(a.average_speed) asc limit 10;
	var minerRetrievalSpeed []MinerRetrievalSpeed
	err := DB.Raw("select a.miner, (avg(a.average_speed)/1000) as \"Average Speed\" from retrieval_success_records a where a.average_speed > 0 group by a.miner order by avg(a.average_speed) asc limit ?", top).Scan(&minerRetrievalSpeed).Error

	if err != nil {
		log.Println("Error getting retrieval average speed: ", err)
		return nil, err
	}
	return minerRetrievalSpeed, nil
}
