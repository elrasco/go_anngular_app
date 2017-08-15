class DownloadsController {
  constructor(NgMap) {
    'ngInject';
    this.name = 'downloads';

    NgMap.getMap().then(map => {
      //I have to resize: https://github.com/allenhwkim/angularjs-google-maps#grey-area-in-google-maps
      let center = map.getCenter();
      google.maps.event.trigger(map, "resize");
      map.setCenter(center);

      this.map = map;
    });

  }

}

export default DownloadsController;
