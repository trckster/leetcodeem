package main

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
	"sort"
)

func roomSortPriority(a, b interface{}) int {
	return utils.IntComparator(a.(int), b.(int))
}

type Release struct {
	Time int
	Room int
}

func releaseSortPriority(a, b interface{}) int {
	releaseA := a.(Release)
	releaseB := b.(Release)

	if releaseA.Time == releaseB.Time {
		return utils.IntComparator(a.(Release).Room, b.(Release).Room)
	}

	return utils.IntComparator(a.(Release).Time, b.(Release).Time)
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	availableRooms := pq.NewWith(roomSortPriority)
	for i := 0; i < n; i++ {
		availableRooms.Enqueue(i)
	}

	bookingsCount := make([]int, n)
	releases := pq.NewWith(releaseSortPriority)
	for _, meeting := range meetings {
		now := meeting[0]
		duration := meeting[1] - meeting[0]

		for !releases.Empty() {
			release, _ := releases.Peek()
			if release.(Release).Time > now {
				break
			}

			releases.Dequeue()
			availableRooms.Enqueue(release.(Release).Room)
		}

		if availableRooms.Empty() {
			nextRelease, _ := releases.Dequeue()
			releases.Enqueue(Release{nextRelease.(Release).Time + duration, nextRelease.(Release).Room})
			bookingsCount[nextRelease.(Release).Room] += 1
		} else {
			room, _ := availableRooms.Dequeue()
			releases.Enqueue(Release{now + duration, room.(int)})
			bookingsCount[room.(int)] += 1
		}
	}

	answer := -1
	maxCount := -1
	for room, cnt := range bookingsCount {
		if cnt > maxCount {
			maxCount = cnt
			answer = room
		}
	}

	return answer
}
