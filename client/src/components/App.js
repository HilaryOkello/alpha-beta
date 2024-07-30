import React, { useEffect, useState } from "react";
import { Switch, Route } from "react-router-dom";

import Login from "./LogIn";
import HomePage from "./HomePage";

function App() {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // auto-login
    fetch("/me").then((r) => {
      if (r.ok) {
        r.json().then((user) => setUser(user));
      }
    });
  }, []);

  return (
    <>
      <main>
        {user ? (
          <Switch>
            <Route path="/">
              <HomePage user={user} />
            </Route>
          </Switch>
        ) : (
          <Switch>
            <Route path="/login">
              <Login setUser={setUser} />
            </Route>

            <Route path="/">
              <HomePage />
            </Route>
          </Switch>
        )}
      </main>
    </>
  );
}

export default App;
