local game = require("game")
local me = require("self")
local player = require("player")

me.FaceTowards(player.GetCoord())
game.Display("I'm a big boy with big boy powers!")
me.Walk("S", "N")