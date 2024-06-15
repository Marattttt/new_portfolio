import { ChangeEvent, useRef, useState } from "react"
import CodeInput from "./CodeInput";
import Button from "../../components/Button";
import RunResult from "./RunResult";

type RunResult = {
	error?: string
	output?: string
}

const Editor = () => {
	const Editor = useRef<HTMLTextAreaElement>(null);
	const handleInput = (e: ChangeEvent<HTMLTextAreaElement>) => {
		console.log(e.target.value)
	}

	const [runResult, setRunResult] = useState<RunResult>({})

	const handleSubmit = async () => {
		let result: RunResult = {}

		try {
			const resp = await fetch('http://localhost:3030/', {
				method: 'POST',
				headers: [
					['Content-Type', 'application/json'],
					['Access-Control-Allow-Origin', 'http:127.0.0.1:3030']
				],

				body: JSON.stringify({ code: Editor.current!.value })

			})
			// TODO: Add validation
			result = (await resp.json()) as RunResult
		}
		catch (error) {
			result = { error: "ehe" }
			console.error(error)
		}

		setRunResult(result)

		console.log(result)
	}

	return <>
		<h1> This is the code editor </h1>
		<CodeInput ref={Editor} onChange={(e) => handleInput(e)} />
		<Button contents="submit" onClick={() => handleSubmit()} />
		<RunResult error={runResult.error} output={runResult.output} />
	</>
}

export default Editor
