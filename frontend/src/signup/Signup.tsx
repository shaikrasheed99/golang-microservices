import { SetStateAction, useState } from "react";
import "./Signup.css"
import { RequestBody } from "../utils/helper";

interface SignupProps {
    setRequestBody: React.Dispatch<SetStateAction<RequestBody | undefined>>
}

const Signup = ({setRequestBody}: SignupProps) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");

    const handleUsername = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setUsername(value);
    }

    const handlePassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setPassword(value);
    }

    const handleEmail = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setEmail(value);
    }

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const requestBody: RequestBody = {
            action: "auth_signup",
            signup: {
                username: username,
                password: password,
                email: email
            }
        }
        setRequestBody(requestBody);

        setUsername("");
        setPassword("");
        setEmail("");
    }

    return ( 
        <form className="signup" onSubmit={handleSubmit}>
            <div className="username">
                <label htmlFor="username">Username:</label>
                <input type="text" id="username" onChange={handleUsername} value={username} />
            </div>

            <div className="password">
                <label htmlFor="password">Password:</label>
                <input type="password" id="password" onChange={handlePassword} value={password}/>
            </div>

            <div className="email">
                <label htmlFor="email">Email:</label>
                <input type="text" id="email" onChange={handleEmail} value={email}/>
            </div>
            
            <input className="submit" type="submit" value="Signup"/>
        </form>
    );
}
 
export default Signup;