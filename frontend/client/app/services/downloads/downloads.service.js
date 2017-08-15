const downloadsService = function($http) {
  "ngInject";
  return {
    get: (from, to) => $http({method: 'GET', url: 'http://localhost:8080/api/downloads', params: {from, to}}).then(result => result.data)
  };
};

export default downloadsService;
