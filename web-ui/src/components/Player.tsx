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
      mediaConstraints: { videoOnly: true },
    })
    // @ts-expect-error EventEmitter interface is only extended on nodejs
    player.on('connect-error', () => {
      const event = new Event('error', { bubbles: true })
      playerRef.current!.dispatchEvent(event)
    })
    // TODO: retry until successfully connects to stream
    player.load(channelUrl)
    return () => {
      player.destroy()
    }
  }, [channelUrl])

  return <video autoPlay muted ref={playerRef} {...props} />
}
