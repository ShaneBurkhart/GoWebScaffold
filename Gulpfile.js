var gulp = require('gulp');
var del = require('del');
var source = require('vinyl-source-stream');
var reactify = require('reactify');
var browserify = require('browserify');

gulp.task('clean', function() {
  del('./app/assets/javascripts/plan_source/dist/*.js');
});

var _bundler = browserify({
  // Only need initial file, browserify finds the deps
  entries: ['./app/assets/javascripts/plan_source/plan_source.js'],
  // We want to convert JSX to normal javascript
  transform: [reactify],
  // Gives us sourcemapping
  debug: true,
  // Requirement of watchify
  cache: {}, packageCache: {}, fullPaths: true
});

gulp.task('js', ['clean'], function() {
  _bundler.bundle()
    .on('error', function(err) {
      process.stdout.write(err.toString());
      this.emit('end');
    })
    .pipe(source('plan_source.js'))
    .pipe(gulp.dest('./app/assets/javascripts/plan_source/dist'));
});

gulp.task('watch', ['js'], function() {
  gulp.watch('./app/assets/javascripts/plan_source/**/*.js', ['js']);
});

gulp.task('default', ['js']);

