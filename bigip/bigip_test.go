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
	"net/http"
	"sync"
	"testing"

	"github.com/Staples-Inc/snap-plugin-collector-bigip/mocks"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/jarcoal/httpmock"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	ENDPOINT = "http://example.com:443/mgmt/tm/ltm/"
	VIRTUAL  = "virtual"
	POOL     = "pool"
	RULE     = "rule"
	NODE     = "node"
)

func TestBigIPCollector(t *testing.T) {
	Convey("Test GetMetricTypes", t, func() {
		bc := NewF5Collector()

		Convey("Collect String", func() {
			mt, err := bc.GetMetricTypes(nil)
			So(err, ShouldBeNil)
			So(len(mt), ShouldEqual, len(NodeMetrics)+len(PoolMetrics)+len(RuleMetrics)+len(VSMetrics))
		})
	})

	Convey("Test GetConfigPolicy", t, func() {
		bc := NewF5Collector()
		_, err := bc.GetConfigPolicy()

		Convey("No error returned", func() {
			So(err, ShouldBeNil)
		})
	})

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", ENDPOINT+"virtual/stats",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mocks.ResponseMap[VIRTUAL])
			return resp, nil
		},
	)
	httpmock.RegisterResponder("GET", ENDPOINT+"node/stats",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mocks.ResponseMap[NODE])
			return resp, nil
		},
	)
	httpmock.RegisterResponder("GET", ENDPOINT+"pool/stats",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mocks.ResponseMap[POOL])
			return resp, nil
		},
	)
	httpmock.RegisterResponder("GET", ENDPOINT+"rule/stats",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, mocks.ResponseMap[RULE])
			return resp, nil
		},
	)
	Convey("Test CollectMetrics", t, func() {
		bc := NewF5Collector()
		mt, _ := bc.GetMetricTypes(nil)
		for i := range mt {
			mt[i].Config = plugin.Config{
				"host":       "example.com",
				"port":       int64(443),
				"username":   "user",
				"password":   "pass",
				"basic_auth": true,
				"https":      false,
			}
		}
		results, err := bc.CollectMetrics(mt)
		So(len(results), ShouldBeGreaterThan, 0)
		Convey("No error returned", func() {
			So(err, ShouldBeNil)
		})
		results, err = bc.CollectMetrics(mt)
		So(len(results), ShouldBeGreaterThan, 0)
		Convey("Ensure no error once initialized", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("Test CollectMetrics", t, func() {
		bc := F5Collector{
			initializedMutex: new(sync.Mutex),
			initialized:      false,
		}
		cfg := plugin.Config{
			"host":       "example.com",
			"port":       int64(443),
			"username":   "user",
			"password":   "pass",
			"basic_auth": true,
			"https":      false,
		}
		err := bc.loadConfigs(cfg)
		Convey("No error returned", func() {
			So(err, ShouldBeNil)
		})
		Convey("Plugin is initialized", func() {
			So(bc.initialized, ShouldBeTrue)
		})
	})
}
