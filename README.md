# Monitoring

Simple monitoring tool with web application written in Go.


## Service checking and notifying

Uses JSON file(s) instead of a MySQL so nothing can break monitoring except purposefully human failure. Even if you delete the json file, monitoring will continue running! These files are used to load checks to be run at really low intervals.

## Statistics gathering

This module has been seperated in this branch. Please see [this branch](https://github.com/mdeheij/monitoring/tree/StatisticsCombined) for a combined version.

Uses a small bash script as a client on a remote server to gather metrics and send them to your monitoring instance. Multiple servers can be supervised easily with clear visualisations. Frontend is based on AngularJS, jQuery, D3 and UIKit.

## Features

-   Thresholds
-   Acknowledging
-   Notifications (e.g. Telegram)
-   Action handling
-   Protection against nuclear bombs and zombie apocalypses

## Screenshots

### Services
![services](http://i.imgur.com/tD9N9l6.png "Services Screenshot")

### Statistics

![stats](http://i.imgur.com/7S2sfvb.png "Statistics Screenshot")

## Note

Recommended to be used with a reverse proxy like nginx
