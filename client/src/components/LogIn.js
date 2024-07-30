import React from "react";

function Login() {

  return (
    <div>
      <form>
        <h1>Login</h1>
        <label htmlFor="username">Username</label>
        <input
          type="text"
          id="username"
          autoComplete="off"
        />
        <label htmlFor="password">Password</label>
        <input
          type="password"
          id="password"
          autoComplete="current-password"
        />
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default Login;
