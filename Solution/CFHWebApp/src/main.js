/**
 * Created by amynam on 1/4/2017.
 */
$ = jQuery = require('jquery');

var React = require('react');
var ReactDom = require('react-dom');
var LoginPage = require('./components/LoginPage');

ReactDom.render(<LoginPage />, document.getElementById('app'));