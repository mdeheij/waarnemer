# Monitoring

Simple monitoring tool with web application written in Go.

## Features

-   Very lightweight
-   Simple
-   Thresholds
-   Low interval checking
-   Acknowledging
-   Notifications (e.g. Telegram)
-   Action handling
-   Protection against nuclear bombs and zombie apocalypses

## Service checking and notifying

Uses JSON file(s) instead of a database such as MySQL! Nothing can break checking except purposefully human failure. Even if you delete the json file containing service checks, monitoring will still continue running! 

![Animation](https://i.imgur.com/7d44ndT.gifv)