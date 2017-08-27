//should show 4 statistic
//if click show I've to see stats #stats-container
//if click hide I've not to see stats #stats-container
//should call statistics


import '../../app.js';
import sinon from 'sinon';

let $compile, parentScope, component, downloadService, $componentController;

const loadAllModules = () => {
  angular.mock.module('app');
};

describe('statistics', () => {
  beforeEach(function() {
    loadAllModules();
  });

  beforeEach(inject(function(_$compile_, _$rootScope_, downloads, _$componentController_) {
    $compile = _$compile_;
    parentScope = _$rootScope_.$new();
    downloadService = downloads;

    //mock the service
    downloads.statistics = sinon.stub().resolves({ statistics: { ByCountry: {}, ByMonth: {}, BySeason: {}, ByWeekDay: {}, ByTime: {} } });

    component = $compile('<statistics ></statistics>')(parentScope);

    $componentController = _$componentController_;
  }));

  it('should call statistics', () => {
    return expect(downloadService.statistics).to.have.been.called;
  });

  it.only('should start statistic component with ByCountry', () => {
    var ctrl = $componentController('statistics', { from: 'from', to: 'to' }, null);

    console.log('------------>', ctrl.chartdata)

    return true;

  });

});