import angular from 'angular';

import Promise from 'promise-polyfill';
if (!window.Promise) { window.Promise = Promise; }

import ngAnimate from 'angular-animate';
import ngAria from 'angular-aria';
import ngMaterial from 'angular-material';
import uiRouter from 'angular-ui-router';
import ngSanitize from 'angular-sanitize';

import Services from './services';
import Views from './views';
import Components from './components';
import AppComponent from './app.component';
import Filters from './filters';

angular
  .module('app', [
    uiRouter,
    ngAnimate,
    ngAria,
    ngMaterial,
    ngSanitize,
    Views,
    Services,
    Components,
    Filters
  ])
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
