/**
 * Created by amynam on 1/7/2017.
 */
"use strict";
var React = require('react');
//var DBServiceCalls = require('../services/DBserviceCalls');

var LoginPageForms = React.createClass({
    createUser: function(event){
        event.preventDefault();

    },
    render:function () {
        return(
            <div>
                <div>
                    <p> Login In</p>
                </div>
                <div className="container">
                    <form className="form-inline">
                    <div className="form-group">
                        <label htmlFor="user_name">User Name:</label>
                        <input type="text" name = "user_name" className="form-control" id="userNameLogin"
                               placeholder="Enter User Name" onChange={this.props.onChange} value={this.props.user.user_name}></input>
                    </div>
                    <div className="form-group">
                        <label htmlFor="pwd">Password:</label>
                        <input type="password" name ="pwd" className="form-control" id="passwordLogin"
                               placeholder="Enter password" onChange={this.props.onChange} value={this.props.user.pwd}></input>
                    </div>
                    <button type="submit" className="btn btn-default">Submit</button>
                    </form>
                </div>
                <div>
                    <p> Sign Up</p>
                </div>
                <div className="container">
                    <form className="form-inline">
                        <div className="form-group">
                            <label htmlFor="usrFName">First Name:</label>
                            <input type="text" name ="usrFName" className="form-control" id="firstName"
                                   placeholder="Enter First Name" onChange={this.props.onChange} value={this.props.user.usrFName}></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="usrLName">Last Name:</label>
                            <input type="text" name = "usrLName" className="form-control" id="lastName"
                                   placeholder="Enter Last Name" onChange={this.props.onChange} value={this.props.user.usrLName}></input>
                        </div>
                    </form>
                    <div className="form-group">
                        <label htmlFor="email">email id:</label>
                        <input type="text" name ="email" className="form-control" id="emailIdSignUp"
                               placeholder="Enter email id" onChange={this.props.onChange} value={this.props.user.email}></input>
                    </div>
                    <form className="form-inline">
                        <div className="form-group">
                            <label htmlFor="pwdCreate"> Create Password:</label>
                            <input type="Password" name="pwdCreate" className="form-control" id="createPswd"
                                   placeholder="Enter password" onChange={this.props.onChange} value={this.props.user.pwdCreate}></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="pwdConfirm">Confirm Password:</label>
                            <input type="Password" name ="pwdConfirm" className="form-control" id="confirmPswd"
                                   placeholder="Confirm password" onChange={this.props.onChange} value={this.props.user.pwdConfirm}></input>
                        </div>
                    </form>
                    <div className="form-group">
                        <label htmlFor="ph_no">Phone Number:</label>
                        <input type="number" name="ph_no" className="form-control" id="phNo"
                               placeholder="Enter Phone Number" onChange={this.props.onChange} value={this.props.user.ph_no}></input>
                    </div>
                    <div className="form-group">
                        <label htmlFor="address">Address:</label>
                        <input type="text" name="address" className="form-control" id="address"
                               placeholder="Enter Address" onChange={this.props.onChange} value={this.props.user.address}></input>
                    </div>
                    <form className="form-inline">
                        <div className="form-group">
                            <label htmlFor="city">City:</label>
                            <input type="text" name ="city" className="form-control" id="city"
                                   placeholder="Enter City" onChange={this.props.onChange} value={this.props.user.city}></input>
                        </div>
                        <div className="form-group">
                            <label htmlFor="zipcode">Zip Code:</label>
                            <input type="number" name ="zipcode" className="form-control" id="zipCode"
                                   placeholder="Enter Zip Code" onChange={this.props.onChange} value={this.props.user.zipcode}></input>
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
module.exports=LoginPageForms;