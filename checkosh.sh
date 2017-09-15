#!/bin/bash
# Source the Check Point profile for library and paths settings
#
export `grep "CPDIR_PATH=" /etc/init.d/firewall1`
[ -f $CPDIR_PATH/tmp/.CPprofile.sh ] || {
    echo "--- Fatal error: cant find CPprofile.sh !!"
    # We are unable to setup essential variables
    echo "Unable to setup needed variables. Stopping." >> $logfile
    exit 2
}
source $CPDIR_PATH/tmp/.CPprofile.sh
cd /var/scripts
# Skip 3 lines and print lines which do not have "-" in 3rd column and it does not start from 0. 
# and second column does not start from 10. 
echo "Writing All Robo info including names, external addresses and policy names to checkosh-lsm" 
LSMcli 1.1.1.1 user password Show -F=nibtp | tail -n +4 | awk '{ if ($3 != "-" && $2 !~ /^10\./ && $3 !~ /^0\./) print $0 }' > checkosh-lsm
echo "Wrting ALL non-Robo to checkosh-gws"
LSMcli 10.254.253.237 CPService IHateMundays1! Show -F=nibtp | tail -n +4 | awk '{ if ($3 != "-" && $2 !~ /^10\./ && $2 !~ /^0\./&& $3 !~ /^0\./) print $0 }' > checkosh-gws
echo "Writing print network_objects to checkosh-gws.dbedit "
cat checkosh-gws | awk '{ print "print network_objects " $1}' > checkosh-gws.dbedit
echo "Writing Topology to checkosh-output.sh"
awk '{if ($2 ~ /^0\./) print "LSMcli 1.1.1.1 user password ShowROBOTopology " $1}' checkosh-lsm > checkosh-output.sh
chmod +x checkosh-output.sh 
echo "Executing checkosh-output.sh with output set to intermediate checkosh-int.tmp"
./checkosh-output.sh | grep 'ROBO\|1)' > checkosh-int.tmp
echo "Parsing checkosh-int.tmp to extract Robo names, internal addresses and comments into checkosh-int"
awk '/^ROBO/ { n = $2 } /^1\)/ { $1 = n; $3 = ""; print }' checkosh-int.tmp> checkosh-int
mdsenv CMA
echo "Printing network_objects"
#dbedit -local -ignore_script_failure -globallock -f checkosh-gws.dbedit | grep 'Object Name:\|manual_encdomain:\|comments: \+.\+' > checkosh-gws.tmp
dbedit -local -ignore_script_failure -globallock -f checkosh-gws.dbedit | grep 'Object Name:\|manual_encdomain:' > checkosh-gws.tmp1
#awk '{line1=$0; getline line2; getline line3; print line1, line3, line2}' checkosh-gws.tmp > checkosh-gws.net
awk '{line1=$0; getline line2; print line1, line2}' checkosh-gws.tmp1 > checkosh-gws.tmp2
awk '{ print $3 " " $6 }' checkosh-gws.tmp2 > checkosh-gws.net
##### NOTES
# PRINT ALL GWs with addresses starting from 10.
#LSMcli LSMcli 1.1.1.1 user password Show -F=nibtp | awk '{ if ($3 != "-" && $2 ~ /^10\./) print $0 }'