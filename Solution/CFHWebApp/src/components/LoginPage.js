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
                <div>
                    <p> </p>
                    <p> </p>
                    <p> </p>
                    <p> Login In</p>
                </div>
                <div className="container">
                    <form className="form-inline">
                        <div className="form-group">
                            <label htmlFor="user_name">User Name:</label>
                            <input type="text" className="form-control" id="email" placeholder="Enter User Name"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="pwd">Password:</label>
                            <input type="password" className="form-control" id="pwd" placeholder="Enter password"></input>
                        </div>
                        <button type="submit" className="btn btn-default">Submit</button>
                    </form>
                </div>
                <div>
                    <p> </p>
                    <p> </p>
                    <p> </p>
                    <p> Sign Up</p>
                </div>
                <div className="container">
                    <form className="form-inline">
                        <div className="form-group">
                            <label htmlFor="usr">First Name:</label>
                            <input type="text" className="form-control" id="First_Name" placeholder="Enter First Name"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="last_name">Last Name:</label>
                            <input type="text" className="form-control" id="last_name" placeholder="Enter Last Name"></input>
                        </div>
                    </form>
                        <div className="form-group">
                            <label htmlFor="email">email id:</label>
                            <input type="text" className="form-control" id="emailId" placeholder="Enter email id"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="ph_no">Phone Number:</label>
                            <input type="text" className="form-control" id="phNo" placeholder="Enter Phone Number"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="address">Address:</label>
                            <input type="text" className="form-control" id="address" placeholder="Enter Address"></input>
                        </div>
                    <form className="form-inline">
                        <div className="form-group">
                            <label htmlFor="city">City:</label>
                            <input type="text" className="form-control" id="city" placeholder="Enter City"></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="zipcode">Zip Code:</label>
                            <input type="text" className="form-control" id="zipCode" placeholder="Enter Zip Code"></input>
                        </div>
                    </form>
                    <div className="dropdown">
                        <button className="btn btn-default dropdown-toggle" type="button" id="menu1" data-toggle="dropdown">State
                            <span className="caret"></span></button>
                        <ul className="dropdown-menu" role="menu" aria-labelledby="menu1">
                            <li role="presentation"><a role="menuitem" tabIndex="-1" href="#">HTML</a></li>
                            <li role="presentation"><a role="menuitem" tabIndex="-1" href="#">CSS</a></li>
                            <li role="presentation"><a role="menuitem" tabIndex="-1" href="#">JavaScript</a></li>
                        </ul>
                    </div>
                    <div className="dropdown">
                        <button className="btn btn-default dropdown-toggle" type="button" id="menu1" data-toggle="dropdown">Country
                            <span className="caret"></span></button>
                        <ul className="dropdown-menu" role="menu" aria-labelledby="menu1">
                            <li role="presentation"><a role="menuitem" tabIndex="-1" href="#">HTML</a></li>
                            <li role="presentation"><a role="menuitem" tabIndex="-1" href="#">CSS</a></li>
                            <li role="presentation"><a role="menuitem" tabIndex="-1" href="#">JavaScript</a></li>
                        </ul>
                    </div>

                </div>
            </div>

        );
    }
});
module.exports=LoginPage;