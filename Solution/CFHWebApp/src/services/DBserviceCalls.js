"use strict";

//This file is mocking a web API by hitting hard coded data.
var _ = require('lodash');

//This would be performed on the server in a real app. Just stubbing in.
var _generateId = function(author) {
    return author.firstName.toLowerCase() + '-' + author.lastName.toLowerCase();
};

var _clone = function(item) {
    return JSON.parse(JSON.stringify(item)); //return cloned copy so that the item is passed by value instead of by reference
};

var userApi = {
    getAllAuthors: function() {
        return _clone(authors);
    },

    getAuthorById: function(id) {
        var author = _.find(authors, {id: id});
        return _clone(author);
    },

    SaveNewUser: function(author) {
        //pretend an ajax call to web api is made here
        console.log('Pretend this just saved the author to the DB via AJAX call...');
        console.log(author);

        return _clone(author);
    },

    deleteAuthor: function(id) {
        console.log('Pretend this just deleted the author from the DB via an AJAX call...');
        _.remove(authors, { id: id});
    }
};

module.exports = userApi;