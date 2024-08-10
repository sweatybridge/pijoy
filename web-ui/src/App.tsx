import { useState } from 'react'
import Arcade from './components/Arcade'
import Joystick from './components/Joystick'
import Player from './components/Player'
import './App.css'

function App() {
  const [fallback, setFallback] = useState('')
  return (
    <>
      <Arcade fallback={fallback}>
        <Player
          src={import.meta.env.VITE_WHEP_URL}
          className="video"
          onError={() => setFallback('WEBCAM OFFLINE')}
        />
      </Arcade>
      <Joystick />
    </>
  )
}

export default App
