package tracer

import (
	"context"
	"fmt"
	"gopkg.in/go-playground/stats.v1"
)

type TracerServer struct {
	Ctx    context.Context
	Server *stats.ServerStats
}

type TracerParams struct {
	Ctx    context.Context
	Domain string
	Port   int
	Debug  bool
}

func NewTracerServer(tracerParams TracerParams) *TracerServer {
	ctx := tracerParams.Ctx
	config := &stats.ServerConfig{
		Domain: tracerParams.Domain,
		Port:   tracerParams.Port,
		Debug:  tracerParams.Debug,
	}

	server, err := stats.NewServer(config)

	if err != nil {
		panic(err)
	}

	return &TracerServer{
		Ctx:    ctx,
		Server: server,
	}

}

// Start
func (t *TracerServer) Start() {
	for stat := range t.Server.Run() {

		// calculate CPU times
		// totalCPUTimes := stat.CalculateTotalCPUTimes()
		// perCoreCPUTimes := stat.CalculateCPUTimes()

		// Do whatever you want with the data
		// * Save to database
		// * Stream elsewhere
		// * Print to console
		//

		fmt.Println(stat)
	}
}
