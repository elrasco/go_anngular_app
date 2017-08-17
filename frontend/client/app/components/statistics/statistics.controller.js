class StatisticsController {
  constructor(downloads) {
    'ngInject';
    this.name = 'statistics';
    this.viewable = false;
    this.downloads = downloads;

    this.chartdata = {};
    this.chartlabels = {};
  }

  $onChanges(changes) {
    this.downloads.statistics(this.from, this.to)
      .then(stat => {
        this.statistics = stat;

        Object.keys(this.statistics.statistics).forEach(key => {
          this.chartdata[key] = [Object.values(this.statistics.statistics[key])];
          this.chartlabels[key] = Object.keys(this.statistics.statistics[key]);
        });


      });
  }

  toggleViewable() {
    this.viewable = !this.viewable;
  }
}

export default StatisticsController;