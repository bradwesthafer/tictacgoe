# TicTacGoe

This repo contains a basic implementation of TicTacToe written in Go. It can
be played on a square board of 3x3 or larger with the option to play with
either 1 or 2 players. If played 1 player, the CPU player uses the following
logic to make moves:

1) If there is a winning move, take it.
2) If you can block the player from winning, do so.
3) Otherwise, pick an open space at random.

The game is run out of a command line interface and includes ASCII art
representing  the current state of the game board.

# License

This code is © 2024 Brad Westhafer and may be used or modified for any
non-commercial purpose. It may not be used for any commercial purpose. It is
provided without any warranties whatsoever whether explicit or implied and
the author makes no claim that it is fit for any purpose whatsoever. Should
the law of any user's jurisdiction prohibit this total disclaimer of
warranty, then said user's license to use the software will be immediately
revoked. All derivative works must include this license.

Should you wish to use this code for commercial purposes including, but not
limited to, selling this game or using it to train an LLM or other
generative AI model, you must contact the author to purchase a license.
Failure to do so is copyright infringement and the author retains all rights
including the right to pursue legal action.