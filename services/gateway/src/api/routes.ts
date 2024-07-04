import { Handler, Request, Response, Router } from "express"
import { AppConfig } from "../config"
import { MultiRunner, RunnerNotConfiguredError } from "../runners/multirunner"
import { GoCode, GoRunner, JsCode, JsRunner } from "../runners/common"
import { GoRunnerGrpc } from "../runners/golang"
import { JsRunnerGrpc } from "../runners/JsRunnerGrpc"

const asyncHandler = (fn: any) => (req: Request, res: Response, next: Handler) => {
	fn(req, res, next).catch(next)
}



const router = (conf: AppConfig): Router => {
	const runner = createMultiRunner(conf)

	const router = Router()
	router.post('/go', asyncHandler(handleGolang(runner)))
	router.post('/js', asyncHandler(handleJs(runner)))

	router.use(catchTopLevelError)

	return router
}

export default router

function handleGolang(runner: GoRunner) {
	return async (req: Request, res: Response) => {
		console.debug(`Received go code ${JSON.stringify(req.body)}`)

		try {
			const code = req.body as GoCode
			const { output, error } = await runner.runGo(code)
			res.send({ output: output, error: error })
		}
		catch (e) {
			if (e instanceof RunnerNotConfiguredError) {
				res.send('Golang is not supported at the mooment')
				return
			}
			throw e
		}
	}
}

function handleJs(runner: JsRunner) {
	return async (req: Request, res: Response) => {
		console.debug(`Received js code ${JSON.stringify(req.body)}`)

		try {
			const code = req.body as JsCode
			const { output, error } = await runner.runJs(code)
			res.send({ output: output, error: error })
		}
		catch (e) {
			if (e instanceof RunnerNotConfiguredError) {
				res.send('Golang is not supported at the mooment')
				return
			}
			throw e
		}
	}
}
function catchTopLevelError(err: any, _req: Request, res: Response, _next: Handler) {
	console.error(err)
	res.status(500).send('internal error')
}

function createMultiRunner(conf: AppConfig) {
	let go: GoRunner | undefined
	let js: JsRunner | undefined
	if (conf.goUrl) {
		go = new GoRunnerGrpc(conf.goUrl)
	}
	if (conf.jsUrl) {
		js = new JsRunnerGrpc(conf.jsUrl)
	}


	return new MultiRunner(go, js)
}
