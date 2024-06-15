import * as grpc from '@grpc/grpc-js'
import { FastRunnerClient } from "../protogen/fastrunner_grpc_pb"
import { promisify } from 'util'
import { GoRequest, RunResponse } from '../protogen/fastrunner_pb'
import { GoCode, GoRunner, RunResult } from './runners'
import { resolve } from 'path'
import { rejects } from 'assert'
import { response } from 'express'

export class GoRunnerGrpc implements GoRunner {
	private url: string

	constructor(runnerServiceUrl: string) {
		this.url = runnerServiceUrl
	}

	async run(code: GoCode): Promise<RunResult> {
		return new Promise<RunResult>((resolve, reject) => {
			const client = new FastRunnerClient(this.url, grpc.credentials.createInsecure())

			const request = new GoRequest()
			request.setCode(code.code)

			client.runGoLang(request, (error, response) => {
				if (error) {
					console.error(error)
					reject(error)
					return
				}

				// FIXME: need special handling for cases of client errors 
				const result: RunResult = {
					output: response.getOutput(),
					error: response.getError(),
				}

				resolve(result)
			})

		})
	}
}

