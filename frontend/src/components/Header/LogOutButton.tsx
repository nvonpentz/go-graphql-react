import React, { useContext } from 'react'
import { withRouter } from 'react-router-dom';
import { gql, useMutation } from '@apollo/client';
import { UserContext } from '../../contexts/UserContext'

const LOG_OUT = gql`
  mutation {
    logOut
  }
`

function LogOutButton(props: any) {
  const { setUser } = useContext(UserContext)
  const [logOut] = useMutation(LOG_OUT)
  const logOutButtonClicked = async function(event: any) {
    event.preventDefault()
    try {
      const response = await logOut()
      if (response.data.logOut === true) {
        localStorage.removeItem('user')
        console.log("setting user to null")
        setUser(null)
        props.history.push('/')
      } else {
        console.log("Received unexpected response from API: " + response)
      }
    } catch (error) {
      console.log(error)
    }
  }

  return (
    // eslint-disable-next-line
    <a href="#" onClick={ logOutButtonClicked }>Log Out</a>
  )
}

export default withRouter(LogOutButton)
