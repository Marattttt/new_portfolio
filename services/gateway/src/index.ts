import "dotenv/config"

import express from "express"
const cors = require('cors')

import router from './api/routes'
import { AppConfig } from './config'

const app = express()
const port = process.env.PORT || 3030

const conf = new AppConfig()
console.log("Using config:")
console.log(conf)


app.use(cors())
app.use(express.json())

app.use('/', router(conf))

app.listen(port, () => console.log(`App listening on port ${port}`))
