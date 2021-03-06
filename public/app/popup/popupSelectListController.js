angular.module('tima').controller('PopupSelectListController',
['$scope', '$uibModalInstance', 'title', 'items', 'acceptButton', 'cancelButton',
function ($scope, $uibModalInstance, title, items, acceptButton, cancelButton) {

    $scope.title = title;
    $scope.items = items;
    $scope.acceptButton = acceptButton;
    $scope.cancelButton = cancelButton;

    $scope.currentPage = 1;
    $scope.totalItems = $scope.items.length;
    $scope.itemsPerPage = 15;

    $scope.$watch('filteredItems', function(newVal, oldVal) {
        $scope.currentPage = 1;
    }, true);

    $scope.accept = function () {
        $uibModalInstance.close(true);
    };

    $scope.cancel = function () {
        $uibModalInstance.dismiss();
    };

}]);
