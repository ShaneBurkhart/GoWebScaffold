var assign = require('object-assign');

var AppDispatcher = require('../dispatchers/app_dispatcher');
var Constants = require('../constants/constants');
var Store = require('./store');

var _currentPath = '';
var _currentPage = null;

var _cleanPath = function(path) {
  if(path.slice(-1) === '/') {
    return path.slice(0, -1);
  }
  return path;
};

var _determinePage = function(path) {
  switch(true) {
    case /^\/app\/jobs$/.test(path):
      return Constants.JOBS_INDEX_PAGE;
    case /^\/app\/jobs\/\d+$/.test(path):
      return Constants.JOBS_SHOW_PAGE;
  }
  return null;
};

var _dispatch = function(payload) {
  switch(payload.actionType) {
    case Constants.UPDATE_APP_PATH:
      var cleanPath = _cleanPath(payload.path);
      var page = _determinePage(cleanPath);
      console.log('Updating app path. ' + payload.path + ' as ' + cleanPath);
      console.log('Updating app page. ' + page);
      _currentPath = cleanPath;
      _currentPage = page;
      AppStore.emitChange();
      break;
  }
};

var AppStore = assign({}, Store, {
  dispatcherIndex: AppDispatcher.register(_dispatch),

  getCurrentPath: function() {
    return _currentPath;
  },

  getCurrentPage: function() {
    return _currentPage;
  }
});

module.exports = AppStore;
