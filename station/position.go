package station

import (
	"bus/constant"
	"bus/util"
	"encoding/json"
	"fmt"
	"log"
)

func GetInPosition(station string) ([]Pretty, error) {
	reps := []Pretty{}
	positive, err := Positive(station, false)
	if err != nil {
		return nil, err
	}
	reps = append(reps, positive...)
	negative, err := Negative(station, true)
	if err != nil {
		return nil, err
	}
	reps = append(reps, negative...)
	return reps, nil
}

func Positive(station string, direction bool) ([]Pretty, error) {
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
		return []Pretty{}, err
	}
	var rep constant.Rep
	err = json.Unmarshal(get, &rep)
	if err != nil {
		return []Pretty{}, err
	}
	if rep.Code != 200 {
		return []Pretty{}, fmt.Errorf("当前站点:%v查询失败:%v", station, rep.Code)
	}
	log.Printf("正向信息:%+v\n", rep)
	Parse(rep)
	return Parse(rep), nil
}

func Negative(station string, direction bool) ([]Pretty, error) {
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
		return []Pretty{}, err
	}
	var rep constant.Rep
	err = json.Unmarshal(get, &rep)
	if err != nil {
		return []Pretty{}, err
	}
	if rep.Code != 200 {
		return []Pretty{}, fmt.Errorf("当前站点:%v查询失败:%v", station, rep.Code)
	}
	log.Printf("反向信息:%+v\n", rep)
	return Parse(rep), nil
}

func Parse(rep constant.Rep) []Pretty {
	/*
	   	2024/12/29 12:02:46 正向信息:{Code:200 City:北京 Line:银河大街 CarCount:15 Data:[{Lines:325路 Price:2~3元 EndSn:衙门口公交场站 BusId:7008005065 Reachtime:2024-12-29 12:13 TravelTime:11分 Surplus:6站} {Lines:325路 Price:2~3元 EndSn:衙门口公交场站 BusId:7008005062 Reachtime:2024-12-29 12:14 TravelTime:12分 Surplus:7站} {Lines:472路 Price:2~3元 EndSn:鲁谷公交场站 BusId:7005095068 Reachtime:2024-12-29 12:06 TravelTime:3分 Surplus:1站} {Lines:472路 Price:2~3元 EndSn:鲁谷公交场站 BusId:7005055960 Reachtime:2024-12-29 12:17 TravelTime:15分 Surplus:6站} {Lines:546路 Price:2~4元 EndSn:富丰桥西 BusId:7005055866 Reachtime:2024-12-29 12:05 TravelTime:3分 Surplus:1站} {Lines:574路 Price:2~4元 EndSn:衙门口东 BusId:4178790580 Reachtime:2024-12-29 12:12 TravelTime:9分 Surplus:3站} {Lines:574路 Price:2~4元 EndSn:衙门口东 BusId:4178790586 Reachtime:2024-12-29 12:25 TravelTime:23分 Surplus:9站} {Lines:597路 Price:2~4元 EndSn:鲁谷公交场站 BusId:7206045660 Reachtime:2024-12-29 12:06 TravelTime:4分 Surplus:2站} {Lines:597路 Price:2~4元 EndSn:鲁谷公交场站 BusId:7206045666 Reachtime:2024-12-29 12:10 TravelTime:7分 Surplus:3站} {Lines:598路 Price:2~3元 EndSn:大瓦窑公交场站 BusId:4571760889 Reachtime:2024-12-29 12:11 TravelTime:9分 Surplus:3站} {Lines:598路 Price:2~3元 EndSn:大瓦窑公交场站 BusId:7203015765 Reachtime:2024-12-29 12:16 TravelTime:14分 Surplus:7站} {Lines:663路 Price:2~4元 EndSn:小马厂 BusId:7002005264 Reachtime:1970-01-01 08:00 TravelTime:-0分 Surplus:已开出站点了哦} {Lines:663路 Price:2~4元 EndSn:小马厂 BusId:7002005068 Reachtime:2024-12-29 12:11 TravelTime:9分 Surplus:4站} {Lines:959路 Price:2~8元 EndSn:大瓦窑公交场站 BusId:7300085166 Reachtime:2024-12-29 12:06 TravelTime:4分 Surplus:2站} {Lines:959路 Price:2~8元 EndSn:大瓦窑公交场站 BusId:7801025162 Reachtime:2024-12-29 12:29 TravelTime:27分 Surplus:12站}]}
	      2024/12/29 12:02:48 反向信息:{Code:200 City:北京 Line:银河大街 CarCount:15 Data:[{Lines:325路 Price:2~3元 EndSn:衙门口公交场站 BusId:7008005065 Reachtime:2024-12-29 12:13 TravelTime:11分 Surplus:6站} {Lines:325路 Price:2~3元 EndSn:衙门口公交场站 BusId:7008005062 Reachtime:2024-12-29 12:14 TravelTime:12分 Surplus:7站} {Lines:472路 Price:2~3元 EndSn:鲁谷公交场站 BusId:7005095068 Reachtime:2024-12-29 12:06 TravelTime:3分 Surplus:1站} {Lines:472路 Price:2~3元 EndSn:鲁谷公交场站 BusId:7005055960 Reachtime:2024-12-29 12:17 TravelTime:15分 Surplus:6站} {Lines:546路 Price:2~4元 EndSn:富丰桥西 BusId:7005055866 Reachtime:2024-12-29 12:05 TravelTime:3分 Surplus:1站} {Lines:574路 Price:2~4元 EndSn:衙门口东 BusId:4178790580 Reachtime:2024-12-29 12:12 TravelTime:9分 Surplus:3站} {Lines:574路 Price:2~4元 EndSn:衙门口东 BusId:4178790586 Reachtime:2024-12-29 12:25 TravelTime:23分 Surplus:9站} {Lines:597路 Price:2~4元 EndSn:鲁谷公交场站 BusId:7206045660 Reachtime:2024-12-29 12:07 TravelTime:4分 Surplus:2站} {Lines:597路 Price:2~4元 EndSn:鲁谷公交场站 BusId:7206045666 Reachtime:2024-12-29 12:10 TravelTime:8分 Surplus:3站} {Lines:598路 Price:2~3元 EndSn:大瓦窑公交场站 BusId:4571760889 Reachtime:2024-12-29 12:11 TravelTime:9分 Surplus:3站} {Lines:598路 Price:2~3元 EndSn:大瓦窑公交场站 BusId:7203015765 Reachtime:2024-12-29 12:16 TravelTime:14分 Surplus:7站} {Lines:663路 Price:2~4元 EndSn:小马厂 BusId:7002005264 Reachtime:1970-01-01 08:00 TravelTime:-0分 Surplus:已开出站点了哦} {Lines:663路 Price:2~4元 EndSn:小马厂 BusId:7002005068 Reachtime:2024-12-29 12:11 TravelTime:9分 Surplus:4站} {Lines:959路 Price:2~8元 EndSn:大瓦窑公交场站 BusId:7300085166 Reachtime:2024-12-29 12:06 TravelTime:4分 Surplus:2站} {Lines:959路 Price:2~8元 EndSn:大瓦窑公交场站 BusId:7801025162 Reachtime:2024-12-29 12:29 TravelTime:27分 Surplus:12站}]}
	      2024/12/29 12:02:50 正向信息:{Code:200 City:北京 Line:京原路口东 CarCount:26 Data:[{Lines:325路 Price:2~3元 EndSn:衙门口公交场站 BusId:7008005065 Reachtime:2024-12-29 12:10 TravelTime:7分 Surplus:3站} {Lines:325路 Price:2~3元 EndSn:衙门口公交场站 BusId:7008005062 Reachtime:2024-12-29 12:11 TravelTime:9分 Surplus:5站} {Lines:337路 Price:2~3元 EndSn:公主坟西 BusId:7005095868 Reachtime:2024-12-29 12:11 TravelTime:9分 Surplus:3站} {Lines:337路 Price:2~3元 EndSn:公主坟西 BusId:4770720986 Reachtime:2024-12-29 12:20 TravelTime:18分 Surplus:8站} {Lines:472路 Price:2~3元 EndSn:鲁谷公交场站 BusId:7005095068 Reachtime:2024-12-29 12:03 TravelTime:1分 Surplus:即将到站} {Lines:472路 Price:2~3元 EndSn:鲁谷公交场站 BusId:7005055960 Reachtime:2024-12-29 12:13 TravelTime:11分 Surplus:4站} {Lines:527路 Price:2~3元 EndSn:吴庄公交场站 BusId:7206065062 Reachtime:2024-12-29 12:04 TravelTime:2分 Surplus:1站} {Lines:527路 Price:2~3元 EndSn:吴庄公交场站 BusId:4178780689 Reachtime:2024-12-29 12:14 TravelTime:12分 Surplus:7站} {Lines:597路 Price:2~4元 EndSn:鲁谷公交场站 BusId:7206045660 Reachtime:2024-12-29 12:03 TravelTime:1分 Surplus:即将到站} {Lines:597路 Price:2~4元 EndSn:鲁谷公交场站 BusId:7206045666 Reachtime:2024-12-29 12:07 TravelTime:5分 Surplus:1站} {Lines:598路 Price:2~3元 EndSn:大瓦窑公交场站 BusId:4571760889 Reachtime:2024-12-29 12:06 TravelTime:4分 Surplus:1站} {Lines:598路 Price:2~3元 EndSn:大瓦窑公交场站 BusId:7203015765 Reachtime:2024-12-29 12:14 TravelTime:11分 Surplus:5站} {Lines:663路 Price:2~4元 EndSn:小马厂 BusId:7002005068 Reachtime:2024-12-29 12:07 TravelTime:5分 Surplus:1站} {Lines:663路 Price:2~4元 EndSn:小马厂 BusId:7002005267 Reachtime:2024-12-29 12:16 TravelTime:13分 Surplus:7站} {Lines:941路 Price:2~7元 EndSn:莲花池南路 BusId:4178720882 Reachtime:2024-12-29 12:08 TravelTime:6分 Surplus:2站} {Lines:941路 Price:2~7元 EndSn:莲花池南路 BusId:7206095766 Reachtime:2024-12-29 12:22 TravelTime:20分 Surplus:8站} {Lines:959路 Price:2~8元 EndSn:大瓦窑公交场站 BusId:7300085166 Reachtime:2024-12-29 12:04 TravelTime:2分 Surplus:即将到站} {Lines:959路 Price:2~8元 EndSn:大瓦窑公交场站 BusId:7801025162 Reachtime:2024-12-29 12:25 TravelTime:23分 Surplus:10站} {Lines:965路 Price:2~5元 EndSn:京原东站 BusId:7201025260 Reachtime:2024-12-29 12:09 TravelTime:7分 Surplus:2站} {Lines:965路 Price:2~5元 EndSn:京原东站 BusId:7201025264 Reachtime:2024-12-29 12:33 TravelTime:31分 Surplus:14站} {Lines:992路 Price:2~4元 EndSn:京原东站 BusId:4775780589 Reachtime:2024-12-29 12:03 TravelTime:1分 Surplus:即将到站} {Lines:992路 Price:2~4元 EndSn:京原东站 BusId:7003005667 Reachtime:2024-12-29 12:20 TravelTime:18分 Surplus:7站} {Lines:专215路 Price:2元 EndSn:老山公交场站 BusId:7009025868 Reachtime:2024-12-29 12:03 TravelTime:1分 Surplus:即将到站} {Lines:专215路 Price:2元 EndSn:老山公交场站 BusId:7009035368 Reachtime:2024-12-29 12:14 TravelTime:11分 Surplus:即将到站} {Lines:专46路 Price:2元 EndSn:京原路口东 BusId:4178710786 Reachtime:2024-12-29 12:09 TravelTime:6分 Surplus:3站} {Lines:专91路 Price:2元 EndSn:京原东站 BusId:7008005366 Reachtime:2024-12-29 12:10 TravelTime:8分 Surplus:3站}]}
	*/
	var ret []Pretty
	station := rep.Line
	for _, data := range rep.Data {
		var p Pretty
		p.Line = data.Lines
		p.EndSn = data.EndSn
		p.ReachTime = data.Reachtime
		p.TravelTime = data.TravelTime
		p.Surplus = data.Surplus
		log.Printf("站点名:%+v\t线路名%v\t终点站%v\t预计到达时间%v\t预计剩余时间%v\t剩余站数%v\n", station, p.Line, p.EndSn, p.ReachTime, p.Surplus, p.ReachTime)
		/*
			站点名:银河大街	线路名597路	终点站鲁谷公交场站	预计到达时间2024-12-29 12:34	预计剩余时间5站	剩余站数2024-12-29 12:34
		*/
		ret = append(ret, p)
	}
	return ret
}

type Pretty struct {
	Line       string // 线路名
	EndSn      string // 终点站
	ReachTime  string // 到达时间
	TravelTime string // 剩余到达时间
	Surplus    string // 剩余站数
}
