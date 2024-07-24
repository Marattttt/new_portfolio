import { credentials } from "@grpc/grpc-js"
import { JsRunRequest, JsRunnerClient } from "../protogen/jsrunner"
import { JsCode, JsRunner, RunResult } from "./common"

export class JsRunnerGrpc implements JsRunner {
	private url: string

	constructor(runnerServiceUrl: string) {
		this.url = runnerServiceUrl
	}

	async runJs(code: JsCode): Promise<RunResult> {
		return new Promise<RunResult>((resolve, reject) => {
			const client = new JsRunnerClient(this.url, credentials.createInsecure())

			const request: JsRunRequest = { code: code.code }

			client.runJs(request, (error, response) => {
				if (error) {
					console.error(error)
					reject(error)
					return
				}

				// FIXME: need special handling for cases of client errors 
				const result: RunResult = {
					output: response.output,
					error: response.error,
				}

				console.debug('received code run result', result)

				resolve(result)
			})

		})
	}
}

