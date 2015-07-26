var React = require('react');
var page = require('page');

var JobActions = require('../actions/job_actions');

var JobsTableRow = React.createClass({
  getDefaultProps: function() {
    return {
      job: {}
    };
  },

  render: function() {
    return (
      <tr onClick={this._onJobClick}>
	<td>{this.props.job.name}</td>
	<td>{this.props.job.user.first_name} {this.props.job.user.last_name}</td>
	<td>
	  <ul className="controls pull-right unstyled">
	    <li><a onClick={this._click}>Message Group</a></li>
	    <li><a>Share</a></li>
	    <li><a>Edit</a></li>
	    <li><a>Delete</a></li>
	  </ul>
	</td>
      </tr>
    );
  },

  _click: function(e) {
    e.stopPropagation();
    console.log('controls click');
  },

  _onJobClick: function(e) {
    e.stopPropagation();
    page('/jobs/' + this.props.job.id);
  }
});

module.exports = JobsTableRow;
