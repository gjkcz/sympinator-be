default:
	@echo "———————————— building ——————————————————————"
	docker-compose build

up: default
	@echo "———————————— starting ———————————————————————"
	# cd frontend; npm run build
	docker-compose up -d

watch-frontend:
	cd frontend ; npm run build

connect-api:
	docker-compose exec symp /bin/bash

connect-db:
	docker-compose exec symp-db /bin/bash

connect-mysql:
	docker-compose exec symp-db /bin/bash -c 'mysql -uroot -proot --init-command="USE internal;"'

logs:
	docker-compose logs -f

down:
	docker-compose down

clean: down
	@echo "———————————— cleaning up ————————————————————"
	rm -f api
	docker system prune -f
	docker volume prune -f
