var request = require('superagent');

var AppDispatcher = require('../dispatchers/app_dispatcher.js');
var Constants = require('../constants/constants.js');

var JobActions = {
  load: function() {
    AppDispatcher.dispatch({
      actionType: Constants.LOAD_JOBS
    });
    request
      .get('/api/jobs')
      .accept('json')
      .end(function(err, res) {
	if(err) {
	  AppDispatcher.dispatch({
	    actionType: Constants.LOAD_JOBS_FAILED,
	    error: err
	  });
	} else {
	  AppDispatcher.dispatch({
	    actionType: Constants.LOAD_JOBS_SUCCESS,
	    jobs: JSON.parse(res.text).jobs
	  });
	}
      });
  },

  loadJob: function(id) {
    AppDispatcher.dispatch({
      actionType: Constants.LOAD_JOBS
    });
    request
      .get('/api/jobs/' + id)
      .accept('json')
      .end(function(err, res) {
	if(err) {
	  AppDispatcher.dispatch({
	    actionType: Constants.LOAD_JOB_FAILED,
	    error: err
	  });
	} else {
	  AppDispatcher.dispatch({
	    actionType: Constants.LOAD_JOB_SUCCESS,
	    job: JSON.parse(res.text).job
	  });
	}
      });
  }
};

module.exports = JobActions;
