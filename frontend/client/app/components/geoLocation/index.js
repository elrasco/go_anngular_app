import angular from 'angular';
import geoLocationComponent from './geoLocation.component';

const geoLocationModule = angular.module('geoLocation.component', [])
  .component('geoLocation', geoLocationComponent)
  .name;

export default geoLocationModule;
