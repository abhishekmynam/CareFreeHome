/**
 * Created by amynam on 1/9/2017.
 */
"use strict";
var Request = require('request');

var Headers = {
    'User-Agent':       'Super Agent/0.0.1',
    'Content-Type':     'application/x-www-form-urlencoded'
}

/*var NewUserDBSaveOptions = {
    url:'http://localhost:8080/savenewuser',
    method:'POST',
    headers: Headers,
    form:newUser
}*/
var getDataSaveOptionsForPost = function(newuser,nodestring){

    var NewUserDBSaveOptions = {
        url:'http://localhost:8080/'+nodestring,
        method:'POST',
        headers: Headers,
        form:newuser}
    return(NewUserDBSaveOptions);

};

var ServiceApis  = {
    SaveNewUser: function (newuser) {
        var NewUserDBSaveOptions = getDataSaveOptionsForPost(newuser,"savenewuser");
        Request(NewUserDBSaveOptions, function (error, response, body) {
            if (!error && response.statusCode == 200) {
                // Print out the response body
                console.log(body)
            }
        })
    }
}
module.exports=ServiceApis;