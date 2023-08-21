package model

type Order struct {
	OrderId  int    `json:"orderId"`
	LineName string `json:"lineName"`
	ShipName string `json:"shipName"`
}

//func (s *Ship) Options() (options []*selectfield.Option, Error error) {
//	getList := []Ship{
//		{
//			LineNo:            "123",
//			LineName:    "T101",
//			ShipName: "2023-09-01 10:00:00",
//		},
//	}
//	for _, v := range getList {
//		option := &selectfield.Option{
//			Label: v.LineName,
//			Value: v.ShipName,
//		}
//		options = append(options, option)
//	}
//	return options, nil
//}
