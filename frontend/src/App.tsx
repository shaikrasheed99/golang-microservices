import { useState } from "react"
import "./App.css"
import Navbar from "./navbar/Navbar"
import Signup from "./signup/Signup"
import Login from "./login/Login"

const App = () => {
  const [hideSignup, setHideSignup] = useState(false);
  const [hideLogin, setHideLogin] = useState(true);

  return (
    <div className="container">
      <div className="header">Microservices Applicaiton</div>

      <Navbar hideSignup={setHideSignup} hideLogin={setHideLogin} />
      
      <div className="content">
        {!hideSignup && <Signup />}
        {!hideLogin && <Login />}
        <hr />
      </div>
    </div>
  )
}

export default App
