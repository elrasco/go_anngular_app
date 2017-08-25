// should call the downloads service
// show the map
// should show statistics

import '../../app.js';

let $compile, parentScope, component, $httpBackend;

const loadAllModules = () => {
  angular.mock.module('app');
};

describe('downloads view', function() {

  beforeEach(function() {
    loadAllModules();
  });

  beforeEach(inject(function(_$compile_, _$rootScope_, _$httpBackend_) {
    $compile = _$compile_;
    $httpBackend = _$httpBackend_;
    parentScope = _$rootScope_.$new();
    parentScope.markers = [];
    component = $compile('<downloads markers="markers"></downloads>')(parentScope);
    $httpBackend.whenGET(new RegExp('/api/statistics'))
      .respond({ statistics: {} });
  }));

  it('should call the downloads service', () => {

    let expectation = $httpBackend.expectGET(new RegExp('/api/downloads'))
      .respond({});

    $httpBackend.flush();
    return expectation;
  });
  it('should show statistics', () => {
    expect(component.html()).contain('statistics');
  });

  it('should show the map', () => {
    expect(component.html()).contain('ng-map');
  });
});