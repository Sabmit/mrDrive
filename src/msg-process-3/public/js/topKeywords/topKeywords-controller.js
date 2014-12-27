var app = angular.module('server');

app.controller('topKeywordsController', ['$scope', '$http', function ($scope, $http) {
    mthis = this;
    mthis.Keywords = [];

    mthis.getTopKeywords = function()  {
        $http.get("http://127.0.0.1:8080/api/topKeywords")
            .success(function(response) {
                console.log(response);
                mthis.Keywords = response;
            });
    };

    mthis.getTopKeywords();
}]);
