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


var config={
    port:9005,
    devBaseUrl:'http://localhost',
    paths:{
        html:'./src/*.html',
        js:'./src/**/*.js',
        dist:'./dist/',
        mainJs:'./src/main.js',
        css:[
            'node_modules/bootstrap/dist/css/bootstrap.min.css',
            'node_modules/bootstrap/dist/css/bootstrap-theme.min.css'
        ],
        images:'./src/images/*'
    }
};

//start local dev server
gulp.task('connect',function(){
    connect.server({
        root:['dist'],
        port: config.port,
        base:config.devBaseUrl,
        livereload: true
    });
});

gulp.task('open',['connect'],function(){
    gulp.src('dist/index.html')
        .pipe(open({ uri: config.devBaseUrl+':'+config.port+'/'})); //settiing url
});

gulp.task('html',function(){
    gulp.src(config.paths.html) //moving files from source to dist
        .pipe(gulp.dest(config.paths.dist))
        .pipe(connect.reload()); //reload browser always after save
});

gulp.task('js',function(){
    browserify(config.paths.mainJs)
        .transform(reactify)  //convert react js to normal js
        .bundle() //bundle to depoly
        .on('error',console.error.bind(console))
        .pipe(source('bundle.js'))
        .pipe(gulp.dest(config.paths.dist + '/scripts'))
        .pipe(connect.reload());
});

gulp.task('css',function(){
    gulp.src(config.paths.css)
        .pipe(concat('bundle.css'))
        .pipe(gulp.dest(config.paths.dist+'/css'));
});

gulp.task('images',function(){
    gulp.src(config.paths.images)
        .pipe(gulp.dest(config.paths.dist+'/images'))
        .pipe(connect.reload());

    gulp.src('./src/favicon.ico')
        .pipe(gulp.dest(config.paths.dist));
});

gulp.task('lint',function(){
    return gulp.src(config.paths.js)
        .pipe(lint({configFile: 'eslint.config.json'}))
        .pipe(lint.format());
});

gulp.task('watch',function(){
    gulp.watch(config.paths.html,['html']);
    gulp.watch(config.paths.js,['js','lint']);
});

gulp.task('default',['html', 'js', 'images','css','lint','open','watch']); //list of tasks to run just by starting gulp without any parameters ir task names