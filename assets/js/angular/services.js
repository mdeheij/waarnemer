running = true;
//DEBUG-console.log(running);
var app = angular.module('myApp', ['ngOrderObjectBy']);
var Timer;
app.controller('customersCtrl', function($scope, $http) {
      $http.get("list.json")
            .success(function(response) {
                  $scope.services = response;
            });
      $scope.start = function() {
            //DEBUG-console.log("Herkansing!");
            var response = $http.get('start');
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Started checking!");
                  $scope.statusMsg = "Started!";
                  UIkit.notify("<i class='uk-icon-check'></i> Started");
            });
            response.error(function(data, status, headers, config) {
                  alert("Error.");
            });
      };
      $scope.stop = function() {

            //DEBUG-console.log("Herkansing!");
            var response = $http.get('stop');
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Stopped checking!");
                  $scope.statusMsg = "Stopped!";
                  UIkit.notify("<i class='uk-icon-check'></i> Stopped checking");
            });
            response.error(function(data, status, headers, config) {
                  alert("Error.");
            });
      };
      $scope.updateList = function() {

            //DEBUG-console.log("updateList!");
            var response = $http.get('updatelist');
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("updateList Success!");
                  $scope.statusMsg = "updateList is uitgevoerd!";
                  UIkit.notify("<i class='uk-icon-check'></i> Updated all checks");
            });
            response.error(function(data, status, headers, config) {
                  alert("Error.");
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
            var response = $http.get('reschedule/' + identifier + "/");
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Rescheduled " + identifier + "!");
                  $scope.statusMsg = "Rescheduled " + identifier + "!";
                  UIkit.notify("<i class='uk-icon-check'></i> Rescheduled checking for " + identifier);
            });
            response.error(function(data, status, headers, config) {
                  alert("Error.");
            });
      };
      $scope.updateCheck = function(identifier) {

            //DEBUG-console.log("Updating!");
            var response = $http.get('update/' + identifier + "/");
            response.success(function(data, status, headers, config) {
                  //DEBUG-console.log("Updated " + identifier + "!");
                  $scope.statusMsg = "Updated " + identifier + "!";
                  UIkit.notify("<i class='uk-icon-check'></i> Updated " + identifier);
            });
            response.error(function(data, status, headers, config) {
                  alert("Error.");
            });
      };

      $scope.mySplit = function(string, nb) {
            $scope.array = string.split(' ');
            return $scope.result = $scope.array[nb];
      }
      $scope.refresh = function() {
            /* //DEBUG-console.log("refresh!");*/
            $http.get("list.json").success(function(response) {
                  $scope.services = response;
            });
      };
      $scope.resetPolling = function() {
            //UIkit.notify("<i class='uk-icon-check'></i> Polling time is set to " + $scope.pollingInterval);
            //DEBUG-console.log("resetPolling -- consolelog - " + $scope.pollingInterval);
            $scope.timerActive = true;
            clearInterval(Timer);
            Timer = setInterval(refreshTable, $scope.pollingInterval)
      };

      function refreshTable() {
            //DEBUG-console.log("refresh!");
            $('#refreshToolbarIcon').addClass("fa-spin uk-text-warning");
            $http.get("list.json").success(function(response) {
                  $scope.services = response;
                  $('#refreshToolbarIcon').removeClass("fa-spin uk-text-warning");
            });
      };
      $scope.pollingInterval = 2000;
      $scope.resetPolling();
      //var Timer=setInterval(refreshTable,$scope.pollingInterval);
      //DEBUG-console.log($scope.pollingInterval);
});
//oh it does
/*UIkit.notify({
      message: 'Polling does not automatically start!',
      status: 'info',
      timeout: 1500,
      pos: 'top-center'
});*/
