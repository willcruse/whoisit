import React, { useState } from "react"
import Layout from "../components/layout";
import SEO from "../components/seo";
import 'bootstrap/dist/css/bootstrap.min.css';
import LoginCheck from "../components/loginCheck";
import PollItem from "../components/pollItem"

const PollPage = () => {

  var [submissions, setSubmissions] = useState([])
  fetch("https://europe-west1-who-is-it-265713.cloudfunctions.net/GetSubmissions").then((resp) => {
    return resp.json()
  }).then((json) => {
    setSubmissions(json.subs);
  }).catch((err) => {console.log(err)})
  return(
    <Layout>
      <SEO title="Polls" />
      <LoginCheck />
      <h1>Open Polls</h1>
      {submissions.map((item) => {
        return <PollItem key={item.SubID} submission={item}/>
      })}
    </Layout>
  )
}


export default PollPage
