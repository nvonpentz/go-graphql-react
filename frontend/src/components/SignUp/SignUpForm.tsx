import { gql, useMutation } from '@apollo/client';
import React from 'react'
import { withRouter } from 'react-router-dom';
import './SignUpForm.css'

const SIGN_UP = gql`
  mutation($input: UserInput!) {
    signUp(input: $input) {
      name
      email
      password
    }
  }
`

function SignUpForm(props:any) {
  const [signUp, { loading: mutationLoading, error: mutationError }] = useMutation(SIGN_UP)
  const submitForm = async function(event: any) {
    event.preventDefault()
    const input = {
      name: event.target.name.value,
      email: event.target.email.value,
      password: event.target.password.value
    }
    try {
      const user = await signUp({ variables: { input } })
      localStorage.setItem("user", JSON.stringify(user))
      props.history.push('/')
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <form className="SignUpForm" onSubmit={submitForm}>
      <div className="form-title">Create your account</div>
      <div><input id="name" type="text" name="name" placeholder="Name" required ></input></div>
      <div><input id="email" type="text" name="email" placeholder="Email" required ></input></div>
      <div><input id="password" type="password" name="password" placeholder="Password" required ></input></div>
      <div><button>Submit</button></div>
      <div className="log-in-alt">Already have an account? Log in. </div>
      {mutationLoading && <p>"Loading.."</p>}
      {mutationError && <p> {mutationError.message} </p>}
    </form>
  )
}

export default withRouter(SignUpForm)
