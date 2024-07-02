export interface GoRunner {
	run(code: GoCode): Promise<RunResult>
}

export class RunnerOptions {
	public url = ""
}

export class GoCode {
	public code = ""
}

export class RunResult {
	public output: string | undefined
	public error: string | undefined
}

