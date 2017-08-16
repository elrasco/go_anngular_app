import template from './statistics.html';
import controller from './statistics.controller';
import './statistics.scss';

const statisticsComponent = {
  bindings: {
    markers: '<'
  },
  template,
  controller
};

export default statisticsComponent;
