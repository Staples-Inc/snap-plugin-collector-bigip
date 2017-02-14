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

type nodeMetric struct {
	metricName  string
	description string
	extract     func(f5.LBNodeStatsInnerEntries) float64
}

type nodeCollector struct {
	bigIP *f5.Device
}

func newNodeCollector(client *f5.Device) nodeCollector {
	return nodeCollector{
		bigIP: client,
	}
}

var (
	// NodeMetrics maps the available node metric functions
	NodeMetrics = []nodeMetric{
		{
			metricName:  "serverside_bytes_out",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_bitsOut.Value / 8)
			},
		},
		{
			metricName:  "serverside_max_conns",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_maxConns.Value)
			},
		},
		{
			metricName:  "serverside_cur_conns",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_curConns.Value)
			},
		},
		{
			metricName:  "serverside_pkts_out",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_pktsOut.Value)
			},
		},
		{
			metricName:  "tot_requests",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.TotRequests.Value)
			},
		},
		{
			metricName:  "serverside_pkts_in",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_pktsIn.Value)
			},
		},
		{
			metricName:  "serverside_tot_conns",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_totConns.Value)
			},
		},
		{
			metricName:  "serverside_bytes_in",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.Serverside_bitsIn.Value / 8)
			},
		},
		{
			metricName:  "cur_sessions",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				return float64(entries.CurSessions.Value)
			},
		},
		{
			metricName:  "status_availability_state",
			description: "",
			extract: func(entries f5.LBNodeStatsInnerEntries) float64 {
				if entries.Status_availabilityState.Description == AVAILABLE {
					return 1
				}
				return 0
			},
		},
	}
)

func getNodeMetricsCatalog() []plugin.Metric {
	nodeMetrics := make([]plugin.Metric, 0, len(NodeMetrics))
	for _, met := range NodeMetrics {
		nodeMetrics = append(nodeMetrics, plugin.Metric{
			Namespace:   createNodeNamespace("*", "*", met.metricName),
			Description: met.description,
		})
	}
	return nodeMetrics
}

func (c nodeCollector) Collect(ch chan<- plugin.Metric) {
	err, nodeStats := c.bigIP.ShowAllNodeStats()
	if err != nil {
		log.Warningf("error collecting node data %v", err)
	} else {
		for key, nodeStat := range nodeStats.Entries {
			keyParts := strings.Split(key, "/")
			path := keyParts[len(keyParts)-2]
			pathParts := strings.Split(path, "~")
			partition := pathParts[1]
			nodeName := pathParts[len(pathParts)-1]

			for _, met := range NodeMetrics {
				ch <- plugin.Metric{
					Namespace:   createNodeNamespace(partition, nodeName, met.metricName),
					Description: met.description,
					Data:        met.extract(nodeStat.NestedStats.Entries),
				}
			}
		}
	}
}

func createNodeNamespace(partition string, node string, metricName string) plugin.Namespace {
	return append(baseNamespace, plugin.NamespaceElement{
		Value: "node",
	},
		plugin.NamespaceElement{
			Value:       partition,
			Name:        "partition",
			Description: "logical container that defines a set of BIG-IP objects",
		},
		plugin.NamespaceElement{
			Value:       node,
			Name:        "node",
			Description: "ip of physical node host",
		},
		plugin.NamespaceElement{
			Value: metricName,
		})
}
