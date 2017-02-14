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

type vsMetric struct {
	metricName  string
	description string
	extract     func(f5.LBVirtualStatsInnerEntries) float64
}

type vsCollector struct {
	bigIP *f5.Device
}

func newVSCollector(client *f5.Device) vsCollector {
	return vsCollector{
		bigIP: client,
	}
}

// VSMetrics maps the available vs metric functions
var VSMetrics = []vsMetric{
	{
		metricName:  "syncookie_accepts",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_accepts.Value)
		},
	},
	{
		metricName:  "ephemeral_bytes_out",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_bitsOut.Value / 8)
		},
	},
	{
		metricName:  "clientside_bytes_out",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_bitsOut.Value / 8)
		},
	},
	{
		metricName:  "five_min_avg_usage_ratio",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.FiveMinAvgUsageRatio.Value)
		},
	},
	{
		metricName:  "five_sec_avg_usage_ratio",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.FiveSecAvgUsageRatio.Value)
		},
	},
	{
		metricName:  "syncookie_syncookies",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_syncookies.Value)
		},
	},
	{
		metricName:  "ephemeral_slow_killed",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_slowKilled.Value)
		},
	},
	{
		metricName:  "ephemeral_pkts_out",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_pktsOut.Value)
		},
	},
	{
		metricName:  "syncookie_rejects",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_rejects.Value)
		},
	},
	{
		metricName:  "syncookie_syncache_curr",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_syncacheCurr.Value)
		},
	},
	{
		metricName:  "cs_min_conn_dur",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.CsMinConnDur.Value)
		},
	},
	{
		metricName:  "cs_mean_conn_dur",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.CsMeanConnDur.Value)
		},
	},
	{
		metricName:  "syncookie_swsyncookie_instance",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_swsyncookieInstance.Value)
		},
	},
	{
		metricName:  "syncookie_syncache_over",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_syncacheOver.Value)
		},
	},
	{
		metricName:  "syncookie_hw_accepts",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_hwAccepts.Value)
		},
	},
	{
		metricName:  "ephemeral_pkts_in",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_pktsIn.Value)
		},
	},
	{
		metricName:  "clientside_tot_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_totConns.Value)
		},
	},
	{
		metricName:  "ephemeral_cur_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_curConns.Value)
		},
	},
	{
		metricName:  "clientside_evicted_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_evictedConns.Value)
		},
	},
	{
		metricName:  "one_min_avg_usage_ratio",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.OneMinAvgUsageRatio.Value)
		},
	},
	{
		metricName:  "ephemeral_evicted_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_evictedConns.Value)
		},
	},
	{
		metricName:  "clientside_slow_killed",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_slowKilled.Value)
		},
	},
	{
		metricName:  "clientside_bytes_in",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_bitsIn.Value / 8)
		},
	},
	{
		metricName:  "ephemeral_max_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_maxConns.Value)
		},
	},
	{
		metricName:  "syncookie_hwsyncookie_instance",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_hwsyncookieInstance.Value)
		},
	},
	{
		metricName:  "clientside_pkts_out",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_pktsOut.Value)
		},
	},
	{
		metricName:  "clientside_cur_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_curConns.Value)
		},
	},
	{
		metricName:  "ephemeral_bytes_in",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_bitsIn.Value / 8)
		},
	},
	{
		metricName:  "clientside_pkts_in",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_pktsIn.Value)
		},
	},
	{
		metricName:  "tot_requests",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.TotRequests.Value)
		},
	},
	{
		metricName:  "cs_max_conn_dur",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.CsMaxConnDur.Value)
		},
	},
	{
		metricName:  "syncookie_hw_syncookies",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Syncookie_hwSyncookies.Value)
		},
	},
	{
		metricName:  "clientside_max_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Clientside_maxConns.Value)
		},
	},
	{
		metricName:  "ephemeral_tot_conns",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			return float64(entries.Ephemeral_totConns.Value)
		},
	},
	{
		metricName:  "status_availability_state",
		description: "",
		extract: func(entries f5.LBVirtualStatsInnerEntries) float64 {
			if entries.Status_availabilityState.Description == AVAILABLE {
				return 1
			}
			return 0
		},
	},
}

func getVSMetricsCatalog() []plugin.Metric {
	metrics := make([]plugin.Metric, 0, len(VSMetrics))
	for _, met := range VSMetrics {
		metrics = append(metrics, plugin.Metric{
			Namespace:   createVSNamespace("*", "*", met.metricName),
			Description: met.description,
		})
	}
	return metrics
}

func (c vsCollector) Collect(ch chan<- plugin.Metric) {
	err, vsStats := c.bigIP.ShowAllVirtualStats()
	if err != nil {
		log.Warningf("error collecting vs data; %v", err)
	} else {
		for key, vsStat := range vsStats.Entries {
			keyParts := strings.Split(key, "/")
			path := keyParts[len(keyParts)-2]
			pathParts := strings.Split(path, "~")
			partition := pathParts[1]
			vsName := pathParts[len(pathParts)-1]

			for _, met := range VSMetrics {
				ch <- plugin.Metric{
					Namespace:   createVSNamespace(partition, vsName, met.metricName),
					Description: met.description,
					Data:        met.extract(vsStat.NestedStats.Entries),
				}
			}
		}
	}
}

func createVSNamespace(partition string, vs string, metricName string) plugin.Namespace {
	return append(baseNamespace, plugin.NamespaceElement{
		Value: "vs",
	},
		plugin.NamespaceElement{
			Value:       partition,
			Name:        "partition",
			Description: "logical container that defines a set of BIG-IP objects",
		},
		plugin.NamespaceElement{
			Value:       vs,
			Name:        "vs",
			Description: "a virtual server is a traffic-management object on the BIG-IP system that is represented by an IP address and a service",
		},
		plugin.NamespaceElement{
			Value: metricName,
		})
}
