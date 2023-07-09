// Switch to 'gameDB' database
db = db.getSiblingDB("gameDB");

// Create a new collection named 'gameStats'
db.createCollection("gameStats");

// Insert documents into the 'gameStats' collection
db.gameStats.insert([
  {
    area_code: 123,
    game_modes: {
      battle_royal: 765,
      team_death_match: 983,
      capture_the_flag: 234,
    },
  },
  {
    area_code: 467,
    game_modes: {
      battle_royal: 765,
      team_death_match: 983,
      capture_the_flag: 234,
    },
  },
  {
    area_code: 789,
    game_modes: {
      battle_royal: 765,
      team_death_match: 983,
      capture_the_flag: 234,
    },
  },
  {
    area_code: 981,
    game_modes: {
      battle_royal: 765,
      team_death_match: 983,
      capture_the_flag: 234,
    },
  },
  {
    area_code: 982,
    game_modes: {
      battle_royal: 265,
      team_death_match: 483,
      capture_the_flag: 734,
    },
  },
  {
    area_code: 183,
    game_modes: {
      battle_royal: 165,
      team_death_match: 583,
      capture_the_flag: 604,
    },
  },
  {
    area_code: 184,
    game_modes: {
      battle_royal: 65,
      team_death_match: 53,
      capture_the_flag: 64,
    },
  },
  {
    area_code: 185,
    game_modes: {
      battle_royal: 365,
      team_death_match: 153,
      capture_the_flag: 764,
    },
  },
]);

db.gameStats.createIndex({ area_code: 1 });
