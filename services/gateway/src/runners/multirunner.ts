import { GoCode, GoRunner, JsCode, JsRunner, RunResult } from "./common";

export class RunnerNotConfiguredError extends Error {
	constructor(runner: string) {
		const message = `Runner for ${runner} was not configured`
		super(message)
		this.name = "RunnerNotSetError"
	}
}

export class MultiRunner implements GoRunner, JsRunner {
	private goRunner?: GoRunner
	private jsRunner?: JsRunner

	public constructor(go?: GoRunner, js?: JsRunner) {
		this.goRunner = go
		this.jsRunner = js
	}

	public async runGo(code: GoCode): Promise<RunResult> {
		if (!this.goRunner) {
			throw new RunnerNotConfiguredError('go')
		}

		return await this.goRunner.runGo(code)
	}

	public async runJs(code: JsCode): Promise<RunResult> {
		if (!this.jsRunner) {
			throw new RunnerNotConfiguredError('js')
		}

		return await this.jsRunner.runJs(code)
	}
}

