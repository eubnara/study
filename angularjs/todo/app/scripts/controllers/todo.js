'use strict';

/**
 * @ngdoc function
 * @name myToDoApp.controller:TodoCtrl
 * @description
 * # TodoCtrl
 * Controller of the myToDoApp
 */
angular.module('myToDoApp')
  .controller('TodoCtrl', function ($scope) {
      $scope.todos = [];
      $scope.addTodo = function () {
        $scope.todos.push($scope.todo);
        $scope.todo = '';
      };
      $scope.removeTodo = function (index) {
        $scope.todos.splice(index, 1);
      }
  });
