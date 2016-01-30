//var scope.stats;
var Timer;
var h1;
google.charts.setOnLoadCallback(function() {
      angular.bootstrap(document.body, ['app']);
});
var app = angular.module('app', []);
google.charts.load('current', {
      'packages': ['line', 'corechart']
});


app.controller('MainCtrl', function($scope, $http, $location) {
      var searchObject = $location.search();
      $scope.graphOptions = {
            chart: {
                  curveType: 'function'
            },
            backgroundColor: {
                  fill: 'transparent'
            },
            chartArea: {
                  backgroundColor: {
                        fill: 'transparent'
                  }
            },
            height: 220,
            axes: {
                  x: {
                        0: {
                              side: 'top'
                        }
                  }
            }
      };
      $scope.charts = [];
      $scope.chartWidth = 900;
      $scope.chartHeight = 320;
      $scope.chartData = [];
      $scope.chartData['traffic'] = [
            [0, 6, 3],
            [1, 1, 3],
            [2, 8, 2],
            [3, 3, 7],
            [4, 1, 13],
            [5, 9, 5],
            [6, 9, 4]
      ];
      //$scope.chartData['traffic']  = getStatistics($scope, "traffic")
      $scope.chartData['ping'] = [
            [0, 70],
            [1, 11],
            [2, 4],
            [2.3, 4],
            [3, 4],
            [4, 7],
            [5, 7],
            [6, 9]
      ];
      $scope.chartData['pie'] = [
            ['Ad Flyers', 33],
            ['Web (Organic)', 4],
            ['Web (PPC)', 4],
            ['Yellow Pages', 7],
            ['Showroom', 3]
      ];
      $scope.bouwGrafiek = function() {
            UIkit.notify("<i class='uk-icon-check'></i> Grafiek gebouwd! ");
      };
      $scope.processToPie = function() {
            var tempArray = [];
            angular.forEach($scope.processesng, function(row) {
                  tempArray.push([row.proc, parseFloat(row.rss)]);
            });
            return tempArray;
      }
      $scope.processToPieCPU = function() {
            var tempArray = [];
            angular.forEach($scope.processesng, function(row) {
                  tempArray.push([row.proc, parseFloat(row.cpu)]);
            });
            return tempArray;
      }

      function draw(target) {
            if (target != undefined && $scope.charts[target] != undefined) {
                  var data = $scope.charts[target].datatable;
                  var label, value;
                  data.removeRows(0, data.getNumberOfRows());
                  angular.forEach($scope.chartData[target], function(row) {
                        /*label = parseFloat(row[0]);
                        value = parseFloat(row[1], 10);
                        if (!isNaN(row[3])) {
                        data.addRow([row[0], value, row[2], row[3]]);
                        } else if (!isNaN(row[2])) {
                        data.addRow([row[0], value, row[2]]);
                        } else if (!isNaN(value)) {
                        data.addRow([row[0], value]);
                        }*/
                        data.addRow(row);
                  });
                  var hoogte;
                  if ($scope.charts[target].height != undefined) {
                        hoogte = $scope.charts[target].height;
                  } else {
                        hoogte = 160;
                  }
                  var options = {
                        'width': $scope.chartWidth,
                        'height': hoogte,
                        backgroundColor: {
                              fill: 'transparent'
                        },
                        curveType: 'function',
                        legend: {
                              position: 'bottom'
                        },
                        chartArea: {
                              left: 30,
                              right: 10, // !!! works !!!
                              bottom: 30, // !!! works !!!
                              top: 30,
                              backgroundColor: {
                                    fill: 'transparent'
                              }
                        }
                  };
                  $scope.charts[target].graph.draw(data, options);
            }
      }
      $scope.updateGraphs = function() {
            draw("traffic");
            draw("loadcpu");
            //draw("loadio");
            draw("load");
            //DEBUG-console.log("[updateGraphs] -> 13:10 ");
            //DEBUG-console.log(getStatistics($scope, "load"));
            draw("ram");
            draw("disk");
            draw("ping");
            draw("pie");
            draw("piecpu");
      }
      $scope.$watch('chartData.traffic', function() {
            $scope.updateGraphs()
      }, true);
      $scope.chartWidth = $('#traffic').parent().width() - 20;
      window.onresize = function() {
            $scope.chartWidth = $('#traffic').parent().width() - 20;
            $scope.updateGraphs();
      }
      setInterval(function() {
            var secs = new Date();
            var timeDiff = (secs - $scope.lastUpdateDate) / 1000;
            var nextUpdate = Math.round($scope.stats[0].Frequency + 2 - timeDiff);
            if (nextUpdate > 0) {
                  $("#nextUpdate").html(nextUpdate);
            } else if (nextUpdate <= 0 && nextUpdate > -10) {
                  $("#nextUpdate").html("Expecting..");
                  $scope.getData();
            } else {
                  $("#nextUpdate").html("<strong>Not responding!</strong>");
            }
      }, 1000);
      $scope.buildGraphs = function() {
            // Create the data table and instantiate the chart
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', '▼ RX (MB)');
            data.addColumn('number', '▲ TX (MB)');
            $scope.charts['traffic'] = [];
            $scope.charts['traffic'].datatable = data;
            $scope.charts['traffic'].graph = new google.visualization.LineChart(document.getElementById('traffic'));
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', 'Latency (ms)');
            $scope.charts['ping'] = [];
            $scope.charts['ping'].datatable = data;
            $scope.charts['ping'].height = 180;
            $scope.charts['ping'].graph = new google.visualization.LineChart(document.getElementById('ping'));
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', 'Load %');
            $scope.charts['loadcpu'] = [];
            $scope.charts['loadcpu'].datatable = data;
            $scope.charts['loadcpu'].graph = new google.visualization.LineChart(document.getElementById('loadcpu'));
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', 'Load %');
            $scope.charts['loadio'] = [];
            $scope.charts['loadio'].datatable = data;
            $scope.charts['loadio'].graph = new google.visualization.LineChart(document.getElementById('loadio'));
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', 'Median');
            data.addColumn('number', 'Short');
            data.addColumn('number', 'Mid');
            data.addColumn('number', 'Long');
            $scope.charts['load'] = [];
            $scope.charts['load'].datatable = data;
            $scope.charts['load'].graph = new google.visualization.LineChart(document.getElementById('load'));
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', 'Total');
            data.addColumn('number', 'Usage');
            $scope.charts['ram'] = [];
            $scope.charts['ram'].datatable = data;
            $scope.charts['ram'].graph = new google.visualization.AreaChart(document.getElementById('ram'));
            var data = new google.visualization.DataTable();
            data.addColumn('datetime', 'Time');
            data.addColumn('number', 'Zero');
            data.addColumn('number', 'Usage (GB)');
            data.addColumn('number', 'Total (GB)');
            $scope.charts['disk'] = [];
            $scope.charts['disk'].datatable = data;
            $scope.charts['disk'].graph = new google.visualization.AreaChart(document.getElementById('disk'));
            var data = new google.visualization.DataTable();
            data.addColumn('string', 'Label');
            data.addColumn('number', 'Value');
            $scope.charts['piecpu'] = [];
            $scope.charts['piecpu'].datatable = data;
            $scope.charts['piecpu'].height = 300;
            $scope.charts['piecpu'].graph = new google.visualization.PieChart(document.getElementById('piecpu'));
            var data = new google.visualization.DataTable();
            data.addColumn('string', 'Label');
            data.addColumn('number', 'Value');
            $scope.charts['pie'] = [];
            $scope.charts['pie'].datatable = data;
            $scope.charts['pie'].height = 350;
            $scope.charts['pie'].graph = new google.visualization.PieChart(document.getElementById('pie'));
      }
      $scope.init = function() {
            $http.get("/stats/server/list").then(function(response) {
                  $scope.servers = response.data;
                  if (searchObject['hostID'] == undefined) {
                        $scope.changeHost($scope.servers[0].HostID);
                  } else {
                        $scope.changeHost(searchObject['hostID']);
                  }
                  $scope.buildGraphs();
            });
      }
      $scope.changeHost = function(hostID) {
            if (hostID != undefined) {
                  $('#contentGrid').hide();
                  //DEBUG-console.log("[changeHost] Hij is defined naar " + hostID + ", dus setten");
                  $scope.selectedHost = hostID;
                  $location.search('hostID', $scope.selectedHost);
                  $scope.getData();
            } else {
                  alert("Undefined hostID!");
            }
      }
      $scope.init();
      $scope.pollingInterval = 2000;
      //DEBUG-console.log($location.search());
      //DEBUG-console.log($location.search().id);
      $scope.getData = function() {
            //DEBUG-console.log("Executed getData()");
            $('#refreshToolbarIcon').addClass("fa-spin uk-text-warning");
            $http.get("/stats/server/view/" + $scope.selectedHost).then(function(response) {
                  if(response.data.obj[10] != undefined) {
                        $('#refreshToolbarIcon').removeClass("fa-spin uk-text-warning");
                        $('#contentGrid').show();
                        $scope.lastUpdateDate = new Date(response.data.obj[0].Date * 1000);
                        $scope.timehuman = new Date(response.data.obj[0].Date * 1000);
                        $scope.processes = response.data.obj[0].Processesarray.split(";");
                        $scope.processesng = processesSplit(response.data.obj[0].Processesarray);
                        $scope.selectedHostIdentifier = response.data.identifier;
                        //DEBUG-console.log("GRAFEGESRGE"+response.data.identifier);
                        //DEBUG-console.log("OKEE");
                        //DEBUG-console.log(response.data);
                        $scope.stats = response.data.obj;
                        $scope.latest = response.data.obj[0];
                        //DEBUG-console.log("getdatadebug");
                        //DEBUG-console.log($scope.chartData['traffic']);
                        //DEBUG-console.log(getStatistics($scope, "traffic"));
                        //DEBUG-console.log("/getdatadebug");
                        $scope.chartData['traffic'] = getStatistics($scope, "traffic");
                        $scope.chartData['loadcpu'] = getStatistics($scope, "loadcpu");
                        $scope.chartData['loadio'] = getStatistics($scope, "loadio");
                        $scope.chartData['load'] = getStatistics($scope, "load");
                        $scope.chartData['disk'] = getStatistics($scope, "disk");
                        $scope.chartData['ram'] = getStatistics($scope, "ram");
                        $scope.chartData['ping'] = getStatistics($scope, "ping");
                        $scope.chartData['pie'] = $scope.processToPie();
                        $scope.chartData['piecpu'] = $scope.processToPieCPU();
                        $scope.updateGraphs();
                        $http.get("/services/list/" + $scope.selectedHostIdentifier).then(function(response) {
                              //DEBUG-console.log("Got services!");
                              //DEBUG-console.log(response);
                              $scope.services = response.data;
                        });
                  } else {
                        UIkit.notify("Sorry, we do not have enough data yet to show some fancy graphs.");
                  }
                  //pingGraph($scope);
                  //loadCpuGraph($scope);
                  //networkGraph($scope);
            });
      }
      processesSplit = function($rawList) {
            processList = [];
            lines = $rawList.split(';').reverse();
            lineNumber = 0;
            for (var i = lines.length - 1; i >= 0; i--) {
                  l = lines[i];
                  lineNumber++;
                  data = l.split(' ');
                  var user = data[0];
                  var cpu = data[1];
                  var rss = data[2];
                  var proc = data[3];
                  processList.push({
                        user: user,
                        cpu: cpu,
                        rss: rss,
                        proc: proc
                  });
            }
            for (var i = processList.length; i--;) {
                  if (processList[i].proc == proc) {
                        //DEBUG-console.log("Duplicate for " + proc + " == " + processList[i].proc + " on " + i);
                  }
            }
            return processList;
      }
      $scope.resetPolling = function() {
            UIkit.notify("<i class='uk-icon-check'></i> Reset polling time to " + $scope.pollingInterval);
            //DEBUG-console.log("resetPolling -- consolelog - " + $scope.pollingInterval);
            $scope.timerActive = true;
            clearInterval(Timer);
            Timer = setInterval($scope.getData, $scope.pollingInterval)
      };
      $scope.direction = false;

      function getStatistics($scope, attr) {
            var tempArray = [];
            var historycount = 25;
            for (var i = 0; i < historycount; i++) {
                  var c = historycount - i - 1;
                  var timestamp = new Date($scope.stats[i].Date * 1000);
                  switch (attr) {
                        case "traffic":
                              tempArray.push([timestamp, parseFloat($scope.stats[i].Rxdiff / 1024 / 1024), parseFloat($scope.stats[i].Txdiff / 1024 / 1024)]);
                              break;
                        case "loadcpu":
                              tempArray.push([timestamp, parseFloat($scope.stats[i].Loadcpu)]);
                              break;
                        case "loadio":
                              tempArray.push([timestamp, parseFloat($scope.stats[i].Loadio)]);
                              break;
                        case "load":
                              loadNumber = $scope.stats[i].Load.split(' ');
                              tempArray.push([timestamp, 1, parseFloat(loadNumber[0]), parseFloat(loadNumber[1]), parseFloat(loadNumber[2])]);
                              break;
                        case "ram":
                              tempArray.push([timestamp, parseFloat($scope.stats[i].Ramtotal / 1024 / 1024), parseFloat($scope.stats[i].Ramusage / 1024 / 1024)]);
                              break;
                        case "disk":
                              tempArray.push([timestamp, 0, parseFloat($scope.stats[i].Diskusage / 1024 / 1024 / 1024), parseFloat($scope.stats[i].Disktotal / 1024 / 1024 / 1024)]);
                              break;
                        case "ping":
                              tempArray.push([timestamp, parseFloat(Math.round($scope.stats[i].Ping))]);
                              break;
                        default:
                              tempArray.push([c, $scope.stats[i][attr]]);
                  }
            }
            //DEBUG-console.log("ATTR: " + attr);
            //DEBUG-console.log(tempArray);
            return tempArray;
      }
});

function responsiveGraph(container) {
      //$(document).ready(function() {
      $(window).resize(function() {
            var margin = chart.margin();
            var margin = 10;
            width = parseInt(d3.select("#test1").style("width")) - margin * 2,
                  height = margin + 100;
            chart
                  .width(width)
                  .height(height);
            d3.select('#test1 svg')
                  .attr('width', width)
                  .attr('height', height)
                  .call(chart);
      });
}
//});
function henkfietspop() {
      h1_source = 'd2';
      h1.lineChart('setDataHolder', h1_source);
      //h1.setViewBox(0,0,$('#chart1').parent().width(),200,false);
      //DEBUG-console.log($('#chart1').parent().width(), 200);
}
/*
angular.filter('orderObjectBy', function () {
    return function (items, field, reverse) {
        // Build array
        var filtered = [];
        for (var key in items) {
            if (field === 'key')
                filtered.push(key);
            else
                filtered.push(items[key]);
        }
        // Sort array
        filtered.sort(function (a, b) {
            if (field === 'key')
                return (a > b ? 1 : -1);
            else
                return (a[field] > b[field] ? 1 : -1);
        });
        // Reverse array
        if (reverse)
            filtered.reverse();
        return filtered;
    };
});
*/