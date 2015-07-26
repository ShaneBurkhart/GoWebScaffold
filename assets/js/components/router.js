var React = require('react');
var page = require('page');

var JobsPage = require('./jobs_page.js');
var JobPage = require('./job_page.js');

var BASE_PATH = '/app';

var _addRoute = function(router, path, component) {
  page(path, function(ctx) {
    router.setState({
      component: React.createElement(component, {routeContext: ctx})
    });
  });
};

var Router = React.createClass({
  getInitialState: function() {
    return {
      component: <div />
    };
  },

  componentDidMount: function() {
    console.log('Init Routes.');
    page.base(BASE_PATH);
    // Redirect root.
    page('/', '/jobs');
    _addRoute(this, '/jobs', JobsPage);
    _addRoute(this, '/jobs/:id', JobPage);
    page();
  },

  componentWillUnmount: function() {
    // TODO remove page.js
  },

  render: function() {
    return this.state.component;
  },
});

module.exports = Router;
