import angular from 'angular';
import uiRouter from 'angular-ui-router';
import downloadsComponent from './downloads.component';


const downloadsModule = angular.module('downloads.view', [uiRouter])
  .config(($stateProvider) => {
    "ngInject";

    $stateProvider
      .state('downloads', {
        url: '/downloads',
        component: 'downloads',
        resolve: {
          markers: downloads => downloads.downloads().then(d => {
            return d.downloads;
          })
        }
      });
  })
  .component('downloads', downloadsComponent)
  .name;

export default downloadsModule;