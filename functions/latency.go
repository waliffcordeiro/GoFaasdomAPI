package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Go_latency(w http.ResponseWriter, r *http.Request) {

	msg := &Message{
		Success: true,
		Payload: Payload{
			Test: "latency test",
		},
	}

	b, err := json.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprint(w, string(b))
}
