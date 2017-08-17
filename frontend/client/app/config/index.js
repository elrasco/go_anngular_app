const mergeLocal = () => {
  try {
    return Object.assign(require('./development'), require('./local'));
  } catch (e) {
    console.warn('local config file does not exist.');
    return require('./development');
  }
};

let config;

switch (process.env.NODE_ENV) {
  case 'production':
    config = require('./production');
    config.env = 'production';
    config.env_prefix = '';
    break;
  default:
    config = mergeLocal();
    config.env = 'dev';
    config.env_prefix = 'dev';
}

export default config;