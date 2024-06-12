export interface GoRunner {
	run(code: GoCode): Promise<RunResult>
}

export type RunnerOptions = {
	url: string
}

export type GoCode = {
	code: string
}

export type RunResult = {
	output: string
	error: string | undefined
}

