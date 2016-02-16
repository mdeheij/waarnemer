//var scope.stats;
var Timer;
var h1;

var app = angular.module('myApp', ['ngOrderObjectBy']);



app.controller('MainCtrl', function($scope, $http, $location) {
      var searchObject = $location.search();

      $scope.charts = [];


      $scope.abstractGraph = function (container, columns, type) {
            var chart = c3.generate({
                bindto: container,
                data: {
                  x: 'timestamp',
                  xFormat: '%Y-%m-%d-%H-%M-%S',
                  columns: columns,
                  type : type, // 'line', 'spline', 'step', 'area', 'area-step' are also available to stack
                },
                transition:  {duration: 0},
                legend: {
                  show: false
                },
                axis : {
                    x : {
                        type : 'timeseries',
                        tick: {
                              format: '   %H:%M:%S'
                        },
                    },
                    y: {
                      tick: {
                       format: d3.format(',')
                      }
                    }

                }
            });
      }
      $scope.trafficGraph = function () {
            $scope.abstractGraph('#traffic', [getStatistics($scope, "timestamp"), getStatistics($scope, "RX"), getStatistics($scope, "TX")], 'area-spline');
            $scope.abstractGraph('#disk', [getStatistics($scope, "timestamp"), getStatistics($scope, "Disk Total"), getStatistics($scope, "Disk Usage")], 'area-spline');
            $scope.abstractGraph('#ram', [getStatistics($scope, "timestamp"), getStatistics($scope, "Memory Total"), getStatistics($scope, "Memory Usage")], 'area-spline');
            $scope.abstractGraph('#ping', [getStatistics($scope, "timestamp"), getStatistics($scope, "ping")], 'spline');

            

            $scope.abstractGraph('#processPieCPU', $scope.processToPie(true), 'donut');
            $scope.abstractGraph('#processPieRSS', $scope.processToPie(false), 'donut');
            $scope.abstractGraph('#cpuGauge', [['timestamp', 0], ["CPU", $scope.latest.Loadcpu]], 'gauge');
            $scope.abstractGraph('#loadGauge', [['timestamp', 0], ["Load", $scope.stats[0].Load.split(' ')[0]]], 'gauge');
            $scope.abstractGraph('#memoryGauge', [['timestamp', 0], ["Memory", ($scope.latest.Ramusage/$scope.latest.Ramtotal)*100]], 'gauge');
            $scope.abstractGraph('#cpu', [getStatistics($scope, "timestamp"), getStatistics($scope, "CPU")], 'spline');
            $scope.abstractGraph('#load', [getStatistics($scope, "timestamp"), getStatistics($scope, "median"), getStatistics($scope, "loadShort"), getStatistics($scope, "loadMid"), getStatistics($scope, "loadLong")], 'spline');
            //$scope.abstractGraph('#cpu', [getStatistics($scope, "RX"), getStatistics($scope, "TX")], 'area-spline');
      }
      $scope.processToPie = function(cpu) {
            var tempArray = [['timestamp', 0]];
            angular.forEach($scope.processesng, function(row) {
                  if (cpu) {
                        tempArray.push([row.proc, parseFloat(row.cpu)]);
                   } else {
                        tempArray.push([row.proc, parseFloat(row.rss)]);
                   }
            });
            //DEBUG-console.log(tempArray);
            return tempArray;
      }
      $scope.processToPieCPU = function() {
            var tempArray = [];
            angular.forEach($scope.processesng, function(row) {
                  tempArray.push([row.proc, parseFloat(row.cpu)]);
            });
            return tempArray;
      }


      $scope.buildModal = function() {
            UIkit.modal('#graphContainerModal').show();
      }
      $scope.updateGraphs = function() {
          console.log("[UpdateGraphs] called");
            if ($scope.latest != undefined) {
                console.log("[UpdateGraphs] proceeded, "+$scope.latest.Loadcpu);
                $scope.trafficGraph();
            }
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
                  var percentageUpdate = 100-((nextUpdate / $scope.stats[0].Frequency) * 100);
                  //DEBUG-console.log("percentageUpdate: "+percentageUpdate);
                  $("#nextUpdateBar").width(percentageUpdate+"%");
                  $("#nextUpdateBar").parent().removeClass( "uk-progress-danger" );


            } else if (nextUpdate <= 0 && nextUpdate > -10) {
                  //$("#nextUpdate").html("Expecting..");


                  $scope.getData();
            } else {
                   $("#nextUpdateBar").width("100%");
                  $("#nextUpdateBar").parent().addClass( "uk-progress-danger" );
                  $("#nextUpdate").html("<strong>Not responding!</strong>");
            }
      }, 1000);
      $scope.init = function() {
            $http.get("/stats/server/list").then(function(response) {
                  $scope.servers = response.data;
                  if (searchObject['hostID'] == undefined) {
                        $scope.changeHost($scope.servers[0].HostID);
                  } else {
                        $scope.changeHost(searchObject['hostID']);
                  }
                  //$scope.buildGraphs();
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
                        $scope.stats = response.data.obj;
                        $scope.latest = response.data.obj[0];
                        $scope.updateGraphs();
                        $http.get("/services/list/" + $scope.selectedHostIdentifier).then(function(response) {
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
            var tempArray = [attr];
            var historycount = 25;
            //console.log("UNDEFINEDCHECKERXXMLG: Checked: "+$scope.stats);
            if ($scope.stats != undefined) {
                //console.log("PROCEEDING: Checked: "+$scope.stats);
                for (var i = 0; i < historycount; i++) {
                      var c = historycount - i - 1;

                      var timestamp = new Date($scope.stats[i].Date * 1000);
                      switch (attr) {
                             case "timestamp":
                                  //var formatted = timestamp.toString();
                                  var date = timestamp;
                                  var year = date.getFullYear();
                                  var month = date.getMonth() + 1;
                                  var day = date.getDate();
                                  var hours = date.getHours();
                                  var minutes = date.getMinutes();
                                  var seconds = date.getSeconds();
                                  var formatted = year + "-" + month + "-" + day + "-" + hours + "-" + minutes+ "-" + seconds;
                                  //var formatted = "2015-05-20-05-10-13";
                                  //console.log(formatted);

                                  tempArray.push(formatted);
                                  break;
                            case "traffic":
                                  tempArray.push([timestamp, parseFloat($scope.stats[i].Rxdiff / 1024 / 1024), parseFloat($scope.stats[i].Txdiff / 1024 / 1024)]);
                                  break;
                            case "median":
                                  tempArray.push(1);
                                  break;
                            case "RX":
                                  tempArray.push(parseFloat(Math.round($scope.stats[i].Rxdiff / 1024) / 1024));
                                  break;
                            case "TX":
                                  tempArray.push(parseFloat(Math.round($scope.stats[i].Txdiff / 1024) / 1024));
                                  break;
                            case "CPU":
                                  tempArray.push(parseFloat($scope.stats[i].Loadcpu));
                                  break;
                            case "loadio":
                                  tempArray.push([timestamp, parseFloat($scope.stats[i].Loadio)]);
                                  break;
                            case "loadShort":
                                  loadNumber = $scope.stats[i].Load.split(' ');
                                  tempArray.push(parseFloat(loadNumber[0]));
                                  break;
                            case "loadMid":
                                  loadNumber = $scope.stats[i].Load.split(' ');
                                  tempArray.push(parseFloat(loadNumber[1]));
                                  break;
                            case "loadLong":
                                  loadNumber = $scope.stats[i].Load.split(' ');
                                  tempArray.push(parseFloat(loadNumber[2]));
                                  break;
                            case "Memory Total":
                                  tempArray.push(parseFloat($scope.stats[i].Ramtotal / 1024 / 1024));
                                  break;
                            case "Memory Usage":
                                  tempArray.push(parseFloat($scope.stats[i].Ramusage / 1024 / 1024));
                                  break;
                            case "Disk Usage":
                                  tempArray.push(Math.round(parseFloat($scope.stats[i].Diskusage / 1024 / 1024 / 1024)*10)/10);
                                  break;
                            case "Disk Total":
                                  tempArray.push(Math.round(parseFloat($scope.stats[i].Disktotal / 1024 / 1024 / 1024)*10)/10);
                                  break;
                            case "ping":
                                  tempArray.push(parseFloat(Math.round($scope.stats[i].Ping)));
                                  break;
                            case "IO":
                                  //DEBUG-console.log("IO", parseFloat($scope.stats[i].Loadio));
                                  var FiftyFifty = Math.round(Math.random());
                                  //tempArray.push(parseFloat($scope.stats[i].Loadio+FiftyFifty));
                                  tempArray.push(parseFloat($scope.stats[i].Loadio));
                                  break;
                            default:
                                  tempArray.push([c, $scope.stats[i][attr]]);
                      }
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
