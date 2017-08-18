import template from './statistic.html';
import controller from './statistic.controller';
import './statistic.scss';

const statisticComponent = {
  bindings: {
    title: '@',
    chartdata: '<',
    chartlabels: '<'
  },
  template,
  controller
};

export default statisticComponent;