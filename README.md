- Game - main entity, belongs to the logged in user; contains list of boards
- Board two dimensional board with player in each cell. Board generate rounds for each player with neighbours. Each player has rounds with up, down, left, rigth and diagonal players(8 rounds total). Board is endless. Players in the very bottom row will compete with player on the top.
- Player entity based on behavior model. Can interact with other player in the round. 
- Model behavior function that response with cooperation(Green) or confrontation(red) based on previous own signals, oponent signals and aproximate maximum interactions quantity.
- Round entity contains two players - left and right. Players interacting with each other by exchangins signals. 
- Interaction base player comunication entity. Represented by two signals from the players in the round.