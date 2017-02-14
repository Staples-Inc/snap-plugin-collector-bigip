package mocks

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

var (
	nodeResponse = `{
  "kind": "tm:ltm:node:nodecollectionstats",
  "selfLink": "https://localhost/mgmt/tm/ltm/node/stats",
  "entries": {
    "https://localhost/mgmt/tm/ltm/node/~Common~2.2.2.1/stats": {
      "nestedStats": {
        "kind": "tm:ltm:node:nodestats",
        "selfLink": "https://localhost/mgmt/tm/ltm/node/~Common~2.2.2.1/stats",
        "entries": {
          "addr": {
            "description": "2.2.2.1"
          },
          "curSessions": {
            "value": 0
          },
          "monitorRule": {
            "description": "none"
          },
          "monitorStatus": {
            "description": "unchecked"
          },
          "tmName": {
            "description": "/Common/2.2.2.1"
          },
          "serverside.bitsIn": {
            "value": 0
          },
          "serverside.bitsOut": {
            "value": 0
          },
          "serverside.curConns": {
            "value": 0
          },
          "serverside.maxConns": {
            "value": 0
          },
          "serverside.pktsIn": {
            "value": 0
          },
          "serverside.pktsOut": {
            "value": 0
          },
          "serverside.totConns": {
            "value": 0
          },
          "sessionStatus": {
            "description": "enabled"
          },
          "status.availabilityState": {
            "description": "unknown"
          },
          "status.enabledState": {
            "description": "enabled"
          },
          "status.statusReason": {
            "description": "Node address does not have service checking enabled"
          },
          "totRequests": {
            "value": 0
          }
        }
      }
    },
    "https://localhost/mgmt/tm/ltm/node/~Common~2.2.2.2/stats": {
      "nestedStats": {
        "kind": "tm:ltm:node:nodestats",
        "selfLink": "https://localhost/mgmt/tm/ltm/node/~Common~2.2.2.2/stats?ver=11.5.3",
        "entries": {
          "addr": {
            "description": "2.2.2.2"
          },
          "curSessions": {
            "value": 0
          },
          "monitorRule": {
            "description": "none"
          },
          "monitorStatus": {
            "description": "unchecked"
          },
          "tmName": {
            "description": "/Common/2.2.2.2"
          },
          "serverside.bitsIn": {
            "value": 0
          },
          "serverside.bitsOut": {
            "value": 0
          },
          "serverside.curConns": {
            "value": 0
          },
          "serverside.maxConns": {
            "value": 0
          },
          "serverside.pktsIn": {
            "value": 0
          },
          "serverside.pktsOut": {
            "value": 0
          },
          "serverside.totConns": {
            "value": 0
          },
          "sessionStatus": {
            "description": "enabled"
          },
          "status.availabilityState": {
            "description": "unknown"
          },
          "status.enabledState": {
            "description": "enabled"
          },
          "status.statusReason": {
            "description": "Node address does not have service checking enabled"
          },
          "totRequests": {
            "value": 0
          }
        }
      }
    }
  }
  }`
	virtualResponse = `{
    "kind": "tm:ltm:virtual:virtualcollectionstats",
    "selfLink": "https://localhost/mgmt/tm/ltm/virtual/stats?ver=11.5.3",
    "entries": {
      "https://localhost/mgmt/tm/ltm/virtual/~Common~test_vs1/stats": {
        "nestedStats": {
          "kind": "tm:ltm:virtual:virtualstats",
          "selfLink": "https://localhost/mgmt/tm/ltm/virtual/~Common~test_vs1/stats?ver=11.5.3",
          "entries": {
            "actualPvaAccel": {
              "description": "none"
            },
            "clientside.bitsIn": {
              "value": 0
            },
            "clientside.bitsOut": {
              "value": 0
            },
            "clientside.curConns": {
              "value": 0
            },
            "clientside.maxConns": {
              "value": 0
            },
            "clientside.pktsIn": {
              "value": 0
            },
            "clientside.pktsOut": {
              "value": 0
            },
            "clientside.totConns": {
              "value": 0
            },
            "cmpEnableMode": {
              "description": "all-cpus"
            },
            "cmpEnabled": {
              "description": "enabled"
            },
            "csMaxConnDur": {
              "value": 0
            },
            "csMeanConnDur": {
              "value": 0
            },
            "csMinConnDur": {
              "value": 0
            },
            "destination": {
              "description": "1.1.1.1:80"
            },
            "ephemeral.bitsIn": {
              "value": 0
            },
            "ephemeral.bitsOut": {
              "value": 0
            },
            "ephemeral.curConns": {
              "value": 0
            },
            "ephemeral.maxConns": {
              "value": 0
            },
            "ephemeral.pktsIn": {
              "value": 0
            },
            "ephemeral.pktsOut": {
              "value": 0
            },
            "ephemeral.totConns": {
              "value": 0
            },
            "fiveMinAvgUsageRatio": {
              "value": 0
            },
            "fiveSecAvgUsageRatio": {
              "value": 0
            },
            "tmName": {
              "description": "/Common/test_vs1"
            },
            "oneMinAvgUsageRatio": {
              "value": 0
            },
            "status.availabilityState": {
              "description": "unknown"
            },
            "status.enabledState": {
              "description": "enabled"
            },
            "status.statusReason": {
              "description": "The children pool member(s) either don't have service checking enabled, or service check results are not available yet"
            },
            "syncookieStatus": {
              "description": "not-activated"
            },
            "syncookie.accepts": {
              "value": 0
            },
            "syncookie.hwAccepts": {
              "value": 0
            },
            "syncookie.hwSyncookies": {
              "value": 0
            },
            "syncookie.hwsyncookieInstance": {
              "value": 0
            },
            "syncookie.rejects": {
              "value": 0
            },
            "syncookie.swsyncookieInstance": {
              "value": 0
            },
            "syncookie.syncacheCurr": {
              "value": 0
            },
            "syncookie.syncacheOver": {
              "value": 0
            },
            "syncookie.syncookies": {
              "value": 0
            },
            "totRequests": {
              "value": 0
            }
          }
        }
      },
      "https://localhost/mgmt/tm/ltm/virtual/~Common~test_vs2/stats": {
        "nestedStats": {
          "kind": "tm:ltm:virtual:virtualstats",
          "selfLink": "https://localhost/mgmt/tm/ltm/virtual/~Common~test_vs2/stats?ver=11.5.3",
          "entries": {
            "actualPvaAccel": {
              "description": "none"
            },
            "clientside.bitsIn": {
              "value": 0
            },
            "clientside.bitsOut": {
              "value": 0
            },
            "clientside.curConns": {
              "value": 0
            },
            "clientside.maxConns": {
              "value": 0
            },
            "clientside.pktsIn": {
              "value": 0
            },
            "clientside.pktsOut": {
              "value": 0
            },
            "clientside.totConns": {
              "value": 0
            },
            "cmpEnableMode": {
              "description": "all-cpus"
            },
            "cmpEnabled": {
              "description": "enabled"
            },
            "csMaxConnDur": {
              "value": 0
            },
            "csMeanConnDur": {
              "value": 0
            },
            "csMinConnDur": {
              "value": 0
            },
            "destination": {
              "description": "1.1.1.2:80"
            },
            "ephemeral.bitsIn": {
              "value": 0
            },
            "ephemeral.bitsOut": {
              "value": 0
            },
            "ephemeral.curConns": {
              "value": 0
            },
            "ephemeral.maxConns": {
              "value": 0
            },
            "ephemeral.pktsIn": {
              "value": 0
            },
            "ephemeral.pktsOut": {
              "value": 0
            },
            "ephemeral.totConns": {
              "value": 0
            },
            "fiveMinAvgUsageRatio": {
              "value": 0
            },
            "fiveSecAvgUsageRatio": {
              "value": 0
            },
            "tmName": {
              "description": "/Common/test_vs2"
            },
            "oneMinAvgUsageRatio": {
              "value": 0
            },
            "status.availabilityState": {
              "description": "unknown"
            },
            "status.enabledState": {
              "description": "enabled"
            },
            "status.statusReason": {
              "description": "The children pool member(s) either don't have service checking enabled, or service check results are not available yet"
            },
            "syncookieStatus": {
              "description": "not-activated"
            },
            "syncookie.accepts": {
              "value": 0
            },
            "syncookie.hwAccepts": {
              "value": 0
            },
            "syncookie.hwSyncookies": {
              "value": 0
            },
            "syncookie.hwsyncookieInstance": {
              "value": 0
            },
            "syncookie.rejects": {
              "value": 0
            },
            "syncookie.swsyncookieInstance": {
              "value": 0
            },
            "syncookie.syncacheCurr": {
              "value": 0
            },
            "syncookie.syncacheOver": {
              "value": 0
            },
            "syncookie.syncookies": {
              "value": 0
            },
            "totRequests": {
              "value": 0
            }
          }
        }
      },
      "https://localhost/mgmt/tm/ltm/virtual/~Common~test_vs3/stats": {
        "nestedStats": {
          "kind": "tm:ltm:virtual:virtualstats",
          "selfLink": "https://localhost/mgmt/tm/ltm/virtual/~Common~test_vs3/stats?ver=11.5.3",
          "entries": {
            "actualPvaAccel": {
              "description": "none"
            },
            "clientside.bitsIn": {
              "value": 0
            },
            "clientside.bitsOut": {
              "value": 0
            },
            "clientside.curConns": {
              "value": 0
            },
            "clientside.maxConns": {
              "value": 0
            },
            "clientside.pktsIn": {
              "value": 0
            },
            "clientside.pktsOut": {
              "value": 0
            },
            "clientside.totConns": {
              "value": 0
            },
            "cmpEnableMode": {
              "description": "all-cpus"
            },
            "cmpEnabled": {
              "description": "enabled"
            },
            "csMaxConnDur": {
              "value": 0
            },
            "csMeanConnDur": {
              "value": 0
            },
            "csMinConnDur": {
              "value": 0
            },
            "destination": {
              "description": "1.1.1.3:80"
            },
            "ephemeral.bitsIn": {
              "value": 0
            },
            "ephemeral.bitsOut": {
              "value": 0
            },
            "ephemeral.curConns": {
              "value": 0
            },
            "ephemeral.maxConns": {
              "value": 0
            },
            "ephemeral.pktsIn": {
              "value": 0
            },
            "ephemeral.pktsOut": {
              "value": 0
            },
            "ephemeral.totConns": {
              "value": 0
            },
            "fiveMinAvgUsageRatio": {
              "value": 0
            },
            "fiveSecAvgUsageRatio": {
              "value": 0
            },
            "tmName": {
              "description": "/Common/test_vs3"
            },
            "oneMinAvgUsageRatio": {
              "value": 0
            },
            "status.availabilityState": {
              "description": "unknown"
            },
            "status.enabledState": {
              "description": "enabled"
            },
            "status.statusReason": {
              "description": "The children pool member(s) either don't have service checking enabled, or service check results are not available yet"
            },
            "syncookieStatus": {
              "description": "not-activated"
            },
            "syncookie.accepts": {
              "value": 0
            },
            "syncookie.hwAccepts": {
              "value": 0
            },
            "syncookie.hwSyncookies": {
              "value": 0
            },
            "syncookie.hwsyncookieInstance": {
              "value": 0
            },
            "syncookie.rejects": {
              "value": 0
            },
            "syncookie.swsyncookieInstance": {
              "value": 0
            },
            "syncookie.syncacheCurr": {
              "value": 0
            },
            "syncookie.syncacheOver": {
              "value": 0
            },
            "syncookie.syncookies": {
              "value": 0
            },
            "totRequests": {
              "value": 0
            }
          }
        }
      }
    }
  }`
	ruleResponse = `{
  "kind": "tm:ltm:rule:rulecollectionstats",
  "selfLink": "https://localhost/mgmt/tm/ltm/rule/stats?ver=11.5.3",
  "entries": {
    "https://localhost/mgmt/tm/ltm/rule/~Common~_sys_https_redirect:HTTP_REQUEST/stats": {
      "nestedStats": {
        "kind": "tm:ltm:rule:rulestats",
        "selfLink": "https://localhost/mgmt/tm/ltm/rule/~Common~_sys_https_redirect:HTTP_REQUEST/stats?ver=11.5.3",
        "entries": {
          "aborts": {
            "value": 0
          },
          "avgCycles": {
            "value": 0
          },
          "eventType": {
            "description": "HTTP_REQUEST"
          },
          "failures": {
            "value": 0
          },
          "maxCycles": {
            "value": 0
          },
          "minCycles": {
            "value": 0
          },
          "tmName": {
            "description": "/Common/_sys_https_redirect"
          },
          "priority": {
            "value": 500
          },
          "totalExecutions": {
            "value": 0
          }
        }
      }
    }
  }
}`
	poolResponse = `{
    "kind": "tm:ltm:pool:poolcollectionstats",
    "selfLink": "https://localhost/mgmt/tm/ltm/pool/stats",
    "entries": {
      "https://localhost/mgmt/tm/ltm/pool/~Common~test_pool1/stats": {
        "nestedStats": {
          "kind": "tm:ltm:pool:poolstats",
          "selfLink": "https://localhost/mgmt/tm/ltm/pool/~Common~test_pool1/stats",
          "entries": {
            "activeMemberCnt": {
              "value": 0
            },
            "connqAll.ageEdm": {
              "value": 0
            },
            "connqAll.ageEma": {
              "value": 0
            },
            "connqAll.ageHead": {
              "value": 0
            },
            "connqAll.ageMax": {
              "value": 0
            },
            "connqAll.depth": {
              "value": 0
            },
            "connqAll.serviced": {
              "value": 0
            },
            "connq.ageEdm": {
              "value": 0
            },
            "connq.ageEma": {
              "value": 0
            },
            "connq.ageHead": {
              "value": 0
            },
            "connq.ageMax": {
              "value": 0
            },
            "connq.depth": {
              "value": 0
            },
            "connq.serviced": {
              "value": 0
            },
            "curSessions": {
              "value": 0
            },
            "minActiveMembers": {
              "value": 0
            },
            "monitorRule": {
              "description": "none"
            },
            "tmName": {
              "description": "/Common/test_pool1"
            },
            "serverside.bitsIn": {
              "value": 0
            },
            "serverside.bitsOut": {
              "value": 0
            },
            "serverside.curConns": {
              "value": 0
            },
            "serverside.maxConns": {
              "value": 0
            },
            "serverside.pktsIn": {
              "value": 0
            },
            "serverside.pktsOut": {
              "value": 0
            },
            "serverside.totConns": {
              "value": 0
            },
            "status.availabilityState": {
              "description": "unknown"
            },
            "status.enabledState": {
              "description": "enabled"
            },
            "status.statusReason": {
              "description": "The children pool member(s) either don't have service checking enabled, or service check results are not available yet"
            },
            "totRequests": {
              "value": 0
            }
          }
        }
      },
      "https://localhost/mgmt/tm/ltm/pool/~Common~test_pool2/stats": {
        "nestedStats": {
          "kind": "tm:ltm:pool:poolstats",
          "selfLink": "https://localhost/mgmt/tm/ltm/pool/~Common~test_pool2/stats",
          "entries": {
            "activeMemberCnt": {
              "value": 0
            },
            "connqAll.ageEdm": {
              "value": 0
            },
            "connqAll.ageEma": {
              "value": 0
            },
            "connqAll.ageHead": {
              "value": 0
            },
            "connqAll.ageMax": {
              "value": 0
            },
            "connqAll.depth": {
              "value": 0
            },
            "connqAll.serviced": {
              "value": 0
            },
            "connq.ageEdm": {
              "value": 0
            },
            "connq.ageEma": {
              "value": 0
            },
            "connq.ageHead": {
              "value": 0
            },
            "connq.ageMax": {
              "value": 0
            },
            "connq.depth": {
              "value": 0
            },
            "connq.serviced": {
              "value": 0
            },
            "curSessions": {
              "value": 0
            },
            "minActiveMembers": {
              "value": 0
            },
            "monitorRule": {
              "description": "none"
            },
            "tmName": {
              "description": "/Common/test_pool2"
            },
            "serverside.bitsIn": {
              "value": 0
            },
            "serverside.bitsOut": {
              "value": 0
            },
            "serverside.curConns": {
              "value": 0
            },
            "serverside.maxConns": {
              "value": 0
            },
            "serverside.pktsIn": {
              "value": 0
            },
            "serverside.pktsOut": {
              "value": 0
            },
            "serverside.totConns": {
              "value": 0
            },
            "status.availabilityState": {
              "description": "unknown"
            },
            "status.enabledState": {
              "description": "enabled"
            },
            "status.statusReason": {
              "description": "The children pool member(s) either don't have service checking enabled, or service check results are not available yet"
            },
            "totRequests": {
              "value": 0
            }
          }
        }
      }
    }
  }`
	// ResponseMap contains example responses for the mock endpoint
	ResponseMap = map[string]string{
		"node":    nodeResponse,
		"virtual": virtualResponse,
		"pool":    poolResponse,
		"rule":    ruleResponse,
	}
)
