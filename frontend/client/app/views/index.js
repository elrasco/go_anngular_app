import angular from 'angular';

import downloads from './downloads';

let viewsModule = angular.module('app.views', [
  downloads
])
  .name;

export default viewsModule;
