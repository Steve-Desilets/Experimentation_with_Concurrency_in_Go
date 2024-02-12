# Experimentation With Concurrency In Golang

For the project included within this Github repository, we imagine that we are a data scientist at a company whose leadership team is interested in leveraging Golang's concurrency features, such as goroutines and channels, in order to reap Golang's parallel processing computational benefits. Specifically, this company's leadership team is interested in studying how parallel processing can impact the company's operational efficiency when conducting data science analyses, such as calculating regression lines for datasets.  Accordingly, the company's leadership requests that the data scientist conduct an experiment to determine the computational efficiency of Golang when calculating regression lines, along with their mean squared error (MSE) and Akaike Information Criterion (AIC), when leveraging and when not leveraging Golang's concurrency features.

For this study, the data scientist leverages a dataset called "boston" that includes information about each of the 506 census tracks in Boston as of the time of the study  (Belsley, Kuh and Welsch 1980; Harrison and Rubinfeld 1978).  The variables included in this dataset are:
1. mv - Median Value of Homes in thousands of 1970 US dollars
2. nox - Air Pollution (nitrogen oxide concentration)
3. crim - Crime Rate
4. zn - Percentage of Land Zoned for Lots
5. indus - Percentage of Business that is Industrial or Nonretail
6. chas - On the Charles River (1) or not (0)
7. rooms - Average Number of Rooms per Home
8. age - Percentage of homes built before 1940
9. dis - Weighted Distance to Employment Centers
10. rad - Accessibility to Radial Highways
11. tax - Tax Rate
12. ptratio - Pupil / Teacher Ratio in Public Schools
13. lstat - Percentage of Population of Lower Socio-economic Status

For this study, the data scientist created the two Golang programs included within this repository - "Boston_Experiment_With_Concurrency.go" and "Boston_Experiment_No_Concurrency.go" - which complete the same task with the difference being that one program leverages goroutines for concurrency and the other does not leverage Golang's concurrency features.  Each program reads and parses the boston.csv file, calculates the regression line to predict median home values with every possible combination of four or more explanatory variables, and prints the results (including the MSE and AIC). Each program completes this exercise 100 times and records the time to calculate all those regression lines for each of the 100 trials. At the end, each program prints the average trial runtime and the sum of all the trial runtimes to an output text file - both of which are included in this Github repository.

After completing this study, we see that the average trial runtimes when leveraging and not leveraging concurrency were 1,053,241 microseconds and 966,108 microseconds, respectively.  Therefore, in this case, the program actually ran 8% more quickly when not leveraging Golang's concurrency features.  This finding underscores the fact that even though concurrency can often help improve the computational efficiency of programs, sometimes concurrency does not really yield strong computational benefits.  If I were the data scientist making a recommendation to the leadership team of this company based on the findings of this study, I would recommend that the company's data scientists take a targeted approach to identifying situations in which concurrency would likely result in significant computational efficiency benefits because the hardware and algorithm allow for parallel processing.  For example, programs that repeatedly read information from or to a disk / network would be ideal candidates for leveraging concurrency since reading and writing such information is often thousands of times slower than other programming processes.  The rationale for recommending that programmers only pursue concurrency in such specific situations is that leveraging Golang's concurrency features like goroutines and channels usually decreases the readability of one's code - even if it does sometimes result in computational benefits in specific situations. By undertaking such an intentional strategy for identifying opportunities to utilize concurrency, programmers could balance the often dueling desires for computational efficiency and code simplicity.


References

Belsley, David A., Edwin Kuh, and Roy E. Welsch. 1980. Regression Diagnostics: Identifying Influential Data and Sources of Collinearity. New York: Wiley. 

Harrison, David and Daniel L. Rubinfeld. 1978. "Hedonic housing prices and the demand for clean air." Journal of Environmental Economics and Management, 5:81–102.

