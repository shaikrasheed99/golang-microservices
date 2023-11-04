import { useEffect, useState } from "react"
import "./App.css"
import Navbar from "./navbar/Navbar"
import Signup from "./signup/Signup"
import Login from "./login/Login"
import APIContent from "./apiContent/APIContent"
import { RequestBody, Response } from "./utils/helper"

const App = () => {
  const [hideSignup, setHideSignup] = useState(false);
  const [hideLogin, setHideLogin] = useState(true);
  const [requestBody, setRequestBody] = useState<RequestBody | undefined>(undefined);
  const [response, setResponse] = useState<Response | undefined>(undefined);

  useEffect(() => {
    if (requestBody !== undefined) {
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(requestBody)
      };

      fetch("http://localhost:8000/handle", requestOptions)
        .then((res) => res.json())
        .then((data) => {
          const res: Response = {
            status: data.status,
            code: data.code,
            message: data.message,
          }

          if (data.status === "success") {
            res.data = data.data;
          } else {
            res.reason = data.reason;
          }

          setResponse(res);
        })
        .catch((error) => {
          console.error("error is", error);
        })
    }
  }, [requestBody]);

  return (
    <div className="container">
      <div className="header">Microservices Applicaiton</div>

      <Navbar hideSignup={setHideSignup} hideLogin={setHideLogin} setRequestBody={setRequestBody} setResponse={setResponse} />

      <div className="content">
        {!hideSignup && <Signup setRequestBody={setRequestBody} />}
        {!hideLogin && <Login setRequestBody={setRequestBody} />}
        <hr />
        <APIContent requestBody={requestBody} response={response} />
      </div>
    </div>
  )
}

export default App;
