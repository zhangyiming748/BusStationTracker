package station

import (
	"bus/constant"
	"bus/util"
	"encoding/json"
	"fmt"
	"log"
)

func GetInPosition(station string) ([]constant.Rep, error) {
	reps := []constant.Rep{}
	positive, err := Positive(station, false)
	if err != nil {
		return nil, err
	}
	reps = append(reps, positive)
	negative, err := Negative(station, true)
	if err != nil {
		return nil, err
	}
	reps = append(reps, negative)
	return reps, nil
}

func Positive(station string, direction bool) (constant.Rep, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	var o string
	if direction {
		o = "1"
	}
	params := map[string]string{
		"type": "json",
		"city": "北京",
		"line": station,
		"o":    o,
	}
	get, err := util.HttpGet(headers, params, "https://api.lolimi.cn/API/che/api.php")
	if err != nil {
		return constant.Rep{}, err
	}
	var rep constant.Rep
	err = json.Unmarshal(get, &rep)
	if err != nil {
		return constant.Rep{}, err
	}
	if rep.Code != 200 {
		return constant.Rep{}, fmt.Errorf("当前站点:%v查询失败:%v", station, rep.Code)
	}
	Parse(rep)
	return constant.Rep{}, nil
}

func Negative(station string, direction bool) (constant.Rep, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	var o string
	if !direction {
		o = "2"
	}
	params := map[string]string{
		"type": "json",
		"city": "北京",
		"line": station,
		"o":    o,
	}
	get, err := util.HttpGet(headers, params, "https://api.lolimi.cn/API/che/api.php")
	if err != nil {
		return constant.Rep{}, err
	}
	var rep constant.Rep
	err = json.Unmarshal(get, &rep)
	if err != nil {
		return constant.Rep{}, err
	}
	if rep.Code != 200 {
		return constant.Rep{}, fmt.Errorf("当前站点:%v查询失败:%v", station, rep.Code)
	}
	Parse(rep)
	return rep, nil
}
func Parse(rep constant.Rep) {
	for _, data := range rep.Data {
		log.Printf("站点名:%s\t线路名:%s\t开往%s\t预计到站时间%s\t预计剩余时间:%s\t预计剩余站数:%s\n", rep.Line, data.Lines, data.EndSn, data.Reachtime, data.TravelTime, data.Surplus)
	}
}
