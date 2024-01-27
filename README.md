# this is price prediction sample app

## Architecture

In This simple example price prediction we will use simple architecture of API based architecture

![Architecture Diagram](./docs/diagram1.png)

### bababos-backend

bababos backend is simple golang application that opena api to the frontend. all the logic of price sales prediction is in this application
all the database connection query and logic is included in repository folder. it could be separated once our code base is getting biger

### bababos-frontend

its little vue js SPA application to showcase the functionality of the applications. the core function is to choose an sku_id and show how much we should sell it with some variables and considerations

## babaos-data 

its just a folder when compose is run


## How to run

### dockerized

this application is dockerized you could use orbstack or docker desktop to run it easyly
```
docker-compose up --force-recreate --build -d
```

that command will run all requirements
- backend (include seeding)
- frontend 
- psql
- migration


## Disable seeding

After first run and try the sample data you could disable the data seeding by commenting line seed in 
*bababos-backend/cmd/server.go*
```
...
seed.Seed()
...
```

