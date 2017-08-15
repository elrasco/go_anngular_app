import sandbox from '../../_testing/sandbox';
import expect from '../../_testing/chai-plus-chai-as-promised-workaround';

describe('GeoLocation', () => {
  beforeEach(function() {
    sandbox.forGeoLocationController();
  });

  it('should expose italian and spanish regions', sandbox.within(tools => {
    const ctrl = tools.createController();

    expect(ctrl).to.have.nested.property('regions.IT');
    expect(ctrl).to.have.nested.property('regions.ES');
  }));

  it('italy and spain checkbox should be selected', sandbox.within(tools => {
    const ctrl = tools.createController();

    expect(ctrl).to.have.nested.property('location.ES.selected', true);
    expect(ctrl).to.have.nested.property('location.IT.selected', true);
  }));

  it('should get italian and spanish regions from FB', sandbox.within(tools => {
    const spy = tools.facebookService
      .on('regions')
      .spy();

    tools.createController();

    expect(spy.withArgs('IT')).to.have.been.calledOnce;
    expect(spy.withArgs('ES')).to.have.been.calledOnce;
  }));

  it('should display italian regions', sandbox.within(tools => {
    tools.facebookService
      .on('regions')
      .returns({ key: '123', name: 'region' });

    const ctrl = tools.createController();

    expect(ctrl.regions.IT).to.deep.include({ key: '123', name: 'region' });
  }));

  it('should display spanish regions', sandbox.within(tools => {
    tools.facebookService
      .on('regions')
      .returns({ key: '123', name: 'region' });

    const ctrl = tools.createController();

    expect(ctrl.regions.ES).to.deep.include({ key: '123', name: 'region' });
  }));
});
