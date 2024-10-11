package createquote

type CreateQuoteInputDTO struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	SimulationType []int        `json:"simulation_type"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Type    int    `json:"type"`
	Country string `json:"country"`
	Zipcode int    `json:"zipcode"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int      `json:"zipcode"`
	TotalPrice       float64  `json:"total_price"`
	Volumes          []Volume `json:"volumes"`
}

type Volume struct {
	Amount        int     `json:"amount"`
	AmountVolumes int     `json:"amount_volumes"`
	Category      string  `json:"category"`
	SKU           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
}

type CreateQuoteOutputDTO struct {
	Carriers []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}

type FreightFastOutputDTO struct {
	Dispatchers []struct {
		Offers []Offer `json:"offers"`
	} `json:"dispatchers"`
}

type Offer struct {
	Carrier struct {
		Name string `json:"name"`
	} `json:"carrier"`
	Service      string `json:"service"`
	DeliveryTime struct {
		Days int `json:"days"`
	} `json:"delivery_time"`
	FinalPrice float64 `json:"final_price"`
}
