import createClient from 'openapi-fetch'
import { useState } from 'react'
import { components, paths } from '../api'

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

export default function Joystick(props: React.HTMLAttributes<HTMLDivElement>) {
  const [inflight, setInflight] = useState(
    {} as Record<Button, AbortController>
  )

  return (
    <div
      {...props}
      onKeyDown={(event) => {
        const button = keyToButton[event.key]
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
      }}
      onKeyUp={(event) => {
        const button = keyToButton[event.key]
        if (button && button in inflight) {
          inflight[button].abort()
          delete inflight[button]
        }
      }}
    />
  )
}
