import React, { useState } from "react";
import Layout from "../components/layout";
import SEO from "../components/seo";
import {navigate} from "gatsby";
import Cookies from "universal-cookie";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Form, Button} from "react-bootstrap";

/*
  REQUEST STRUCTURE
	UserEmail string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"pwd"`
 */

const SignUpPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [firstName, setFirstname] = useState('');
  const [lastName, setLastname] = useState('');
  const cookies = new Cookies();
  const signUpHandler = (event) => {
    event.preventDefault();
    fetch("https://europe-west1-who-is-it-265713.cloudfunctions.net/NewUser", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(
          {
            email: email,
            pwd: password,
            firstName: firstName,
            lastName: lastName
          })
      }
    ).then((resp) => {
      return resp.json();
    }).then((json) => {
      /*
        RESPONSE STRUCTURE
        {
	        Success bool   `json:"success"`
	        Error   string `json:"error"`
	        Hash string `json:"hash"`
        }
      */
      console.log(json);
      if (json.success === true) {
        cookies.set("loginHash", json.hash);
        navigate("/home");
      }
    }).catch((err) => {
      console.log(err);
    });
  };

  return(
    <Layout>
      <SEO title="Signup" />
      <h1>Sign Up to Who is it</h1>
      <Form onSubmit={signUpHandler}>
        <Form.Group controlID="formBasicEmail" >
          <Form.Label>Email Address</Form.Label>
          <Form.Control type="email" placeholder="Enter email" onInput={e => setEmail(e.target.value)} />
        </Form.Group>
        <Form.Group controlID="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control type="password" placeholder="Password" onInput={e => setPassword(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Form.Label>First Name</Form.Label>
          <Form.Control type="text" placeholder="First Name" onInput={e => setFirstname(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Form.Label>Last Name</Form.Label>
          <Form.Control type="text" placeholder="Last Name" onInput={e => setLastname(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Button variant="primary" type="submit">
            Sign Up!
          </Button>
        </Form.Group>
      </Form>
    </Layout>
  )
}


export default SignUpPage
