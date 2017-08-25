//here we define the component

import template from './statistics.html';
import controller from './statistics.controller';
import './statistics.scss';

const statisticsComponent = {
  bindings: {
    from: '<',
    to: '<'
  },
  template,
  controller
};

export default statisticsComponent;