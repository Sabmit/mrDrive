var app = angular.module('server');

app.controller('topIpsController', ['$scope', '$http', function ($scope, $http) {
    mthis = this;
    mthis.Ips = [];

    mthis.getTopIps = function()  {
        $http.get("http://127.0.0.1:8080/api/topIps")
            .success(function(response) {
                console.log(response);
                mthis.Ips = response;
            });
    };

    mthis.getTopIps();
}]);
