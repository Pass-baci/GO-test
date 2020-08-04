package main

import (
	"fmt"
	"time"
)

type Bar struct {
	percent int64
	cur int64
	total int64
	rate string
	graph string
}

func (b *Bar)NewOptionWithGraph(start, total int64, graph string) {
	b.graph = graph
	b.NewOption(start,total)
}

func (b *Bar)NewOption(start, total int64){
	b.cur = start
	b.total = total
	if b.graph == "" {
		b.graph = "█"
	}
	b.percent = b.getPercent()
	for i := 0; i < int(b.percent); i += 2 {
		b.rate += b.graph
	}
}

func (b *Bar) getPercent() int64 {
	return int64(float64(b.cur) / float64(b.total) * 100)
}

func (b *Bar) Play(cur int64) {
	b.cur = cur
	last := b.percent
	b.percent = b.getPercent()
	if b.percent != last && b.percent%2 == 0 {
		b.rate += b.graph
	}
	fmt.Printf("\r[%-49s]%3d%%  %8d/%d", b.rate, b.percent, b.cur, b.total)
}

func (b *Bar) Finish(){
	fmt.Println()
}

func main() {
	var bar Bar
	bar.NewOption(0, 100)
	for i:= 0; i<=100; i++{
		time.Sleep(100*time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}
