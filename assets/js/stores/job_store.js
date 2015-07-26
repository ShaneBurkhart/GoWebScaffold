var assign = require('object-assign');

var AppDispatcher = require('../dispatchers/app_dispatcher');
var Constants = require('../constants/constants');
var Store = require('./store');

var _jobs = [];
var _selectedJobId = null;

var _dispatch = function(payload) {
  switch(payload.actionType) {
    case Constants.LOAD_JOBS:
      console.log('Loading jobs.');
      JobStore.emitChange();
      break;
    case Constants.LOAD_JOBS_SUCCESS:
      console.log('Successfully loaded jobs.');
      _jobs = payload.jobs;
      JobStore.emitChange();
      break;
    case Constants.LOAD_JOBS_FAILED:
      console.log('Failed loading jobs.');
      JobStore.emitChange();
      break;
    case Constants.LOAD_JOB_SUCCESS:
      console.log('Successfully loaded job.');
      var job = payload.job;
      var updated = false;
      for(var i = 0; i < _jobs.length; i++) {
	if(_jobs[i].id === job.id) {
	  console.log('Job with id ' + job.id + ' already exists.  Updating...');
	  _jobs[i] = job;
	  updated = true;
	}
      }
      if(updated === false) {
	console.log('Adding job with id ' + job.id + '.');
	_jobs.push(job);
      }
      JobStore.emitChange();
      break;
    case Constants.LOAD_JOB_FAILED:
      console.log('Failed loading job.');
      JobStore.emitChange();
      break;
  }
};

var JobStore = assign({}, Store, {
  dispatcherIndex: AppDispatcher.register(_dispatch),

  getJobs: function() {
    return _jobs;
  },

  getJob: function(id) {
    for(var i = 0; i < _jobs.length; i++) {
      if(_jobs[i].id === id) {
	return _jobs[i];
      }
    }
    return null;
  }
});

module.exports = JobStore;
