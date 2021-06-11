import React from "react"
import { Navbar, Nav} from "react-bootstrap"
import {Link} from "react-router-dom"

const Header = () => {

    return (
        <div>
            <Navbar bg="light" expand="lg">
            <Navbar.Brand href="#home">Book List App</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="mr-auto">
                <Nav.Link ><Link>Register</Link></Nav.Link>
                <Nav.Link ><Link>Login</Link></Nav.Link>
                </Nav>
            </Navbar.Collapse>
            </Navbar>
        </div>
    )
}

export default Header