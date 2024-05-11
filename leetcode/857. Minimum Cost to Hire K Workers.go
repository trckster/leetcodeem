package main

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"sort"
)

type Worker struct {
	quality int
	wage    int
}

func byMinQuality(a, b interface{}) int {
	priorityA := a.(Worker).quality
	priorityB := b.(Worker).quality

	return -utils.IntComparator(priorityA, priorityB)
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]Worker, len(quality))
	for i := 0; i < len(quality); i++ {
		workers[i] = Worker{quality[i], wage[i]}
	}

	sort.Slice(workers, func(i, j int) bool {
		aRatio := float64(workers[i].wage) / float64(workers[i].quality)
		bRatio := float64(workers[j].wage) / float64(workers[j].quality)
		return aRatio < bRatio
	})

	queue := pq.NewWith(byMinQuality)

	slidingQualitySum := float64(0)
	for i := 0; i < k; i++ {
		queue.Enqueue(workers[i])
		slidingQualitySum += float64(workers[i].quality)
	}

	answer := slidingQualitySum * float64(workers[k-1].wage) / float64(workers[k-1].quality)
	for i := k; i < len(workers); i++ {
		w, _ := queue.Dequeue()
		worker := w.(Worker)

		slidingQualitySum -= float64(worker.quality)
		slidingQualitySum += float64(workers[i].quality)

		queue.Enqueue(workers[i])
		answer = min(answer, slidingQualitySum*float64(workers[i].wage)/float64(workers[i].quality))
	}

	return answer
}
