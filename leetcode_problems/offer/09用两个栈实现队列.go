package offer

type CQueue09 struct {
	headers, tails []int
}

func Constructor09() CQueue09 {
	return CQueue09{}
}

func (q *CQueue09) AppendTail09(value int) {
	q.tails = append(q.tails, value)
}

func (q *CQueue09) DeleteHead09() int {
	// header为空，则将tails的内容全部append过去
	if len(q.headers) == 0 && len(q.tails) == 0 {
		return -1
	} else if len(q.headers) == 0 {
		q.headers, q.tails = q.tails, q.headers
	}
	res := q.headers[0]
	q.headers = q.headers[1:]
	return res
}
