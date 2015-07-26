var React = require('react');
var JobPlansTableRow = require('./job_plans_table_row');

var JobStore = require('../stores/job_store');
var JobActions = require('../actions/job_actions');

var _getJobId = function(ctx) {
  // TODO check for number
  return parseInt(ctx.params.id);
};

var _getState = function(id) {
  return {
    job: JobStore.getJob(id)
  };
};

var _back = function() {
  window.history.back();
};

var JobPage = React.createClass({
  getInitialState: function() {
    return _getState(_getJobId(this.props.routeContext));
  },

  componentDidMount: function() {
    JobStore.registerChangeListener(this._change);
    JobActions.loadJob(_getJobId(this.props.routeContext));
  },

  componentWillUnmount: function() {
    JobStore.removeChangeListener(this._change);
  },

  _change: function() {
    this.setState(_getState(_getJobId(this.props.routeContext)));
  },

  render: function() {
    var tableBody, jobName;
    if(this.state.job) {
      jobName = this.state.job.name;
      tableBody = this.state.job.plans.map(function(plan, i){
	return <JobPlansTableRow key={plan.id} plan={plan} />
      });
    }

    return (
      <section>
	<h2>{jobName}</h2>
	<p><a onClick={_back}>Back to your jobs</a></p>
	<table className="table table-hover">
	  <thead>
	    <tr>
	      <th>#</th>
	      <th>Plan Name</th>
	      <th>Updated</th>
	      <th>Download File</th>
	      <th className="pull-right">Controls</th>
	    </tr>
	  </thead>
	  <tbody>
	    {tableBody}
	  </tbody>
	</table>
      </section>
    );
  }
});

module.exports = JobPage;
