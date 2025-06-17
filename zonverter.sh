!/bin/bash

tmux new-session -d -s zonvert

tmux send-keys "nvim ~/Code/Zonverter" C-m
tmux rename-window "Code"

tmux new-window -t zonvert:2 -n "terminal"
tmux send-keys "nvim ~/Code/Zonverter -c 'terminal'" C-m

tmux attach -t zonvert
