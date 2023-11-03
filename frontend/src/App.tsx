import { useState } from "react"
import "./App.css"
import Navbar from "./navbar/Navbar"
import Signup from "./signup/Signup"

const App = () => {
  const [hideSignup, setHideSignup] = useState(false);

  return (
    <div className="container">
      <div className="header">Microservices Applicaiton</div>

      <Navbar setHideSignup={setHideSignup} />
      
      <div className="content">
        {!hideSignup && <Signup />}
        <hr />
      </div>
    </div>
  )
}

export default App
