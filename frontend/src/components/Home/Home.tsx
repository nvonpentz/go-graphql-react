import { gql, useQuery } from '@apollo/client';
import React from 'react'
import './Home.css'

const GET_USER =gql`
  query {
    user {
      id
      createdAt
      name
      email
      password
    }
  }
`

function Home() {
  const { loading, error, data } = useQuery(GET_USER);

  if (loading) return (<div> 'Loading...' </div>)
  if (error) return (<div> {error.message} </div>)

  if (!data.user) {
    return (<div> nothing </div>)
  }

  return (
    <div className="Home">
      <div> welcome to the homepage </div>
      <div> {data.user.id} </div>
      <div> {data.user.createdAt} </div>
      <div> {data.user.name} </div>
      <div> {data.user.email} </div>
    </div>
  )
}

export default Home
