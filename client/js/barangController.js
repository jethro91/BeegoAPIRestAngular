MainControllers.controller('BarangListCtrl', ['$scope', '$routeParams', '$location', 'Restangular',
    function ($scope, $routeParams, $location, Restangular) {
        // Get List Barang
        $scope.title = 'Barang List';
        Restangular.all('barang').getList()
            .then(function(data){
                $scope.barangs = data;
        });
        //New Event
        $scope.newItem = function(){
            $location.path('/barang/new');
        };
        // Detail Event
        $scope.detailItem = function(idx, selected){
          $location.path('/barang/detail/'+ selected.Id);
        };
        // Edit Event
        $scope.editItem = function(idx, selected){
            $location.path('/barang/edit/'+ selected.Id);
        };
    }]);
MainControllers.controller('BarangNewCtrl', ['$scope','$routeParams','$location','Restangular',
    function ($scope,$routeParams,$location, Restangular) {
        $scope.title = 'Barang New';
        $scope.barang = {
            Nama: '',
            Stock: 0,
            Keterangan: ''
        };
        $scope.back = function(){
            $location.path('/barang');
        };

        $scope.createItem = function(){
            Restangular.all('barang').post($scope.barang)
                .then(function() {
                $location.path('/barang');
            });
        };
    }]);

MainControllers.controller('BarangDetailCtrl', ['$scope','$routeParams','$location','Restangular',
    function ($scope,$routeParams,$location, Restangular) {
        $scope.title = 'Barang Detail - ID:' +  $routeParams.Id;
        // Get List Barang
        Restangular.one('barang/'+ $routeParams.Id).get()
            .then(function(data){
            $scope.barang = data;
        });
        $scope.back = function(){
            $location.path('/barang');
        };


    }]);

MainControllers.controller('BarangEditCtrl', ['$scope','$routeParams','$location','Restangular',
    function ($scope,$routeParams,$location, Restangular) {
        $scope.title = 'Barang Edit - ID:' +  $routeParams.Id;
        // Get List Barang
        Restangular.one('barang/'+$routeParams.Id).get()
            .then(function(data){
                $scope.selectedbarang = data;
            });
        $scope.back = function(){
            $location.path('/barang');
        };
        $scope.updateItem = function(){
            console.log($scope.selectedbarang);
            $scope.selectedbarang.put().then(updateResp);
        };

        var updateResp = function(response){
            console.log(response);
            if(response.indexOf('OK') > -1){
                console.log('update sukses ');
                $location.path('/barang');
            }else {
                console.log('update gagal');
            }
        };
    }]);