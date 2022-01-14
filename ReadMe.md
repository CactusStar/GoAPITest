# About the structure
1. v1 use excel to manage test case, in the following, we can try to use xml, json. yaml, database to store test cases
    - we can also try if JMeter files(.jmx) can be imported as test case
2. The test case should contain multiple things include the api to test, request data, expected response
3. encapsulate the net/http library to send data/recieve data
4. The library to generate report
5. CommonFunctions used to encapsulate some existing functions which will be used personally include: 
    - Get
    - Post
    - Update
    - Delete
    - Read test case
        - excel
        - json
        - xml, yaml
        - database
    - Log
    - Generate report
6. Configuration used to store the config.ini or other config settings
7. Scripts is used for store test scripts
8. TestCases is used for store test cases in different format
9. TestReport is used for store the generated test report

# Question
1. Easy way to manage the test case(excel/json/yaml/database)?
2. Do we need to introduce database in the automation testing?

# To do 2022/01/14
1. read the document on testerhome for api testing. They are stored in the API_Test folder in the browser