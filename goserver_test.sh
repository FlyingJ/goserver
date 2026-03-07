#!/usr/bin/env bash

curl --verbose --request GET http://localhost:8080/admin/metrics
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > GET /admin/metrics HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 200 OK
# < Content-Type: text/html
# < Date: Sun, 01 Mar 2026 21:33:20 GMT
# < Content-Length: 112
# < 
# <html>
#   <body>
#     <h1>Welcome, Chirpy Admin</h1>
#     <p>Chirpy has been visited 0 times!</p>
#   </body>
# * Connection #0 to host localhost left intact
# </html>

curl --verbose --request GET http://localhost:8080/app/
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > GET /app/ HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 200 OK
# < Accept-Ranges: bytes
# < Content-Length: 65
# < Content-Type: text/html; charset=utf-8
# < Last-Modified: Sun, 22 Feb 2026 03:37:30 GMT
# < Date: Sun, 01 Mar 2026 21:34:21 GMT
# < 
# <html>
#   <body>
#     <h1>Welcome to Chirpy</h1>
#   </body>
# </html>
# * Connection #0 to host localhost left intact

curl --verbose http://localhost:8080/app/
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > GET /app/ HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 200 OK
# < Accept-Ranges: bytes
# < Content-Length: 65
# < Content-Type: text/html; charset=utf-8
# < Last-Modified: Sun, 22 Feb 2026 03:37:30 GMT
# < Date: Sun, 01 Mar 2026 21:34:21 GMT
# < 
# <html>
#   <body>
#     <h1>Welcome to Chirpy</h1>
#   </body>
# </html>
# * Connection #0 to host localhost left intact

curl --verbose http://localhost:8080/admin/metrics
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > GET /admin/metrics HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 200 OK
# < Content-Type: text/html
# < Date: Sun, 01 Mar 2026 21:34:51 GMT
# < Content-Length: 112
# < 
# <html>
#   <body>
#     <h1>Welcome, Chirpy Admin</h1>
#     <p>Chirpy has been visited 1 times!</p>
#   </body>
# * Connection #0 to host localhost left intact
# </html>

curl --verbose http://localhost:8080/admin/reset
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > GET /admin/reset HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 405 Method Not Allowed
# < Allow: POST
# < Content-Type: text/plain; charset=utf-8
# < X-Content-Type-Options: nosniff
# < Date: Sun, 01 Mar 2026 21:35:18 GMT
# < Content-Length: 19
# < 
# Method Not Allowed
# * Connection #0 to host localhost left intact

curl --verbose --request POST http://localhost:8080/admin/reset
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > POST /admin/reset HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 200 OK
# < Content-Type: text/plain; charset=utf-8
# < Date: Sun, 01 Mar 2026 21:37:52 GMT
# < Content-Length: 15
# < 
# * Connection #0 to host localhost left intact
# Hits reset to 0

curl --verbose --request POST http://localhost:8080/admin/reset
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > POST /admin/reset HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > 
# < HTTP/1.1 200 OK
# < Content-Type: text/plain; charset=utf-8
# < Date: Sun, 01 Mar 2026 21:37:52 GMT
# < Content-Length: 15
# < 
# * Connection #0 to host localhost left intact
# Hits reset to 0

curl --verbose --data '{}' http://localhost:8080/api/validate_chirp
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > POST /api/validate_chirp HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > Content-Length: 2
# > Content-Type: application/x-www-form-urlencoded
# > 
# < HTTP/1.1 400 Bad Request
# < Content-Type: application/json
# < Date: Sun, 01 Mar 2026 21:25:36 GMT
# < Content-Length: 25
# < 
# * Connection #0 to host localhost left intact
# {"error":"empty payload"}

curl --verbose --data '{"body":"This is my Chirp. There are many like it, but this one is mine."}' http://localhost:8080/api/validate_chirp
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > POST /api/validate_chirp HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > Content-Length: 74
# > Content-Type: application/x-www-form-urlencoded
# > 
# < HTTP/1.1 200 OK
# < Content-Type: application/json
# < Date: Sun, 01 Mar 2026 21:27:40 GMT
# < Content-Length: 14
# < 
# * Connection #0 to host localhost left intact
# {"valid":true}

curl --verbose --data '{"body":"This is my Chirp. There are many like it, but this one is mine.","payload":"this is some sneaky stuff"}' http://localhost:8080/api/validate_chirp
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > POST /api/validate_chirp HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > Content-Length: 112
# > Content-Type: application/x-www-form-urlencoded
# > 
# < HTTP/1.1 200 OK
# < Content-Type: application/json
# < Date: Sun, 01 Mar 2026 21:28:16 GMT
# < Content-Length: 14
# < 
# * Connection #0 to host localhost left intact
# {"valid":true}

curl --verbose --data 'body=abcdefg' http://localhost:8080/api/validate_chirp
# * Host localhost:8080 was resolved.
# * IPv6: ::1
# * IPv4: *********
# *   Trying [::1]:8080...
# * Connected to localhost (::1) port 8080
# > POST /api/validate_chirp HTTP/1.1
# > Host: localhost:8080
# > User-Agent: curl/8.5.0
# > Accept: */*
# > Content-Length: 12
# > Content-Type: application/x-www-form-urlencoded
# > 
# < HTTP/1.1 500 Internal Server Error
# < Content-Type: application/json
# < Date: Sun, 01 Mar 2026 21:30:16 GMT
# < Content-Length: 41
# < 
# * Connection #0 to host localhost left intact
# {"error":"unable to decode request body"}