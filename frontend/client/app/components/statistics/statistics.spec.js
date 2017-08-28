//should show 4 statistic
//if click show I've to see stats #stats-container
//if click hide I've not to see stats #stats-container
//should call statistics


import '../../app.js';
import sinon from 'sinon';

let $compile, parentScope, downloadService, element, elCtrl;

const loadAllModules = () => {
  angular.mock.module('app');
};

describe('statistics', () => {
  beforeEach(function() {
    loadAllModules();
  });

  beforeEach(inject(function(_$compile_, _$rootScope_, downloads) {
    $compile = _$compile_;
    parentScope = _$rootScope_.$new();
    downloadService = downloads;

    element = angular.element('<statistics ></statistics>');

    //mock the statistics service
    downloadService.statistics = sinon.stub().resolves({ statistics: { ByCountry: {}, ByMonth: {}, BySeason: {}, ByWeekDay: {}, ByTime: {} } });
    $compile(element)(parentScope);

    //get the element controller
    elCtrl = element.controller('statistics');

    parentScope.$digest();

  }));

  it('should call statistics', () => {
    return expect(downloadService.statistics).to.have.been.called;
  });

  it('should start statistic component with ByCountry', () => {
    elCtrl.setChartDataWith = sinon.spy();
    return downloadService.statistics().then(() => {
      return expect(elCtrl.setChartDataWith).to.have.been.calledWith('ByCountry');
    });
  });

  it('should start statistic component with ByMonth', () => {
    elCtrl.setChartDataWith = sinon.spy();
    return downloadService.statistics().then(() => {
      return expect(elCtrl.setChartDataWith).to.have.been.calledWith('ByMonth');
    });
  });

  it('should start statistic component with BySeason', () => {
    elCtrl.setChartDataWith = sinon.spy();
    return downloadService.statistics().then(() => {
      return expect(elCtrl.setChartDataWith).to.have.been.calledWith('BySeason');
    });
  });

  it('should be visible show stat botton', () => {
    let showstats = angular.element(element[0].querySelector('#show-stats-btn'));
    expect(showstats.hasClass('ng-hide')).to.be.false;
  });

  it('should be hidden hide stat botton', () => {
    let showstats = angular.element(element[0].querySelector('#hide-stats-btn'));
    expect(showstats.hasClass('ng-hide')).to.be.true;
  });

  it('should hide show stat button if click show stat button', () => {
    let showstats = angular.element(element[0].querySelector('#show-stats-btn'));

    showstats.triggerHandler('click');
    parentScope.$digest();

    expect(showstats.hasClass('ng-hide')).to.be.true;
  });

});