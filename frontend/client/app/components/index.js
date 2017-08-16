import angular from 'angular';

import statistics from './statistics';

const commonModuleName = angular.module('app.common', [statistics  ])
  .name;

export default commonModuleName;
