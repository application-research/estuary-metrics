package core

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"net/http"
	"time"
)

// Heartbeat import launch heartbeat checker of all connected nodes.

type Heartbeat struct {
	ServerLookup []ServerLookup
	Badger       *badger.DB
}

type ServerLookup struct {
	Name          string
	Endpoint      string
	LastHeartbeat int64
}

type ServerCheck struct {
	ServerLookup  ServerLookup
	Status        string
	CheckDateTime time.Time
}

func NewHeartbeat(serverLookup []ServerLookup) (hb Heartbeat, err error) {
	hb = Heartbeat{
		ServerLookup: serverLookup,
	}
	//	initialize badger.
	db, err := badger.Open(badger.DefaultOptions("/tmp/est-metrics"))
	if err != nil {
		return
	}
	hb.Badger = db
	return hb, nil
}

func (m Heartbeat) Run() (err error) {

	// 	check every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				m.checkHeartBeat(m.ServerLookup)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return
}

func (m Heartbeat) checkHeartBeat(serverLookup []ServerLookup) (err error) {

	for _, server := range serverLookup {
		//	check heartbeat
		fmt.Println("check heartbeat: ", server.Name)s
		client := &http.Client{}
		req, errRq := http.NewRequest("GET", server.Endpoint, nil)
		if errRq != nil {
			fmt.Println(errRq)
			continue
		}
		res, errRes := client.Do(req)
		if errRes != nil {
			fmt.Println(errRes)
		}
		if res.StatusCode == 200 {
			fmt.Println("success")
			//	save state to badger
			m.Badger.Update(func(txn *badger.Txn) error {
				err := txn.Set([]byte(server.Name), []byte("success"))
				return err
			})
			return
		}
		if res.StatusCode == 500 {
			fmt.Println("failed")
			//	save state to badger
			return
		}
	}
	return
}
