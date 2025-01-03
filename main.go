package main

import (
	"bus/constant"
	"bus/station"
	"bus/util"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	util.SetLog("bus.log")
	tally()
}

func tally() {
	var positions = []string{
		constant.BAJIAOROADEAST,
		constant.GALAXYSTREET,
		constant.JINGYUANROADEAST,
		constant.LAOSHAN,
		constant.NATIONALBOTANIXALGARDEN,
		constant.SHIJINGSHANSCULPTUREPARK,
		constant.SHIJINGSHANTECHNOLOGYMUSEUM,
		constant.TIMESGARDENNORTH,
	}
	for _, position := range positions {
		station.GetInPosition(position)
	}
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
	if position, err := station.GetInPosition(constant.JINGYANHOTEL); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
	if position, err := station.GetInPosition(constant.BAJIAO); err == nil {
		reps = append(reps, position...)
		time.Sleep(1 * time.Second)
	}
}
