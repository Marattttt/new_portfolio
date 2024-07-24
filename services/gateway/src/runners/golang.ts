import * as grpc from '@grpc/grpc-js'
import { GoCode, GoRunner, RunResult } from './common'
import { GoRunRequest, GoRunnerClient } from '../protogen/gorunner'

export class GoRunnerGrpc implements GoRunner {
	private url: string

	constructor(runnerServiceUrl: string) {
		this.url = runnerServiceUrl
	}

	async runGo(code: GoCode): Promise<RunResult> {
		return new Promise<RunResult>((resolve, reject) => {
			const client = new GoRunnerClient(this.url, grpc.credentials.createInsecure())

			const request: GoRunRequest = { code: code.code }

			client.runGo(request, (error, response) => {
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

