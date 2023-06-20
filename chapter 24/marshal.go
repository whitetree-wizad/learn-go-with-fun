package main

import (
	"encoding/json"
	"fmt"
)

type coordinate struct {
	D float64
	M float64
	S float64
	H rune
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Decimal    float64 `json:"decimal"`
		Dms        string  `json:"dms"`
		Degrees    float64 `json:"degrees"`
		Minutes    float64 `json:"minutes"`
		Seconds    float64 `json:"seconds"`
		Hemisphere rune    `json:"hemisphere"`
	}{
		c.decimal(),
		c.String(),
		c.D,
		c.M,
		c.S,
		c.H,
	})
}

type location struct {
	lat, long coordinate
}

func (c coordinate) String() string {
	return fmt.Sprintf(`:%vo%v'%0.1f" %c`, c.D, c.M, c.S, c.H)
}

func (l location) String() string {
	return fmt.Sprintf("%v %v", l.lat, l.long)
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.H {
	case 'S', 'W':
		sign = -1
	}
	return sign * (c.D + c.M/60 + c.S/3600)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}

	lat, _ := json.Marshal(elysium.lat)
	long, _ := json.Marshal(elysium.long)

	fmt.Println(string(lat), string(long))

}
