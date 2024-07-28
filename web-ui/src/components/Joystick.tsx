import createClient from 'openapi-fetch'
import { useEffect, useState } from 'react'
import { components, paths } from '../api'
import './Joystick.css'

type Button = components['schemas']['Button']

const keyToButton: Record<string, Button> = {
  ArrowUp: 'up',
  ArrowDown: 'down',
  ArrowLeft: 'left',
  ArrowRight: 'right',
}

const client = createClient<paths>({
  baseUrl: import.meta.env.VITE_API_URL,
})

export default function Joystick() {
  const [inflight, setInflight] = useState(
    {} as Record<Button, AbortController>
  )

  const press = (button: Button) => {
    if (button && !(button in inflight)) {
      const abort = new AbortController()
      client
        .POST('/joystick/{button}', {
          params: {
            path: { button },
          },
          signal: abort.signal,
        })
        .catch((e) => {
          if (!abort.signal.aborted) {
            console.error(e)
          }
        })
      setInflight({ ...inflight, [button]: abort })
    }
  }

  const release = (button: Button) => {
    if (button && button in inflight) {
      inflight[button].abort()
      delete inflight[button]
      setInflight({ ...inflight })
    }
  }

  useEffect(() => {
    // Pressing single key should ignore repeats
    const onKeyDown = (event: KeyboardEvent) =>
      !event.repeat && press(keyToButton[event.key])
    // Releasing single key only cancels one request
    const onKeyUp = (event: KeyboardEvent) => release(keyToButton[event.key])
    // Releasing primary mouse button anywhere cancels action
    const onMouseUp = (event: MouseEvent) => {
      if (event.button === 0) {
        Object.values(inflight).forEach((v) => v.abort())
        setInflight({} as Record<Button, AbortController>)
      }
    }
    document.addEventListener('keydown', onKeyDown)
    document.addEventListener('keyup', onKeyUp)
    document.addEventListener('mouseup', onMouseUp)
    return () => {
      document.removeEventListener('keydown', onKeyDown)
      document.removeEventListener('keyup', onKeyUp)
      document.removeEventListener('mouseup', onMouseUp)
    }
  }, [inflight])

  return (
    <>
      <div>
        <button
          className="kbc-button"
          onTouchStart={() => press('up')}
          onTouchEnd={() => release('up')}
          onTouchCancel={() => release('up')}
          onMouseDown={(event) => event.button === 0 && press('up')}
        >
          &uarr;
        </button>
      </div>
      <button
        className="kbc-button"
        onTouchStart={() => press('left')}
        onTouchEnd={() => release('left')}
        onTouchCancel={() => release('left')}
        onMouseDown={(event) => event.button === 0 && press('left')}
      >
        &larr;
      </button>
      <button
        className="kbc-button"
        onTouchStart={() => press('down')}
        onTouchEnd={() => release('down')}
        onTouchCancel={() => release('down')}
        onMouseDown={(event) => event.button === 0 && press('down')}
      >
        &darr;
      </button>
      <button
        className="kbc-button"
        onTouchStart={() => press('right')}
        onTouchEnd={() => release('right')}
        onTouchCancel={() => release('right')}
        onMouseDown={(event) => event.button === 0 && press('right')}
      >
        &rarr;
      </button>
      <p>Press arrow keys to control the claw.</p>
    </>
  )
}
