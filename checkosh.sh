#!/bin/bash
cd /var/scripts
# Skip 3 lines and print lines which do not have "-" in 3rd column and 3rd column does not start from 0. and second column does not start from 10. 
echo "Writing All Robo to checkosh1 " 
LSMcli 1.1.1.1 user password Show -F=nibtp | tail -n +4 | awk '{ if ($3 != "-" && $2 !~ /^10\./ && $3 !~ /^0\./) print $0 }' > checkosh1
echo "Writing Topology to checkosh-output.sh"
awk '{if ($2 ~ /^0\./) print "LSMcli 1.1.1.1 user password ShowROBOTopology " $1}' checkosh1 > checkosh-output.sh
echo "Running checkosh-output.sh with output set to checkosh2"
chmod +x checkosh-output.sh 
./checkosh-output.sh | grep 'ROBO\|1)' > checkosh2
##### NOTES
# PRINT ALL GWs with addresses starting from 10.
#LSMcli LSMcli 1.1.1.1 user password Show -F=nibtp | awk '{ if ($3 != "-" && $2 ~ /^10\./) print $0 }'