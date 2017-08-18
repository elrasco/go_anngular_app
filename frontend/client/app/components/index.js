import angular from 'angular';

import statistics from './statistics';
import statistic from './statistic';

const commonModuleName = angular.module('app.common', [statistics, statistic])
  .name;

export default commonModuleName;