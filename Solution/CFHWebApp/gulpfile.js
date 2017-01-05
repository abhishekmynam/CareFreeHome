/**
 * Created by amynam on 1/4/2017.
 */
"use strict";
var gulp = require('gulp');
var connect = require('gulp-connect'); //Runs local dev server
var open = require('gulp-open'); //open url in a web browser
var browserify = require('browserify');//bundle js
var reactify = require('reactify');// converts react js to normal js
var source = require('vinyl-source-stream');// use conventional text stream with gulp
var concat = require('gulp-concat');//concatenate files
var lint = require('gulp-eslint');//lint js and jsx files, maintain quality of code we are writing