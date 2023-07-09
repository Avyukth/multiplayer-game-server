// Switch to 'gameDB' database
db = db.getSiblingDB('gameDB');

// Create a new collection named 'gameStats'
db.createCollection('gameStats');

// Insert documents into the 'gameStats' collection
db.gameStats.insert([
  {
    "area_code": 123,
    "game_modes": {
      "battle_royal": 10,
      "team_death_match": 20,
      "capture_the_flag": 30,
    }
  }])
