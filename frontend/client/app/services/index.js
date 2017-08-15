import angular from 'angular';

import downloads from './downloads';

const names = [downloads];
const servicesModule = angular
  .module('app.services', names)
  .name;

export default servicesModule;
