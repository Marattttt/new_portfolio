import { ChangeEvent, forwardRef, useState } from "react"
import styles from './CodeInput.module.css'

interface ButtonProps {
	code?: string;
	onChange: (e: ChangeEvent<HTMLTextAreaElement>) => void;
}

const CodeInput = forwardRef<HTMLTextAreaElement, ButtonProps>((props, ref) => {
	const [code, setCode] = useState("")
	const onChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
		setCode(e.target.value)
		props.onChange(e)
	}

	return <textarea
		className={`${styles.codeInput}`}
		ref={ref}
		defaultValue={props.code}
		value={code} onChange={(e) => onChange(e)}
	/>
})

export default CodeInput
