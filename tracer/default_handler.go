package tracer

import (
	"gopkg.in/go-playground/stats.v1"
)

func InitializeMyOwnServerForHandler() {
	server := NewTracerServer(TracerParams{
		Domain: "localhost",
	})
	server.SetHandler(TracerStatsHandler{
		func(stats *stats.Stats) {
			// handle any incoming stats
		},
	})
	server.Start()
}

func (h *MyCustomHandler) HandleStats(stats *stats.Stats) {
	h.Handler(stats)
}
