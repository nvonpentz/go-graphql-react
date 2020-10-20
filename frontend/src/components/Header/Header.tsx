import { Link } from 'react-router-dom'
import React, { useContext } from 'react'
import './Header.css'
import LogOutButton from './LogOutButton'
import { UserContext } from '../../contexts/UserContext'

function Header(props: any) {
  const { user, setUser } = useContext(UserContext)
  console.log("user is " + JSON.stringify(user))
  console.log("setUser is ")
  console.log(setUser)

  let link
  if (user) {
    link = <LogOutButton />
  } else {
    link = <Link to="/log_in">Log In</Link>
  }

  return (
    <header>
      <div className="logo header-item">Helium</div>
      <nav className="header-item">
        <ul>
          <li> { link }</li>
        </ul>
      </nav>
    </header>
  )
}

export default Header
