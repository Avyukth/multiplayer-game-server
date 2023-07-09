# lila-app Service Protocol Documentation

## Overview

The `lila-app` service provides statistical data on games based on the geographic area code.

## Service Definition

The service `lila-app` is defined with the following methods:

- `GetGameStats(GameStatsRequest) returns (GameStatsResponse)`: This method fetches game statistics based on the provided area code.

## Message Definitions

### GameStatsRequest

This message is sent to request game statistics. The message includes the following field:

- `area_code (int32)`: This field represents the geographic area code for which the game statistics are requested.

### GameStatsResponse

This message is received as a response to the `GameStatsRequest`. It includes the following fields:

- `area_code (int32)`: This field indicates the geographic area code for which the game statistics are provided.

- `game_modes (repeated GameMode)`: This field contains a list of `GameMode` objects that include the statistics for each game mode.

### GameMode

The `GameMode` message is used to encapsulate game mode-specific statistics. It includes the following fields:

- `mode (string)`: This field represents the type or mode of the game.

- `players (int32)`: This field represents the number of players playing the game in the particular mode.

## Usage

You can use the `GetGameStats` method to fetch game statistics for a particular area code. The response will include a list of all game modes, along with the number of players for each mode.

Please note that this is a gRPC protocol, and you will need to use a gRPC client to interact with this service.

#### NB: you can run clients.go to test the service , to run it use the command `cd clients && go run client.go` in the terminal

it is a concurrent client that sends 7 requests to the server concurrently and prints the response to the terminal

## Example

Here's an example of how to use this protocol:

Request:

```
{
  "area_code": 123
}
```

Response:

```
{
  "area_code": 123,
  "game_modes": [
    {
      "mode": "Battle Royale",
      "players": 765
    },
    {
      "mode": "Team Death Match",
      "players": 983
    },
    {
      "mode": "Capture The Flag",
      "players": 234
    }
  ]
}
```

In this example, for area code 123, there are 765 players playing in Battle Royale mode, 983 in Team Death Match mode, and 234 in Capture The Flag mode.
