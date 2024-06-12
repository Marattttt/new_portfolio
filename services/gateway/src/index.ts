import express from "express"
import { GoCode, GoRunner } from "./runners/runners"
import { GoRunnerGrpc } from "./runners/golang"

const app = express()
const port = process.env.PORT || 3030

const runner: GoRunner = new GoRunnerGrpc('localhost:3001')

app.use(express.json())

const asyncHandler = (fn: any) => (req: any, res: any, next: any) => {
	fn(req, res, next).catch(next)
}
//
app.get('/', (req, res) => {
	console.log(req.body)
	res.send(req.body)
})

app.post('/', asyncHandler(async (req: any, res: any) => {
	console.log(`Received: ${req.body}`)

	try {
		const code = req.body as GoCode
		const { output, error } = await runner.run(code)
		if (error) {
			res.send(error)
			return
		}

		res.send(output)
	}
	catch (e) {
		console.log("NOOOOOO EXCEPTIOUNN HAPPEDNED!!!!")
		console.log(e)
		res.send("unexpected error")
	}
}))

app.use((err: any, req: any, res: any, next: any) => {
	console.error(err)
	res.status(500).send('internal errorororoororoor')
})

app.listen(port, () => console.log(`app listening on port ${port}`))
