package config

import "github.com/biancheng347/bitget-sdk-api/constants"

const (
	BaseUrl = "https://api.bitget.com"
	//WsUrl   = "wss://ws.bitget.com/mix/v1/stream"
	WsUrl = "wss://ws.bitget.com/v2/ws/public"

	ApiKey        = ""
	SecretKey     = ""
	PASSPHRASE    = ""
	TimeoutSecond = 30
	SignType      = constants.SHA256
)
