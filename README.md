# **HTTP mocker configuration service**

This service persists the mocked responses to the requests received from the user.
<br>
<br>
___
### **Notes**

1. Body always as JSON
2. **Configuration**
   1. **CREATE**
      1. Tag (to group with other configurations)
      2. Request
         - Endpoint URI
         - Method
         - URI params
         - Body response
           - Status code
           - Headers
           - Body
   2. **List**
      1. Tag (optional)
      2. Method
   3. **UPDATE**
      1. Id
      2. Tag (to group with other configurations)
      3. Request
          - Endpoint URI
          - Method
          - URI params
          - Body response
            - Status code
            - Headers
            - Body

   4. **DELETE**
       1. Id