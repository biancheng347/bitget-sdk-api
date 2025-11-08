package main

import (
	"fmt"
	"github.com/biancheng347/bitget-sdk-api/internal/model"
	v2 "github.com/biancheng347/bitget-sdk-api/pkg/client/v2"
	"github.com/biancheng347/bitget-sdk-api/pkg/client/ws"
)

func main() {
	//spSymbols()
	fuSymbols()
}

func spSymbols() {
	client := new(v2.MixMarketClient).Init()
	rsp, err := client.Contracts(map[string]string{
		"productType": "usdt-futures",
	})
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	fmt.Println(rsp)
}

func fuSymbols() {
	client := new(ws.BitgetWsClient).Init(false, func(message string) {
		fmt.Println("message: " + message)

	}, func(message string) {
		fmt.Println("error:" + message)
	})

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "USDT-FUTURES",
		Channel:  "candle1m",
		InstId:   "BTCUSDT",
	}
	channelsDef = append(channelsDef, subReqDef1)
	client.SubscribeDef(channelsDef)

	//var channels []model.SubscribeReq
	//subReq1 := model.SubscribeReq{
	//	InstType: "USDT-FUTURES",
	//	Channel:  "candle5m",
	//	InstId:   "BTCUSDT",
	//}
	//channels = append(channels, subReq1)
	//client.Subscribe(channels, func(message string) {
	//	fmt.Println("appoint:" + message)
	//})
	<-client.WsClient().ExitCh
	fmt.Println("end")
}
