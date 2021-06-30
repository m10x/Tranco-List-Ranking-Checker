# Tranco-List-Ranking-Checker
[![Release](https://img.shields.io/github/release/m10x/Tranco-List-Ranking-Checker.svg?color=brightgreen)](https://github.com/m10x/Tranco-List-Ranking-Checker/releases)
[![Exe](https://img.shields.io/badge/.exe-Windows%2Famd64-orange)](https://github.com/m10x/Tranco-List-Ranking-Checker/releases/)
[![Elf](https://img.shields.io/badge/elf-Linux%2Famd64-orange)](https://github.com/m10x/Tranco-List-Ranking-Checker/releases/)

# What does this do?
The Tranco-List-Ranking-Checker takes two files as input.  
A [tranco-list](https://tranco-list.eu/) and a list with URLs (e.g. a bugbounty URL list).  
The programm iterates through the URL list and checks if those URLs are on the tranco-list.  
It also checks if the tranco-list rank of this URL is within the specified threshold.  
The default threshold is 1000 - to only receive URLs within the Top 1000.  
At the end of the iteration a sorted list (best to worst rank) of the eligible URLs is printed.

# Info
The tranco-list has the following format:  
1,www.google.com  
2,www.facebook.com

The URL list can have the following formats:  
www.google.com  
http://www.google.com/  
https://www.google.com/  
http://www.facebook.com  
https://www.facebook.com

# Usage:
## Windows (amd64)
[Download here from releases](https://github.com/m10x/Tranco-List-Ranking-Checker/releases/)  
.\tranco-checker.exe -u C:\path\to\urllist.txt -t C:\path\to\tranco.csv

## Linux (amd64)
[Download here from releases](https://github.com/m10x/Tranco-List-Ranking-Checker/releases/)  
chmod +x .\tranco-checker  
.\tranco-checker -u path/to/urllist.txt -t /path/to/tranco.csv

## Go
go run path/to/tranco-checker.go -u path/to/urllist.txt -t /path/to/tranco.csv

## Optional flags:
- -threshold int => Ranking threshold. Only print URLs whose ranking is within the threshold. Default is 1000.
- -threads int => Threads to use. Default is 10.

# Possible TODOs:
- write sorted output to file
- ?
