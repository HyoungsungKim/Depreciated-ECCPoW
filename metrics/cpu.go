// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package metrics

<<<<<<< HEAD:swarm/api/testapi.go
import (
	"github.com/Onther-Tech/go-ethereum/swarm/network"
)

type Control struct {
	api  *API
	hive *network.Hive
}

func NewControl(api *API, hive *network.Hive) *Control {
	return &Control{api, hive}
}

//func (self *Control) BlockNetworkRead(on bool) {
//	self.hive.BlockNetworkRead(on)
//}
//
//func (self *Control) SyncEnabled(on bool) {
//	self.hive.SyncEnabled(on)
//}
//
//func (self *Control) SwapEnabled(on bool) {
//	self.hive.SwapEnabled(on)
//}
//
func (c *Control) Hive() string {
	return c.hive.String()
=======
import "github.com/elastic/gosigar"

// CPUStats is the system and process CPU stats.
type CPUStats struct {
	GlobalTime int64 // Time spent by the CPU working on all processes
	GlobalWait int64 // Time spent by waiting on disk for all processes
	LocalTime  int64 // Time spent by the CPU working on this process
}

// ReadCPUStats retrieves the current CPU stats.
func ReadCPUStats(stats *CPUStats) {
	global := gosigar.Cpu{}
	global.Get()

	stats.GlobalTime = int64(global.User + global.Nice + global.Sys)
	stats.GlobalWait = int64(global.Wait)
	stats.LocalTime = getProcessCPUTime()
>>>>>>> upstream/master:metrics/cpu.go
}
