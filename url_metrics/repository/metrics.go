package repository

import (
	"fmt"
	"sort"
	"sync"
	"url-shortner/domain"
)

type Metric struct {
	Domain string
	Count  int
}

type inMemoryMetricStore struct {
	countMetricStore map[string]Metric
	lock             sync.Mutex
}

// return new in memory store for metrics
func NewInMemoryMetricStore() domain.DomainMetricsRepository {
	return &inMemoryMetricStore{
		countMetricStore: make(map[string]Metric),
	}
}

// increements domain shorten count in store
func (store *inMemoryMetricStore) IncreementDomainCountMetric(domain string) error {
	//check if store exists
	if store.countMetricStore == nil {
		return fmt.Errorf("store not initialized")
	}

	//acquire a lock on map to avoid simultaneous read write
	store.lock.Lock()
	//release lock on return
	defer store.lock.Unlock()

	if metric, ok := store.countMetricStore[domain]; !ok {
		//if not found then set the count to 1
		store.countMetricStore[domain] = Metric{Domain: domain, Count: 1}
	} else {
		//if found then increement the count
		metric.Count += 1
		store.countMetricStore[domain] = metric
	}
	return nil
}

// return top shortened domain list from store
func (store *inMemoryMetricStore) GetTopDomains(limit int) (map[string]int, error) {
	//check if store exists
	if store.countMetricStore == nil {
		return nil, fmt.Errorf("store not initialized")
	}

	//acquire a lock on map to avoid simultaneous read write
	store.lock.Lock()

	//get the values in array
	metricArray := []Metric{}
	for _, metric := range store.countMetricStore {
		metricArray = append(metricArray, metric)
	}

	//release lock
	store.lock.Unlock()

	//sort the metric array
	sort.Slice(metricArray, func(i, j int) bool {
		return metricArray[i].Count > metricArray[j].Count
	})

	topNDomainMap := make(map[string]int)

	if limit == 0 {
		limit = 3
	}

	if limit > len(metricArray) {
		limit = len(metricArray)
	}

	for _, metric := range metricArray[:limit] {
		topNDomainMap[metric.Domain] = metric.Count
	}
	return topNDomainMap, nil
}
