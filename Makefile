run:
	docker-compose up

run-local:
	docker-compose -f docker-compose.local.yml up

stop:
	docker-compose down
