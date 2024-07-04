export class RunResult {
	public output: string | undefined
	public error: string | undefined
}

export interface GoRunner {
	runGo(code: GoCode): Promise<RunResult>
}

export class RunnerOptions {
	public runnerUrl = ""
}

export class GoCode {
	public code = ""
}

export interface JsRunner {
	runJs(code: JsCode): Promise<RunResult>
}

export class JsCode {
	public code = ""
}

