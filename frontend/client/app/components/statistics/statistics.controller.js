class StatisticsController {
  constructor(downloads) {
    'ngInject';
    this.name = 'statistics';
    this.viewable = false;
    this.downloads = downloads;

    this.chartdata = {};
    this.chartlabels = {};
  }

  $onChanges() {
    this.downloads.statistics(this.from, this.to)
      .then(stat => {
        this.statistics = stat;
        Object.keys(this.statistics.statistics).forEach(key => {
          this.setChartDataWith(key, Object.values(this.statistics.statistics[key]));
          this.setLabelDataWith(key, Object.keys(this.statistics.statistics[key]));
        });
      });
  }

  setChartDataWith(key, values) {
    this.chartdata[key] = [values];
  }

  setLabelDataWith(key, values) {
    this.chartlabels[key] = values;
  }

  toggleViewable() {
    this.viewable = !this.viewable;
  }
}

export default StatisticsController;