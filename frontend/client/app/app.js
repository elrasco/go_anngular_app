import angular from 'angular';

import Promise from 'promise-polyfill';
if (!window.Promise) { window.Promise = Promise; }

import ngAnimate from 'angular-animate';
import ngAria from 'angular-aria';
import ngMaterial from 'angular-material';
import uiRouter from 'angular-ui-router';
import ngSanitize from 'angular-sanitize';
import ngMap from 'ngmap';

import Services from './services';
import Views from './views';
import Components from './components';
import AppComponent from './app.component';
import Filters from './filters';
import ngChart from 'angular-chart.js';

import config from './config';

angular
  .module('app', [
    uiRouter,
    ngAnimate,
    ngAria,
    ngMaterial,
    ngSanitize,
    ngMap,
    Views,
    Services,
    Components,
    Filters,
    ngChart
  ])
  .run(() => {
    console.log(config);
  })
  .config(function($httpProvider) {
    "ngInject";
    $httpProvider.defaults.withCredentials = true;
  })
  .config(($locationProvider, $urlRouterProvider) => {
    "ngInject";
    $locationProvider.html5Mode(true).hashPrefix('!');
    $urlRouterProvider.otherwise('/downloads');
  })
  .config($compileProvider => {
    "ngInject";
    $compileProvider.preAssignBindingsEnabled(true);
    $compileProvider.debugInfoEnabled(false);
  })
  .component('app', AppComponent);