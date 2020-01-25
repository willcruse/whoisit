import React from "react";
import { Container, Jumbotron, Row, Col } from "react-bootstrap"

const PollItem = (submission) => {
  return(
    <Jumbotron>
      <Container>
        <Row>
          <Col><p>{"User: " + submission.submission.User}</p></Col>
          <Col><p>{"Points: " + submission.submission.Points}</p></Col>
          <Col><p>{"Thing: " + submission.submission.Thing}</p></Col>
          <Col><p>{"Justification: " + submission.submission.Justification}</p></Col>
        </Row>
      </Container>
    </Jumbotron>
  )
}


export default PollItem
