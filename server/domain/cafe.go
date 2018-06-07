package domain

import "time"

type timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// TODO: エリアとかもいれる
type Cafe struct {
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Address       string        `json:"address"`
	Comment       string        `json:"comment"`
	Visitors      int           `json:"visitors"`
	Price         Price         `json:"price"`
	BusinessHours BusinessHours `json:"business_hours"`
	Facility      Facility      `json:"facility"`
	IsPublic      bool          `json:"is_public"`
	IsClosed      bool          `json:"is_closed"`
	timestamps
}

type Price struct {
	Low  int `json:"low"`
	High int `json:"high"`
}

type BusinessHours struct {
	Open  time.Time `json:"open"`
	Close time.Time `json:"close"`
}

type Facility struct {
	Socket    string  `json:"socket"`
	WifiSpeed float64 `json:"wifi_speed"`
}

// NewCafe returns new cafe object
// TODO 全プロパティに対応する
func NewCafe(name string, lowPrice, highPrice int) *Cafe {
	c := &Cafe{
		Name: name,
		Price: Price{
			Low:  lowPrice,
			High: highPrice,
		},
		BusinessHours: BusinessHours{
			Open: time.Now(),
			Close: time.Now().Add(3 * time.Hour),
		},
	}

	if isNil(&c.CreatedAt) {
		c.CreatedAt = time.Now()
	}

	c.UpdatedAt = time.Now()

	return c
}

// isNil judge timestamp is initialized
func isNil(timestamp *time.Time) bool {
	return timestamp.Unix() == time.Time{}.Unix()
}
