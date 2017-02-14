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

type ruleMetric struct {
	metricName  string
	description string
	extract     func(f5.LBRuleStatsInnerEntries) float64
}

type ruleCollector struct {
	bigIP *f5.Device
}

func newRuleCollector(client *f5.Device) ruleCollector {
	return ruleCollector{
		bigIP: client,
	}
}

var (
	// RuleMetrics contains the available rule metrics
	RuleMetrics = []ruleMetric{
		{
			metricName:  "priority",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.Priority.Value)
			},
		},
		{
			metricName:  "failures",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.Failures.Value)
			},
		},
		{
			metricName:  "total_executions",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.TotalExecutions.Value)
			},
		},
		{
			metricName:  "aborts",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.Aborts.Value)
			},
		},
		{
			metricName:  "min_cycles",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.MinCycles.Value)
			},
		},
		{
			metricName:  "max_cycles",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.MaxCycles.Value)
			},
		},
		{
			metricName:  "avg_cycles",
			description: "",
			extract: func(entries f5.LBRuleStatsInnerEntries) float64 {
				return float64(entries.AvgCycles.Value)
			},
		},
	}
)

func getRuleMetricsCatalog() []plugin.Metric {
	metrics := make([]plugin.Metric, 0, len(RuleMetrics))
	for _, met := range RuleMetrics {
		metrics = append(metrics, plugin.Metric{
			Namespace:   createRuleNamespace("*", "*", "*", met.metricName),
			Description: met.description,
		})
	}
	return metrics
}

func (c ruleCollector) Collect(ch chan<- plugin.Metric) {
	err, nodeStats := c.bigIP.ShowAllRuleStats()
	if err != nil {
		log.Warningf("error collecting rule data; %v", err)
	} else {
		for key, stat := range nodeStats.Entries {
			keyParts := strings.Split(key, "/")
			path := keyParts[len(keyParts)-2]
			pathParts := strings.Split(path, "~")
			partition := pathParts[1]
			eventParts := strings.Split(pathParts[len(pathParts)-1], ":")
			ruleName := eventParts[0]
			event := eventParts[1]

			for _, met := range RuleMetrics {
				ch <- plugin.Metric{
					Namespace:   createRuleNamespace(partition, ruleName, event, met.metricName),
					Description: met.description,
					Data:        met.extract(stat.NestedStats.Entries),
				}
			}
		}
	}
}

func createRuleNamespace(partition string, rule string, event string, metricName string) plugin.Namespace {
	return append(baseNamespace, plugin.NamespaceElement{
		Value: "rule",
	},
		plugin.NamespaceElement{
			Value:       partition,
			Name:        "partition",
			Description: "logical container that defines a set of BIG-IP objects",
		},
		plugin.NamespaceElement{
			Value:       rule,
			Name:        "rule",
			Description: "",
		},
		plugin.NamespaceElement{
			Value:       event,
			Name:        "event",
			Description: "event that triggers rule",
		},
		plugin.NamespaceElement{
			Value: metricName,
		})
}
