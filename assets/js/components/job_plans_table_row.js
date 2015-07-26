var React = require('react');

var JobPlansTableRow = React.createClass({
  getDefaultProps: function() {
    return {
      plan: {}
    };
  },

  render: function() {
    return (
      <tr>
	<td>{this.props.plan.plan_num}</td>
	<td>{this.props.plan.plan_name}</td>
	<td>{this.props.plan.plan_updated_at}</td>
	<td>{this.props.plan.filename}</td>
	<td>
	  <ul className="controls pull-right unstyled">
	    <li><a onClick={this._click}>View</a></li>
	  </ul>
	</td>
      </tr>
    );
  },

  _click: function(e) {
    e.stopPropagation();
    console.log('plan controls click');
  }
});

module.exports = JobPlansTableRow;
