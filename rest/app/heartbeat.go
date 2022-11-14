package app

import "github.com/application-research/estuary-metrics/core"

func InitHeartbeat() {
	hb, err := core.NewHeartbeat([]core.ServerLookup{
		{
			Name:          "Main Estuary website",
			Endpoint:      "https://estuary.tech",
			LastHeartbeat: 0,
		},
		{
			Name:          "API Node",
			Endpoint:      "https://api.estuary.tech",
			LastHeartbeat: 0,
		},
		{
			Name:          "Upload Node",
			Endpoint:      "https://upload.estuary.tech",
			LastHeartbeat: 0,
		},
		{
			Name:          "Backup Node",
			Endpoint:      "https://backup.estuary.tech",
			LastHeartbeat: 0,
		},
		{
			Name:          "Shuttle6",
			Endpoint:      "https://shuttle-6.estuary.tech",
			LastHeartbeat: 0,
		},
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
