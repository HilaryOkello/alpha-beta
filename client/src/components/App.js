import React from "react";
import { Switch, Route } from "react-router-dom";

import Login from "./LogIn";


function App() {
 
  return (
    <>
      <main>
        <Switch>
          <Route path="/">
          <Login />
          </Route>
        </Switch>
      </main>
    </>
  );
}

export default App;
