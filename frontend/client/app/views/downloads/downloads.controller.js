class DownloadsController {
  constructor(NgMap, downloads) {
    'ngInject';
    this.name = 'downloads';

    this.googleMapsUrl = "https://maps.googleapis.com/maps/api/js?key=AIzaSyCgUcK5V8fDzPEzwUynWQ2kZe7UDvEmXVo";

    downloads.downloads().then(d => {
      this.markers = d.downloads;
    });

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