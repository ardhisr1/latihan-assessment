import React from "react"
import {BrowserRouter as Router, Switch, Route} from "react-router-dom"
import LoginUser from "./Pages/LoginUser";
import RegisterUser from "./Pages/RegisterUser"
import Dashboard from "./Pages/Dashboard"
import Profile from "./Pages/Profile"
import BookDetail from "./Pages/BookDetail"
import BookUpdate from "./Pages/UpdateBook"
import Home from "./Pages/Home"

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/login">
            <LoginUser/>
          </Route>
          <Route path="/register">
            <RegisterUser/>
          </Route>
          <Route path="/dashboard">
            <Dashboard/>
          </Route>
          <Route path="/profile">
            <Profile/>
          </Route>
          <Route path="/book-detail">
            <BookDetail/>
          </Route>
          <Route path="/book-update">
            <BookUpdate/>
          </Route>
          <Route path="/">
            <Home/>
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
