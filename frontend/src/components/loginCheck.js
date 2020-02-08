import React from "react";
import {navigate} from "gatsby";
import Cookies from "universal-cookie";

const LoginCheck = () => {
    var cookies = new Cookies();
    var loginHash = cookies.get("loginHash");

    if (loginHash == undefined) {
        navigate("/").then(r => console.log("nav"));
    }

    return (<></>)
}

export default LoginCheck;