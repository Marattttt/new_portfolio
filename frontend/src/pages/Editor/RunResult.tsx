import { FC } from "react"

interface OutputProps {
	output?: string
	error?: string
}

const RunResult: FC<OutputProps> = ({ output, error }) => {
	return <>
		<p>Output:  {output} </p>
		{error ? <p style={{ color: "red" }}> Error: {error} </p> : undefined}
	</>
}

export default RunResult
