package app

import "github.com/application-research/estuary-metrics/core"

func InitHeartbeat() {
	hb, err := core.NewHeartbeat([]core.ServerLookup{
		{
			Name:          "Shuttle7",
			Endpoint:      "https://shuttle-7.estuary.tech",
			LastHeartbeat: 0,
		},
		{
			Name:          "Shuttle8",
			Endpoint:      "https://shuttle-8.estuary.tech",
			LastHeartbeat: 0,
		},
	})

	if err != nil {
		panic(err)
	}
	hb.Run() // run it
}
