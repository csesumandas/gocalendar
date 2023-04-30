# A simple calendar cli tool written in golang
Written this cli tool using golang. Used some basic features of golang.


#### Examples
##### To see help use the below command
```
$ gocalendar -h
Simple Command Line Calendar Tool Utility. Inspired by cal tool in Linux.
Author: Suman Das
Year: 2022

Usage; gocalendar
  -m int
    	Calendar Month Input
  -y int
    	Calendar Year Input

```

##### To see current month use the below command
```
$ gocalendar
    January 2022
Su Mo Tu We Th Fr Sa 
                   1 
 2  3  4  5  6  7  8 
 9 10 11 12 13 14 15 
16 17 18 19 20 21 22 
23 24 25 26 27 28 29 
30 31 

```

##### To see a particular year use the below command
```
$ gocalendar -y 2020
                               2020
       January               February                March          
Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   
          1  2  3  4                      1    1  2  3  4  5  6  7   
 5  6  7  8  9 10 11    2  3  4  5  6  7  8    8  9 10 11 12 13 14   
12 13 14 15 16 17 18    9 10 11 12 13 14 15   15 16 17 18 19 20 21   
19 20 21 22 23 24 25   16 17 18 19 20 21 22   22 23 24 25 26 27 28   
26 27 28 29 30 31      23 24 25 26 27 28 29   29 30 31               
                                                                     

        April                   May                   June          
Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   
          1  2  3  4                   1  2       1  2  3  4  5  6   
 5  6  7  8  9 10 11    3  4  5  6  7  8  9    7  8  9 10 11 12 13   
12 13 14 15 16 17 18   10 11 12 13 14 15 16   14 15 16 17 18 19 20   
19 20 21 22 23 24 25   17 18 19 20 21 22 23   21 22 23 24 25 26 27   
26 27 28 29 30         24 25 26 27 28 29 30   28 29 30               
                       31                                            

        July                 August               September        
Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   
          1  2  3  4                      1          1  2  3  4  5   
 5  6  7  8  9 10 11    2  3  4  5  6  7  8    6  7  8  9 10 11 12   
12 13 14 15 16 17 18    9 10 11 12 13 14 15   13 14 15 16 17 18 19   
19 20 21 22 23 24 25   16 17 18 19 20 21 22   20 21 22 23 24 25 26   
26 27 28 29 30 31      23 24 25 26 27 28 29   27 28 29 30            
                       30 31                                         

       October               November              December        
Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   
             1  2  3    1  2  3  4  5  6  7          1  2  3  4  5   
 4  5  6  7  8  9 10    8  9 10 11 12 13 14    6  7  8  9 10 11 12   
11 12 13 14 15 16 17   15 16 17 18 19 20 21   13 14 15 16 17 18 19   
18 19 20 21 22 23 24   22 23 24 25 26 27 28   20 21 22 23 24 25 26   
25 26 27 28 29 30 31   29 30                  27 28 29 30 31         

```

##### To see a particular month for a year use the below command
```
$ gocalendar -y 2021 -m 12
   December 2020
Su Mo Tu We Th Fr Sa 
       1  2  3  4  5 
 6  7  8  9 10 11 12 
13 14 15 16 17 18 19 
20 21 22 23 24 25 26 
27 28 29 30 31 

```
