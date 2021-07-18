# Schedule
An API for managing bus locations in real time.

## How it works
The whole system consists of multiple clients and one main server. Each bus will have a client, which sends updates about its location to the main server. The main server stores data of the location of each bus and supports a RESTful API which can be queried to get the location of buses.

## Requirements

### Client
- [ ] POST `/auth` To be sent to the main server for authentication. Once this in complete, the client is now authenticated to send location updates to the main server. (basic auth)
- [ ] POST `/location` Updates the main server on the location of the client. id, e.g. timestamp, longitude, latitude, status = NOT_STARTED, IN_PROGRESS, FINISHED, DISCONTINUED

### Server
- [ ] GET `/location` Returns locations of all buses
- [ ] GET `/location?id=a,b,c` Returns the location of the bus with the given ids, comma separated
- [ ] POST `/location` Update location of buses
- [ ] POST `/add` Add the buses and returns the data added (basic auth)
- [ ] PUT `/update` Update the bus with the given ids (basic auth, admin only)
- [ ] DELETE `/delete?id=a,b,c` Delete the bus with the given id (basic auth, admin only)

## Data structure
```json
{
    "id": "id1",
    "metadata": {
        "manufacturer": "man" 
    },
    "location": {
        "longitude": 10.101,
        "latitude": 5.304
    }
}
```

