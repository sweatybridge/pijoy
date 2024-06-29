import Joystick from './components/Joystick'
import Player from './components/Player'
import './App.css'

function App() {
  return (
    <>
      <Joystick className="card" tabIndex={-1}>
        <Player src={import.meta.env.VITE_WHEP_URL} className="video" />
        <p>Press direction keys to control the claw.</p>
      </Joystick>
    </>
  )
}

export default App
