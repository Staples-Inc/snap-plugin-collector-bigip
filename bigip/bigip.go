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
	"strconv"
	"strings"
	"sync"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/pr8kerl/f5er/f5"
)

const (
	// AVAILABLE description to match available resources
	AVAILABLE = "available"
)

var baseNamespace = plugin.Namespace{plugin.NamespaceElement{Value: "staples"}, plugin.NamespaceElement{Value: "bigip"}}

// Collector allows for mocking the data from the bigip device
type Collector interface {
	Collect(chan<- plugin.Metric)
}

// NewF5Collector returns an F5Collector for the snap plugin main process
func NewF5Collector() F5Collector {
	imutex := new(sync.Mutex)
	return F5Collector{
		initializedMutex: imutex,
		initialized:      false,
		collectors:       []Collector{},
	}
}

// F5Collector interfaces with the BigIP api
type F5Collector struct {
	collectors       []Collector
	initializedMutex *sync.Mutex
	initialized      bool
}

func (c *F5Collector) loadConfigs(cfg plugin.Config) error {
	c.initializedMutex.Lock()
	defer c.initializedMutex.Unlock()
	if c.initialized {
		return nil
	}
	var (
		host  string
		port  int64
		user  string
		pass  string
		err   error
		auth  bool
		https bool
	)
	if host, err = cfg.GetString("host"); err != nil {
		return err
	}
	if port, err = cfg.GetInt("port"); err != nil {
		port = 443
	}
	if user, err = cfg.GetString("username"); err != nil {
		user = ""
	}
	if pass, err = cfg.GetString("password"); err != nil {
		pass = ""
	}
	if auth, err = cfg.GetBool("basic_auth"); err != nil {
		auth = false
	}
	if https, err = cfg.GetBool("https"); err != nil {
		https = true
	}
	authMethod := f5.TOKEN
	if auth {
		authMethod = f5.BASIC_AUTH
	}
	endpoint := host + ":" + strconv.Itoa(int(port))
	client := f5.New(endpoint, user, pass, authMethod)
	if !https {
		client = f5.NewInsecure(endpoint, user, pass, authMethod)
	}
	node := newNodeCollector(client)
	pool := newPoolCollector(client)
	rule := newRuleCollector(client)
	vs := newVSCollector(client)
	c.collectors = append(c.collectors, &node, &pool, &rule, &vs)
	c.initialized = true
	return nil
}

// GetConfigPolicy returns the configuration of this plugin
func (c F5Collector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	cfg := plugin.NewConfigPolicy()
	cfg.AddNewStringRule([]string{"staples", "bigip"}, "host", true)
	cfg.AddNewIntRule([]string{"staples", "bigip"}, "port", false, plugin.SetDefaultInt(443), plugin.SetMinInt(0), plugin.SetMaxInt(65535))
	cfg.AddNewStringRule([]string{"staples", "bigip"}, "username", false)
	cfg.AddNewStringRule([]string{"staples", "bigip"}, "password", false)
	cfg.AddNewBoolRule([]string{"staples", "bigip"}, "basic_auth", false, plugin.SetDefaultBool(false))
	cfg.AddNewBoolRule([]string{"staples", "bigip"}, "https", false, plugin.SetDefaultBool(true))
	cfg.AddNewBoolRule([]string{"staples", "bigip"}, "sanitize", false, plugin.SetDefaultBool(false))
	return *cfg, nil
}

// GetMetricTypes returns the available metrics from this plugin
func (c F5Collector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := getNodeMetricsCatalog()
	metrics = append(metrics, getPoolMetricsCatalog()...)
	metrics = append(metrics, getRuleMetricsCatalog()...)
	metrics = append(metrics, getVSMetricsCatalog()...)

	return metrics, nil
}

func sanitizeNodes(mts []plugin.Metric) error {
	r := strings.NewReplacer(".", "_")
	for i := range mts {
		for n := range mts[i].Namespace {
			mts[i].Namespace[n].Value = r.Replace(mts[i].Namespace[n].Value)
		}
	}
	return nil
}

// CollectMetrics collects the requested metrics from this plugin
func (c F5Collector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	if err := c.loadConfigs(mts[0].Config); err != nil {
		return nil, err
	}
	var sanitize bool
	var err error
	if sanitize, err = mts[0].Config.GetBool("sanitize"); err != nil {
		sanitize = false
	}
	metricChan := make(chan plugin.Metric)
	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		wg.Wait()
		close(metricChan)
	}()
	for _, collect := range c.collectors {
		go func(coll Collector) {
			coll.Collect(metricChan)
			defer wg.Done()
		}(collect)
	}

	metrics := []plugin.Metric{}
	for mt := range metricChan {
		metrics = append(metrics, mt)
	}

	if sanitize {
		err = sanitizeNodes(metrics)
		if err != nil {
			return nil, err
		}
	}
	return metrics, nil
}
