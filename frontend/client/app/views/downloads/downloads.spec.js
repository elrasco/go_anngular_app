// should call the downloads service
// should build the component
// show the map and have markers
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
    //parentScope.$digest();

    $httpBackend.whenGET(new RegExp('/api/statistics'))
      .respond({ statistics: {} });
  }));

  it('should call the downloads service', () => {

    let expectation = $httpBackend.expectGET(new RegExp('/api/downloads'))
      .respond({});

    $httpBackend.flush();
    return expectation;
  });
  it('should build the component', () => {

    parentScope.markers = [{ App_id: "1", Latitude: 33.165085, Longitude: 72.63186, Downloaded_at: "2017-07-26T23:18:13Z" }];

    parentScope.$digest();

    expect(component.html()).contain('statistics');

  });
});