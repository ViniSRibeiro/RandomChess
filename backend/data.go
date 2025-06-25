package main

type ticker struct {
	Last    float64 `json:"last"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Volume  float64 `json:"vol"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Updated int64   `json:"updated"`
}

type user struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}
