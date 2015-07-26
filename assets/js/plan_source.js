var React = require('react');
var App = require('./components/app');

window.renderApp = function() {
  React.render(<App jobName="asdf" />, document.getElementById('job-app'));
};
