import React, { ReactNode } from "react"

export type ButtonProps = {
  onClick: () => void
  contents: ReactNode
  className?: string
}

const Button: React.FC<ButtonProps> = ({ onClick, contents, className }) => {
  return <button onClick={onClick} className={`btn ${className}`}>
    {contents}
  </button>
}

export default Button

