running = true;
//DEBUG-console.log(running);
var app = angular.module('myApp', ['ngOrderObjectBy']);
var Timer;
var token;

function setCookie(cname, cvalue, exdays) {
    var d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue + "; " + expires;
}

app.controller('customersCtrl', function($scope, $http) {
      $http.get("/api/service/list")
            .success(function(response) {
                  $scope.services = response;
            });

            //CSRF token code
      $http.get("/admin/token")
                  .success(function(response) {
                        console.log("GOT TOKEN! "+response);
                        token = response;
                        setCookie("XSRF-TOKEN", response, 1);
                  });
      $scope.start = function() {
            //DEBUG-console.log("Herkansing!");
            var response = $http.post('/api/service/start');
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Started checking!");
                  $scope.statusMsg = "Started!";
                  notify("<i class='uk-icon-check'></i> Started");
            });
            response.error(function(data, status, headers, config) {
                  alert("Error: "+data);
            });
      };
      $scope.stop = function() {

            //DEBUG-console.log("Herkansing!");
            var response = $http.post('/api/service/stop');
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Stopped checking!");
                  $scope.statusMsg = "Stopped!";
                  notify("<i class='uk-icon-check'></i> Stopped checking");
            });
            response.error(function(data, status, headers, config) {
                  alert("Error: "+data);
            });
      };
      $scope.updateList = function() {

            //DEBUG-console.log("updateList!");
            var response = $http.post('/api/service/updatelist');
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("updateList Success!");
                  $scope.statusMsg = "updateList is uitgevoerd!";
                  notify("<i class='uk-icon-check'></i> Updated all checks");
            });
            response.error(function(data, status, headers, config) {
                  alert("Error: "+data);
            });
      };
      $scope.openDetails = function(x) {
          console.log("Info over X:");
          console.log(x);

            UIkit.modal("#service-info").show();
            $scope.selectedService = x;
      };
      $scope.rescheduleCheck = function(identifier) {

            //DEBUG-console.log("Rescheduling!");
            var response = $http.post('/api/service/reschedule/' + identifier + "/");
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Rescheduled " + identifier + "!");
                  $scope.statusMsg = "Rescheduled " + identifier + "!";
                  notify("<i class='uk-icon-check'></i> Rescheduled checking for " + identifier);
            });
            response.error(function(data, status, headers, config) {
                  alert("Error: "+data);
            });
      };
      $scope.updateCheck = function(identifier) {

            //DEBUG-console.log("Updating!");
            var response = $http.post('/api/service/update/' + identifier + "/");
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Updated " + identifier + "!");
                  $scope.statusMsg = "Updated " + identifier + "!";
                  notify("<i class='uk-icon-check'></i> Updated " + identifier);
            });
            response.error(function(data, status, headers, config) {
                  alert("Error: "+data);
            });
      };

      $scope.mySplit = function(string, nb) {
            $scope.array = string.split(' ');
            return $scope.result = $scope.array[nb];
      }
      $scope.refresh = function() {
            /* //DEBUG-console.log("refresh!");*/
            $http.get("/api/service/list").success(function(response) {
                  $scope.services = response;
            });
      };
      $scope.disableTimer = function() {
            console.log("[disableTimer] refreshing should be disabled now");
            $scope.timerActive = false; //visualise auto refreshing is not active
            clearInterval(Timer);
      };
      $scope.resetPolling = function() {
            notify("<i class='uk-icon-check'></i> Polling time is set to " + $scope.pollingInterval);
            //DEBUG-console.log("resetPolling -- consolelog - " + $scope.pollingInterval);
            $scope.timerActive = true;
            clearInterval(Timer);
            Timer = setInterval(refreshTable, $scope.pollingInterval)
      };

      function refreshTable() {
            //DEBUG-console.log("refresh!");
            $('#refreshToolbarIcon').addClass("fa-spin uk-text-warning");

                $http.get("/api/service/list").success(function(response) {
                      $scope.services = response;
                      $('#refreshToolbarIcon').removeClass("fa-spin uk-text-warning");
                });

      };
      $scope.pollingInterval = 2000;
      $scope.resetPolling();
      //var Timer=setInterval(refreshTable,$scope.pollingInterval);
      //DEBUG-console.log($scope.pollingInterval);
});

function notify(message) {

     //UIkit.notify(message+"sent by notify()");
     UIkit.notify({
           message: message,
           status: 'info',
           timeout: 700,
           pos: 'bottom-right'
     });
}
//oh it does
/*UIkit.notify({
      message: 'Polling does not automatically start!',
      status: 'info',
      timeout: 1500,
      pos: 'top-center'
});*/
