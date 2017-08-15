module.exports = {
  "env": {
    "browser": true,
    "es6": true,
    "node": true,
    "mocha": true
  },
  "extends": "eslint:recommended",
  "parser": "babel-eslint",
  "parserOptions": {
    "sourceType": "module"
  },
  "globals": {
    "sa": true,
    'Chatra': true,
    'expect': true,
    'sinon': true,
    'inject': true,
    'angular': true
  },
  "rules": {
    "indent": [
      "warn", 2,
      { "SwitchCase": 1 }
    ],
    "semi": [
      "error", "always"
    ],
    "no-unused-vars": [
      "warn", {
        "vars": "local",
        "args": "after-used"
      }
    ],
    "no-console": [
      "warn", {
        allow: ["warn", "error"]
      }
    ]
  }
};
