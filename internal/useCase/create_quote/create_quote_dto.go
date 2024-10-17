package createquote

type CreateQuoteInputDTO struct {
	Recipient Recipient `json:"recipient"`
	Volumes   []Volume  `json:"volumes"`
}

type Recipient struct {
	Address Address `json:"address"`
}

type Address struct {
	Zipcode string `json:"zipcode"`
}

type Volume struct {
	Category      int     `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight float64 `json:"unitary_weight"`
	Price         float64 `json:"price"`
	SKU           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

type CreateQuoteOutputDTO struct {
	Carriers   []Carrier `json:"carrier"`
	NoCarriers bool      `json:"no_carriers"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}

type FreightFastOutputDTO struct {
	Dispatchers []Dispatcher `json:"dispatchers"`
}

type Dispatcher struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	Carrier      Carrier      `json:"carrier"`
	Service      string       `json:"service"`
	DeliveryTime DeliveryTime `json:"delivery_time"`
	FinalPrice   float64      `json:"final_price"`
}

type DeliveryTime struct {
	Days int `json:"days"`
}
