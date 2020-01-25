import React from "react";
import Layout from "../components/layout";
import SEO from "../components/seo";
import Cookies from "universal-cookie";
import 'bootstrap/dist/css/bootstrap.min.css';
import LoginCheck from "../components/loginCheck";

const IndexPage = () => {
    var cookies = new Cookies();
    var loginHash = cookies.get("loginHash");
    
    return(
        <Layout>
            <SEO title="Login" />
            <LoginCheck />
            <h1>HomePage</h1>
        </Layout>
  )
}


export default IndexPage
