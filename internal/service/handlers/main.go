package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
)

func InfoAddress(w http.ResponseWriter, r *http.Request) {
	ape.Render(w, map[string]any{
		"is_ok": true,
	})
}
