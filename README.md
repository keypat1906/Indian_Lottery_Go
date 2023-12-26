                                 Indian lottery Problem

The problem

This problem is related to the Indian lottery. In case you are not familiar with it: players pick 5 distinct numbers from 1 to 90. There is a weekly lottery picking event when the lotto organization picks 5 distinct numbers randomly between 1 and 90 – just like the players did. The player’s reward then depends on how many of the player’s numbers match with the ones selected at the lotto picking. A player wins if he/she has 2, 3, 4 or 5 matching numbers.
Now, the problem: at the lottery event, right after picking the numbers, a computer shall be able to report quickly how many winners are in each category, for example:
This report shall be generated within a couple of seconds after picking the winner number. The player’s numbers are known in advance – at least 1 hour ahead of the show. In peak periods there are about 5 million players, but to be sure we can assume it does not exceed 10 million.

Technical specification

Write a console application in a freely chosen programming language that can be compiled on Linux. Your application will be called like this:
./yourapp input.txt
where input.txt file exists in the same folder and is an ascii file, in which each line contains 5 space separated integers (in the range of 1-90) representing one player’s numbers.

When your application finished processing the player’s dataset from the file, it should write a line to
the standard output like this: 
READY

Note that it should be newline terminated. After that, the program may receive multiple lines (identical to the file’s lines) representing the lottery’s picks and it should be able to report 4 space separated numbers on the standard output as fast as possible (line should be newline terminated). The four numbers shall mean the number of winners with 2, 3, 4 and 5 matches respectively.

                                   Solution
Introduction:
I have created this program in Go and it is using argument as a file which contains multiple lines of numbers. My program parses each line and calculate the total number of winners who won 2, 3, 4 and 5 numbers.


How to run:
You can run using Go source file which I have attached in the src folder. You need Go installed on your machine.

    pick_winner.py test.txt


Input validation:
If you enter other than number, it throws error and return. 
It only accepts 5 numbers so there is validation to check if numbers are not more than 5, otherwise it returns the program. You can enter less than 5 numbers though. 


Test results:
This program is heavily CPU bound so result may vary based on the machine we run. I developed and tested this code on my 2015 Macbook Pro which has CPU configuration 2.8 GHz Quad-Core Intel Core i7 with 4 CPU Cores and it has 16 GB RAM.

This is sample screenshot of my tests. You can see it also lists how much time it takes to pull the results. Based on the sample file provided which has 10 million records, it takes average 5 to 6 seconds to calculate the results of 10 million records. 

This is sample output once you enter the lottery winners

78 6 19 10 62
Winner Numbers: ['78', '6', '19', '10', '62']

--------------------------------------
 Numbers matching   |      Winners   

--------------------------------------

 5                  |       1

 4                  |       61

 3                  |       4056

 2                  |       112196

--------------------------------------
Total time taken to pull winners 5.23 seconds


READY
 

Time and Space complexity

Time complexity is easily O(N) where N is number of input lines we feed to the program. In addition I used dictionary to calculate and store the results which is O(1) best case and worst case is O(N). That means total time complexity is O(N) + O(1)  best case scenario.

For space, I am using dictionary to store 5 numbers it so it takes O(N) space complexity where N=5


Improvement 

This is most optimized code in the Go language. I have used memory map to process the large file which is 145 MB size. Any input or criticism welcome. 

