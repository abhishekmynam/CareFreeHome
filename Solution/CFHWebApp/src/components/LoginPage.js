/**
 * Created by amynam on 1/4/2017.
 */
"use strict";

var React = require('react');
var LoginForms = require('./LoginSignUpForms');

var LoginPage= React.createClass({
    getInitialState: function(){
        return{
            user:{
                user_name:"",
                pwd:"",
                usrFName:"",
                usrLName:"",
                email:"",
                pwdCreate:"",
                pwdConfirm:"",
                ph_no:"",
                address:"",
                city:"",
                zipcode:""

        }
        };
    },
    setUserState: function(event){
        this.state.user[event.target.name]=event.target.value;
        return this.setState({user:this.state.user})
    },
    render: function() {
        return (
            <div>
                <div className="jumbotron">
                    <h1> Welcome to Care Free Home</h1>
                </div>
                <LoginForms
                    user={this.state.user}
                    onChange={this.setUserState}
                />
            </div>
        );
    }
});
module.exports=LoginPage;