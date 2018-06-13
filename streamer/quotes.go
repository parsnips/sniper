package streamer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type HandshakeRequest struct {
	Advice                   AdviceType `json:"advice"`
	Channel                  string     `json:"channel"`
	Ext                      ExtType    `json:"ext"`
	ID                       string     `json:"id"`
	MinimumVersion           string     `json:"minimumVersion"`
	SupportedConnectionTypes []string   `json:"supportedConnectionTypes"`
	Version                  string     `json:"version"`
}

type AdviceType struct {
	Interval int `json:"interval"`
	Timeout  int `json:"timeout"`
}

type ExtType struct {
	StreamerToken string `json:"com.devexperts.auth.AuthToken"`
}

type HandshakeResponse struct {
	ClientID   string `json:"clientId"`
	Successful bool   `json:"successful"`
	Channel    string `json:"channel"`
}

type ResetRequest struct {
	Channel  string `json:"channel"`
	ClientID string `json:"clientId"`
	Data     struct {
		Reset bool `json:"reset"`
	} `json:"data"`
	ID string `json:"id"`
}

type WebsocketConnectRequest struct {
	Advice         AdviceType `json:"advice"`
	Channel        string     `json:"channel"`
	ClientID       string     `json:"clientId"`
	ConnectionType string     `json:"connectionType"`
	ID             string     `json:"id"`
}

type WebsocketConnectRequest2 struct {
	Channel        string `json:"channel"`
	ClientID       string `json:"clientId"`
	ConnectionType string `json:"connectionType"`
	ID             string `json:"id"`
}

func quoteRequest(clientID string, symbols []string) []byte {
	req := []QuoteRequest{
		QuoteRequest{
			Channel:  "/service/sub",
			ClientID: clientID,
			Data: DataType{
				Add: AddType{
					Summary: symbols,
					Quote:   symbols,
					Greeks:  symbols,
				},
			},
			ID: "100",
		},
	}

	bytes, _ := json.Marshal(req)
	return bytes
}

type QuoteRequest struct {
	Channel  string   `json:"channel"`
	ClientID string   `json:"clientId"`
	Data     DataType `json:"data"`
	ID       string   `json:"id"`
}

type DataType struct {
	Add AddType `json:"add"`
}

type AddType struct {
	Summary []string `json:"Summary"`
	Quote   []string `json:"Quote"`
	Greeks  []string `json:"Greeks"`
}

var clientID string

func makeRequest(obj interface{}) []byte {
	bytes, err := json.Marshal([]interface{}{obj})
	if err != nil {
		log.Fatal("makeRequest:", err)
	}
	return bytes
}

func TwQuotes(streamerToken string, symbols []string) {
	u := url.URL{Scheme: "wss", Host: "tasty.dxfeed.com", Path: "/live/cometd"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	handler := func(message []byte) {
		var thunk []map[string]interface{}

		json.Unmarshal(message, &thunk)

		item := thunk[0]
		if clid, ok := item["clientId"].(string); ok {
			clientID = clid
			connect := WebsocketConnectRequest{
				ID:      "3",
				Channel: "/meta/connect",
				Advice: AdviceType{
					Timeout: 0,
				},
				ClientID:       clientID,
				ConnectionType: "websocket",
			}

			c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`[{"id":"2","channel":"/service/sub","data":{"reset":true},"clientId":"%s"}]`, clientID)))
			c.WriteMessage(websocket.TextMessage, makeRequest(connect))
			c.WriteMessage(websocket.TextMessage, quoteRequest(clientID, symbols))
		} else {
			fmt.Printf("%s\n", message)
		}
	}

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			handler(message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	handshakeRequest := HandshakeRequest{
		Ext: ExtType{
			StreamerToken: streamerToken,
		},
		ID:             "1",
		Version:        "1.0",
		MinimumVersion: "1.0",
		Channel:        "/meta/handshake",
		SupportedConnectionTypes: []string{
			"websocket",
			"long-polling",
			"callback-polling",
		},
		Advice: AdviceType{
			Timeout:  60000,
			Interval: 0,
		},
	}

	c.WriteMessage(websocket.TextMessage, makeRequest(handshakeRequest))

	// c.WriteMessage(websocket.TextMessage, []byte(`[{"id":"2","channel":"/service/sub","data":{"reset":true},"clientId":"3nc51don848o6r079r0jimav13og7"}]`))
	// c.WriteMessage(websocket.TextMessage, []byte(`[{"id":"3","channel":"/meta/connect","connectionType":"websocket","advice":{"timeout":0},"clientId":"3nc51don848o6r079r0jimav13og7"}]`))
	// c.WriteMessage(websocket.TextMessage, []byte(`[{"id":"4","channel":"/meta/connect","connectionType":"websocket","clientId":"3nc51don848o6r079r0jimav13og7"}]`))
	// c.WriteMessage(websocket.TextMessage,

	for {
		select {
		case <-done:
			return
			// case t := <-ticker.C:
			//	err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			//	if err != nil {
			//		log.Println("write:", err)
			//		return
			//	}
		}
	}
}
