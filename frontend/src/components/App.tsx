import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import './App.css'
import Header from './Header/Header'
import Home from './Home/Home'
import LogIn from './LogIn/LogIn'
import ProtectedRoute from './ProtectedRoute'
import React, { useState } from 'react'
import SignUp from './SignUp/SignUp'
import { UserContext } from '../contexts/UserContext'

const gqlClient = new ApolloClient({
  uri: 'http://localhost:3000/graphql',
  cache: new InMemoryCache(),
})

function App() {
  const userFromStorage = JSON.parse(localStorage.getItem('user') || '{}')
  let defaultUser
  if (Object.keys(userFromStorage).length === 0) {
    defaultUser = null
  } else {
    defaultUser = userFromStorage
  }

  const [user, setUser] = useState(defaultUser)
  return (
    <Router>
      <ApolloProvider client={gqlClient}>
        <UserContext.Provider value={{user, setUser}}>
          <div className="App">
            <Header/>
            <main>
              <Switch />
                <ProtectedRoute exact={ true } path="/" component={ Home }/>
                <Route exact path="/sign_up" component={ SignUp } />
                <Route exact path="/log_in" component={ LogIn } />
              <Switch />
            </main>
          </div>
        </UserContext.Provider>
      </ApolloProvider>
    </Router>
  )
}

export default App
