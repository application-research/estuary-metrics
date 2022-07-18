package core

import (
	"github.com/application-research/estuary-metrics/core/generated/query"
)

type StorageDealSuccessRate struct {
	ContentWithDeal    int64
	ContentWithoutDeal int64
	SuccessRate        float64
	FailureRate        float64
}

type StorageAcceptanceRate struct {
	SuccessRate float64
	FailureRate float64
}

func (m Metrics) GetStorageDealSuccessRate() (StorageDealSuccessRate, error) {
	// estuary + deal with miner
	contentDeal := query.Use(DB).ContentDeal
	countWithDeal, errWithDeal := contentDeal.WithContext(m.Context).Where(contentDeal.DealID.Neq(0)).Count()
	countWithoutDeal, errWithoutDeal := contentDeal.WithContext(m.Context).Where(contentDeal.DealID.Eq(0)).Count()

	if errWithDeal != nil {
		return StorageDealSuccessRate{}, errWithDeal
	}

	if errWithoutDeal != nil {
		return StorageDealSuccessRate{}, errWithoutDeal
	}

	totalRecordWithDeal := countWithDeal + countWithoutDeal
	successRate := (float64(countWithDeal) / float64(totalRecordWithDeal)) * 100
	failureRate := (float64(countWithoutDeal) / float64(totalRecordWithDeal)) * 100

	return StorageDealSuccessRate{
		ContentWithDeal:    countWithDeal,
		ContentWithoutDeal: countWithoutDeal,
		SuccessRate:        successRate,
		FailureRate:        failureRate,
	}, nil
}

func (m Metrics) GetStorageAcceptanceRate() (StorageAcceptanceRate, error) {
	// accepted to estuary + deal record
	// estuary + deal with miner
	content := query.Use(DB).Content
	contentFailed, errFailed := content.WithContext(m.Context).Where(content.Failed.Is(true)).Count()
	contentSuccess, errSuccess := content.WithContext(m.Context).Where(content.Failed.Is(false)).Count()

	if errFailed != nil {
		return StorageAcceptanceRate{}, errFailed
	}

	if errSuccess != nil {
		return StorageAcceptanceRate{}, errSuccess
	}

	totalRecordContent := contentFailed + contentSuccess
	contentFailedRate := (float64(contentFailed) / float64(totalRecordContent)) * 100
	contentSuccessRate := (float64(contentSuccess) / float64(totalRecordContent)) * 100

	return StorageAcceptanceRate{
		SuccessRate: contentSuccessRate,
		FailureRate: contentFailedRate,
	}, nil
}

func (m Metrics) GetTotalNumberOfStorageDealsAttempted() (int64, error) {
	// total deals attempted (total)
	return 0, nil
}

func (m Metrics) GetTotalNumberOfStorageDealsSuccessful() (int64, error) {
	// total deals success (not null miner)
	return 0, nil
}
func (m Metrics) GetDistributionDataSizeUploadedPerUser() (int64, error) {
	// total data uploaded per user
	return 0, nil
}
