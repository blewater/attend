import * as React from "react";

import {
  BrowserRouter as Router,
  Switch,
  Route,
  useRouteMatch
} from "react-router-dom";
import Members from '../components/Members/Members'
import Account from '../components/Account'




function NoMatch() {
    return (
      <div>
        <h3>404</h3>
      </div>
    );
  }

const DashBoardRoutes=()=> {
    let { path, url } = useRouteMatch();
    return (
      <>
        <div>
          <Switch>
          <Route path={`${path}/members`}>
              <Members />
            </Route>
            <Route path={`${path}/account`}>
              <Account />
            </Route>
          </Switch>
        </div>
      </>
    );
  }

export default DashBoardRoutes