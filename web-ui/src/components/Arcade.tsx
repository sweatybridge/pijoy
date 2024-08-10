import './Arcade.css'

interface ArcadeProps {
  fallback: string
}

export default function Arcade({
  fallback,
  children,
}: React.PropsWithChildren<ArcadeProps>) {
  return (
    <div className="arcade-machine">
      <div className="shadow"></div>
      <div className="top">
        <div className="stripes"></div>
      </div>
      <div className="screen-container">
        <div className="shadow"></div>
        <div className="screen">
          {fallback ? <Screen text={fallback} /> : children}
        </div>
        <div className="joystick">
          <div className="stick"></div>
        </div>
      </div>
      <div className="board">
        <div className="button button-a"></div>
        <div className="button button-b"></div>
        <div className="button button-c"></div>
      </div>
      <div className="bottom">
        <div className="stripes"></div>
      </div>
    </div>
  )
}

function Screen({ text }: { text: string }) {
  return (
    <>
      <div className="screen-display"></div>
      <h2>{text}</h2>
      <div className="alien-container">
        <div className="alien">
          <div className="ear ear-left"></div>
          <div className="ear ear-right"></div>
          <div className="head-top"></div>
          <div className="head">
            <div className="eye eye-left"></div>
            <div className="eye eye-right"></div>
          </div>
          <div className="body"></div>
          <div className="arm arm-left"></div>
          <div className="arm arm-right"></div>
        </div>
      </div>
    </>
  )
}
