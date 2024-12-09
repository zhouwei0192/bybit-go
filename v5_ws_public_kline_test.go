package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebsocketV5Public_Kline(t *testing.T) {
	//respBody := map[string]interface{}{
	//	"topic": "kline.5.BTCUSDT",
	//	"type":  "snapshot",
	//	"ts":    1672324988882,
	//	"data": []map[string]interface{}{
	//		{
	//			"start":     1672324800000,
	//			"end":       1672325099999,
	//			"interval":  "5",
	//			"open":      "16649.5",
	//			"close":     "16677",
	//			"high":      "16677",
	//			"low":       "16608",
	//			"volume":    "2.081",
	//			"turnover":  "34666.4005",
	//			"confirm":   false,
	//			"timestamp": 1672324988882,
	//		},
	//	},
	//}
	//bytesBody, err := json.Marshal(respBody)
	//require.NoError(t, err)
	//
	//category := CategoryV5Linear
	//
	//server, teardown := testhelper.NewWebsocketServer(
	//	testhelper.WithWebsocketHandlerOption(V5WebsocketPublicPathFor(category), bytesBody),
	//)
	//defer teardown()
	var proxyURL, _ = url.Parse("http://127.0.0.1:7890")
	dialer := websocket.DefaultDialer
	dialer.Proxy = http.ProxyURL(proxyURL)
	//"wss://stream.bybit.com/v5/public/linear"
	//"wss://stream.bybit.com/v5/public/spot/v5/public/spot"
	wsClient := NewWebsocketClient().WithBaseURL("wss://stream.bybit.com").WithDialer(dialer)
	svc, err := wsClient.V5().Public(CategoryV5Spot)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//wsClient := NewTestWebsocketClient().
	//	WithBaseURL(server.URL)
	//
	//svc, err := wsClient.V5().Public(category)
	//require.NoError(t, err)

	{
		params := make([]V5WebsocketPublicKlineParamKey, 0)
		params = append(params, V5WebsocketPublicKlineParamKey{Interval: Interval1, Symbol: SymbolV5BTCUSDT})
		_, err := svc.SubscribeKline(
			params,
			func(response V5WebsocketPublicKlineResponse) error {
				marshal, _ := json.Marshal(response)
				fmt.Println(string(marshal))
				return nil
			},
		)
		require.NoError(t, err)
	}
	errHandler := func(isWebsocketClosed bool, err error) {
		fmt.Println("error: ", err.Error())
	}
	svc.Start(context.Background(), errHandler)
	//svc.Ping()
	select {}
	//assert.NoError(t, svc.Run())
	//assert.NoError(t, svc.Ping())
	//assert.NoError(t, svc.Close())
}
