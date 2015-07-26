var React = require('react');

var Router = require('./router');

var App = React.createClass({
  render: function() {
    return (
      <Router />
    );
  }
});

module.exports = App;
