var app = angular.module('server');

app.controller('SearchController', ['$scope', '$http', function ($scope, $http) {
    mthis = this;
    mthis.Keyword = {};
    mthis.KeywordForm = "";

    mthis.getKeywordData = function() {
        $http.get("http://127.0.0.1:8080/api/keywords/" + mthis.KeywordForm)
            .success(function(response) {
                console.log(response);
//                mthis.Keywords = response;
            });
        console.log(mthis.KeywordForm);
    }

    // mthis.getTopKeywords();
    // console.log(mthis.Keywords);
}]);
