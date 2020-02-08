import React from "react";
import Layout from "../components/layout";
import SEO from "../components/seo";
import Cookies from "universal-cookie";
import 'bootstrap/dist/css/bootstrap.min.css';
import LoginCheck from "../components/loginCheck";
import { Card, Jumbotron } from "react-bootstrap"

const HomePage = () => {
    var cookies = new Cookies();
    var loginHash = cookies.get("loginHash");


    return(
        <Layout>
            <SEO title="Home" />
            <LoginCheck />
            <h1>HomePage</h1>
            <Card>
              <Card.Header as="h5">Your Profile</Card.Header>
              <Card.Body>
              <Card.Title>Will Cruse</Card.Title>
              <Card.Text>Total Points: 10</Card.Text>
              </Card.Body>
            </Card>
        </Layout>
  )
}


export default HomePage
