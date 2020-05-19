# Notelet Server
A simple note manegemnt system for creating, storing, viewing and searching useful notes. 
The service provides REST Api for management of the notes and template suitable for SPA (single-page application) client.  
It is written in Go programming language.
## Routes
Notelet server service can be acessed:
* > '<domain_name> /

for SPA template;
* > <domain_name>/api/

for REST Api.
## REST Api
REST Api provides the following services:
| HTTP verb     | Path          | Service |
| ------------- |:-------------:| -----:|
| GET    | /notes | Recieve the list of all notes |
| GET    | /notes/{id} | Retrieve a note with the given id |
| POST   | /notes      |    Add a new note |
| DELETE | /notes/{id} | Remove a note with the given id