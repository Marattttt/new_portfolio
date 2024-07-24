import { ChangeEvent, useRef, useState } from "react"
import CodeInput from "../../components/CodeInput";
import Button from "../../components/Button";
import RunResult from "./RunResult";
import Selector from "../../components/Selector";

import styles from './editor.module.css'

type RunResult = {
	error?: string
	output?: string
}

const Editor = () => {
	const langs = ['go', 'js']
	const [lang, setLang] = useState(langs[0])
	const selectLang = (lang: string) => { setLang(lang) };

	const Editor = useRef<HTMLTextAreaElement>(null);
	const handleInput = (e: ChangeEvent<HTMLTextAreaElement>) => {
		console.log(e.target.value)
	}

	const [runResult, setRunResult] = useState<RunResult>({})

	const handleSubmit = async () => {
		let result: RunResult = {}

		try {
			const resp = await fetch(`http://localhost:3030/${lang}`, {
				method: 'POST',
				headers: [['Content-Type', 'application/json']],

				body: JSON.stringify({ code: Editor.current!.value })

			})
			// TODO: Add validation
			result = (await resp.json()) as RunResult
		}
		catch (error) {
			result.error = (error as Error).message
			console.error(error)
		}

		setRunResult(result)

		console.log(result)
	}

	return <>
		<h1> This is the code editor </h1>
		<div className={`${styles.container}`}>
			<div className={`${styles.langSelect}`}>
				<Selector items={langs} onSelect={(l) => selectLang(l)} />
			</div>

			<div className={`${styles.editor}`}>
				<CodeInput ref={Editor} onChange={(e) => handleInput(e)} />
			</div>

			<div className={`${styles.submit}`}>
				<Button contents="submit" onClick={() => handleSubmit()} />
			</div>
			<div className={`${styles.result}`}>
				<RunResult error={runResult.error} output={runResult.output} />
			</div>
		</div >
	</>
}

export default Editor
