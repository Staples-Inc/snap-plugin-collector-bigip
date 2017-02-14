package bigip

/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 Staples, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/pr8kerl/f5er/f5"
)

type poolMetric struct {
	metricName  string
	description string
	extract     func(f5.LBPoolStatsInnerEntries) float64
}

type poolCollector struct {
	bigIP *f5.Device
}

func newPoolCollector(client *f5.Device) poolCollector {
	return poolCollector{
		bigIP: client,
	}
}

var (
	// PoolMetrics contains the available pool metric functions
	PoolMetrics = []poolMetric{
		{
			metricName:  "connq_all_age_max",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ConnqAll_ageMax.Value / 1000)
			},
		},
		{
			metricName:  "serverside_cur_conns",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_curConns.Value)
			},
		},
		{
			metricName:  "min_active_members",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.MinActiveMembers.Value)
			},
		},
		{
			metricName:  "serverside_bytes_in",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_bitsIn.Value / 8)
			},
		},
		{
			metricName:  "connq_all_serviced",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ConnqAll_serviced.Value)
			},
		},
		{
			metricName:  "serverside_pkts_in",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_pktsIn.Value)
			},
		},
		{
			metricName:  "serverside_max_conns",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_maxConns.Value)
			},
		},
		{
			metricName:  "connq_depth",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Connq_depth.Value)
			},
		},
		{
			metricName:  "connq_all_depth",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ConnqAll_depth.Value)
			},
		},
		{
			metricName:  "connq_all_age_head",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ConnqAll_ageHead.Value / 1000)
			},
		},
		{
			metricName:  "cur_sessions",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.CurSessions.Value)
			},
		},
		{
			metricName:  "connq_serviced",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Connq_serviced.Value)
			},
		},
		{
			metricName:  "tot_requests",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.TotRequests.Value)
			},
		},
		{
			metricName:  "connq_all_age_edm",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ConnqAll_ageEdm.Value / 1000)
			},
		},
		{
			metricName:  "connq_age_head",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Connq_ageHead.Value / 1000)
			},
		},
		{
			metricName:  "connq_age_max",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Connq_ageMax.Value / 1000)
			},
		},
		{
			metricName:  "connq_age_edm",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Connq_ageEdm.Value)
			},
		},
		{
			metricName:  "serverside_bytes_out",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_bitsOut.Value / 8)
			},
		},
		{
			metricName:  "connq_age_ema",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Connq_ageEma.Value / 1000)
			},
		},
		{
			metricName:  "connq_all_age_ema",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ConnqAll_ageEma.Value / 1000)
			},
		},
		{
			metricName:  "active_member_cnt",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.ActiveMemberCnt.Value)
			},
		},
		{
			metricName:  "serverside_pkts_out",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_pktsOut.Value)
			},
		},
		{
			metricName:  "serverside_tot_conns",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				return float64(entries.Serverside_totConns.Value)
			},
		},
		{
			metricName:  "status_availability_state",
			description: "",
			extract: func(entries f5.LBPoolStatsInnerEntries) float64 {
				if entries.Status_availabilityState.Description == AVAILABLE {
					return 1
				}
				return 0
			},
		},
	}
)

func getPoolMetricsCatalog() []plugin.Metric {
	poolMetrics := make([]plugin.Metric, 0, len(PoolMetrics))
	for _, met := range PoolMetrics {
		poolMetrics = append(poolMetrics, plugin.Metric{
			Namespace:   createPoolNamespace("*", "*", met.metricName),
			Description: met.description,
		})
	}
	return poolMetrics
}

func (c poolCollector) Collect(ch chan<- plugin.Metric) {
	err, poolStats := c.bigIP.ShowAllPoolStats()
	if err != nil {
		log.Warningf("error collecting pool data; %v", err)
	} else {
		for key, poolStat := range poolStats.Entries {
			keyParts := strings.Split(key, "/")
			path := keyParts[len(keyParts)-2]
			pathParts := strings.Split(path, "~")
			partition := pathParts[1]
			poolName := pathParts[len(pathParts)-1]

			for _, met := range PoolMetrics {
				ch <- plugin.Metric{
					Namespace:   createPoolNamespace(partition, poolName, met.metricName),
					Description: met.description,
					Data:        met.extract(poolStat.NestedStats.Entries),
				}
			}
		}
	}
}

func createPoolNamespace(partition string, pool string, metricName string) plugin.Namespace {
	return append(baseNamespace, plugin.NamespaceElement{
		Value: "pool",
	},
		plugin.NamespaceElement{
			Value:       partition,
			Name:        "partition",
			Description: "logical container that defines a set of BIG-IP objects",
		},
		plugin.NamespaceElement{
			Value:       pool,
			Name:        "pool",
			Description: "A logical set of devices grouped together to receive and process traffic.",
		},
		plugin.NamespaceElement{
			Value: metricName,
		})
}
