var app = angular.module("Nomen");

// Balance
app.factory("Balance", ['$http', function($http) {
  var myBalance = 0.0;
  return {
    getBalance:function(){
      return myBalance;
    },
    refreshBalance:function(callback) {
      $http.get('/api/balance').success(function(data) {
        myBalance = data.balance;
        if (callback) {
          callback(myBalance);
        }
      });
    }
  };
}]);

app.controller("BalanceCtrl", ['$scope', 'Balance', function($scope, Balance) {
  $scope.refreshBalance = function() {
    Balance.refreshBalance(function(bal) {
      $scope.myBalance = bal;
    });
    //$scope.myBalance = Balance.getBalance();
  };
}]);

// Ids List
app.factory("Ids", ['$http', function($http) {
  return $http.get('/api/ids');
}]);

app.controller("IdsCtrl", ['$scope', '$http', 'Ids', function($scope, $http, Ids) {
  $scope.getIds = function() {
    Ids.success(function(data) {
      $scope.myIds = data;
    });
  };
}]);

// Id
app.controller("IdCtrl", ['$scope', '$http', '$routeParams', function($scope, $http, $routeParams) {
  $scope.initId = function() {
    $http.get('/api/ids/' + $routeParams.identifier)
      .success(function(data) {
        $scope.activeId = data;
      });
  };
}]);

// Domains List
app.factory("Domains", ['$http', function($http) {
  return $http.get('/api/domains');
}]);

app.controller("DomainsCtrl", ['$scope', '$http', 'Domains', function($scope, $http, Domains) {
  $scope.getDomains = function() {
    Domains.success(function(data) {
      $scope.myDomains = data;
    });
  };
}]);

// Domain
app.controller("DomainCtrl", ['$scope', '$http', '$routeParams', function($scope, $http, $routeParams) {
  $scope.initDomain = function() {
    $http.get('/api/domains/' + $routeParams.identifier)
      .success(function(data) {
        $scope.activeDomain = data;
      });
  };
}]);

// Expired
app.factory("Expired", ['$http', function($http) {
  return $http.get('/api/expired');
}]);

app.controller("ExpiredCtrl", ['$scope', '$http', 'Expired', function($scope, $http, Expired) {
  $scope.getExpired = function() {
    Expired.success(function(data) {
      $scope.expiredNames = data;
    });
  };
}]);

// Search
app.factory("Search", ['$http', function($http) {
  var searchQuery = "";
  return {
    setQuery:function(query) {
      searchQuery = query;
    },
    getResults:function() {
      return $http.post('/api/search', {regexp: searchQuery});
    }
  };
}]);

app.controller("SearchCtrl", ['$scope', '$location', '$route', 'Search', function($scope, $location, $route, Search) {
  $scope.doSearch = function() {
    Search.getResults().success(function(data) {
      $scope.searchResults = data;
    });
  };
  $scope.doQuery = function() {
    Search.setQuery($scope.searchQuery);
    $location.path('/search');
    $route.reload();
  };
}]);
