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
