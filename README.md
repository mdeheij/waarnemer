# Monitoring
<img align="right" src="https://mdeheij.github.io/monitoring_logo.png">
Simple monitoring tool with web application written in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/mdeheij/monitoring)](https://goreportcard.com/report/github.com/mdeheij/monitoring) [![Build Status](https://travis-ci.org/mdeheij/monitoring.svg?branch=master)](https://travis-ci.org/mdeheij/monitoring)

## Features

-   Very lightweight
-   Simple
-   Thresholds
-   Stable
-   Low interval checking
-   Nagios Plugin compatible
-   Acknowledging
-   Notifications (e.g. Telegram)
-   Action handling

## Service checking and notifying

Uses Yaml file(s) instead of a database such as MySQL.

![Animation](https://i.imgur.com/7d44ndT.gif)

## Sample configuration

**a-sample-service.yaml**

```yaml
- identifier: server.sample-service.ping
  host: 127.0.0.1
  command: ping -H $HOST$ -t $TIMEOUT$
  timeout: 2
  interval: 4
  threshold: 3
  enabled: true
  action:
    name: telegram
```

**config.yaml**

```yaml
actions:
    telegram:
        bot: '1234567:AAgsafaSFASFASFAfafasfasfafAsdad'
        target: '90011337'
paths:
    checks: /etc/monitoring/checks
    services: /etc/monitoring/services
api:
    address: "0.0.0.0:4200"
```
