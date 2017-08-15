import angular from 'angular';

const names = [
];
const servicesModule = angular
  .module('app.services', names)
  .name;

export default servicesModule;
