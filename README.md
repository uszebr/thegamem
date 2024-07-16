# Game theory Simulation

- Game - main entity, belongs to the logged in user. One user can have only one game(for now). Game contains list of boards
- Board two dimensional board with player in each cell. Board generate rounds for each player with neighbours. Each player has rounds with up, down, left, rigth and diagonal players(8 rounds total). Board is endless. Players in the very bottom row will compete with player on the top.
- Player entity based on behavior model. Can interact with other player in the round. 
- Model behavior function that response with cooperation(Green) or confrontation(red) based on previous own signals, oponent signals and aproximate maximum interactions quantity.
- Round entity contains two players - left and right. Players interacting with each other by exchangins signals. 
- Interaction: Base player communication entity. Represented by two signals from the players in the round.
- After Interaction each player get score based on oponent and own signal combination. Cooperation-Cooperation Confrontation-Confrontation Cooperation-Confrontation.
- Players exchange signals based on their previous history of interactions, which includes a list of signals and an approximate quantity of interactions. Models do not receive the exact quantity of interactions to avoid the "Last Round Always Red" issue.
- Rotation: An integer representing the number of players that win or lose on the board. The new board will be created with all players except the losers. Instead each loser will generate a new player based on the model of the corresponding winner.
- The Supabase URL and key are stored in environment variables. These are necessary for authentication and storing games (in the future).