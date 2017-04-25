'use strict';

describe('Controller: TodoCtrl', function () {

  // load the controller's module
  beforeEach(module('myToDoApp'));

  var TodoCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    TodoCtrl = $controller('TodoCtrl', {
      $scope: scope
      // place here mocked dependencies
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(TodoCtrl.awesomeThings.length).toBe(3);
  });
});
