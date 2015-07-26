var AppDispatcher = require('../dispatchers/app_dispatcher.js');
var Constants = require('../constants/constants.js');

var AppActions = {
  updatePath: function(path) {
    AppDispatcher.dispatch({
      actionType: Constants.UPDATE_APP_PATH,
      path: path
    });
  }
};

module.exports = AppActions;
