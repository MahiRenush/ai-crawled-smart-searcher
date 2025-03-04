import {Link} from "react-router-dom";
import logo from '../images/secure-search-ai-inverted.svg';
import user from '../images/useravatar.png';


function Header() {
  return (
    <header className="navbar navbar-dark bg-color text-white">      
      <div className="logo">
        <Link to='/'><img src={logo} alt="logo"/></Link>
      </div>
      <button className="user" type="button">
        <div>
          {/* TODO: Should be dynamic based on user profile */}
          <img src={user} alt="user"></img>
        </div>
      </button>
    </header>
  )
}

export default Header