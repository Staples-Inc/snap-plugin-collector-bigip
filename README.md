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
[![Build Status](https://travis-ci.org/Staples-Inc/snap-plugin-collector-bigip.svg?branch=master)](https://travis-ci.org/Staples-Inc/snap-plugin-collector-bigip)
# Snap collector plugin - BigIp

This plugin collects BigIp statistics using the IControl rest API.

It's used in the [Snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
3. [Community Support](#community-support)
4. [License](#license)

## Getting Started
### Operating systems
* Linux/amd64
* Darwin/amd64

### Installation
#### To build the plugin binary:
Fork https://github.com/Staples-Inc/snap-plugin-collector-bigip
Clone repo into `$GOPATH/src/github.com/Staples-Inc/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-bigip.git
```

Build the Snap bigip plugin by running make within the cloned repo:
```
$ make
```
It may take a while to pull dependencies if you haven't had them already.
This builds the plugin in `./build/<GOOS>/<GOARCH>/snap-plugin-collector-bigip`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Load the plugin and create a task, see example in [Examples](#examples).

#### Configuration parameters
It's possible to provide configuration to plugin via task manifest.

    workflow:
      collect:
        config:
          /staples/bigip:
            host: "<BIGIP-HOST>"
            port: <BIGIP-PORT>
            username: "user"
            password: "pass"
            basic_auth: <BOOL>

## Documentation
### Collected Metrics
The list of collected metrics is described in [METRICS.md](METRICS.md).

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/Staples-Inc/snap-plugin-collector-bigip/issues) and/or submit a [pull request](https://github.com/Staples-Inc/snap-plugin-collector-bigip/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [Slack](http://slack.snap-telemetry.io).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).
