var EventEmitter = require('events').EventEmitter;
var assign = require('object-assign');

var Constants = require('../constants/constants');

var Store = assign({}, EventEmitter.prototype, {
  registerChangeListener: function(callback) {
    this.on(Constants.CHANGE, callback);
  },

  removeChangeListener: function(callback) {
    this.removeListener(Constants.CHANGE, callback);
  },

  emitChange: function() {
    this.emit(Constants.CHANGE);
  },
});

module.exports = Store;
