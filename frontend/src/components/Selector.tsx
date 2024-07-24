import { useState } from "react"
import styles from './CodeInput.module.css'

interface SelectorOpts {
  items: string[]
  onSelect: (i: string) => void
}

const Selector = ({ items, onSelect }: SelectorOpts) => {
  const [selected, setSelected] = useState(items[0]);
  const handleSelect = (i: string) => {
    setSelected(i)
    onSelect(i)
  }
  return (
    <div className={styles.langBtnBox}>
      {items.map((item: string) => (
        <button
          key={item}
          className={`${styles.langBtn} ${item === selected ? styles.selectedLangBtn : ''}`}
          onClick={() => handleSelect(item)}
        >
          {item}
        </button>
      ))}
    </div>
  )
}

export default Selector
