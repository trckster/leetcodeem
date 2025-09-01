package main

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type Fraction struct {
	numerator   int
	denominator int
	benefit     float64
}

func NewFraction(numerator int, denominator int) *Fraction {
	f := &Fraction{numerator: numerator, denominator: denominator}
	f.updateBenefit()
	return f
}

func (f *Fraction) Bump() {
	f.numerator++
	f.denominator++
	f.updateBenefit()
}

func (f *Fraction) updateBenefit() {
	now := float64(f.numerator) / float64(f.denominator)
	willBe := float64(f.numerator+1) / float64(f.denominator+1)

	f.benefit = willBe - now
}

func byBenefit(a, b interface{}) int {
	return -utils.Float64Comparator(a.(*Fraction).benefit, b.(*Fraction).benefit)
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	queue := pq.NewWith(byBenefit)

	for _, class := range classes {
		queue.Enqueue(NewFraction(class[0], class[1]))
	}

	for extraStudents > 0 {
		queueItem, _ := queue.Peek()
		f := queueItem.(*Fraction)

		if f.benefit <= 0 {
			break
		}

		_, _ = queue.Dequeue()

		f.Bump()
		queue.Enqueue(f)
		extraStudents--
	}

	ratio := float64(0)
	for !queue.Empty() {
		queueItem, _ := queue.Dequeue()
		f := queueItem.(*Fraction)

		ratio += float64(f.numerator) / float64(f.denominator)
	}

	return ratio / float64(len(classes))
}
