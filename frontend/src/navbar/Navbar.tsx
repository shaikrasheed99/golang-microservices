import { SetStateAction } from "react";
import "./Navbar.css"

interface NavbarProps {
    hideSignup: React.Dispatch<SetStateAction<boolean>>
    hideLogin: React.Dispatch<SetStateAction<boolean>>
}

const Navbar = ({hideSignup, hideLogin} : NavbarProps) => {
    const handleSignupClick = () => {
        hideSignup(false);
        hideLogin(true);
    }

    const handleLoginClick = () => {
        hideLogin(false);
        hideSignup(true);
    }

    return (
        <div className="btn-container">
            <button className="btn" onClick={handleSignupClick}>Signup</button>
            <button className="btn" onClick={handleLoginClick}>Login</button>
        </div>
    );
}
 
export default Navbar;