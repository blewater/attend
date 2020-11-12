import * as React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useHistory
} from "react-router-dom";
import PrivateRoute from './PrivateRoute'
import SignIn from '../components/SignIn/SignIn'
import Dashboard from '../components/DashBoard/DashBorard'
function NoMatch() {
    return (
      <div>
        <h3>404</h3>
      </div>
    );
  }

const IntroRoutes=()=> {
    return (
      <>
      <Router>
        <div>
          <Switch>
          <PrivateRoute path="/dashboard">
              <Dashboard />
            </PrivateRoute>
            <Route exact path="/">
              <SignIn />
            </Route>
            {/* Always last */}
            <Route path="*">
              <NoMatch />
            </Route>
          </Switch>
        </div>
      </Router>
      </>
    );
  }

export default IntroRoutes