import { gql, useMutation } from '@apollo/client';
import React, { useContext } from 'react'
import { withRouter, Link } from 'react-router-dom';
import './LogInForm.css'
import { UserContext } from '../../contexts/UserContext'

const LOG_IN = gql`
  mutation($input: UserInput!) {
    logIn(input: $input) {
      email
      password
    }
  }
`

function LogInForm(props: any) {
  const { setUser } = useContext(UserContext)
  const [logIn] = useMutation(LOG_IN)
  const submitForm = async function(event: any) {
    event.preventDefault()
    const input = {
      email: event.target.email.value,
      password: event.target.password.value
    }
    try {
      const user = await logIn({ variables: { input } })
      localStorage.setItem("user", JSON.stringify(user))
      setUser(user)
      props.history.push('/')
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <form className="LogInForm" onSubmit={submitForm}>
      <div className="form-title">Log In</div>
      <div><input id="email" type="text" name="email" placeholder="Email" required></input></div>
      <div><input id="password" type="password" name="password" placeholder="Password" required ></input></div>
      <div><button>Submit</button></div>
      <div className="log-in-alt">Don't have an account? <Link to="/sign_up"> Sign up. </Link> </div>
    </form>
  )
}

export default withRouter(LogInForm)
