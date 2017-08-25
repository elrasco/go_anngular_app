//here we define the module and load the component definition

import angular from 'angular';
import statisticsComponent from './statistics.component';

const statisticsModule = angular.module('statistics.component', [])
  .component('statistics', statisticsComponent)
  .name;

export default statisticsModule;