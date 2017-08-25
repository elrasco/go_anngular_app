import Config from '../../config';

const downloadsService = function($http) {
  "ngInject";
  return {
    downloads: (from, to) => $http({ method: 'GET', url: `${Config.API_URL}/api/downloads`, params: { from, to } }).then(result => result.data),
    statistics: (from, to) => $http({ method: 'GET', url: `${Config.API_URL}/api/statistics`, params: { from, to } }).then(result => result.data),
  };
};

export default downloadsService;