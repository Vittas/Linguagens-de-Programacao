import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

const hello = (word: string, sjsh: number) => {
  return <div>
    {word}
  </div>
}

const App = () => (
  <>
    {hello('vitor')}
  </>
)

export default App
