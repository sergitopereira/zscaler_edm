# zscaler_edm v1.0
As per [ creating exact data match](https://help.zscaler.com/zia/creating-exact-data-match-template)
Alphanumeric is Any combination of digits (0-9) or letters (a-z, A-Z). Including hyphens and underscores as character delimiters is supported. You must make sure that the data includes at least 3 characters but is no longer than 24 characters. 

Therefore, the present script removes invalid characters from the alphanumeric column (which is defined in main.go). Moreober, removes fields that are less than 3 characters long and removed exceeding charactes when the alphanumeric word is longer than 24 characters.


# Installation Instructions

Download the binary from git

Linux OSX: 
wget -O zscaler_edm https://raw.githubusercontent.com/sergitopereira/zscaler_edm/master/zscaler_edm

WINDOWS:
wget -O zscaler_edm.ese https://raw.githubusercontent.com/sergitopereira/zscaler_edm/master/zscaler_edm.exe


# usage
```golang

./zscaler_edm {PATH_TO_CSV_FILE}

```
# Logs

Logs are saved in logs.txt file

# Author 
Sergio Pereira 
Zscaler Professional Services

