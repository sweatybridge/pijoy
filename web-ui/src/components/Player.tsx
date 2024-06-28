import { WebRTCPlayer } from '@eyevinn/webrtc-player'
import { useEffect, useRef } from 'react'

export default function Player({
  src,
  ...props
}: {
  src: string
} & React.VideoHTMLAttributes<HTMLVideoElement>) {
  const playerRef = useRef<HTMLVideoElement>(null)

  const channelUrl = new URL(src)
  useEffect(() => {
    const player = new WebRTCPlayer({
      video: playerRef.current!,
      iceServers: [{ urls: 'stun:stun.cloudflare.com:3478' }],
      type: 'whep',
    })
    // TODO: retry until successfully connects to stream
    player.load(channelUrl).catch((e) => console.error(e))
    return () => {
      player.destroy()
    }
  }, [channelUrl])

  return <video autoPlay muted ref={playerRef} {...props} />
}
