import angular from 'angular';

import geoLocation from './geoLocation';

const commonModuleName = angular.module('app.common', [geoLocation  ])
  .name;

export default commonModuleName;
