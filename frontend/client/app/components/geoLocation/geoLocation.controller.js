import R from 'ramda';

class GeoLocationController {
  constructor(FB) {
    'ngInject';
    this.name = 'geoLocation';
    this.regions = {};
    ['IT', 'ES'].forEach(country => {
      this.regions[country] = [];
      FB.regions(country).then(regions => {
        this.regions[country] = regions;
      });

      this.location = this.location || {};
      this.location = this.location[country] ? this.location : R.assocPath([country], {}, this.location);
      this.location = this.location[country].selected ? this.location : R.assocPath([country, 'selected'], true, this.location || {});
      this.location = this.location[country].regions ? this.location : R.assocPath([country, 'regions'], [], this.location || {});
    });

    this.printSelectedRegion = country => this.regions[country].filter(region => this.location[country].regions.includes(region.key)).map(R.prop('name')).join(', ');

    this.toggleCheckbox = country => {
      if (this.location[country].selected === false) {
        this.location = R.assocPath([country, 'regions'], [], this.location);
      }
    };
  }
}

export default GeoLocationController;