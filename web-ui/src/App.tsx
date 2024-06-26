import React from "react";
import Joystick from "./components/Joystick";
import logo from "./logo.svg";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Press direction keys to control the claw.</p>
        <Joystick className="joystick" tabIndex={-1} />
      </header>
    </div>
  );
}

export default App;
