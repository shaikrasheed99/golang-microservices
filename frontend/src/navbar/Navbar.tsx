import { SetStateAction } from "react";
import "./Navbar.css"

interface NavbarProps {
    setHideSignup: React.Dispatch<SetStateAction<boolean>>
}

const Navbar = ({setHideSignup} : NavbarProps) => {
    const handleSignupClick = () => {
        setHideSignup(false)
    }

    return (
        <div className="btn-container">
            <button className="btn" onClick={handleSignupClick}>Signup</button>
            <button className="btn">Login</button>
        </div>
    );
}
 
export default Navbar;