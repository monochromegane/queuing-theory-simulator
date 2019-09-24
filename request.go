package queuing_theory_simulator

import "container/list"

type request struct {
	arriveAt    int
	serviceTime int
	beginAt     int
	endAt       int
}

func (r request) waitTime() int {
	return r.beginAt - r.arriveAt
}

type requests struct {
	queue *list.List
	list  []*request
}

func newRequests() *requests {
	return &requests{
		queue: list.New(),
		list:  []*request{},
	}
}

func (rs *requests) push(r *request) {
	rs.list = append(rs.list, r)
	rs.queue.PushBack(len(rs.list) - 1)
}

func (rs *requests) pop() (int, *request) {
	elem := rs.queue.Front()
	defer rs.queue.Remove(elem)

	reqId, _ := elem.Value.(int)
	return reqId, rs.list[reqId]
}

func (rs requests) lenQueue() int {
	return rs.queue.Len()
}
