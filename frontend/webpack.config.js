const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const precss = require('precss');
const autoprefixer = require('autoprefixer');
const copy = require('copy-webpack-plugin');

module.exports = {

  devtool: 'source-map',
  entry: {},
  module: {
    loaders: [
      {
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
  postcss: function() {
    return [autoprefixer, precss, ];
  },
  plugins: [
    new webpack.DefinePlugin({
      'process.env': {
        'NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'development')
      }
    }),
    // Injects bundles in your index.html instead of wiring all manually.
    // It also adds hash to all injected assets so we don't have problems
    // with cache purging during deployment.
    new HtmlWebpackPlugin({ template: 'client/index.html', inject: 'body', hash: true }),

    // Automatically move all modules defined outside of application directory to vendor bundle.
    // If you are using more complicated project structure, consider to specify common chunks manually.
    new webpack
      .optimize
      .CommonsChunkPlugin({
      name: 'vendor',
      minChunks: function(module, count) {
        return module.resource && module
          .resource
          .indexOf(path.resolve(__dirname, 'client')) === -1;
      }
    })
  ]

};
