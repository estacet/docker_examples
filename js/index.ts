const fastify = require('fastify')({
	logger: true
})

fastify.get('/', async () => {
	return { hello: "Hello from container" }
})

fastify.listen({ port: 3000 }, (err: Error | null) => {
	if (err) {
		fastify.log.error(err)
		process.exit(1)
	}
	console.log("Starting server...");
})
