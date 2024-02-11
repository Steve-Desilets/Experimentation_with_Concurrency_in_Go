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








References

Belsley, David A., Edwin Kuh, and Roy E. Welsch. 1980. Regression Diagnostics: Identifying Influential Data and Sources of Collinearity. New York: Wiley. 

Harrison, David and Daniel L. Rubinfeld. 1978. "Hedonic housing prices and the demand for clean air." Journal of Environmental Economics and Management, 5:81–102.





Accordingly, the data scientist in this study created the "web_crawling_and_scraping" Go script included within this repository to crawl across the Wikipedia pages for the following topics and scrape information from them: Robotics, Robot Reinforcement Learning, Robot Operating System, Intelligent Agent, Software Agent, Robotic Process Automation, Chatbot, Applications of Artificial Intelligence, and Android (Robot).  For each of these webpages, the data scientist extracted the webpage's title, the webpage's URL, and the text from the body of the webpage.  The Go program then compiled this data into a JSON for each webpage and exported all the scraped data into the "scrapedOutputGo" output julia file. 

If I were the data scientist making a recommendation to the leadership team of this company based on the findings of this study, I likely would recommend that the company begin leveraging Go instead of Python for its data engineering projects.  That's because previous experiments at this company (as described in the Command_Line_Applications and Benchmarking-Project Github repositories) have proven that Go is a much faster language computationally than Python is.  Given that Go can also produce equally excellent outputs when crawling and scraping the internet, Go should be the data engineering language of choice for this company.
