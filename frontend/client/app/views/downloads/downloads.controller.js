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
      this.dynMarkers = [];
      this.markers.downloads.forEach(download => {
        const latLng =new google.maps.LatLng(download.Latitude, download.Longitude);
        this.dynMarkers.push(new google.maps.Marker({position:latLng}))
      });
      this.markerClusterer = new MarkerClusterer(map, this.dynMarkers, {});
    });
  }
}

export default DownloadsController;
