import Joystick from './components/Joystick'
import Player from './components/Player'
import './App.css'

function App() {
  return (
    <>
      <Player src={import.meta.env.VITE_WHEP_URL} className="video" />
      <Joystick />
    </>
  )
}

export default App
