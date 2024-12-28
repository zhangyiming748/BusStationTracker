package main

import (
	"bus/constant"
	"bus/station"
	"fmt"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	tally()
}

func tally() {
	reps := []constant.Rep{}
	if position, err := station.GetInPosition(constant.BAJIAOROADEAST); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.GALAXYSTREET); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.JINGYUANROADEAST); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.LAOSHAN); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.NATIONALBOTANIXALGARDEN); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.SHIJINGSHANSCULPTUREPARK); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.SHIJINGSHANTECHNOLOGYMUSEUM); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.TIMESGARDENNORTH); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	for _, rep := range reps {
		fmt.Println(rep)
	}
}
