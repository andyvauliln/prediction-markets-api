package types

import (
	"encoding/json"
	"testing"

	"github.com/andyvauliln/prediction-markets-api/utils/testutils"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
)

func Compare(t *testing.T, expected interface{}, value interface{}) {
	expectedBytes, _ := json.Marshal(expected)
	bytes, _ := json.Marshal(value)

	assert.JSONEqf(t, string(expectedBytes), string(bytes), "")
}

func CompareStructs(t *testing.T, expected interface{}, order interface{}) {
	diff := deep.Equal(expected, order)
	if diff != nil {
		t.Errorf("\n%+v\nGot: \n%+v\n\n", expected, order)
	}
}

func TestWebSocketMessageJSON(t *testing.T) {
	market := testutils.GetMockMarket()

	msg := &WebSocketMessage{
		Channel: "markets",
		Payload: WebSocketPayload{
			Type: "NEW_MARKET",
			Data: order,
		},
	}

	encoded, err := json.Marshal(msg)
	if err != nil {
		t.Errorf("Error encoding order: %v", err)
	}

	decoded := &WebSocketMessage{}
	err = json.Unmarshal([]byte(encoded), &decoded)
	if err != nil {
		t.Errorf("Could not unmarshal payload: %v", err)
	}

	//we re-encode the order because the order is unmarshalled as a map[string]interface
	encodedOrder, err := json.Marshal(decoded.Payload.Data)
	if err != nil {
		t.Errorf("Error encoding order: %v", err)
	}

	decodedOrder := &Order{}
	err = json.Unmarshal(encodedOrder, decodedOrder)
	if err != nil {
		t.Errorf("Error decoding order; %v", err)
	}

	testutils.Compare(t, msg, decoded)
	testutils.CompareStructs(t, order, decodedOrder)
}

