import React, { useState } from "react";
import Layout from "../components/layout";
import SEO from "../components/seo";
import {navigate} from "gatsby";
import Cookies from "universal-cookie";
import 'bootstrap/dist/css/bootstrap.min.css';
import {Form, Button} from "react-bootstrap";

const SubmitPage = () => {
  const [user, setUser] = useState('');
  const [thing, setThing] = useState('');
  const [score, setScore] = useState(0);
  const [just, setJust] = useState('');
  const submitHandler = (event) => {
    event.preventDefault();
    fetch("https://europe-west1-who-is-it-265713.cloudfunctions.net/SubmitScore", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          // 'Access-Control-Allow-Origin': '*',
      },
        body: JSON.stringify(
          {
            userID: user,
            thing: thing,
            score: Number(score),
            just: just
          })
      }
    ).then((resp) => {
      return resp.json();
    }).then((json) => {
      /*
        RESPONSE STRUCTURE
        {
	        Success bool   `json:"success"`
        }
      */
      console.log(json);
      if (json.success === true) {
        console.log("success");
      }
    }).catch((err) => {
      console.log(err);
    });
  };

  return(
    <Layout>
      <SEO title="Submit" />
      <h1>Submit new score</h1>
      <Form onSubmit={submitHandler}>
        <Form.Group>
          <Form.Label>What is their email address?</Form.Label>
          <Form.Control type="email" placeholder={"keith@me.com"} onInput={e => setUser(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Form.Label>What have they done?</Form.Label>
          <Form.Control type="text" placeholder="Had an encounter with a bus" onInput={e => setThing(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Form.Label>What score do they deserve?</Form.Label>
          <Form.Control type="number" placeholder={-60} onInput={e => setScore(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Form.Label>Why do they deserve this score?</Form.Label>
          <Form.Control type="text" placeholder="It was a very special encounter" onInput={e => setJust(e.target.value)}/>
        </Form.Group>
        <Form.Group>
          <Button variant="primary" type="submit">
            Submit
          </Button>
        </Form.Group>
      </Form>
    </Layout>
  )
}


export default SubmitPage
