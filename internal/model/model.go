package model

// CustomerDB - customer substructure representation in database
type CustomerDB struct {
	CustomerID string `yaml:"customer_id"`
	Name       string `yaml:"name"`
	Phone      string `yaml:"phone"`
	Zip        string `yaml:"zip"`
	City       string `yaml:"city"`
	Address    string `yaml:"address"`
	Region     string `yaml:"region"`
	Email      string `yaml:"email"`
}

// ItemDB - item substructure representation in database
type ItemDB struct {
	ChrtID int     `yaml:"chrt_id"`
	Price  float64 `yaml:"price"`
	Rid    string  `yaml:"rid"`
	Name   string  `yaml:"name"`
	Sale   int     `yaml:"sale"`
	Size   string  `yaml:"size"`
	NmId   int     `yaml:"nm_id"`
	Brand  string  `yaml:"brand"`
}

// PaymentDB - payment substructure representation in database
type PaymentDB struct {
	Transaction  string  `yaml:"transaction"`
	RequestID    string  `yaml:"request_id"`
	Currency     string  `yaml:"currency"`
	Provider     string  `yaml:"provider"`
	PaymentDt    int     `yaml:"payment_dt"`
	Bank         string  `yaml:"bank"`
	DeliveryCost float64 `yaml:"delivery_cost"`
	GoodsTotal   float64 `yaml:"goods_total"`
	CustomFee    float64 `yaml:"custom_fee"`
}

// OrderDB - order structure representation in database
type OrderDB struct {
	OrderUID          string `yaml:"order_uid"`
	TrackNumber       string `yaml:"track_number"`
	Entry             string `yaml:"entry"`
	CustomerID        string `yaml:"customer_id"`
	Locale            string `yaml:"locale"`
	DeliveryService   string `yaml:"delivery_service"`
	ShardKey          string `yaml:"shardkey"`
	SmID              int    `yaml:"sm_id"`
	DateCreated       string `yaml:"date_created"`
	OofShard          string `yaml:"oof_shard"`
	InternalSignature string `yaml:"internal_signature"`
}

// OrderItemsDB - order representation of items in database
type OrderItemsDB struct {
	OrderID string `yaml:"order_id"`
	ItemID  int    `yaml:"item_id"`
	Amount  int    `yaml:"amount"`
	Status  int    `yaml:"status"`
}

// Delivery - substructure
type Delivery struct {
	Name    string `json:"name" yaml:"name"`
	Phone   string `json:"phone" yaml:"phone"`
	Zip     string `json:"zip" yaml:"zip"`
	City    string `json:"city" yaml:"city"`
	Address string `json:"address" yaml:"address"`
	Region  string `json:"region" yaml:"region"`
	Email   string `json:"email" yaml:"email"`
}

// Payment - substructure
type Payment struct {
	Transaction  string  `json:"transaction" yaml:"transaction"`
	RequestID    string  `json:"request_id" yaml:"request_id"`
	Currency     string  `json:"currency" yaml:"currency"`
	Provider     string  `json:"provider" yaml:"provider"`
	Amount       float64 `json:"amount" yaml:"amount"`
	PaymentDt    int     `json:"payment_dt" yaml:"payment_dt"`
	Bank         string  `json:"bank" yaml:"bank"`
	DeliveryCost float64 `json:"delivery_cost" yaml:"delivery_cost"`
	GoodsTotal   float64 `json:"goods_total" yaml:"goods_total"`
	CustomFee    float64 `json:"custom_fee" yaml:"custom_fee"`
}

// Items - substructure
type Items struct {
	ChrtID      int     `json:"chrt_id" yaml:"chrt_id"`
	TrackNumber string  `json:"track_number" yaml:"track_number"`
	Price       float64 `json:"price" yaml:"price"`
	Rid         string  `json:"rid" yaml:"rid"`
	Name        string  `json:"name" yaml:"name"`
	Sale        int     `json:"sale" yaml:"sale"`
	Size        string  `json:"size" yaml:"size"`
	TotalPrice  float64 `json:"total_price" yaml:"total_price"`
	NmId        int     `json:"nm_id" yaml:"nm_id"`
	Brand       string  `json:"brand" yaml:"brand"`
	Status      int     `json:"status" yaml:"status"`
}

// Data - basic structure
type Data struct {
	OrderUID          string   `json:"order_uid" yaml:"order_uid"`
	TrackNumber       string   `json:"track_number" yaml:"track_number"`
	Entry             string   `json:"entry" yaml:"entry"`
	Delivery          Delivery `json:"delivery" yaml:"delivery"`
	Payment           Payment  `json:"payment" yaml:"payment"`
	Items             []Items  `json:"items" yaml:"items"`
	Locale            string   `json:"locale" yaml:"locale"`
	InternalSignature string   `json:"internal_signature" yaml:"internal_signature"`
	CustomerID        string   `json:"customer_id" yaml:"customer_id"`
	DeliveryService   string   `json:"delivery_service" yaml:"delivery_service"`
	ShardKey          string   `json:"shardkey" yaml:"shardkey"`
	SmID              int      `json:"sm_id" yaml:"sm_id"`
	DateCreated       string   `json:"date_created" yaml:"date_created"`
	OofShard          string   `json:"oof_shard" yaml:"oof_shard"`
}