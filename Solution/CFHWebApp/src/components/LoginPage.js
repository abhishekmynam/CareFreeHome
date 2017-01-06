/**
 * Created by amynam on 1/4/2017.
 */
"use strict";

var React = require('react');

var LoginPage= React.createClass({
    render: function() {
        return (
            <div>
                <div className="jumbotron">
                    <h1> Welcome to Care Free Home</h1>
                </div>
                <div className="container">
                    <div class="form-group">
                        <label for="usr">User Name</label>
                        <input type="text" class="form-control" id="usr"></input>
                    </div>
                    <div class="form-group">
                        <label for="pwd">Password </label>
                        <input type="password" class="form-control" id="pwd"></input>
                    </div>
                </div>
            </div>

        );
    }
});


module.exports=LoginPage;