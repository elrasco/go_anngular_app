import angular from 'angular';
import downloadsService from './downloads.service';

const downloadsModule = angular.module('downloads.service', [])
  .factory('downloads', downloadsService)
  .name;

export default downloadsModule;
