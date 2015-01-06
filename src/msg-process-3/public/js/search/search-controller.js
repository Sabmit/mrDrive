var app = angular.module('server');

app.controller('SearchController', ['$scope', '$http', function ($scope, $http) {
    host = "http://" + window.location.host;
    mthis = this;
    mthis.Keyword = {};
    mthis.Found = false;
    mthis.KeywordForm = '';
    mthis.notFoundWord = '';
    mthis.total_used = 0;

    mthis.getKeywordData = function() {
        $http.get(host + '/api/keywords/' + mthis.KeywordForm)
            .success(function(response) {
                if (response.Keyword != '') {
                    mthis.total_used = 0;
                    mthis.Found = true;
                    mthis.Keyword = response;
                    for (var i = 0, len = mthis.Keyword.Ips.length; i < len; i++) {
                        mthis.total_used += mthis.Keyword.Ips[i]["used"];
                    }
                    mthis.notFoundWord = '';
                } else {
                    mthis.Found = false;
                    mthis.notFoundWord = mthis.KeywordForm;
                }
            });
    }
}]);
