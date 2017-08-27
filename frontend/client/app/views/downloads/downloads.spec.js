// should call the downloads service
// show the map
// should show statistics

import '../../app.js';
import sinon from 'sinon';

let $compile, parentScope, component, downloadService;

const loadAllModules = () => {
  angular.mock.module('app');
};

describe('downloads view', function() {

  beforeEach(function() {
    loadAllModules();
  });

  beforeEach(inject(function(_$compile_, _$rootScope_, downloads) {
    $compile = _$compile_;
    parentScope = _$rootScope_.$new();
    parentScope.markers = [];
    downloadService = downloads;

    //mock the service
    downloads.downloads = sinon.stub().resolves({ downloads: [] });

    component = $compile('<downloads markers="markers"></downloads>')(parentScope);
  }));

  it('should call the downloads service', () => {
    return expect(downloadService.downloads).to.have.been.called;
  });

  it('should show statistics', () => {
    expect(component.html()).contain('statistics');
  });

  it('should show the map', () => {
    expect(component.html()).contain('ng-map');
  });
});