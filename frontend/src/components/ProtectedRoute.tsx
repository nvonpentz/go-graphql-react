import React, { useContext } from "react"
import { Route, Redirect } from "react-router-dom"
import { UserContext } from '../contexts/UserContext'

function ProtectedRoute({
  component: Component,
  ...rest
}: any) {
  const { user } = useContext(UserContext)
  return (
    <Route
      {...rest}
      render={props => {
        if (user) {
          return <Component {...props} />
        } else {
          return (
            <Redirect
              to={{
                pathname: "/log_in",
              }}
            />
          )
        }
      }}
    />
  )
}

export default ProtectedRoute
