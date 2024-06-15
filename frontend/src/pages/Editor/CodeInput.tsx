import { ChangeEvent, forwardRef, useState } from "react"

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

	return <textarea ref={ref}
		defaultValue={props.code}
		value={code} onChange={(e) => onChange(e)}
		cols={60}
		rows={10}
	/>
})

export default CodeInput
