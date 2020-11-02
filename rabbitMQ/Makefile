init-docker-management:
	@docker run -d --hostname my-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
	@echo 'Rabbit MQ management running on http://localhost:15672/'
	@echo 'Username: guest'
	@echo 'Password: guest'
	
get-dependencies:
	@go get github.com/streadway/amqp