import { useState } from "react"
import styles from './Selector.module.css'

interface SelectorOpts {
  items: string[]
  onSelect: (i: string) => void
}

const Selector = ({ items, onSelect }: SelectorOpts) => {
  const [selected, setSelected] = useState({ name: items[0], index: 0 });
  const [isPicking, setIsPicking] = useState(false);

  const togglePicking = () => setIsPicking(!isPicking)

  const select = (index: number) => {
    setSelected({ name: items[index], index: index })
    setIsPicking(false)
    onSelect(items[index])
  }

  return (
    <div className={`${styles.container}`}>
      <div className={`${styles.selected}`} onClick={togglePicking}>
        {selected.name}
      </div>
      {
        isPicking && (
          <ul className={`${styles.itemsList}`}>
            {items.map((item, index) => (
              <li
                key={index}
                className={styles.item}
                onClick={() => select(index)}
              >
                {item}
              </li>
            ))}
          </ul>
        )
      }
    </div >
  )
}

export default Selector
