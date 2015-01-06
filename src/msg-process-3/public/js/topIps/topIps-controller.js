var app = angular.module('server');

app.controller('topIpsController', ['$scope', '$http', function ($scope, $http) {
    host = "http://" + window.location.host;
    mthis = this;
    mthis.Ips = [];

    mthis.getTopIps = function()  {
        $http.get(host + "/api/topIps")
            .success(function(response) {
                console.log(response);
                mthis.Ips = response;
            });
    };

    mthis.getTopIps();
}]);
