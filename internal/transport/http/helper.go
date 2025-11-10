package transporthttp

import (
	"encoding/json"
	stdhttp "net/http"
)

func writeJSON(w stdhttp.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
