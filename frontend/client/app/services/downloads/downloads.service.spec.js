import '.';

const loadAllModules = () => {
  angular.mock.module('downloads.service');
};

describe('downloads service', () => {

  beforeEach(function() {
    loadAllModules();
  });

  describe('get downloads', () => {
    // should call the downloads api route
    it('should call the /api/downloads route', inject((_$httpBackend_, downloads) => {
      _$httpBackend_.expectGET(new RegExp('/api/downloads'))
        .respond({});

      downloads.downloads().then(() => {
        return expect(true).to.be.true;
      });

      _$httpBackend_.flush();
    }));
  });

  describe('statistics', () => {
    // should call the statistics api route
    it('should call the /api/statistics route', inject((_$httpBackend_, downloads) => {
      _$httpBackend_.expectGET(new RegExp('/api/statistics'))
        .respond({});

      downloads.statistics().then(() => {
        return expect(true).to.be.true;
      });

      _$httpBackend_.flush();

    }));
  });
});