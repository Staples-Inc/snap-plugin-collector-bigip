<!-- http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 Staples, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. -->
# snap plugin collector - BigIp

## Collected Metrics
This plugin has the ability to gather the following metrics:

a) **BigIp Node Metrics**

The prefix of metric's namespace is `/staples/bigip/node/<partition>/<node>`

Namespace                                         | Data Type
--------------------------------------------------|-------------
/staples/bigip/node/*/*/cur_sessions              | conn
/staples/bigip/node/*/*/serverside_bytes_in       | B
/staples/bigip/node/*/*/serverside_bytes_out      | B
/staples/bigip/node/*/*/serverside_cur_conns      | conn
/staples/bigip/node/*/*/serverside_max_conns      | conn
/staples/bigip/node/*/*/serverside_pkts_in        | pckt
/staples/bigip/node/*/*/serverside_pkts_out       | pckt
/staples/bigip/node/*/*/serverside_tot_conns      | conn
/staples/bigip/node/*/*/status_availability_state | availability
/staples/bigip/node/*/*/tot_requests              | req

b) **BigIp Pool Metrics**

The prefix of metric's namespace is `/staples/bigip/pool/<partition>/<pool>`

Namespace                                         | Data Type
--------------------------------------------------|----------
/staples/bigip/pool/*/*/active_member_cnt         | counter
/staples/bigip/pool/*/*/connq_age_edm             | s
/staples/bigip/pool/*/*/connq_age_ema             | s
/staples/bigip/pool/*/*/connq_age_head            | s
/staples/bigip/pool/*/*/connq_age_max             | s
/staples/bigip/pool/*/*/connq_all_age_edm         | s
/staples/bigip/pool/*/*/connq_all_age_ema         | s
/staples/bigip/pool/*/*/connq_all_age_head        | s
/staples/bigip/pool/*/*/connq_all_age_max         | s
/staples/bigip/pool/*/*/connq_all_depth           | conn
/staples/bigip/pool/*/*/connq_all_serviced        | conn
/staples/bigip/pool/*/*/connq_depth               | s
/staples/bigip/pool/*/*/connq_serviced            | conn
/staples/bigip/pool/*/*/cur_sessions              | conn
/staples/bigip/pool/*/*/min_active_members        | counter
/staples/bigip/pool/*/*/serverside_bytes_in       | B
/staples/bigip/pool/*/*/serverside_bytes_out      | B
/staples/bigip/pool/*/*/serverside_cur_conns      | conn
/staples/bigip/pool/*/*/serverside_max_conns      | conn
/staples/bigip/pool/*/*/serverside_pkts_in        | pckt
/staples/bigip/pool/*/*/serverside_pkts_out       | pckt
/staples/bigip/pool/*/*/serverside_tot_conns      | conn
/staples/bigip/pool/*/*/status_availability_state | bool
/staples/bigip/pool/*/*/tot_requests              | req

c) **BigIp Rule Metrics**

The prefix of metric's namespace is `/staples/bigip/rule/<partition>/<rule>/<event>/`

Namespace                                  | Data Type
-------------------------------------------|----------
/staples/bigip/rule/*/*/*/aborts           | event
/staples/bigip/rule/*/*/*/avg_cycles       | event
/staples/bigip/rule/*/*/*/failures         | event
/staples/bigip/rule/*/*/*/max_cycles       | event
/staples/bigip/rule/*/*/*/min_cycles       | event
/staples/bigip/rule/*/*/*/priority         | event
/staples/bigip/rule/*/*/*/total_executions | event


d) **BigIp VS Metrics**

The prefix of metric's namespace is `/staples/bigip/rule/<partition>/<rule>/<event>/`

Namespace                                            | Data Type
-----------------------------------------------------|-------------
/staples/bigip/vs/*/*/clientside_bytes_in            | B
/staples/bigip/vs/*/*/clientside_bytes_out           | B
/staples/bigip/vs/*/*/clientside_cur_conns           | conn
/staples/bigip/vs/*/*/clientside_evicted_conns       | conn
/staples/bigip/vs/*/*/clientside_max_conns           | conn
/staples/bigip/vs/*/*/clientside_pkts_in             | pckt
/staples/bigip/vs/*/*/clientside_pkts_out            | pckt
/staples/bigip/vs/*/*/clientside_slow_killed         | conn
/staples/bigip/vs/*/*/clientside_tot_conns           | conn
/staples/bigip/vs/*/*/cs_max_conn_dur                | s
/staples/bigip/vs/*/*/cs_mean_conn_dur               | s
/staples/bigip/vs/*/*/cs_min_conn_dur                | s
/staples/bigip/vs/*/*/ephemeral_bytes_in             | B
/staples/bigip/vs/*/*/ephemeral_bytes_out            | B
/staples/bigip/vs/*/*/ephemeral_cur_conns            | B
/staples/bigip/vs/*/*/ephemeral_evicted_conns        | conn
/staples/bigip/vs/*/*/ephemeral_max_conns            | conn
/staples/bigip/vs/*/*/ephemeral_pkts_in              | pckt
/staples/bigip/vs/*/*/ephemeral_pkts_out             | pckt
/staples/bigip/vs/*/*/ephemeral_slow_killed          | conn
/staples/bigip/vs/*/*/ephemeral_tot_conns            | conn
/staples/bigip/vs/*/*/five_min_avg_usage_ratio       | event
/staples/bigip/vs/*/*/five_sec_avg_usage_ratio       | event
/staples/bigip/vs/*/*/one_min_avg_usage_ratio        | event
/staples/bigip/vs/*/*/status_availability_state      | availability
/staples/bigip/vs/*/*/syncookie_accepts              | event
/staples/bigip/vs/*/*/syncookie_hw_accepts           | event
/staples/bigip/vs/*/*/syncookie_hw_syncookies        | event
/staples/bigip/vs/*/*/syncookie_hwsyncookie_instance | event
/staples/bigip/vs/*/*/syncookie_rejects              | event
/staples/bigip/vs/*/*/syncookie_swsyncookie_instance | event
/staples/bigip/vs/*/*/syncookie_syncache_curr        | event
/staples/bigip/vs/*/*/syncookie_syncache_over        | event
/staples/bigip/vs/*/*/syncookie_syncookies           | event
/staples/bigip/vs/*/*/tot_requests                   | req
