import angular from 'angular';
import statisticComponent from './statistic.component';

const statisticModule = angular.module('statistic.component', [])
  .component('statistic', statisticComponent)
  .name;

export default statisticModule;