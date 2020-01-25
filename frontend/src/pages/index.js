import React, { useState } from "react";
import Layout from "../components/layout";
import SEO from "../components/seo";
import {navigate} from "gatsby";
import Cookies from "universal-cookie";
import 'bootstrap/dist/css/bootstrap.min.css';
import { Form, Button} from "react-bootstrap";

const IndexPage = () => {
  var cookies = new Cookies();
  var loginHash = cookies.get("loginHash");

  if (loginHash !== undefined) {
    //TODO authenticate
    navigate("/home");
  }

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const loginHandler = (event) => {
    event.preventDefault();
    fetch("https://europe-west1-who-is-it-265713.cloudfunctions.net/Login", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({email: email, pwd: password})
      }
    ).then((resp) => {
      return resp.json();
    }).then((json) => {
      /*
        RESPONSE STRUCTURE
        {
          success: bool
          error: string
          auth: bool
          hash: string
        }
      */
      console.log(json);
      if (json.success === true) {
        if (json.auth === true) {
          cookies.set("loginHash", json.hash);
          navigate("/home");
        } 
      }
    }).catch((err) => {
      console.log(err);
    });
  };

  return(
    <Layout>
    <SEO title="Login" />
    <Form onSubmit={loginHandler}>
      <Form.Group controlID="formBasicEmail" >
        <Form.Label>Email Address</Form.Label>
        <Form.Control type="email" placeholder="Enter email" onInput={e => setEmail(e.target.value)} />
      </Form.Group>
      <Form.Group controlID="formBasicPassword">
        <Form.Label>Password</Form.Label>
        <Form.Control type="password" placeholder="Password" onInput={e => setPassword(e.target.value)}/>
      </Form.Group>
      <Form.Group>
      <Button variant="primary" type="submit">
        Login
      </Button>
      </Form.Group>
      <Form.Group>
      <Button variant="secondary" onClick={() => {navigate("/signup")}}>
        Sign Up
      </Button>
      </Form.Group>
    </Form>
    </Layout>
  )
}


export default IndexPage
