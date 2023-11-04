import { SetStateAction } from "react";
import "./Navbar.css"
import { RequestBody, Response } from "../utils/helper";

interface NavbarProps {
    hideSignup: React.Dispatch<SetStateAction<boolean>>
    hideLogin: React.Dispatch<SetStateAction<boolean>>
    setRequestBody: React.Dispatch<SetStateAction<RequestBody | undefined>>
    setResponse: React.Dispatch<SetStateAction<Response | undefined>>
}

const Navbar = ({ hideSignup, hideLogin, setRequestBody, setResponse }: NavbarProps) => {
    const handleSignupClick = () => {
        hideSignup(false);
        hideLogin(true);
        setRequestBody(undefined);
        setResponse(undefined);
    }

    const handleLoginClick = () => {
        hideLogin(false);
        hideSignup(true);
        setRequestBody(undefined);
        setResponse(undefined);
    }

    return (
        <div className="btn-container">
            <button className="btn" onClick={handleSignupClick}>Signup</button>
            <button className="btn" onClick={handleLoginClick}>Login</button>
        </div>
    );
}

export default Navbar;