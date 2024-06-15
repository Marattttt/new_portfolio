import express from "express"
import { GoCode, GoRunner } from "./runners/runners"
import { GoRunnerGrpc } from "./runners/golang"

const cors = require('cors')
const app = express()
const port = process.env.PORT || 3030

const runner: GoRunner = new GoRunnerGrpc('localhost:3001')

app.use(cors())
app.use(express.json())

const asyncHandler = (fn: any) => (req: express.Request, res: express.Response, next: express.Handler) => {
	fn(req, res, next).catch(next)
}

app.post('/', asyncHandler(async (req: express.Request, res: express.Response) => {
	console.log(`Received: `, req.body)
	try {
		const code = req.body as GoCode
		const { output, error } = await runner.run(code)
		res.send({ output: output, error: error })
	}
	catch (e) {
		console.log("NOOOOOO EXCEPTIOUNN HAPPEDNED!!!!")
		console.log(e)
		res.send("unexpected error")
	}
}))

app.use((err: any, req: express.Request, res: express.Response, next: express.Handler) => {
	console.error(err)
	res.status(500).send('internal error')
})

app.listen(port, () => console.log(`app listening on port ${port}`))
