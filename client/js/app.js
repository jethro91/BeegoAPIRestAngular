

var App = angular.module('app', [
'restangular',
'ngRoute',
'MainControllers'
]);

//Controller
var MainControllers = angular.module('MainControllers', []);


App.config(['$routeProvider','$locationProvider', 'RestangularProvider',
  function($routeProvider,$locationProvider, RestangularProvider) {
    $routeProvider.
      when('/', {
        templateUrl: 'views/home.html',
        controller: 'HomeCtrl'
      }).
        when('/barang', {
            templateUrl: 'views/barang.html',
            controller: 'BarangListCtrl'
        }).
        when('/barang/new', {
            templateUrl: 'views/barang-new.html',
            controller: 'BarangNewCtrl'
        }).
        when('/barang/detail/:Id', {
            templateUrl: 'views/barang-detail.html',
            controller: 'BarangDetailCtrl'
        }).
        when('/barang/edit/:Id', {
            templateUrl: 'views/barang-edit.html',
            controller: 'BarangEditCtrl'
        }).

      otherwise({
        redirectTo: '/'
      });


      RestangularProvider.setBaseUrl('v1/');

  }]);


