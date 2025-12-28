// Package conf implements the Conf type and methods.
package conf

// Conf is a container of parsed configuration data.
type Conf struct {
	Addr    string `json:"addr"`
	Dire    string `json:"dire"`
	BodyMax int64  `json:"body_max"`
	PairTTL int64  `json:"pair_ttl"`
	RateMax int64  `json:"rate_max"`
	UserMax int64  `json:"user_max"`
}
