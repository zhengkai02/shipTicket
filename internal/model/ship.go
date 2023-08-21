package model

type Ship struct {
	LineNo   int    `json:"LineNo"`
	LineName string `json:"LineName"`
	ShipName string `json:"ShipName"`
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
