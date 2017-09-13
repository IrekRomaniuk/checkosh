#!/bin/bash
cd /var/scripts
# Skip 3 lines and print lines which do not have "-" in 3rd column and it does not start from 0. 
# and second column does not start from 10. 
echo "Writing All Robo info including names, external addresses and policy names to checkosh-lsm" 
LSMcli 1.1.1.1 user password Show -F=nibtp | tail -n +4 | awk '{ if ($3 != "-" && $2 !~ /^10\./ && $3 !~ /^0\./) print $0 }' > checkosh-lsm
echo "Wrting ALL non-Robo to checkosh-gws"
LSMcli 10.254.253.237 CPService IHateMundays1! Show -F=nibtp | tail -n +4 | awk '{ if ($3 != "-" && $2 !~ /^10\./ && $2 !~ /^0\./&& $3 !~ /^0\./) print $0 }' > checkosh-gws
echo "Writing Topology to checkosh-output.sh"
awk '{if ($2 ~ /^0\./) print "LSMcli 1.1.1.1 user password ShowROBOTopology " $1}' checkosh-lsm > checkosh-output.sh
chmod +x checkosh-output.sh 
echo "Executing checkosh-output.sh with output set to intermediate checkosh-int.tmp"
./checkosh-output.sh | grep 'ROBO\|1)' > checkosh-int.tmp
echo "Parsing checkosh-int.tmp to extract Robo names, internal addresses and comments into checkosh-int"
awk '/^ROBO/ { n = $2 } /^1\)/ { $1 = n; $3 = ""; print }' checkosh-int.tmp> checkosh-int
##### NOTES
# PRINT ALL GWs with addresses starting from 10.
#LSMcli LSMcli 1.1.1.1 user password Show -F=nibtp | awk '{ if ($3 != "-" && $2 ~ /^10\./) print $0 }'