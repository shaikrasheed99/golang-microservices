import { SetStateAction, useState } from "react";
import { RequestBody } from "../utils/helper";

interface LoginProps {
    setRequestBody: React.Dispatch<SetStateAction<RequestBody | undefined>>
}

const Login = ({ setRequestBody }: LoginProps) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleUsername = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setUsername(value);
    }

    const handlePassword = (event: React.ChangeEvent<HTMLInputElement>) => {
        const value = event.target.value;
        setPassword(value);
    }

    const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        const requestBody: RequestBody = {
            action: "auth_login",
            login: {
                username: username,
                password: password
            }
        }
        setRequestBody(requestBody);

        setUsername("");
        setPassword("");
    }

    return (
        <form className="signup" onSubmit={handleSubmit}>
            <div className="username">
                <label htmlFor="username">Username:</label>
                <input type="text" id="username" onChange={handleUsername} value={username} />
            </div>

            <div className="password">
                <label htmlFor="password">Password:</label>
                <input type="password" id="password" onChange={handlePassword} value={password} />
            </div>

            <input className="submit" type="submit" value="Login" />
        </form>
    );
}

export default Login;