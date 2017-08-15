const downloadsService = function($http) {
  "ngInject";
  return {
    get: (from, to) => $http({method: 'GET', url: 'http://localhost:8080/api/downloads', params: {from, to}})
  };
};

export default downloadsService;
