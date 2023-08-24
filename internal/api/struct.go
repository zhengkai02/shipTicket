package api

/**
*@Auther kaikai.zheng
*@Date 2023-08-22 10:27:29
*@Name struct
*@Desc // 结构体定义
**/

type LoginResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	AccountTypeId int    `json:"accountTypeId"`
	UserId        int    `json:"userId"`
	Token         string `json:"token"`
}

type UserInfoResp struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    *UserData `json:"data"`
}

type UserData struct {
	Id            int         `json:"id"`
	AccountTypeId int         `json:"accountTypeId"`
	UserName      interface{} `json:"userName"`
	NickName      string      `json:"nickName"`
	PhoneNum      string      `json:"phoneNum"`
	OpenId        string      `json:"openId"`
	HeadimgUrl    string      `json:"headimgUrl"`
	CreateTime    string      `json:"createTime"`
	IoPush        int         `json:"ioPush"`
}
type TicketReq struct {
	StartPortNo int64  `json:"startPortNo"`
	EndPortNo   int64  `json:"endPortNo"`
	StartDate   string `json:"startDate"`
}
type TicketResp struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []*TicketData `json:"data"`
}

type TicketData struct {
	Dwh              int            `json:"dwh"`
	LineNum          int            `json:"lineNum"`
	LineName         string         `json:"lineName"`
	LineNo           int            `json:"lineNo"`
	ShipNo           interface{}    `json:"shipNo"`
	ShipName         string         `json:"shipName"`
	StartPortNo      int            `json:"startPortNo"`
	StartPortName    string         `json:"startPortName"`
	EndPortNo        int            `json:"endPortNo"`
	EndPortName      string         `json:"endPortName"`
	IoSelectSeat     int            `json:"ioSelectSeat"`
	SailDate         string         `json:"sailDate"`
	SailTime         string         `json:"sailTime"`
	BusStartTime     string         `json:"busStartTime"`
	Sx               int            `json:"sx"`
	LineDirect       int            `json:"lineDirect"`
	SaleBeginTime    string         `json:"saleBeginTime"`
	SaleEndTime      string         `json:"saleEndTime"`
	StopSaleTime     int            `json:"stopSaleTime"`
	CarStopSaleTime  interface{}    `json:"carStopSaleTime"`
	OnSale           bool           `json:"onSale"`
	OffSaleMsg       interface{}    `json:"offSaleMsg"`
	SeatClasses      []*SeatClasses `json:"seatClasses"`
	DriverSeatClass  []*DriverClass `json:"driverSeatClass"`
	BuyTicketType    int            `json:"buyTicketType"`
	Clxh             int            `json:"clxh"`
	Clxm             string         `json:"clxm"`
	Hxlxh            int            `json:"hxlxh"`
	Hxlxm            string         `json:"hxlxm"`
	Bus              int            `json:"bus"`
	Bus2             int            `json:"bus2"`
	LineState        int            `json:"lineState"`
	LineStateName    string         `json:"lineStateName"`
	EmbarkPortName   string         `json:"embarkPortName"`
	FreeChildCount   int            `json:"freeChildCount"`
	CandidateTimeEnd string         `json:"candidateTimeEnd"`
	IsVisible        interface{}    `json:"isVisible"`
	PortMemos        string         `json:"portMemos"`
	TimeMemos        string         `json:"timeMemos"`
}

type FerryTicketResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Dwh             int         `json:"dwh"`
		LineNum         int         `json:"lineNum"`
		LineName        string      `json:"lineName"`
		LineNo          int         `json:"lineNo"`
		ShipNo          interface{} `json:"shipNo"`
		ShipName        string      `json:"shipName"`
		StartPortNo     int         `json:"startPortNo"`
		StartPortName   string      `json:"startPortName"`
		EndPortNo       int         `json:"endPortNo"`
		EndPortName     string      `json:"endPortName"`
		IoSelectSeat    interface{} `json:"ioSelectSeat"`
		SailDate        string      `json:"sailDate"`
		SailTime        string      `json:"sailTime"`
		BusStartTime    string      `json:"busStartTime"`
		Sx              int         `json:"sx"`
		LineDirect      int         `json:"lineDirect"`
		SaleBeginTime   string      `json:"saleBeginTime"`
		SaleEndTime     string      `json:"saleEndTime"`
		StopSaleTime    int         `json:"stopSaleTime"`
		CarStopSaleTime interface{} `json:"carStopSaleTime"`
		OnSale          bool        `json:"onSale"`
		OffSaleMsg      *string     `json:"offSaleMsg"`
		SeatClasses     []struct {
			ClassNum            int         `json:"classNum"`
			ClassName           string      `json:"className"`
			LocalCurrentCount   int         `json:"localCurrentCount"`
			PubCurrentCount     int         `json:"pubCurrentCount"`
			TotalCount          int         `json:"totalCount"`
			FerryPassTotalCount int         `json:"ferryPassTotalCount"`
			OriginPrice         float64     `json:"originPrice"`
			TotalPrice          float64     `json:"totalPrice"`
			HalfPrice           float64     `json:"halfPrice"`
			LocalPrice          float64     `json:"localPrice"`
			LocalHalfPrice      float64     `json:"localHalfPrice"`
			SeatState           int         `json:"seatState"`
			SeatStateName       string      `json:"seatStateName"`
			TotalOriginCount    interface{} `json:"totalOriginCount"`
			CandidateCount      int         `json:"candidateCount"`
			Floor               interface{} `json:"floor"`
			FloorSortNo         interface{} `json:"floorSortNo"`
			Xlen                interface{} `json:"xlen"`
			Ylen                interface{} `json:"ylen"`
			EndSeatPortNo       interface{} `json:"endSeatPortNo"`
		} `json:"seatClasses"`
		DriverSeatClass  []*DriverClass `json:"driverSeatClass"`
		BuyTicketType    int            `json:"buyTicketType"`
		Clxh             int            `json:"clxh"`
		Clxm             string         `json:"clxm"`
		Hxlxh            int            `json:"hxlxh"`
		Hxlxm            string         `json:"hxlxm"`
		Bus              int            `json:"bus"`
		Bus2             int            `json:"bus2"`
		LineState        int            `json:"lineState"`
		LineStateName    string         `json:"lineStateName"`
		EmbarkPortName   interface{}    `json:"embarkPortName"`
		FreeChildCount   int            `json:"freeChildCount"`
		CandidateTimeEnd interface{}    `json:"candidateTimeEnd"`
		IsVisible        interface{}    `json:"isVisible"`
		PortMemos        string         `json:"portMemos"`
		TimeMemos        string         `json:"timeMemos"`
	} `json:"data"`
}

type SeatClasses struct {
	ClassNum            int         `json:"classNum"`
	ClassName           string      `json:"className"`
	LocalCurrentCount   int         `json:"localCurrentCount"`
	PubCurrentCount     int64       `json:"pubCurrentCount"`
	TotalCount          int         `json:"totalCount"`
	FerryPassTotalCount interface{} `json:"ferryPassTotalCount"`
	OriginPrice         float64     `json:"originPrice"`
	TotalPrice          float64     `json:"totalPrice"`
	HalfPrice           float64     `json:"halfPrice"`
	LocalPrice          float64     `json:"localPrice"`
	LocalHalfPrice      float64     `json:"localHalfPrice"`
	SeatState           int         `json:"seatState"`
	SeatStateName       string      `json:"seatStateName"`
	TotalOriginCount    int         `json:"totalOriginCount"`
	CandidateCount      int         `json:"candidateCount"`
	Floor               int         `json:"floor"`
	FloorSortNo         int         `json:"floorSortNo"`
	Xlen                int         `json:"xlen"`
	Ylen                int         `json:"ylen"`
	EndSeatPortNo       interface{} `json:"endSeatPortNo"`
}

type DriverClass struct {
	ClassNum            int         `json:"classNum"`
	ClassName           string      `json:"className"`
	LocalCurrentCount   int         `json:"localCurrentCount"`
	PubCurrentCount     int64       `json:"pubCurrentCount"`
	TotalCount          int         `json:"totalCount"`
	FerryPassTotalCount int         `json:"ferryPassTotalCount"`
	OriginPrice         float64     `json:"originPrice"`
	TotalPrice          float64     `json:"totalPrice"`
	HalfPrice           float64     `json:"halfPrice"`
	LocalPrice          float64     `json:"localPrice"`
	LocalHalfPrice      float64     `json:"localHalfPrice"`
	SeatState           int         `json:"seatState"`
	SeatStateName       string      `json:"seatStateName"`
	TotalOriginCount    interface{} `json:"totalOriginCount"`
	CandidateCount      int         `json:"candidateCount"`
	Floor               interface{} `json:"floor"`
	FloorSortNo         interface{} `json:"floorSortNo"`
	Xlen                interface{} `json:"xlen"`
	Ylen                interface{} `json:"ylen"`
	EndSeatPortNo       interface{} `json:"endSeatPortNo"`
}

type SaveReq struct {
	UserId              int                  `json:"userId"`
	BuyTicketType       int                  `json:"buyTicketType"`
	ContactNum          string               `json:"contactNum"`
	LineNum             int                  `json:"lineNum"`
	LineName            string               `json:"lineName"`
	LineNo              int                  `json:"lineNo"`
	ShipName            string               `json:"shipName"`
	StartPortNo         int                  `json:"startPortNo"`
	StartPortName       string               `json:"startPortName"`
	EndPortNo           int                  `json:"endPortNo"`
	EndPortName         string               `json:"endPortName"`
	SailDate            string               `json:"sailDate"`
	SailTime            string               `json:"sailTime"`
	LineDirect          int                  `json:"lineDirect"`
	TotalFee            float64              `json:"totalFee"`
	TotalPayFee         float64              `json:"totalPayFee"`
	Sx                  int                  `json:"sx"`
	OrderItemRequests   []*OrderItemRequests `json:"orderItemRequests"`
	BusStartTime        string               `json:"busStartTime"`
	Clxm                string               `json:"clxm"`
	Clxh                int                  `json:"clxh"`
	Hxlxh               int                  `json:"hxlxh"`
	Hxlxm               string               `json:"hxlxm"`
	Bus                 int                  `json:"bus"`
	Bus2                int                  `json:"bus2"`
	Dwh                 int                  `json:"dwh"`
	IoAssignDefaultSeat int                  `json:"ioAssignDefaultSeat"`
}

type SaveResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderId string `json:"orderId"`
	} `json:"data"`
}

type OrderItemRequests struct {
	PassName       string  `json:"passName"`
	PlateNum       string  `json:"plateNum"`
	CredentialType int     `json:"credentialType"`
	PassId         int     `json:"passId"`
	SeatClassName  string  `json:"seatClassName"`
	SeatClass      int     `json:"seatClass"`
	TicketFee      float64 `json:"ticketFee"`
	RealFee        float64 `json:"realFee"`
	FreeChildCount int     `json:"freeChildCount"`
	PassType       string  `json:"passType"`
	SeatNumber     int     `json:"seatNumber"`
}

type CheckToeknResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PassengersResp struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []*Passenger `json:"data"`
}

type Passenger struct {
	Id               int64       `json:"id"`
	UserId           int64       `json:"userId"`
	PassName         string      `json:"passName"`
	CredentialTypeId int         `json:"credentialTypeId"`
	CredentialNum    string      `json:"credentialNum"`
	PhoneNum         interface{} `json:"phoneNum"`
	PlateNum         string      `json:"plateNum"`
	PassType         int         `json:"passType"`
	IoDefault        int         `json:"ioDefault"`
	IoLocal          int         `json:"ioLocal"`
	IoDriver         interface{} `json:"ioDriver"`
	PriceType        int         `json:"priceType"`
	Bz               int         `json:"bz"`
	CreateTime       string      `json:"createTime"`
	UpdateTime       interface{} `json:"updateTime"`
	Xzh              int         `json:"xzh"`
}

type VehicleResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []*Vehicle `json:"data"`
}

type Vehicle struct {
	Id         int         `json:"id"`
	UserId     int         `json:"userId"`
	PlateNum   string      `json:"plateNum"`
	Bz         int         `json:"bz"`
	IoLocal    int         `json:"ioLocal"`
	CreateTime string      `json:"createTime"`
	UpdateTime interface{} `json:"updateTime"`
}
