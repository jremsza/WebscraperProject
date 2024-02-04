# Webscraper Project
#Progam Use
To use the program navigate to the directory on a bash command line where robotwiki.exe is stored and run the following command:
./robotwiki.exe
items.jl file populates in that directory and contains the text scraped from the web.

#Program Walktrhough
This project is a web scraper built in Go using the Colly framework.

1. Creates a JSON Lines file, ensuring the file is closed properly at the end of operations using a defer statement.
2. Instantiates the Colly web scraping framework, denoted as c. 
3. A call back function is created to capture the title of the page.
4. A second call back is used to capture the body of the webpage. Within the call back function the stored title is retrived, the pages' URL and body text. The body text is cleaned of HTML tags and whitespace by calling the cleanHTML function that uses regex and trim.
5. A PageInfo instance is used for storing the URL, title and body content.
6. The writeJSON function marshals PageInfo into JSON format.

#CPU Times
Measurements of CPU execution times for both the Go and Python web scraping programs were conducted on the command line, with the results documented in the time_report.txt file. It is evident from the results that Go executes these operations significantly faster than Python, demonstrating its superior efficiency for this operation. It is recomended to perform web scraping opertaions using Go as the primary appraoch based off of this project result.
