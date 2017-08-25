const precss = require('precss');
const autoprefixer = require('autoprefixer');
const pluginsNotIncludingTestHearness = [
  'karma-chrome-launcher',
  'karma-mocha',
  'karma-mocha-reporter',
  'karma-sourcemap-loader',
  'karma-webpack',
  'karma-chai'
];

module.exports = function(config) {
  config.set({
    basePath: '.',
    frameworks: ['mocha', 'chai'],
    files: [{ pattern: 'spec.bundle.js' }],
    client: {
      chai: {
        includeStack: true
      }
    },
    plugins: pluginsNotIncludingTestHearness,
    preprocessors: { 'spec.bundle.js': ['webpack', 'sourcemap'] },
    webpack: {
      devtool: 'inline-source-map',
      module: {
        loaders: [{
            test: /\.js$/,
            exclude: [
              /app\/lib/, /node_modules/,
            ],
            loader: 'ng-annotate!babel'
          }, {
            test: /\.html$/,
            loader: 'raw'
          }, {
            test: /\.styl$/,
            loader: 'style!css!postcss!stylus'
          }, {
            test: /\.scss$/,
            loader: 'style!css!postcss!sass'
          }, {
            test: /\.css$/,
            loader: 'style!css'
          },
          {
            test: /\.json$/,
            loader: 'json'
          }
        ]
      },
      postcss: [autoprefixer, precss],
    },
    webpackServer: {
      noInfo: true
    },
    reporters: ['mocha'],
    colors: true,
    logLevel: config.LOG_INFO,
    browsers: [
      'ChromeNoSandbox'
    ],
    customLaunchers: {
      ChromeNoSandbox: {
        base: 'Chrome',
        flags: ['--no-sandbox']
      }
    },
    autoWatch: true,
    singleRun: false
  });
};