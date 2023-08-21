package admin

type TicketReq struct {
	StartPortNo string `json:"startPortNo"`
	EndPortNo   string `json:"endPortNo"`
	StartDate   string `json:"startDate"`
}

type TicketResp struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []*TicketData `json:"data"`
}

type TicketData struct {
	ID               int            `json:"id"`
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

type SeatClasses struct {
	ClassNum            int         `json:"classNum"`
	ClassName           string      `json:"className"`
	LocalCurrentCount   int         `json:"localCurrentCount"`
	PubCurrentCount     int         `json:"pubCurrentCount"`
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
}

type OrderDetailResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		OrderId          string      `json:"orderId"`
		Edition          int         `json:"edition"`
		UserId           int         `json:"userId"`
		AccountTypeId    int         `json:"accountTypeId"`
		MchId            string      `json:"mchId"`
		ContactNum       string      `json:"contactNum"`
		Dwh              int         `json:"dwh"`
		LineNum          int         `json:"lineNum"`
		LineName         string      `json:"lineName"`
		LineNo           int         `json:"lineNo"`
		Sx               int         `json:"sx"`
		ShipName         string      `json:"shipName"`
		StartPortNo      int         `json:"startPortNo"`
		StartPortName    string      `json:"startPortName"`
		EndPortNo        int         `json:"endPortNo"`
		EndPortName      string      `json:"endPortName"`
		SailDate         string      `json:"sailDate"`
		SailTime         string      `json:"sailTime"`
		LineDirect       int         `json:"lineDirect"`
		BuyTicketType    int         `json:"buyTicketType"`
		TotalFee         float64     `json:"totalFee"`
		TotalPayFee      float64     `json:"totalPayFee"`
		ChannelType      string      `json:"channelType"`
		PayChannel       string      `json:"payChannel"`
		OrderState       int         `json:"orderState"`
		Yzm              string      `json:"yzm"`
		CreateTime       string      `json:"createTime"`
		ExpireTime       string      `json:"expireTime"`
		PayTime          string      `json:"payTime"`
		CancelTime       interface{} `json:"cancelTime"`
		AlterTime        interface{} `json:"alterTime"`
		RefundTime       interface{} `json:"refundTime"`
		UpdateTime       string      `json:"updateTime"`
		BusStartTime     string      `json:"busStartTime"`
		Clxh             int         `json:"clxh"`
		Clxm             string      `json:"clxm"`
		Hxlxh            int         `json:"hxlxh"`
		Hxlxm            string      `json:"hxlxm"`
		Bus              int         `json:"bus"`
		Bus2             int         `json:"bus2"`
		ExpireTimeDiif   interface{} `json:"expireTimeDiif"`
		IoTotalRefund    interface{} `json:"ioTotalRefund"`
		EmbarkPortName   string      `json:"embarkPortName"`
		GroupReplaceTime int         `json:"groupReplaceTime"`
		OrderItemList    []struct {
			Id                 int         `json:"id"`
			OrderId            string      `json:"orderId"`
			Edition            int         `json:"edition"`
			PassName           string      `json:"passName"`
			CredentialType     int         `json:"credentialType"`
			CredentialNum      string      `json:"credentialNum"`
			PhoneNum           interface{} `json:"phoneNum"`
			LineNum            int         `json:"lineNum"`
			LineName           string      `json:"lineName"`
			LineNo             int         `json:"lineNo"`
			ShipName           string      `json:"shipName"`
			SailDate           string      `json:"sailDate"`
			SailTime           string      `json:"sailTime"`
			BusStartTime       string      `json:"busStartTime"`
			SeatClass          int         `json:"seatClass"`
			SeatClassName      string      `json:"seatClassName"`
			IoLocal            int         `json:"ioLocal"`
			PriceType          int         `json:"priceType"`
			TicketFee          float64     `json:"ticketFee"`
			DiscountFee        interface{} `json:"discountFee"`
			RealFee            float64     `json:"realFee"`
			RefundCharge       interface{} `json:"refundCharge"`
			RefundFee          interface{} `json:"refundFee"`
			PlateNum           interface{} `json:"plateNum"`
			ItemState          int         `json:"itemState"`
			CreateTime         string      `json:"createTime"`
			AlterTime          interface{} `json:"alterTime"`
			RefundTime         interface{} `json:"refundTime"`
			UpdateTime         string      `json:"updateTime"`
			RefundNum          interface{} `json:"refundNum"`
			RefundState        interface{} `json:"refundState"`
			ObtainTicketNum    string      `json:"obtainTicketNum"`
			OldObtainTicketNum interface{} `json:"oldObtainTicketNum"`
			SeatNumber         string      `json:"seatNumber"`
			OldSeatNumber      interface{} `json:"oldSeatNumber"`
			RefundScale        interface{} `json:"refundScale"`
			RealRefundCharge   interface{} `json:"realRefundCharge"`
			RealRefundFee      interface{} `json:"realRefundFee"`
			RealRefundScale    interface{} `json:"realRefundScale"`
			RealRefundTime     interface{} `json:"realRefundTime"`
			Sx                 int         `json:"sx"`
			Xzh                int         `json:"xzh"`
			Clxh               int         `json:"clxh"`
			Clxm               string      `json:"clxm"`
			Hxlxh              int         `json:"hxlxh"`
			Hxlxm              string      `json:"hxlxm"`
			Bus                int         `json:"bus"`
			Bus2               int         `json:"bus2"`
			FreeChildCount     int         `json:"freeChildCount"`
			RefundSource       interface{} `json:"refundSource"`
			Status             int         `json:"status"`
		} `json:"orderItemList"`
		IoAssignDefaultSeat  int         `json:"ioAssignDefaultSeat"`
		WgroupReplaceTime    interface{} `json:"wgroupReplaceTime"`
		WgroupReplacePercent interface{} `json:"wgroupReplacePercent"`
	} `json:"data"`
}
