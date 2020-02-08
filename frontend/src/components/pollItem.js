import React from "react";
import { Container, Jumbotron, Row, Col, Button } from "react-bootstrap"
import Cookies from "universal-cookie"

const PollItem = (submission) => {

  const cookies = new Cookies();
  var loginHash = cookies.get("loginHash");


  const upvoteHandler = () => {
    fetch("https://europe-west1-who-is-it-265713.cloudfunctions.net/ReceivePoll",
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          subID: submission.submission.SubID,
          value: 1,
          userID: loginHash
        })
      }).then((resp) => {return resp.json()}).then((json) => console.log(json)).catch((err) => console.log(err));
  }

  const downvoteHandler = () => {
    fetch("https://europe-west1-who-is-it-265713.cloudfunctions.net/ReceivePoll",
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          subID: submission.submission.SubID,
          value: 0,
          userID: loginHash
        })
      }).then((resp) => {return resp.json()}).then((json) => console.log(json)).catch((err) => console.log(err));
  }
  return(
    <Jumbotron>
      <Container>
        <Row>
          <Col><p>{"User: " + submission.submission.User}</p></Col>
          <Col><p>{"Points: " + submission.submission.Points}</p></Col>
          <Col><p>{"Thing: " + submission.submission.Thing}</p></Col>
          <Col><p>{"Justification: " + submission.submission.Justification}</p></Col>
        </Row>
        <Row>
          <Col>
            <Button onClick={upvoteHandler} style={styles.button}>Upvote</Button>
            <Button onClick={downvoteHandler} style={styles.button}>Downvote</Button>
          </Col>
        </Row>
      </Container>
    </Jumbotron>
  )
}

const styles = {
  button: {
    marginRight: 10
  }
}


export default PollItem
