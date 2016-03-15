# Monitoring

Simple monitoring tool with web application written in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/mdeheij/monitoring)](https://goreportcard.com/report/github.com/mdeheij/monitoring) [![Build Status](https://travis-ci.org/mdeheij/monitoring.svg?branch=master)](https://travis-ci.org/mdeheij/monitoring)

## Features

-   Very lightweight
-   Simple
-   CSRF tokens
-   Thresholds
-   Low interval checking
-   Acknowledging
-   Notifications (e.g. Telegram)
-   Action handling
-   Embed functionality (for example in Grafana)
-   Protection against nuclear bombs and zombie apocalypses

## Service checking and notifying

Uses JSON file(s) instead of a database such as MySQL! Nothing can break checking except purposefully human failure. Even if you delete the json file containing service checks, monitoring will still continue running!

![Animation](https://i.imgur.com/7d44ndT.gif)

## Sample configuration

**services.json**

```json
[
    {
        "identifier": "github.web",
        "host": "github.com",
        "command": "curl -H $HOST$ -t $TIMEOUT$ ",
        "timeout": 5,
        "interval": 15,
        "threshold": 3,
        "enabled": true,
        "acknowledged": false,
        "action": {
            "name": "telegram",
            "telegramtarget": [
                9001,
                -1337
            ]
        }
    },
    {
        "identifier": "localhost.ping",
        "host": "localhost",
        "command": "ping -H $HOST$ -t $TIMEOUT$ ",
        "timeout": 5,
        "interval": 15,
        "threshold": 3,
        "enabled": true,
        "acknowledged": false,
        "action": {
            "name": "none"
        }
    }
]
```

**config.json**

```json
{
    "TelegramNotificationTarget": 0,
    "TelegramBotToken": "",
    "BaseFolder": "/etc/monitoring/",
    "ResourceFolder": "/usr/share/monitoring/",
    "ServerAddress": "0.0.0.0",
    "ServerPort": 4200,
    "SecureCookieName": "MonitoringSession",
    "SecureCookie": "ReplaceMeWithSomethingMoreSecure",
    "Users": [
        {
            "Username": "admin",
            "Hash": "$2a$10$Ctop420kek13379001d52"
        }
    ],
    "Cookieconfig": {
        "Path": "/",
        "Domain": "localhost",
        "MaxAge": 86400,
        "Secure": false,
        "HttpOnly": true
    },
}
```
