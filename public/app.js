var app = angular.module('Nomen', ['ngRoute']);

app.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider
      .when('/ids', { // List ids
        templateUrl: '/views/ids_list.html',
        controller: 'IdsCtrl'
      })
      .when('/id/:identifier', { // Show id
        templateUrl: '/views/id.html',
        controller: 'IdCtrl'
      })
      .when('/domains', { // List domains
        templateUrl: '/views/domains_list.html',
        controller: 'DomainsCtrl'
      })
      .when('/d/:identifier', { // Show domain
        templateUrl: '/views/domain.html',
        controller: 'DomainCtrl'
      })
      .when('/expired', { // List expired names
        templateUrl: '/views/expired.html',
        controller: 'ExpiredCtrl'
      })
      .when('/search', {
        templateUrl: '/views/search.html',
        controller: 'SearchCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
}]);
