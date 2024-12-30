package constant

type Rep struct {
	Code     int    `json:"code"`      // http code 如 200
	City     string `json:"city"`      // 城市名
	Line     string `json:"line"`      // 站点名
	CarCount int    `json:"car_count"` // 此时存在的线路数
	Data     Data   `json:"data"`
}
type Data []struct {
	Lines      string `json:"lines"`      // 598路
	Price      string `json:"price"`      // 2~3元
	EndSn      string `json:"endSn"`      // 终点站 刘娘府北街北口
	BusId      string `json:"busId"`      // 车辆id 7302019660
	Reachtime  string `json:"reachtime"`  // 预计到站时间
	TravelTime string `json:"travelTime"` // 预计剩余时间
	Surplus    string `json:"surplus"`    // 剩余站数
}

const (
	BAJIAOROADEAST              = "八角路东口"
	GALAXYSTREET                = "银河大街"
	JINGYUANROADEAST            = "京原路口东"
	LAOSHAN                     = "老山"
	NATIONALBOTANIXALGARDEN     = "国家植物园"
	SHIJINGSHANSCULPTUREPARK    = "石景山雕塑公园"
	SHIJINGSHANTECHNOLOGYMUSEUM = "石景山科技馆"
	TIMESGARDENNORTH            = "时代花园北路"
	JINGYANHOTEL                = "京燕饭店"
	BAJIAO                      = "八角"
)
