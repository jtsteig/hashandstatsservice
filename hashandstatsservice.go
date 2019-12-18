package hashandstatsservice

import (
	encodedhash "github.com/jtsteig/encodedhash"
	hashmodel "github.com/jtsteig/hashmodels"
)

// HashStatsService is the orchestration entry point for storing entries and stats.
type HashStatsService struct {
	hashRepository *hashmodel.HashRepository
}

// StoreValue takes a plain string value, calculates the hash and stores it and the elapsed time to persistence.
func (service *HashStatsService) StoreValue(value string) (int, error) {
	hash, duration := encodedhash.CalculateHash(value)

	countID, err := service.hashRepository.StoreHash(hash, duration.Milliseconds())
	if err != nil {
		return -1, err
	}

	return countID, nil
}

// GetHash returns the HashStats for a countID or an error if not found.
func (service *HashStatsService) GetHash(countID int) (hashmodel.HashStat, error) {
	hashStat, err := service.hashRepository.GetHashStat(countID)
	if err != nil {
		return hashmodel.HashStat{}, err
	}
	return hashStat, err
}

// GetTotalStats returns the total stats for all runs and returns an error if anything goes amiss.
func (service *HashStatsService) GetTotalStats() (hashmodel.TotalStats, error) {
	totalStats, err := service.hashRepository.GetTotalStats()
	if err != nil {
		return hashmodel.TotalStats{}, err
	}
	return totalStats, err
}