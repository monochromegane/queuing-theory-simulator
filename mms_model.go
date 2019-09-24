package queuing_theory_simulator

import (
	"math"
)

type MMSModel struct {
	timeArrival Exper
	timeService Exper

	requests    *requests
	servers     map[int]*request
	currentStep int
	arrivalStep int
}

func NewMMSModel(seed int64, lambda, mu func(int) float64) *MMSModel {
	return &MMSModel{
		timeArrival: NewChangeExp(seed, lambda),
		timeService: NewChangeExp(seed+1, mu),
		requests:    newRequests(),
		servers:     map[int]*request{},
	}
}

func (m *MMSModel) Progress(s int) (int, int, int, []int) {
	arrival := 0
	responseTimes := []int{}
	if m.currentStep == m.arrivalStep {
		// Queue request
		req := &request{
			arriveAt:    m.currentStep,
			serviceTime: int(math.Ceil(m.timeService.Exp(m.currentStep))),
		}
		m.requests.push(req)
		arrival++

		// Calculate next arrival step
		m.arrivalStep = m.currentStep + int(math.Ceil(m.timeArrival.Exp(m.currentStep)))
	}

	// Start service from request queue
	for i := 0; i < s; i++ {
		if s <= len(m.servers) {
			break
		}
		if m.requests.lenQueue() == 0 {
			break
		}
		reqId, req := m.requests.pop()
		req.beginAt = m.currentStep
		req.endAt = m.currentStep + req.serviceTime - 1
		m.servers[reqId] = req
	}

	processing := len(m.servers)
	waiting := m.requests.lenQueue()

	// Stop service
	for reqId, req := range m.servers {
		if req.endAt == m.currentStep {
			delete(m.servers, reqId)
			responseTimes = append(responseTimes, req.waitTime()+req.serviceTime)
		}
	}
	m.currentStep++

	return arrival, processing, waiting, responseTimes
}

func (m *MMSModel) History() [][]int {
	history := make([][]int, len(m.requests.list))
	for i, req := range m.requests.list {
		history[i] = []int{
			req.arriveAt,
			req.serviceTime,
			req.beginAt,
			req.endAt,
		}
	}
	return history
}
