package main

import "fmt"

type IDGenerator struct{
	AvailableIds []uint64
	ch_proc chan chan uint64
}

// todo 反复重试
func (gen *IDGenerator) GetNewIds(category string) {
	bound := GetCurrentBound(category)
	Acquire(category, bound, CAPACITY)
	gen.AvailableIds = append(gen.AvailableIds, NewSlice(bound, CAPACITY, 1)...)
	fmt.Println(gen.AvailableIds)
}

func (gen *IDGenerator) Init() {
	// 连接Mysql
	InitDB()
	gen.GetNewIds(CATEGORY)

	// 构建channel
	gen.ch_proc = make(chan chan uint64, NUM_CHANNELS)
	go gen.uuid_task()
}

func (gen *IDGenerator) GetNewId() (uint64, error) {
	req := make(chan uint64, 1)
	gen.ch_proc <- req
	return <-req, nil
}

func (gen *IDGenerator) uuid_task() {
	for {
		ret := <-gen.ch_proc  //收到新请求
		x := gen.AvailableIds[0]
		gen.AvailableIds = gen.AvailableIds[1:]
		ret <- x
		if len(gen.AvailableIds) < CAPACITY * CAP_RATE {
			gen.GetNewIds(CATEGORY)
		}
	}
}

func NewSlice(start uint64, count, step uint64) []uint64 {
	s := make([]uint64, count)
	start += 1
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}
