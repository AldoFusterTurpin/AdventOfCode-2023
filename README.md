# [Advent of Code 2023](https://adventofcode.com/2023)

### Warning ⚠ 
I do Advent of Code to practice different approaches to problem solving and enjoy some Go coding. I really focus in code clarity and writing solutions that everybody could understand (even if you are new to Go, at least I try it), rather than looking for the most optimized solution. Sometimes I spend a bit more in the problem to leave a better code rather than writing the shortest possible solution in the minimum time since I am enjoying this, not competing.

If you want to follow my mental process, check the 
```sh
git log
```
and the README.

I do not claim to be an expert in any subject, I just love Go and I try to follow the craftsmanship philosophy.

Cheers, 

Aldo. 

🙂 👨‍💻 

### How to run the tests
```sh
go test ./...
```

### How to run dayN (where N is 1, 2 ...)
```sh
cd dayN
go build
./dayN
```

### Go is awesome
I use the great [embed](https://pkg.go.dev/embed) package to include a huge string in my unit tests, which is pretty nice for better readability and ensure my unit tests are not touching I/O (as the string is included in the binary when using the embed package 👌).

Do ctrl + f with "//go:embed" to see examples.

### Considerations
#### Day1 considerations
For the day1, I followed an inside-out approach, starting from the details of the problem and building small functions in a TDD way from scratch, obtaining a fast feedback loop between the Red and the Green steps (advantage). 

It has the drawback that some unit tests could be erased because once you ensure your function A (which invokes multiple times B()) is correct, you could delete B as it is just a substep of A, but I prefer to keep those unit tests as they don't hurt. You can also understand better what does each function with those redundant tests as they are fine grained. 


#### Day2 considerations
For the day2, and to change the previous approach, I decided to follow an outside-in strategy for the TDD lifecycle. I just created the test "TestIsGamePossible" and I started building the solution from the highest function ("isGamePossible") narrowing down at each step to the more specific one ("isSingleGuessPossible"). You can clearly see that if you read the day2.go from top to bottom.

I usually like to sort the code from highest level of abstraction to lowest level so you can read the functions from the top of the file to the tail without thinking about the details of the how and just focus on the what.

In this case, the outside-in approach has the disadvantatge that the feddback loop you get from the Red-Green-Refactor cycle is worse as the amount of code you need to write before the test becomes green is more, compared to inside-out (in a different type of problem using Stubs/Mocks and components interactions that would be a different story). 

We could even considered that this time I just did Test Frist Development rather than TDD as after having the first red test, I just wrote stright the whole solution for the first test case but in a general form. After that, I included the other test cases for the function, ensured for each one they were red due to the expected reason, and then changing the expected value to make them become green, one by one.

After all the tests where green, I performed a series of refactors, and after each one, I just executed all the tests, having confidence in the path I was following.

I decided to sort the test cases for "TestIsGamePossible" puting first all the "is possible" test cases and then the "not possible" as you can follow easily with the problem explanation on the web.

I did not created a unit test for "sumOfIDsOfThePossibleGames" as I thought the implementation was really traightforward and I just decided to write the function and execute the main(), passing it the sample statement as input to verify the output. If that had failed, I would probably have written a unit tests for that and started building it (but it was trivial).

For the whole part 2, I just decided to revisit my solution for part 1 and duplicate and modify the code to incldue the counter for the colors. I just wrote the whole day2_part2.go file and run the main with the statement as input file .It worked at the first try so I just used the input.txt file and provide the solution in the web. Part2 worked at the first try, but it does not always happen.

I thik despite the fact I was not doing TDD in a strict form, the fact that I was following an outside-in approach writing really small functions at each step just focusing on one thing at a time helped me to get the solution at the first try. 

Last note: I could have used regex to extract the values of the subsets, but I found the iterative approach of the "split" really straightforward and easy to follow. I could have used structs to represent each of the games instead of plain strings but it was really easy to follow using simple strings. Let's see what I decide for the next problem....

#### Day3 considerations
Sometimes I use as input for my unit tests an string when I should/could use directly primitives from the language (like [][]rune or []int) but it is a lot easier to copy + paste from the problem statement rather than converting it. In other words, I could simply call convertStringToMatrix() just once and pass arround [][]rune rather than converting twice from string to [][]rune, but I want to copy + paste easily (or use the embed package) and the effort is not worth right now (not enough time to do so now...).

In the getGearRatioOfGear() function (check also getGearRatioOfGear.png), I am sure I could have extracted the common behavior to a function and then just iterate over a number of relative indexes and delete a lot of the lines that exist. Nonetheless, as the code right now maps directly to the mental model to solve the problem, I thing it is good enough as I need to keep working in the problem of the next days. It does not affect the time complexity, so it is not that bad. Also, I know there are some redundant assignments in that function regarding x and y, but I prefer to repeat the assignment of those rather than look upper in the code to knwo what value had (eg: on line 85 of day3/day3_part2.go, I do not need to assign again the x, but I prefer to have both x and y assigned again before any check as it is more straightforward and clear).

Lastly, I am almost sure there more clever ways to implement a solution to the problem, but the one I followed is the one I found easy to understand. Iterate once over the input to construct the map of Coordinates to the numbers and then iterate again the input and use that map to get the gear ratio of each gear.

### Interesting links
[Runes in Go](https://exercism.org/tracks/go/concepts/runes)