import PropTypes from "prop-types";
import React from "react";
import {Navbar, Nav} from "react-bootstrap";
import {Link} from "gatsby";

const Header = () => (
  <Navbar bg="light" expand="lg">
    <Navbar.Brand><Link to="/">Who is it?</Link></Navbar.Brand>
    <Navbar.Toggle aria-controls="basic-navbar-nav" />
    <Navbar.Collapse id="basic-navbar-nav">
      <Nav className="mr-auto">
        <Nav.Item><Link to="/home">Your, Profile</Link></Nav.Item>
        <Nav.Item><Link to="/submit">Submit Score</Link></Nav.Item>
      </Nav>
    </Navbar.Collapse>
  </Navbar>
)

Header.propTypes = {
  siteTitle: PropTypes.string,
}

Header.defaultProps = {
  siteTitle: `Who is it?`,
}

export default Header
