# About the structure
1. v1 use database to store test cases
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
        - database
    - Log
    - Generate report
6. Configuration used to store the config.ini or other config settings
7. Scripts is used for store test scripts
8. TestCases is used for store test cases in different format
9. TestReport is used for store the generated test report
10. should try to auto generate the struct based on the test case expected result and convert the response to that struct

# To do 2022/01/14
1. read the document on testerhome for api testing. They are stored in the API_Test folder in the browser