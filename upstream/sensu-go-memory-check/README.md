[![Bonsai Asset Badge](https://img.shields.io/badge/Sensu%20Go%20Memory%20Checks-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/asachs01/sensu-go-memory-check) [![TravisCI Build Status](https://travis-ci.org/asachs01/sensu-go-memory-check.svg?branch=master)](https://travis-ci.org/asachs01/sensu-go-memory-check)

# Sensu Go Memory Check

- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Asset definition](#asset-definition)
  - [Check definition](#resource-definition)
- [Installation from source and contributing](#installation-from-source-and-contributing)
- [Additional notes](#additional-notes)

## Overview

This plugin provides a check for system memory utilization for Sensu Go. The `sensu-go-memory-check` check takes the flags `-w` (warning) and `-c` (critical) and a desired  duration utilization percentage after each flag. By default, these are a warning value of 75% and a critical value of 90%. This check also outputs data as `nagios_perfdata`(for more information, see [this Nagios documentation article](https://assets.nagios.com/downloads/nagioscore/docs/nagioscore/3/en/perfdata.html). This allows for the check to be used as both a status check and a metric check. You can see an example of this in the [example check definition](#check-definition) below.

## Usage Examples

### Command line help

```
The Sensu Go check for system memory usage

Usage:
  sensu-go-memory-check [flags]

Flags:
  -c, --critical string   Critical used percentage for system memory (default "90")
  -h, --help              help for sensu-go-memory-check
  -w, --warning string    Warning used percentage for system memory. (default "75")

```

### Example Output

```bash
./sensu-go-memory-check
CheckMem OK - value = 63.84 | system_memory_used=63.84
```

## Configuration

### Asset Registration

Assets are the best way to make use of this check. If you're not using an asset, please consider doing so! If you're using Sensu 5.13 or later, you can install this plugin as an asset by running:

`sensuctl asset add asachs01/sensu-go-memory-check`

Else, you can find this asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/asachs01/sensu-go-memory-check).

### Asset definition

You can download the asset definition there, or you can do a little bit of copy/pasta and use the one below:

```yml
---
type: Asset
api_version: core/v2
metadata:
  name: sensu-go-memory-check
  namespace: CHANGEME
  labels: {}
  annotations: {}
spec:
  url: https://github.com/asachs01/sensu-go-memory-check/releases/download/0.0.1/sensu-go-memory-check_0.0.1_linux_amd64.tar.gz
  sha512: 
  filters:
  - entity.system.os == 'linux'
  - entity.system.arch == 'amd64'
```

**NOTE**: PLEASE ENSURE YOU UPDATE YOUR URL AND SHA512 BEFORE USING THE ASSET. If you don't, you might just be stuck on a super old version. Don't say I didn't warn you ¯\\_(ツ)_/¯

### Check definition

Example Sensu Go definition:

**sensu-go-memory-check**
```yml
type: CheckConfig
api_version: core/v2
metadata:
  name: sensu-go-memory-check
  namespace: CHANGEME
spec:
  command: sensu-go-memory-check
  runtime_assets:
  - asachs01/sensu-go-memory-check
  interval: 60
  publish: true
  output_metric_format: nagios_perfdata
  output_metric_handlers:
  - influxdb
  handlers:
  - slack
  subscriptions:
  - system
```

## Installation from source and contributing

While it's generally recommended to use an asset, you can download a copy of the handler plugin from [releases][1],
or create an executable script from this source.

From the local path of the sensu-go-memory-check repository:

**sensu-go-memory-check**
```
go build -o /usr/local/bin/sensu-go-memory-check main.go
```

To contribute to this repo, see https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md.

## Additional notes

### Supported Operating Systems

This project uses `gopsutil`, and is thus largely dependent on the systems that it supports. For this plugin, the following operating systems are supported:

* Linux
* FreeBSD
* OpenBSD
* Mac OS X
* Windows
* Solaris

[1]: https://github.com/asachs01/sensu-go-memory-check/releases
