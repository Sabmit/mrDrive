// Declare app level module which depends on filters, and services

var app = angular.module('server', ['ngResource', 'ngRoute', 'ui.bootstrap', 'ui.date', 'server-directives']);

app.config(['$routeProvider', function ($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'views/search/search.html',
            controller: 'SearchController',
            controllerAs: 'searchCtrl'})
        .when('/topKeywords', {
            templateUrl: 'views/topKeywords/topKeywords.html',
            controller: 'topKeywordsController',
            controllerAs: 'topKeyCtrl'})
        .when('/topIps', {
            templateUrl: 'views/topIps/topIps.html',
            controller: 'topIpsController',
            controllerAs: 'topIpsCtrl'})
        .otherwise({redirectTo: '/'});
}]);

var app = angular.module('server-directives', []);
app.directive("leftSidebar", ["$location", function($location) {
    return {
        restrict: 'E',
        templateUrl: "directives/left-sidebar.html",
        controller: function() {
            iniTab = function (titles) {
                        for (var k in titles) {
                            if (titles[k].route === $location.path()) {
                                return +k;}
                        }
                        return 0;
            };

            this.titles = [{"name": 'Search', "route": '/'},
                           {"name": 'Top keywords', "route": '/topKeywords'},
                           {"name": 'Top Ips', "route": '/topIps'}];


            this.tab = iniTab(this.titles);

            this.isSet = function(checkTab) {
                return this.tab === checkTab;
            };

            this.getTitle = function() {
                return this.titles[this.tab].name;
            };

            this.setTab = function(activeTab) {
                this.tab = activeTab;
            };

            this.isActive = function (viewLocation) {
                return viewLocation === $location.path();
            };
        },
        controllerAs: "tab"
    };
}]);
