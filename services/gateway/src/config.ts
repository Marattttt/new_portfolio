import * as cnst from "./constants"

export class AppConfig {
	public port: Number
	public goUrl: string
	public jsUrl: string

	constructor() {
		this.port = getEnvInt(cnst.APP_PORT_ENV_KEY) || cnst.APP_PORT_DEFAULT
		this.goUrl = process.env[cnst.GORUNNER_URL_ENV_KEY] || cnst.GORUNNER_URL_DEFAULT
		this.jsUrl = process.env[cnst.JSRUNNER_URL_ENV_KEY] || cnst.JSRUNNER_URL_DEFAULT
	}
}

function getEnvInt(name: string): Number | undefined {
	if (!process.env[name]) {
		return
	}

	const num = Number(process.env[name])
	if (num % 1) {
		console.error(`${name} env variable is not an integer`)
		return
	} else {
		return num
	}
}
