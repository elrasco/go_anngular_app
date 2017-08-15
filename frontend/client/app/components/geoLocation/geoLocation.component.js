import template from './geoLocation.html';
import controller from './geoLocation.controller';
import './geoLocation.scss';

const geoLocationComponent = {
  bindings: {
    location: '=',
    ro: '<'
  },
  template,
  controller
};

export default geoLocationComponent;