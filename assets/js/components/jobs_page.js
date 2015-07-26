var React = require('react');
var JobsTableRow = require('./jobs_table_row');

var JobStore = require('../stores/job_store');
var JobActions = require('../actions/job_actions');

var _getState = function() {
  return {
    jobs: JobStore.getJobs()
  };
};

var JobsPage = React.createClass({
  getInitialState: function() {
    return _getState();
  },

  componentDidMount: function() {
    JobStore.registerChangeListener(this._change);
    JobActions.load();
  },

  componentWillUnmount: function() {
    JobStore.removeChangeListener(this._change);
  },

  _change: function() {
    this.setState(_getState());
  },

  render: function() {
    return (
      <section>
	<h2>My Jobs</h2>
	<table className="table table-hover">
	  <thead>
	    <tr>
	      <th>Job Name</th>
	      <th>Owner</th>
	      <th className="pull-right">Controls</th>
	    </tr>
	  </thead>
	  <tbody>
	    {this.state.jobs.map(function(job, i){
	      return <JobsTableRow key={job.id} job={job} />
	    })}
	  </tbody>
	</table>
      </section>
    );
  }
});

module.exports = JobsPage;
