
import * as React from "react";
import {
    Route,
    Redirect,
  } from "react-router-dom";
  


const PrivateRoute=({ children, ...rest })=> {
    const auth = localStorage.getItem("AppToken");
    return (
      <Route
        {...rest}
        render={({ location }) =>
          auth && auth !== "undefined" ? (
            children
          ) : (
            <Redirect
              to={{
                pathname: "/",
                state: { from: location },
              }}
            />
          )
        }
      />
    );
  }
  export default PrivateRoute
