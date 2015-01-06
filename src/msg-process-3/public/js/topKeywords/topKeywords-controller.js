var app = angular.module('server');

app.controller('topKeywordsController', ['$scope', '$http', function ($scope, $http) {
    host = "http://" + window.location.host;
    mthis = this;
    mthis.Keywords = [];

    mthis.getTopKeywords = function()  {
        $http.get(host + "/api/topKeywords")
            .success(function(response) {
                console.log(response);
                mthis.Keywords = response;
            });
    };

    mthis.getTopKeywords();
}]);
