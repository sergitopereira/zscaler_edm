# zscaler_edm v1.0
As per [ creating exact data match](https://help.zscaler.com/zia/creating-exact-data-match-template)
Alphanumeric is Any combination of digits (0-9) or letters (a-z, A-Z). Including hyphens and underscores as character delimiters is supported. You must make sure that the data includes at least 3 characters but is no longer than 24 characters. 

Numeric: Any combination of digits (0-9). Including spaces and hyphens as character delimiters is supported. You must make sure that the data includes at least 3 characters but is no longer than 24 characters.

The present script removes invalid characters from the alphanumeric and numeric columns (which are defined in main.go). Moreover, removes fields that are less than 3 characters long and exceeding characters when they are longer than 24 characters.



# Installation Instructions

Download the binary from git

Linux OSX: 
wget -O zscaler_edm https://raw.githubusercontent.com/sergitopereira/zscaler_edm/master/zscaler_edm

WINDOWS:
wget -O zscaler_edm.exe https://raw.githubusercontent.com/sergitopereira/zscaler_edm/master/zscaler_edm.exe


# usage
```golang

./zscaler_edm {PATH_TO_CSV_FILE}

from the index tool:
.\/zscaler_edm {PATH_TO_CSV_FILE}

```

The script will generate  the following two files:

 - updated_file.csv: Updated cvs to be imported to index tool 
 - logs.txt: Contains all the changes made

# Logs

Logs are saved in logs.txt file

# Author 
Sergio Pereira 
Zscaler Professional Services


