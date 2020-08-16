# Soldier Magazine Problem

- A Soldier want magazine needed to be full ammo before war.
- A Soldier will load bullets randomly into the magazine in gun.
- A Soldier can put multiple magazine into one gun.
- Only one magazine will be mark as verified, and soldier will stop put/remove magazine until it
  verified.
- Verified will happen when the magazine is full.
- You must test the magazine is full or not with firing it on a gun.

## Requirements

- Go 1.13+
- Go mod (for dependencies)

update dependencies:
- go mod download
- go mod verify
- go mod tidy

## Technology
- Go 1.13+
- SQLlite3 
- Go Gin Framework
- Go Mock 

## Run Application
ensuring all dependencies with go mod

build application with script `go build main.go`

run application `./main`

run app with auto build `go run main.go`

## Run Tests
### Using Framework
execute this command to run test with go test framework

`go test github.com/mariojuzar/soldier-magazine/tests`

## API Documentation
1. Create Soldier
   `POST /api/v1/soldier`
        
   request body
        
    ```
   {"name" : "desired soldier name"}
   ```

2. Update Soldier
   `PUT /api/v1/soldier`
        
   request body
        
    ```
   {"id": id, name" : "desired soldier name"}
   ```

3. Get All Soldier
   `GET /api/v1/soldier`

4. Get Soldier
   `GET /api/v1/soldier/:id`
   
5. Delete Soldier
   `DELETE /api/v1/soldier/:id`
   
6. Put Magazine Into Soldier 
   `POST /api/v1/magazine`

   Magazine must give its max capacity. This magazine put into Magazine Packs in Soldier. All magazine will be in unverified state.
   
   request body
           
   ```
    {
        "soldier_id": 1,
        "magazines": [
            {"max_capacity": 6},
            {"max_capacity": 4}
        ]
    }
   ```
   
   sample expected response
   
   ```
   {
       "server_time": "2020-08-17T06:27:49.81794+07:00",
       "code": 200,
       "message": "OK",
       "data": {
           "soldier_id": 1,
           "soldier_name": "Soldier Name",
           "magazine_packs": {
               "magazines": [
                   {
                       "capacity": 0,
                       "max_capacity": 6,
                       "is_verified": false
                   },
                   {
                       "capacity": 0,
                       "max_capacity": 4,
                       "is_verified": false
                   }
               ]
           }
       }
   }
   ```
   
7. Load Bullet Into Magazine 
   `POST /api/v1/magazine/load`
   
   Given list of bullets that will be insert into magazine. You can set bullet randomly. Soldier will put it in magazine by order from its Magazine Pack.
   
   request body
   
   ```
   {
       "soldier_id": 1,
       "bullets": [3, 2, 5, 1]
   }
   ```
   
   sample expected response indicate that still have no verified magazine
   
   ```
   {
       "server_time": "2020-08-17T06:29:38.13423+07:00",
       "code": 200,
       "message": "OK",
       "data": {
           "soldier_id": 1,
           "is_verified_magazine": false
       }
   }
   ```
   
   sample expected response if there is a verified magazine
   
   ```
   {
       "server_time": "2020-08-17T06:31:17.593377+07:00",
       "code": 200,
       "message": "OK",
       "data": {
           "soldier_id": 1,
           "is_verified_magazine": true
       }
   }
   ```
   
   to make sure, you can check it into soldier data by using get soldier by id. Sample response:
   
   ```
   {
       "server_time": "2020-08-17T06:32:25.545783+07:00",
       "code": 200,
       "message": "OK",
       "data": {
           "soldier_id": 1,
           "soldier_name": "Soldier Name",
           "magazine_packs": {
               "magazines": [
                   {
                       "capacity": 3,
                       "max_capacity": 6,
                       "is_verified": false
                   },
                   {
                       "capacity": 4,
                       "max_capacity": 4,
                       "is_verified": true
                   }
               ]
           }
       }
   }
   ```
   
   Only one magazine that will be marked as verified, and then soldier stop to load bullet into magazine. 
   Everytime soldier put bullet into magazine he check it by firing it into a gun to know its verified or not.