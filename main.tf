provider "example"{
    url_address="https://api.zoom.us/v2/users"
    token="eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImxOR0pCSGp1Uk9PRktDTTY4TGpIMGciLCJleHAiOjE2MTY1Njc1MTAsImlhdCI6MTYxNTk2MjcxMX0.qjKLu13bhsk47pMrJeECZdLjN9LSwJx4DzDDruaz658"
}

resource "example_user" "user1"{
   email="ekansh0786@gmail.com"
   first_name ="Ekansh"
   last_name="Singh"
} 